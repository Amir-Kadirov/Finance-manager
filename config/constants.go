package config

const (
	ERR_INFORMATION     = "The server has received the request and is continuing the process"
	SUCCESS             = "The request was successful"
	ERR_REDIRECTION     = "You have been redirected and the completion of the request requires further action"
	ERR_BADREQUEST      = "Bad request"
	ERR_INTERNAL_SERVER = "While the request appears to be valid, the server could not complete the request"
	SmtpServer          = "smtp.gmail.com"
	SmtpPort            = "587"
	SmtpUsername        = "amirjonqodirov28@gmail.com"
	SmtpPassword        = "bgwc hetf jyxo otkk"

	ExchengeRateApi		="cur_live_eE6yAHP3sCsWzfpJREGkKdEjEEiXApJzR0vqT28W"
)

type CardType int

const (
	AmericanExpress CardType = iota + 1
	DinersClub
	DinersClubUS
	Discover
	JCB
	Laser
	Maestro
	Mastercard
	Solo
	Unionpay
	Visa
	Mir
)