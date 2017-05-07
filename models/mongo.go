package models

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
)

var mongo *mgo.Session

func InitMongo(m *mgo.Session) error {
	if m == nil {
		return errors.New("mongo session is nil")
	}
	mongo = m
	return nil
}

func GetMongo() (*mgo.Session, error) {
	if mongo != nil {
		s := mongo.Clone()
		return s, nil
	}
	return nil, errors.New("mongo session is nil")
}
