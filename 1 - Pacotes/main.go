package main

import (
	"fmt"
	"module/helper"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá mundo!")

	helper.ReadText()

	erro := checkmail.ValidateFormat("123123")

	fmt.Println(erro)
}
