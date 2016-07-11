package main

import (
	mgo "gopkg.in/mgo.v2"
	"github.com/herman-rogers/gogeta/logger"
)

func ConnectToMongoDB() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017/gogeta")
	if err != nil {
		logger.Error("MongoDB Error: " + err.Error())
	}
	return session
}
