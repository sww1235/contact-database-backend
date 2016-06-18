package contactDatabase

//address structs are meant for storing the address of an individual or orginization
type address struct {
	streetAddress string
	town          string
	province      string
	postalCode    string
	country       string
}

type email struct {
	category string
	address  string
}

type im struct {
	category string
	username string
}

type phone struct {
	category string
	number   string
}
