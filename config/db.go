package config

import (
	"gocroot/helper"
	"os"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: helper.SRVLookup(MongoString),
	DBName:   "bukupedia",
}

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)
