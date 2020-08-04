![muthofun](muthofun.png)

# GO Package For Muthofun SMS gateway

Send sms with less code with muthofun sms gateway


### Installation

Install the package using
```go
$ go get github.com/mhshajib/muthofun
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/mhshajib/muthofun"
```

### Example

***Note: Replace `username` and `password` With your Credential***

***You can add multiple mobile numbers as array***

```go

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



```

You'll get response as map


### **License**
The **muthofun** is an open-source software licensed under the [MIT License](LICENSE)