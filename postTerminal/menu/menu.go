package menu

import (
	"fmt"
)

var Reqem int

func Menu() {

	fmt.Println("istifade ede bileceyiniz xidemtler :  ::: 1 istfadeci melumatlari  ::: 2 balansa pul elave etmek   ::: 3 balansdan pul cixarmaq ")
	fmt.Println("secminizi daxil edin  :  ")
	fmt.Scan(&Reqem)

}
