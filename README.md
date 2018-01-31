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

	mongo "github.com/spiderorg/mgo-cs"
)

func main() {
    //
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
