package stub_contacts

func isSameContact(c1, c2 Contact) bool {
	return c1.FirstName == c2.FirstName && c1.LastName == c2.LastName && c1.Phone == c2.Phone
}
