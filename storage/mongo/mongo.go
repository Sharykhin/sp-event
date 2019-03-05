package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	hosts = "localhost:27017"
)

func NewSession() *mgo.Session {
	info := mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: "spevent",
		Username: "spevent",
		Password: "spevent",
	}

	session, err := mgo.DialWithInfo(&info)
	if err != nil {
		panic(fmt.Errorf("could not connect to mongodb: %v", err))
	}

	return session
}
