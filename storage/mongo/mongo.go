package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

func NewSession() *mgo.Session {
	info := mgo.DialInfo{
		Addrs:    []string{os.Getenv("MONGO_HOSTS")},
		Timeout:  60 * time.Second,
		Database: os.Getenv("MONGO_DB"),
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}

	session, err := mgo.DialWithInfo(&info)
	if err != nil {
		panic(fmt.Errorf("could not connect to mongodb: %v", err))
	}

	return session
}
