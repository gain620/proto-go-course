package main

import (
	pb "github.com/gain620/proto-go-course/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func makeRandomPerson() *pb.Person {
	return &pb.Person{
		Id:    rand.Int31n(1000),
		Email: randStringBytes(10) + "@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "123-4241-123",
				Type:   pb.Person_MOBILE,
			},
		},
		LastUpdated: timestamppb.Now(),
	}
}

func makeAddressBook() *pb.AddressBook {
	return &pb.AddressBook{
		BookId: rand.Int31n(100),
		People: nil,
	}
}

func addPerson(book *pb.AddressBook, person *pb.Person) *pb.AddressBook {
	book.People = append(book.People, person)
	return book
}
