package conv

import (
	"github.com/wshaman/stub_contacts"
	proto "github.com/wshaman/web_practice/grpc/proto"
)

func ContactP2B(c proto.Contact) stub_contacts.Contact {
	return stub_contacts.Contact{
		ID:        uint(c.Id),
		Phone:     c.Phone,
		FirstName: c.FirstName,
		LastName:  c.LastName,
	}
}
