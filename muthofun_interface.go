package muthofun

//Interface of Sms
type ISms interface {
	Init(username string, password string)
	Send(mobiles []string, sms string, unicode string) (map[string]interface{}, error)
}

//use this variable to access the implementation of this interface
//It is also for singleton pattern
var Sms ISms
