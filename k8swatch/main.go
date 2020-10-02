package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

type deploymentInfo struct {
	name              string
	revision          string
	availableReplicas int32
}

var (
	k8sClientset *kubernetes.Clientset
	unAuth       string = "Unauthorized"
	kubeconfig   string = "/root/.kube/config"
	K8SNameSpace string = "dev-ithome"
	bot          *tgbotapi.BotAPI
)

const (
	chatID   = 123456
	youToken = "123456:AABBBCC"
)

func main() {
	osc := make(chan os.Signal, 1)

	logrus.Info("k8s watch start")
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "severity",
		},
	})
	logrus.SetOutput(os.Stdout)

	setTgBotApi()
	setK8SClient()

	go getPods(k8sClientset)

	signal.Notify(osc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	select {
	case <-osc:
		logrus.Printf("k8s watch 退出訊號: %s", <-osc)
		for t := 3; t > 0; t-- {
			logrus.Printf("%d秒後退出", t)
			time.Sleep(time.Duration(1) * time.Second)
		}
		return
	}
}

func sendMsg(msg string) {
	NewMsg := tgbotapi.NewMessage(chatID, msg)
	NewMsg.ParseMode = tgbotapi.ModeHTML //傳送html格式的訊息
	_, err := bot.Send(NewMsg)
	if err == nil {
		logrus.Info("Send telegram message success")
	} else {
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

//設定k8s叢集文件
func setK8SClient() {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		logrus.Fatalf("Failed to create k8s clientcmd: %v", err)
	}
	k8sClientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Fatalf("Failed to create k8s k8sClient: %v", err)
	}
}

func getK8SClient() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		logrus.Error(err.Error())
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Error(err.Error())
	}
	return c
}

func getPods(c *kubernetes.Clientset) {
	ctx := context.Background()
	//取得 初始pods 資料
	pods, err := c.CoreV1().Pods(K8SNameSpace).List(ctx, metav1.ListOptions{})
	if err != nil {
		logrus.Errorf("Failed to create k8s getPods: %v", err)
	}
	podRestartList := make(map[string]int32)
	newPodRestartList := make(map[string]int32)
	for _, podList := range pods.Items {
		if podList.Status.Phase == "Running" {
			podRestartList[podList.GetName()] = podList.Status.ContainerStatuses[0].RestartCount
		}
	}
	//迴圈定時檢查pods
	for {
		pods, err = c.CoreV1().Pods(K8SNameSpace).List(ctx, metav1.ListOptions{})
		if err != nil {
			logrus.Errorf("Failed to get Pod List: %v", err)
			if err.Error() == unAuth {
				c = getK8SClient()
			}
			time.Sleep(10 * time.Second)
			continue
		}
		for _, podList := range pods.Items {
			if podList.Status.Phase == "Running" {
				//暫存列表有的pod才檢查
				if val, ok := podRestartList[podList.GetName()]; ok {
					if val != podList.Status.ContainerStatuses[0].RestartCount {
						msg := fmt.Sprintf("偵測到Pods: [%s] 重啟記錄，請進行確認。", podList.GetName())
						//發信
						logrus.Info(msg)
						sendMsg(msg)
					}
				}
				newPodRestartList[podList.GetName()] = podList.Status.ContainerStatuses[0].RestartCount
			}
		}
		//更新廢棄掉的pod
		podRestartList = newPodRestartList
		newPodRestartList = make(map[string]int32)
		logrus.Info("k8s api: monitoring pods")
		time.Sleep(30 * time.Second)
	}
}
