package main

//Individual is a struct that represents one individual person. 
//They can be associated with an orginization and/or family or neither.
type individual struct {
  add
  
  family string //Optional
  
  phoneNumbers []string
  emails []string
  
  
  
}
