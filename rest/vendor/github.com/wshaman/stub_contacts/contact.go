package stub_contacts

import "fmt"

func isSameContact(c1, c2 Contact) bool {
	return c1.FirstName == c2.FirstName && c1.LastName == c2.LastName && c1.Phone == c2.Phone
}

func (c *Contact) String() string {
	return fmt.Sprintf("Name: %s %s, Phone: %s", c.FirstName, c.LastName, c.Phone)
}
