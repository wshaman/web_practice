package stub_contacts

import "errors"

type YellowPages interface {
	List() ([]Contact, error)              //Returns all saved contacts
	Load(uint) (Contact, error)            //Tries to load saved contact by ID
	Save(Contact) (uint, error)            //Tries to save if ID is set, insert if not
	Delete(uint) error                     //Tries to delete given contact
	FindByPhone(string) ([]Contact, error) //Tries to find all contacts by part of phone
}

type Contact struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

var EmptyContact = Contact{} // just an alias allows to avoid pointers
var ErrNotFound = errors.New("no record found for given request")
var ErrDuplicate = errors.New("contact already exists in DB")
