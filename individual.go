package contactDatabase

//Individual is a struct that represents one individual person.
//They can be associated with an orginization and/or family or neither.
//individual is mapped to a vCard of individual type
type individual struct {
	name        string //no separation between first and last names to avoid 'issues' with international names, Optional if Institution flag is set
	department  string //Optional, used for institutions/campuses
	institution string //Optional, this is an orginization that the individual is affilated with such as their place of work
	family      family //Optional

	addresses    map[string]address
	phoneNumbers map[string]string
	emails       map[string]string
}

//Methods declared w/ pointers b/c some methods might need to modify the struct directly.
//also this can be cheaper w/ large structs

func (i *individual) String() string {
	return i.name
}

func (i *individual) address(key string) address {
	return i.addresses[key]

}

func (i *individual) phone(key string) string {
	return i.phoneNumbers[key]

}

func (i *individual) email(key string) string {
	return i.emails[key]

}
