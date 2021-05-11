package config

import (
	mgo "gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {
	host := "mongodb://admin:admin@117.53.44.15:27017"
	dbName := "digileaps"

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil
}
