package Continuee

import (
	"fmt"
	"os"
	"time"
)

var nomber int

func Continuee() {
	fmt.Println()
	fmt.Println("emliyyata davam etmek ucun   ::: 2   emliyyati dayandirmaq ucun ::: 1 daxil edin    ")
	fmt.Scan(&nomber)

	if nomber == 1 {
		fmt.Println("emelliyat sona catdi : ")
		os.Exit(1)
	} else if nomber == 2 {
		fmt.Println("Zehmet olmasa gozleyin  ")
		for i := 0; i < 5; i++ {
			fmt.Print(".")
			time.Sleep(1 * time.Second)

		}
		fmt.Println()

	} else {
		fmt.Println("duzgun reqem daxil etmediniz")
		fmt.Println("emeliyyat sona catdi :")
		os.Exit(1)

	}

}
