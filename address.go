package contactDatabase

//address structs are meant for storing the address of an individual or orginization
type address struct {
	streetAddress string
	town          string //
	province      string
	postalCode    string
	country       string
}
