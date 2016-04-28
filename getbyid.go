package guestdata

import "log"

func ReturnGuestsById(id string) (interface{}, bool) {
	log.Println("in the connect..")
	key := "guest::" + id

	var guest map[string]interface{}
	cas, _ := guestBucket.Get(key, &guest)

	log.Println(cas)

	return guest, (guest != nil)
}
