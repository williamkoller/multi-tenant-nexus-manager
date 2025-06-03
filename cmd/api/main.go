package main

import (
	"fmt"
	"log"

	"github.com/williamkoller/multi-tenant-nexus-manager/internal/shared/domain/value_objects"
	domain_user "github.com/williamkoller/multi-tenant-nexus-manager/internal/user/domain"
)

func main() {
	emailVO, err := value_objects.NewEmail("william@mail.com")
	if err != nil {
		log.Fatal(err)
	}

	cpfVO, err := value_objects.NewCPF("06819091966")
	if err != nil {
		log.Fatal(err)
	}

	phoneVO, err := value_objects.NewPhone("41998682343")
	if err != nil {
		log.Fatal(err)
	}
	userD := domain_user.User{
		Email: emailVO,
		CPF:   cpfVO,
		Phone: phoneVO,
	}

	user, err := domain_user.NewUser(&userD)

	if err != nil {
		log.Fatal(err)
	}

	user.Activate()
	fmt.Println("GetDomainEvents(): ", user.GetDomainEvents())
	fmt.Println("ID: ", user.GetID())
	fmt.Println("Email: ", user.Email)
	fmt.Println("CPF: ", user.CPF.Formatted())
	fmt.Println("Phone: ", user.Phone.Formatted())

}
