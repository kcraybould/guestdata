package guestdata

import (
	"log"
	"reflect"

	"github.com/couchbase/gocb"
)

func ReturnGuestsView() (interface{}, bool) {
	log.Println("in the all connect..")
	myQuery := gocb.NewViewQuery("list", "list")

	rows, err := guestBucket.ExecuteViewQuery(myQuery)
	var guests = map[string]Name{}

	// so this is kinda messy.  Basically, we get a ViewResults type back from the gocb.v1
	// the typing gets a bit confusing, but essentially the slice is not a slice of string but a slice
	// of empty interfaces
	var row map[string]interface{}
	for rows.Next(&row) {
		slice := reflect.ValueOf(row["value"])
		var person = Name{FirstName: (slice.Index(1).Interface()).(string), LastName: (slice.Index(0).Interface()).(string), MiddleInit: (slice.Index(2).Interface()).(string)}
		index := reflect.ValueOf(row["key"])
		guests[index.String()] = person
	}

	return guests, (err == nil)
}
