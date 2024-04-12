package password

import (
	"Resul/postTerminalProject/users"
	"fmt"
	"strings"
)

func Password() {

	Use := users.UsersOne{
		Name:       "Resul",
		Username:   "Resul99",
		Password:   "9999",
		CardNomber: "45645645654",
		Money:      456,
	}

	var pas string
	fmt.Println("sifrenizi daxil edin : ")
	fmt.Scan(&pas)

	if strings.EqualFold(pas, Use.Password) {
		fmt.Printf("sizin melumatlariniz  %+v\n", Use)

	} else {
		fmt.Println("duzgun sifre daxil etmediniz")

	}
	fmt.Println(" ")

}
