package contactDatabase

import (
	"bufio"
	"strings"

	"github.com/sww1235/fileHandling"
)

//TODO: need to look at which attributes take a type and make them structs

//VCard represents a vCard file in memory. All fields may not be filled
type VCard struct {
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
	FormattedName  string //FN in vCard formal definition
	FreeBusyURL    string
	Gender         string
	LatLong        string   // GEO in vCard formal definition
	InstMess       []IMPro  //IMPP in vCard formal definition
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
	Version        string
}

//ImportFromFile takes the filepath of a vCard file and adds its data to the
//VCard it is called on. This means that you must have an existing, ideally
//blank VCard struct to call this function on and that there is somewhere to
//store the data.
func (v *VCard) ImportFromFile(filename string) {
	f := fileHandling.CreateFileHandle(filename)
	defer fileHandling.CloseFileHandle(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ADR") {
			lineparts := strings.Split(line, ";")
			//ADR;TYPE=home:;;123 Main St.;Springfield;IL;12345;USA

		}
		if strings.HasPrefix(line, "AGENT") {

		}
		if strings.HasPrefix(line, "ANIVERSARY") {

		}
		if strings.HasPrefix(line, "BDAY") {

		}
		if strings.HasPrefix(line, "CALADRURI") {

		}
		if strings.HasPrefix(line, "CALURI") {

		}
		if strings.HasPrefix(line, "CATEGORIES") {

		}
		if strings.HasPrefix(line, "Class") {

		}
		if strings.HasPrefix(line, "CLIENTPIDMAP") {

		}
		if strings.HasPrefix(line, "EMAIL") {

		}
		if strings.HasPrefix(line, "FN") {

		}
		if strings.HasPrefix(line, "FREEBUSYURL") {

		}
		if strings.HasPrefix(line, "GENDER") {

		}
		if strings.HasPrefix(line, "GEO") {

		}
		if strings.HasPrefix(line, "IMPP") {

		}
		if strings.HasPrefix(line, "KEY") {

		}
		if strings.HasPrefix(line, "KIND") {

		}
		if strings.HasPrefix(line, "LANG") {

		}
		if strings.HasPrefix(line, "LOGO") {

		}
		if strings.HasPrefix(line, "MAILER") {

		}
		if strings.HasPrefix(line, "MEMBER") {

		}
		if strings.HasPrefix(line, "N") {

		}
		if strings.HasPrefix(line, "NAME") {

		}
		if strings.HasPrefix(line, "NICKNAME") {

		}
		if strings.HasPrefix(line, "NOTE") {

		}
		if strings.HasPrefix(line, "ORG") {

		}
		if strings.HasPrefix(line, "PHOTO") {

		}
		if strings.HasPrefix(line, "PRODID") {

		}
		if strings.HasPrefix(line, "PROFILE") {

		}
		if strings.HasPrefix(line, "RELATED") {

		}
		if strings.HasPrefix(line, "REV") {

		}
		if strings.HasPrefix(line, "ROLE") {

		}
		if strings.HasPrefix(line, "SOUND") {

		}
		if strings.HasPrefix(line, "SOURCE") {

		}
		if strings.HasPrefix(line, "TEL") {

		}
		if strings.HasPrefix(line, "TITLE") {

		}
		if strings.HasPrefix(line, "TZ") {

		}
		if strings.HasPrefix(line, "UID") {

		}
		if strings.HasPrefix(line, "URL") {

		}
		if strings.HasPrefix(line, "VERSION") {

		}
	}
}

//ExportToFile takes the filepath of a vCard file and adds its data to the
//vCard it is called on.  This means that you must have an existing, ideally
//blank vCard struct to call this function on and that there is somewhere to
//store the data.
func (v *VCard) ExportToFile() {

}

func (v VCard) String() string {
	var s string
	s = "VCard version: " + v.Version + "\n"
	s += "FormattedName: " + v.FormattedName + "\n"
	//s += "Address: " + displayStrings(v.Address) + "\n"
	return s
}

func displayStrings(ss []string) (display string) {
	for _, s := range ss {
		display += s + ", "
	}
	return display
}
