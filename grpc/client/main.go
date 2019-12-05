package main

import (
	"context"
	"fmt"
	"log"

	"github.com/wshaman/web_practice/grpc/common/conv"
	"github.com/wshaman/web_practice/grpc/common/proto"
	"google.golang.org/grpc"
)

const address = "127.0.0.1:8082"

func show(cl proto.ContactBookClient, id int64) error {
	v, err := cl.Load(context.Background(), &proto.ID{Id: id})
	if err != nil {
		return err
	}
	c := conv.ContactP2B(*v.People)
	fmt.Println(c.String())
	return nil
}

func findByPhone(cl proto.ContactBookClient, phone string) error {
	list, err := cl.ListByPhone(context.Background(), &proto.Phone{Phone: phone})
	if err != nil {
		return err
	}
	for _, v := range list.People {
		c := conv.ContactP2B(*v)
		fmt.Println(c.String())
	}
	return nil
}

func main() {

	// Установим соединение с сервером
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewContactBookClient(conn)
	if err = show(client, int64(3)); err != nil {
		log.Fatal(err)
	}
	if err = findByPhone(client, ""); err != nil {
		log.Fatal(err)
	}
}
