package main

import (
	"log"
	"math/rand"
	"time"

	"fmt"

	"flag"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/pborman/uuid"
)

type SomeStruct struct {
	A    int         `as:"a"` // alias the field to a
	Self *SomeStruct `as:"-"` // will not persist the field
}

type OtherStruct struct {
	I           int        `as:"i"`
	OtherObject SomeStruct `as:"otherObject"`
}

func main() {
	cleanUp := flag.Bool("cleanUp", false, "a bool")
	dropIndex := flag.Bool("dropIndex", true, "a bool")
	flag.Parse()

	// Truncate table
	// asinfo -v "truncate:namespace=test;set=aerospike;lut=time"

	ns := "test"
	set := "aerospike"

	policy := as.NewWritePolicy(0, 0)
	policy.SendKey = true
	client, err := as.NewClient("127.0.0.1", 3000)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	obj := OtherStruct{
		I:           random(0, 20),
		OtherObject: SomeStruct{A: 18},
	}

	// Create an object
	key, _ := as.NewKey(ns, set, uuid.New())
	log.Printf("Put->key:%v, obj:%v\n", key, obj)
	err = client.PutObject(policy, key, obj)
	if err != nil {
		log.Fatal(err)
	}

	// Get an object
	rObj := OtherStruct{}
	err = client.GetObject(nil, key, &rObj)
	if err != nil {
		// handle error here
		log.Fatal(err)
	}
	log.Printf("Get<-key:%v, obj:%v\n", key, rObj)

	// Update an object
	rObj.I = rObj.I + random(0, 10)
	err = client.PutObject(policy, key, rObj)
	log.Printf("Put->key:%v, obj:%v\n", key, rObj)
	if err != nil {
		log.Fatal(err)
	}

	// Rebuild index
	// asinfo -v "sindex-repair:ns=test;indexname=ind_name;set=set_name;"

	binName1 := "i"
	binName2 := "otherObject"
	indexName := fmt.Sprintf("idx_%v_%v_%v", ns, set, binName1)

	// Drop index
	if *dropIndex == true {
		client.DropIndex(nil, ns, set, indexName)
		if err != nil {
			log.Println(err)
		}
	}

	// Create index
	task, err := client.CreateIndex(nil, ns, set, indexName, binName1, as.NUMERIC)
	for err := range task.OnComplete() {
		if err != nil {
			log.Fatal(err)
		}
	}

	// Do query
	stmt := as.NewStatement(ns, set, binName1, binName2)
	stmt.Addfilter(as.NewRangeFilter(binName1, 0, 30))

	rs, err := client.Query(nil, stmt)
	if err != nil {
		log.Fatal(err)
	}

	for res := range rs.Results() {
		if res.Err != nil {
			// handle error here
			// if you want to exit, cancel the recordset to release the resources
			log.Fatal(res.Err)
		} else {
			// process record here
			log.Println(res.Record.Bins)
		}
	}

	// Do scan
	spolicy := as.NewScanPolicy()
	spolicy.ConcurrentNodes = true
	spolicy.Priority = as.LOW
	spolicy.IncludeBinData = true

	recs, err := client.ScanAll(spolicy, ns, set, binName1, binName2)
	// deal with the error here
	if err != nil {
		log.Fatal(err)
	}

	for res := range recs.Results() {
		if res.Err != nil {
			// handle error here
			// if you want to exit, cancel the recordset to release the resources
			log.Fatal(res.Err)
		} else {
			// process record here
			log.Println(res.Record.Bins)
		}
	}

	// Delete a record
	if *cleanUp == true {
		existed, err := client.Delete(policy, key)
		if err != nil {
			log.Fatal(err)
		}
		if existed == true {
			log.Printf("Delete->key:%v\n", key)
		}
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
