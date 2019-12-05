package main

import (
	"context"
	"log"
	"net"

	//"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	//
	//"github.com/wshaman/stub_contacts"
	//"github.com/wshaman/web_practice/grpc/conv"
	//proto "github.com/wshaman/web_practice/grpc/proto"

	"github.com/wshaman/stub_contacts"
	"github.com/wshaman/web_practice/grpc/conv"
	"github.com/wshaman/web_practice/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type service struct {
	repo stub_contacts.YellowPages
}

func (s *service) Delete(_ context.Context, id *proto.ID) (*proto.DeleteReply, error) {
	if err := s.repo.Delete(uint(id.Id)); err != nil {
		return nil, err
	}
	return &proto.DeleteReply{Status: "OK"}, nil
}

func (s *service) Load(_ context.Context, in *proto.ID) (*proto.LoadReply, error) {
	contact, err := s.repo.Load(uint(in.Id))
	if err != nil {
		return nil, err
	}
	return &proto.LoadReply{People: conv.ContactB2P(contact)}, nil
}

func (s *service) Save(_ context.Context, in *proto.Contact) (*proto.SaveReply, error) {
	uid, err := s.repo.Save(conv.ContactP2B(*in))
	if err != nil {
		return nil, err
	}
	return &proto.SaveReply{Id: int64(uid)}, nil
}

func (s *service) ListByPhone(_ context.Context, req *proto.Phone) (*proto.ListReply, error) {
	list, err := s.repo.FindByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	reply := &proto.ListReply{
		People: make([]*proto.Contact, len(list), len(list)),
	}
	for i, v := range list {
		reply.People[i] = conv.ContactB2P(v)
	}
	return reply, nil
}

func main() {

	// Стартуем наш gRPC сервер для прослушивания tcp
	tcp, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bookRepo, err := stub_contacts.NewYellowPages()
	if err != nil {
		panic(err)
	}
	if err = stub_contacts.Populate(bookRepo); err != nil {
		panic(err)
	}

	proto.RegisterContactBookServer(s, &service{repo: bookRepo})
	reflection.Register(s)
	if err := s.Serve(tcp); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
