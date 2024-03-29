package db

import (
	"vc.svc/models"

	"gopkg.in/mgo.v2"
)

var (
	MongoSession *mgo.Session
	DialInfo     mgo.DialInfo
)

func Init(config models.MongodbConfig) {
	dialInfo := mgo.DialInfo{Addrs: []string{config.Host}, Database: config.DB, ReplicaSetName: config.ReplicaSetName, Username: config.User, Password: config.Pass}
	session, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		panic(err)
	}
	// defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	MongoSession = session
	DialInfo = dialInfo
}
