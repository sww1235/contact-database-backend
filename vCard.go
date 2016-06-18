package contactDatabase

// 1      | Exactly one instance per vCard MUST be present.  |
//  |      *1     | Exactly one instance per vCard MAY be present.   |
//  |      1*     | One or more instances per vCard MUST be present. |
//  |      *      | One or more instances per vCard MAY be present.
//BEGIN:vCard - needed at begining of vCard file - 1
//END:VCARD - needed at end of vCard file - 1
//SOURCE: uri - location for potentially more up to date information, such as
//an ldap link or a website - *
//KIND: string value, can be individual, group(contains memebers), org
//(does not contain memebers), location (named geographical location), 1
//VERSION: 4.0 for all practical purposes, but needs to be handled for import
//purposes 1
//ADR:
//FN: - this is a text string for name representaion

//TODO: need to look at which attributes take a type and make them structs

//VCard represents a vCard file in memory. All fields may not be filled
type VCard struct {
	Version        string
	FormattedName  string //FN in vCard formal definition
	Gender         string
	Address        []Address //ADR in vCard formal definition
	Agent          []VCard
	Aniversary     string
	Birthday       string //BDAY in vCard formal definition
	CalAdrURI      string
	CalURI         string
	Categories     []string
	Class          string
	ClientPIDMap   string
	Email          []Email
	FreeBusyURL    string
	LatLong        string   // GEO in vCard formal definition
	InstMess       []IM     //IMPP in vCard formal definition
	PubKey         []string //KEY in vCard formal definition
	Kind           string   //value to be selected from set {'application', 'individual, 'group', 'location' or 'organization'; 'x-*' values may be used for experimental purposes}
	LanguageSpoken []string //LANG in vCard formal definition, of form fr-CA
	Logo           string   //TODO: decide if struct would be better
	Mailer         string
	Member         []VCard //kind must be group to use this property
	Name           string  //N in vCard formal definition
	SourceName     string  //NAME in vCard formal definition
	Nickname       string
	Note           string
	Org            []VCard
	Photo          string //TODO: decide if struct would be better
	ProdID         string
	Profile        string //always is PROFILE:VCARD
	Related        string
	Revision       string //REV in vCard formal definition
	Role           string
	Sound          string
	Source         string
	Tel            []Phone
	Title          string
	Tz             string
	UID            string
	URL            string
}

//ImportFromFile takes the filepath of a vCard file and adds its data to the
//VCard it is called on. This means that you must have an existing, ideally
//blank VCard struct to call this function on and that there is somewhere to
//store the data.
func (v *VCard) ImportFromFile(filename string) {

}

//ExportToFile takes the filepath of a vCard file and adds its data to the
//vCard it is called on.  This means that you must have an existing, ideally
//blank vCard struct to call this function on and that there is somewhere to
//store the data.
func (v *VCard) ExportToFile() {

}
