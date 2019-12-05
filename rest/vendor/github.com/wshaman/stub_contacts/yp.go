package stub_contacts

import (
	"strings"
	"sync"
)

type yellowPages struct {
	knownContacts map[uint]Contact
	locker        sync.RWMutex
}

func NewYellowPages() (YellowPages, error) {
	yp := &yellowPages{
		knownContacts: map[uint]Contact{},
	}
	return yp, nil
}

//Load Tries to load saved contact by ID
func (y *yellowPages) Load(id uint) (resp Contact, err error) {
	y.locker.RLock()
	defer y.locker.RUnlock()
	var ok bool
	if resp, ok = y.knownContacts[id]; !ok {
		return EmptyContact, ErrNotFound
	}
	return resp, nil
}

//List returns all saved contacts
func (y *yellowPages) List() (resp []Contact, err error) {
	y.locker.RLock()
	defer y.locker.RUnlock()
	if len(y.knownContacts) == 0 {
		return nil, ErrNotFound
	}
	resp = make([]Contact, 0, len(y.knownContacts))
	for _, v := range y.knownContacts {
		resp = append(resp, v)
	}
	return resp, nil
}

//Save tries to save if ID is set, insert if not
func (y *yellowPages) Save(c Contact) (uint, error) {
	y.locker.Lock()
	defer y.locker.Unlock()
	if c.ID == 0 {
		return y.addNewContact(c)
	}
	if err := y.updateContact(c); err != nil {
		return 0, err
	}
	return c.ID, nil
}

//Delete tries to delete given contact
func (y *yellowPages) Delete(id uint) error {
	if _, ok := y.knownContacts[id]; !ok {
		return ErrNotFound
	}
	delete(y.knownContacts, id)
	return nil
}

//FindByPhone tries to find all contacts by part of phone
func (y *yellowPages) FindByPhone(part string) ([]Contact, error) {
	contacts := make([]Contact, 0)
	for _, v := range y.knownContacts {
		if strings.Index(v.Phone, part) >= 0 {
			contacts = append(contacts, v)
		}
	}
	if len(contacts) == 0 {
		return nil, ErrNotFound
	}
	return contacts, nil
}

func (y *yellowPages) updateContact(contact Contact) error {
	if _, ok := y.knownContacts[contact.ID]; !ok {
		return ErrNotFound
	}
	if isSameContact(y.knownContacts[contact.ID], contact) {
		return ErrDuplicate
	}
	y.knownContacts[contact.ID] = contact
	return nil
}

func (y *yellowPages) addNewContact(contact Contact) (uint, error) {
	for _, v := range y.knownContacts {
		if isSameContact(contact, v) {
			return 0, ErrDuplicate
		}
	}
	newID := y.getNextID()
	contact.ID = newID
	y.knownContacts[newID] = contact
	return newID, nil
}

func (y *yellowPages) getNextID() (res uint) {
	for i := range y.knownContacts {
		if res < i {
			res = i
		}
	}
	res++
	return res
}
