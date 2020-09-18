package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type userStrcut struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping ok"))
}

func userAdd(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": //使用POST新增使用者
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"trace": "userAdd_1",
				"err":   err,
			}).Error("Error Log")
			w.Write([]byte("參數錯誤"))
			return
		}
		var user userStrcut
		err = json.Unmarshal(payload, &user)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"trace": "userAdd_2",
				"err":   err,
			}).Error("Error Log")
			w.Write([]byte("參數錯誤"))
			return
		}
		if AddUser(user.Username, user.Password) {
			w.Write([]byte("使用者新增成功"))
			return
		}
		w.Write([]byte("使用者已新增過"))
		return
	}
	w.Write([]byte("Received Post request"))
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT": //使用Put使用者登入
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"trace": "userLogin_1",
				"err":   err,
			}).Error("Error Log")
			w.Write([]byte("參數錯誤"))
			return
		}
		var user userStrcut
		err = json.Unmarshal(payload, &user)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"trace": "userLogin_2",
				"err":   err,
			}).Error("Error Log")
			w.Write([]byte("參數錯誤"))
			return
		}
		if LoginUser(user.Username, user.Password) {
			w.Write([]byte("使用者登入成功"))
			return
		}
		w.Write([]byte("使用者登入失敗"))
		return
	}
	w.Write([]byte("Received Put request"))
}
