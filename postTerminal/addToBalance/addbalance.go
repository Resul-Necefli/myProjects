package addtobalance

import (
	"Resul/postTerminalProject/users"
	"fmt"
	"strings"
	"time"
)

var Money int

func Addtobalance() {

	Use := users.UsersOne{
		Name:       "resul",
		Username:   "Resul99",
		Password:   "456122",
		CardNomber: "5555555566666666",
		Money:      456,
	}
	var pas string
	fmt.Print(" card melumatlarinizi  daxil edin : ")
	fmt.Scan(&pas)
	for true {
		if len(pas) < 16 {
			fmt.Println("card nomreniz tam doldurun ")
			fmt.Scan(&pas)

		} else {
			fmt.Println("zehmet olmasa  gozleyin  ")
			for i := 0; i < 5; i++ {
				fmt.Print(".")
				time.Sleep(1 * time.Second)

			}
			break
		}
	}

	if strings.EqualFold(pas, Use.CardNomber) {
		fmt.Println()
		fmt.Print("elave etmek isdediyiniz meblegi daxil edin : ")
		fmt.Scan(&Money)

		Use.Money = Use.Money + Money
		fmt.Println("  daxil etdyiniz mebleg balasnsiniza elave edildi cari balans : ", Use.Money)

	} else {
		fmt.Println()
		fmt.Println("duzgun cart nomresi  etmediniz daxil etmediniz")

	}

}
