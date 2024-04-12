package exitmoneybalanec

import (
	"Resul/postTerminalProject/users"
	"fmt"
	"strings"
)

var mon int

func ExitMoney() {
	Use := users.UsersOne{
		Name:       "resul",
		Username:   "Resul99",
		Password:   "9999",
		CardNomber: "45645645654",
		Money:      456,
	}
	var password string
	fmt.Print("sifrenizi daxil  edin daxil edin  :  ")
	fmt.Scan(&password)

	if strings.EqualFold(password, Use.Password) {

		fmt.Print("cixarmaq isdedyiniz meblegi daxil edin : ")
		fmt.Scan(&mon)
		if mon < Use.Money {
			Use.Money = Use.Money - mon
			fmt.Println("cari balans  : ", Use.Money)

		} else {
			fmt.Println()
			fmt.Println("daxil etdiyin  mebleg balansinizda movcud deyil")
		}

	} else {
		fmt.Println("sifreni yanlis daxil etdiniz : ")
	}

}
