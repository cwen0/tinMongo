package models

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

type Auth struct {
	HostName string `json:"hostname" binding:"required"`
	Port     int    `json:"port" binding:"required"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (a *Auth) Connect() (*mgo.Session, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", a.UserName, a.Password, a.UserName, a.Port, a.Database)
	if a.UserName == "" {
		url = fmt.Sprintf("mongodb://%s:%d", a.UserName, a.Port)
	}
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return session, nil
}
