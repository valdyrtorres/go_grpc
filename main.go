package main

import (
	"exemplopb/src/pb/users"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	readData()
}

func readData() {
	var user1 users.User
	dados, err := os.ReadFile("dados.txt")
	if err != nil {
		log.Fatalln("error on read file. Error: ", err)
	}

	err = proto.Unmarshal(dados, &user1)
	if err != nil {
		log.Fatalln("error on jnmarshal. Error: ", err)
	}

	fmt.Printf("user: %+v\n", user1)

}

func createData() {
	user1 := users.User{
		Id:       1,
		Name:     "Valdir",
		Email:    "valdir@email.com",
		Password: "123",
	}

	out, err := proto.Marshal(&user1)
	if err != nil {
		log.Fatalln("error on marshal. Error: ", err)
	}

	err = os.WriteFile("dados.txt", out, 0644)
	if err != nil {
		log.Fatalln("error on write file. Error: ", err)
	}

	//fmt.Printf("user: %s\n", out)
}
