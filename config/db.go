package config

import (
	"gocroot/helper"
	"gocroot/model"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = model.DBInfo{
	DBString: helper.SRVLookup(MongoString),
	DBName:   "bukupedia",
}

var Mongoconn, _ = helper.MongoConnect(mongoinfo)
