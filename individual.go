package main

//Individual is a struct that represents one individual person. 
//They can be associated with an orginization and/or family or neither.
type individual struct {
  name string //no separation between first and last names to avoid 'issues' with international names, Optional if Institution flag is set
  department string //Optional, used for institutions/campuses
  institution string //Optional, this is an orginization that the individual is affilated with such as their place of work
  family family //Optional
  
  addresses []address 
  phoneNumbers []string
  emails []string
  
  
  
}
