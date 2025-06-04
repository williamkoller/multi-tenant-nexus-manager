package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/williamkoller/multi-tenant-nexus-manager/internal/core/domain/value_objects"
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

	// Prints originais
	fmt.Println("GetDomainEvents(): ", user.GetDomainEvents())
	fmt.Println("ID: ", user.GetID())
	fmt.Println("Email: ", user.Email)
	fmt.Println("CPF: ", user.CPF.Formatted())
	fmt.Println("Phone: ", user.Phone.Formatted())

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("JSON :")
	fmt.Println(strings.Repeat("=", 50))

	apiResponse := map[string]interface{}{
		"success": true,
		"message": "User created successfully",
		"data": map[string]interface{}{
			"user": map[string]interface{}{
				"id":    user.GetID(),
				"email": user.Email.String(),
				"cpf":   user.CPF.Formatted(),
				"phone": user.Phone.Formatted(),
			},
			"events": user.GetDomainEvents(),
		},
		"timestamp": time.Now().Format(time.RFC3339),
	}

	apiJSON, err := json.MarshalIndent(apiResponse, "", "  ")
	if err != nil {
		log.Printf("Erro ao serializar API response: %v", err)
	} else {
		fmt.Println("Response Format (JSON):")
		fmt.Println(string(apiJSON))
	}
}
