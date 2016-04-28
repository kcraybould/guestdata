package guestdata

import (
	"encoding/json"
	"fmt"

	"github.com/couchbase/gocb"
)

func ReturnGuestsSearch(args ...interface{}) (interface{}, bool) {
	var myParams []interface{}
	var guests []Guest
	var query string

	//which kind  of search are we doing?
	for i, p := range args {
		switch i {
		case 0: // last name
			param, ok := p.(string)
			if !ok {
				fmt.Println("Name must be string")
				return guests, false
			}
			myParams = append(myParams, param)

		case 1: //first name
			param, ok := p.(string)
			if !ok {
				fmt.Println("Name must be string")
				return guests, false
			}
			myParams = append(myParams, param)

		default:
		}
	}

	//which query do we use?
	fmt.Println("params", myParams)
	switch len(myParams) {
	case 1:
		query = "SELECT guestId, personalInfo FROM guest WHERE personalInfo.name.lastName=$1"
	case 2:
		query = "SELECT guestId, personalInfo FROM guest WHERE personalInfo.name.lastName=$1 AND personalInfo.name.firstName=$2"
	default:
	}

	myQuery := gocb.NewN1qlQuery(query)
	rows, err := guestBucket.ExecuteN1qlQuery(myQuery, myParams)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after search query")
	var row Guest
	for rows.Next(&row) {
		fmt.Println("row:", row)
		//try the json decode
		guests = append(guests, row)
		bytes, err := json.Marshal(row)
		fmt.Println("err: ", err)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		fmt.Printf("bytes: %s", bytes)
	}

	_ = rows.Close()

	return guests, (err == nil)
}

func ReturnGuestEmailSearch(email string) (interface{}, bool) {
	var myParams []interface{}
	var guests []Guest

	fmt.Println("email:", email)
	myParams = append(myParams, email)
	myQuery := gocb.NewN1qlQuery("SELECT guestId, personalInfo FROM guest WHERE ANY email IN personalInfo.emails SATISFIES email.emailAddress=$1 END")
	rows, err := guestBucket.ExecuteN1qlQuery(myQuery, myParams)
	if err != nil {
		fmt.Println(err)
		return guests, false
	}

	fmt.Println("after search query")
	var row Guest
	for rows.Next(&row) {
		fmt.Println("row:", row)
		//try the json decode
		guests = append(guests, row)
		bytes, err := json.Marshal(row)
		fmt.Println("err: ", err)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		fmt.Printf("bytes: %s", bytes)
	}

	_ = rows.Close()

	return guests, (err == nil)
}
