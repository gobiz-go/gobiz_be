package config

import (
	"gocroot/helper"
	"gocroot/model"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = model.DBInfo{
	DBString: MongoString,
	DBName:   "gobizdevelop",
}

var Mongoconn, ErrorMongoconn = helper.MongoConnect(mongoinfo)
