#Stub data provider 

**This is a testing purposes stub in-memory contact book lib.**

It's only purpose to help students focus on writing REST/RPC service around data provider.

*Please, be sure to keep Delete and Update methods idempotent*

this lib provides a simple interface:

````go
List() ([]Contact, error)              //Returns all saved contacts
Load(uint) (Contact, error)            //Tries to load saved contact by ID
Save(Contact) (uint, error)            //Tries to save if ID is set, insert if not
Delete(uint) error                     //Tries to delete given contact
FindByPhone(string) ([]Contact, error) //Tries to find all contacts by part of phone

````
As this is in-memory storage, it will lose all data on restart. For testing purposes `Populate` can be used
````go
	c, err := NewYellowPages()
	err = Populate(c)
````
It seeds DB with some valid data you could check with `List`