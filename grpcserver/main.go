package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "github.com/bigsai95/ithome12th/mypb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServerService struct{}

var UserData sync.Map

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
	grpcError := make(chan error)

	logrus.Info("grpc server start")

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyprotoServiceServer(s, &ServerService{})
	reflection.Register(s)

	go func() {
		grpcError <- s.Serve(lis)
	}()

	signal.Notify(osc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	select {
	case <-grpcError:
		log.Fatalf("grpc.Serve Error: %v", <-grpcError)
		return
	case <-osc:
		logrus.Printf("grpc server退出訊號: %s", <-osc)
		for t := 3; t > 0; t-- {
			logrus.Printf("%d秒後退出", t)
			time.Sleep(time.Duration(1) * time.Second)
		}
		return
	}
}

func (s *ServerService) AddUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	var result string = "error"
	_, ok := UserData.Load(in.UserName)
	if !ok {
		UserData.Store(in.UserName, in.UserPwd)
		result = "ok"
		logrus.Infof("user add ok: %s", in.UserName)
	} else {
		logrus.Warnf("user add failed: %s", in.UserName)
	}

	return &pb.UserResponse{
		Result: result,
	}, nil
}

func (s *ServerService) LoginUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	var result string = "error"
	pwd, ok := UserData.Load(in.UserName)
	if ok {
		if in.UserPwd == pwd {
			result = "ok"
			logrus.Infof("user login ok: %s", in.UserName)
		} else {
			logrus.Warnf("user login password error: %s", in.UserName)
		}
	} else {
		logrus.Warnf("user not found: %s", in.UserName)
	}
	return &pb.UserResponse{
		Result: result,
	}, nil
}

func (s *ServerService) UsersList(ctx context.Context, in *pb.UsersListRequest) (*pb.UsersListResponse, error) {
	var users []string
	var userpasswd []string

	f := func(k, v interface{}) bool {
		name := k.(string)
		passwd := v.(string)
		users = append(users, name)
		userpasswd = append(userpasswd, passwd)
		return true
	}
	UserData.Range(f)
	return &pb.UsersListResponse{
		Result:   "ok",
		UserName: users,
	}, nil
}
