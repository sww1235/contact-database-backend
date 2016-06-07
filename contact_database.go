package main

//address structs are meant for storing the address of an individual or orginization
type address struct {
  department string //Optional, used for institutions/campuses
  institution string //Optional, this is an orginization that the individual is affilated with such as their place of work
  streetAddress string //Normally will be the company address if institution is filled in unless the personal relationship is more important
  town string // 
  province string
  postalCode string
  country string
  
}
func main {

}
