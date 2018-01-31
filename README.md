## Description
A mongodb-pool tool base on go-mgo
## Development
 ```go
 go get github.com/spiderorg/mgo-cs

 ```
 ## Opening a database

 ```
 package main

import (
	"log"

	"github.com/spiderorg/mgo-cs/mongo"
)

func main() {

    // init first
    // param explain:
    // 1 : connection's username
    // 2 : connection's password
    // 3 : connection's addr ----format like "127.0.0.1:27017/dbname"
    // 4 : max size of the pool
    // 5 : idl time of the per link
	mongo.Refresh("dbUser","dbPassword","dbConnect",1000,100)

	defer db.Close()

    // find
    result := make(bson.M, 1)
	err := mongo.Mgo(result, "find", map[string]interface{}{
		"Database":   "dbName",
		"Collection": "dbCollection",
		"Query":      query,
		"Sort":       sort,
	})
	if err != nil {
		fmt.Printf(" *Fail %v\n", err)
	}
    .
    .
    .
	...
}

 ```
