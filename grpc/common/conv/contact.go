package conv

import (
	"github.com/wshaman/stub_contacts"
	proto "github.com/wshaman/web_practice/grpc/common/proto"
)

func ContactB2P(c stub_contacts.Contact) *proto.Contact {
	return &proto.Contact{
		Id:        int64(c.ID),
		Phone:     c.Phone,
		FirstName: c.FirstName,
		LastName:  c.LastName,
	}
}

func ContactP2B(c proto.Contact) stub_contacts.Contact {
	return stub_contacts.Contact{
		ID:        uint(c.Id),
		Phone:     c.GetPhone(),
		FirstName: c.GetFirstName(),
		LastName:  c.GetLastName(),
	}
}
