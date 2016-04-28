package guestdata

import (
	"github.com/couchbase/gocb"
)

type PersonalInfo struct {
	Name      Name      `json:"name"`
	Addresses []Address `json:"addresses"`
	Phones    []Phone   `json:"phones"`
	Emails    []Email   `json:"emails"`
	Payments  []Payment `json:"payments"`
}

type Address struct {
	AddressId    int    `json:"addressId"`
	AddressType  string `json:"addressType"`
	Preferred    bool   `json:"preferred"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	AddressLine3 string `json:"addressLine3"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	PostalCode   string `json:"postalCode"`
	Company      string `json:"company"`
}

type Phone struct {
	PhoneId        int    `json:"phoneId"`
	PhoneType      string `json:"phoneType"`
	Preferred      bool   `json:"preferred"`
	PhoneNumber    string `json:"phoneNumber"`
	PhoneExtension string `json:"phoneExtension"`
}

type Email struct {
	EmailId      int    `json:"emailId"`
	EmailAddress string `json:"emailaddress"`
	Preferred    bool   `json:"preferred"`
}

type Payment struct {
	PaymentId  int    `json:"paymentId"`
	CardNumber string `json:"cardNumber"`
	CardCode   string `json:"cardCode"`
	ExpireDate string `json:"expireDate"`
	Preferred  bool   `json:"preferred"`
}

type Name struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleInit string `json:"middleInit"`
	Title      string `json:"title"`
}
type Guest struct {
	GuestId      int          `json:"guestId"`
	PersonalInfo PersonalInfo `json:"personalInfo"`
}

// oh bucket, my bucket
// pointer to allow for passing by reference
// which allows for better performance
var guestBucket *gocb.Bucket

// lets fire this library up
// init the cluster and bucket, for sharing
// go SDK handles the concurrency.  I think
func init() {
	guestCluster, _ := gocb.Connect("couchbase://127.0.0.1")
	guestBucket, _ = guestCluster.OpenBucket("guest", "")

}
