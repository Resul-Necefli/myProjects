package process

import (
	Continuee "Resul/postTerminalProject/Contniueee"
	addtobalance "Resul/postTerminalProject/addToBalance"
	exitmoneybalanec "Resul/postTerminalProject/exitMoneyBalanec"
	password "Resul/postTerminalProject/melumatlariniz"
	"Resul/postTerminalProject/menu"
	"fmt"
)

func Process() {
	for true {
		fmt.Println()
		menu.Menu()

		switch menu.Reqem {

		case 1:
			password.Password()
			Continuee.Continuee()

		case 2:
			addtobalance.Addtobalance()
			Continuee.Continuee()

		case 3:
			exitmoneybalanec.ExitMoney()
			Continuee.Continuee()

		default:
			fmt.Println("bele bir xidmet movcud deyil : ")
			break

		}

	}

}
