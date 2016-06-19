package contactDatabase

//Address structs are meant for storing the address of an individual or orginization
type Address struct {
	StreetAddress string
	Town          string
	Province      string
	PostalCode    string
	Country       string
	Label         string
}

//Email structs store the address and category of an email address
type Email struct {
	Category string
	Address  string
}

//IMPro structs store the username and category of an email address
type IMPro struct {
	Category string
	Username string
}

//Phone structs store the category and number of a telephone number
type Phone struct {
	Category string
	Number   string
}
