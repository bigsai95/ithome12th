package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buger/jsonparser"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/olivere/elastic/v7"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

var (
	elkclient *elastic.Client
	elkip     string = "http://elasticsearch:9200/"
	elkindex  string = "index*" //索引名字
	bot       *tgbotapi.BotAPI
)

const (
	chatID   = 123456
	youToken = "123456:AABBBCC"
)

func main() {
	osc := make(chan os.Signal, 1)

	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
		},
	})
	logrus.SetOutput(os.Stdout)
	logrus.Info("elk watch start")

	setTgBotApi()
	setElkClient()

	go cronExec()

	signal.Notify(osc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	select {
	case <-osc:
		logrus.Printf("elk watch 退出訊號: %s", <-osc)
		for t := 3; t > 0; t-- {
			logrus.Printf("%d秒後退出", t)
			time.Sleep(time.Duration(1) * time.Second)
		}
		return
	}
}

//設定排程
func cronExec() {
	c := cron.New()
	c.AddFunc("0 * * * * *", func() {
		go getELKLog()
	})
	c.Start()
	select {}
}

func sendMsg(msg string) {
	NewMsg := tgbotapi.NewMessage(chatID, msg)
	NewMsg.ParseMode = tgbotapi.ModeHTML //傳送html格式的訊息
	_, err := bot.Send(NewMsg)
	if err != nil {
		logrus.Error("Send telegram message error")
	}
}

func setTgBotApi() {
	var err error
	bot, err = tgbotapi.NewBotAPI(youToken)
	if err != nil {
		logrus.Fatal(err)
	}
	bot.Debug = false
}

func setElkClient() {
	var err error
	elkclient, err = elastic.NewClient(
		elastic.SetURL(elkip),
		elastic.SetSniff(false),
	)
	if err != nil {
		logrus.Fatalf("Failed to create elkclient: %v", err)
	}
}

func getELKLog() {
	var (
		t  string
		nt string
	)
	status := make([]interface{}, 5) //要監控的log等級
	status[0] = "400"
	status[1] = "500"
	status[2] = "600"
	status[3] = "700"
	status[4] = "800"

	t = time.Now().UTC().Add(-1 * time.Minute).Format(time.RFC3339)
	nt = time.Now().UTC().Format(time.RFC3339)
	logrus.Infof("elk檢查開始: %s ~ %s", t, nt)

	rQuery := elastic.NewRangeQuery("@timestamp") //查詢的時間區間
	rQuery.Gte(t)
	rQuery.Lte(nt)

	query := elastic.NewBoolQuery()
	query.Must(rQuery)
	query.Filter(elastic.NewTermsQuery("content.severity", status...)) //查詢的log等級
	res, err := elkclient.Search().Index(elkindex).
		Query(query).
		Size(20). //查詢的筆數，預設為10筆
		Do(context.Background())
	if err != nil {
		logrus.Error("elk failed: ", elkindex, err)
	}
	for _, hit := range res.Hits.Hits {
		d, _ := hit.Source.MarshalJSON()
		msg, _ := jsonparser.GetString(d, "message") //取得錯誤訊息
		logrus.Info("elk msg: ", msg)
		sendMsg(msg)
	}
}
