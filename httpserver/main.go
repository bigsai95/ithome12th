package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/bigsai95/ithome12th/mypb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var grpcClient pb.MyprotoServiceClient

func init() {
	//log輸出為json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//輸出設定為標準輸出(預設為stderr)
	logrus.SetOutput(os.Stdout)
	//設定要輸出的log等級
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	osc := make(chan os.Signal, 1)

	logrus.Info("http server start")

	//GRPC 連線
	grpconn, err := grpc.Dial("gs:8081", grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("grpc create error: %v", err)
	}
	defer grpconn.Close()

	grpcClient = pb.NewMyprotoServiceClient(grpconn)

	go webserver()

	signal.Notify(osc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	select {
	case <-osc:
		logrus.Printf("http server退出訊號: %s", <-osc)
		for t := 3; t > 0; t-- {
			logrus.Printf("%d秒後退出", t)
			time.Sleep(time.Duration(1) * time.Second)
		}
		return
	}
}

func webserver() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!"))
	})
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/user/add", userAdd)
	http.HandleFunc("/user/login", userLogin)
	http.HandleFunc("/user/list", userList)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
