package ds

import (
	"log"

	"github.com/wshaman/stub_contacts"
)

//Singleton connection
var d stub_contacts.YellowPages

func GetYP() stub_contacts.YellowPages {
	var err error
	if d == nil {
		if d, err = stub_contacts.NewYellowPages(); err != nil {
			log.Fatal(err)
		}
		if err = stub_contacts.Populate(d); err != nil {
			log.Fatal(err)
		}
	}
	return d

}
