package main

//Individual is a struct that represents one individual person. 
//They can be associated with an orginization or not
type Individual struct {
  name string //no separation between first and last names to avoid 'issues' with international names, Optional if Institution flag is set
  department string //Optional, used for institutions/campuses
  institution string //Optional, this is an orginization that the individual is affilated with such as their place of work
  streetAddress string //Normally will be the company address if institution is filled in unless the personal relationship is more important
  town string // 
  province string
  postalCode string
  country string
  
  family string
  
  
  emails []string
  
  
  
}
