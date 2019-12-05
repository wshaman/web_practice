package conv

import (
	"github.com/wshaman/stub_contacts"
	proto "github.com/wshaman/web_practice/grpc/proto"

)

func ContactB2P(c stub_contacts.Contact) *proto.Contact {
	return &proto.Contact{
		Id:        int64(c.ID),
		Phone:     c.Phone,
		FirstName: c.FirstName,
		LastName:  c.LastName,
	}
}
