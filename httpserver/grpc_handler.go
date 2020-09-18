package main

import (
	"context"

	pb "github.com/bigsai95/ithome12th/mypb"
	"github.com/sirupsen/logrus"
)

func AddUser(username, userpwd string) bool {
	var result bool
	res, err := grpcClient.AddUser(context.Background(),
		&pb.UserRequest{
			UserName: username,
			UserPwd:  userpwd,
		})
	if err != nil {
		logrus.Errorf("AddUser Error: %v", err)
		return result
	}
	if res.Result == "ok" {
		result = true
	}
	return result
}

func LoginUser(username, userpwd string) bool {
	var result bool
	res, err := grpcClient.LoginUser(context.Background(),
		&pb.UserRequest{
			UserName: username,
			UserPwd:  userpwd,
		})
	if err != nil {
		logrus.Errorf("LoginUser Error: %v", err)
		return result
	}
	if res.Result == "ok" {
		result = true
	}
	return result
}

func UserList() []string {
	res, err := grpcClient.UsersList(context.Background(), &pb.UsersListRequest{})
	if err != nil {
		logrus.Fatalf("UsersList Error: %v", err)
	}
	return res.UserName
}
