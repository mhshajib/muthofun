package main

import (
	"fmt"

	"github.com/mhshajib/muthofun"
)

func init() {
	//Setting Sms Interface
	muthofun.Sms = &muthofun.MuthofunConnection{}

	//Initializing Muthofun as sms gateway
	muthofun.Sms.Init("username", "password")
}

func main() {
	//Setting mobile numbers for bulk sms
	mobiles := []string{
		"017xxxxxxxx",
		"018xxxxxxxx",
		"019xxxxxxxx",
	}

	//Sms for message body
	sms := "Thank you for using github.com/mhshajib/muthofun"

	//Flag for unicode
	unicode := "0"

	//Sendding sms
	response, err := muthofun.Sms.Send(mobiles, sms, unicode)
	if err != nil {
		//If something happen before sending sms
		panic(err)
	}

	//API called successfully. Printing API Response in map.
	fmt.Println(response)
}
