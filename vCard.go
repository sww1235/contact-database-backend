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
		line := scanner.Text() //TODO: detect and implement line unfolding
		if strings.HasPrefix(line, "ADR") {
			lineparts := strings.Split(line, ":") //split up into
			//ADR;TYPE=home:;;123 Main St.;Springfield;IL;12345;USA

		}
		if strings.HasPrefix(line, "AGENT") {

		}
		if strings.HasPrefix(line, "ANIVERSARY") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Aniversary = lineparts[1]
		}
		if strings.HasPrefix(line, "BDAY") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Birthday = lineparts[1]
		}
		if strings.HasPrefix(line, "CALADRURI") {
			lineparts := strings.SplitN(line, ":", 1)
			v.CalAdrURI = lineparts[1]
		}
		if strings.HasPrefix(line, "CALURI") {
			lineparts := strings.SplitN(line, ":", 1)
			v.CalURI = lineparts[1]
		}
		if strings.HasPrefix(line, "CATEGORIES") {

		}
		if strings.HasPrefix(line, "Class") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Class = lineparts[1]
		}
		if strings.HasPrefix(line, "CLIENTPIDMAP") {

		}
		if strings.HasPrefix(line, "EMAIL") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Email = lineparts[1]
		}
		if strings.HasPrefix(line, "FN") {
			lineparts := strings.SplitN(line, ":", 1)
			v.FormattedName = lineparts[1]
		}
		if strings.HasPrefix(line, "FREEBUSYURL") {
			lineparts := strings.SplitN(line, ":", 1)
			v.FreeBusyURL = lineparts[1]
		}
		if strings.HasPrefix(line, "GENDER") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Gender = lineparts[1]
		}
		if strings.HasPrefix(line, "GEO") {

		}
		if strings.HasPrefix(line, "IMPP") {

		}
		if strings.HasPrefix(line, "KEY") {

		}
		if strings.HasPrefix(line, "KIND") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Kind = lineparts[1]
		}
		if strings.HasPrefix(line, "LANG") {
			lineparts := strings.SplitN(line, ":", 1)
			v.LanguageSpoken = lineparts[1]
		}
		if strings.HasPrefix(line, "LOGO") {

		}
		if strings.HasPrefix(line, "MAILER") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Mailer = lineparts[1]
		}
		if strings.HasPrefix(line, "MEMBER") {

		}
		if strings.HasPrefix(line, "N") {

		}
		if strings.HasPrefix(line, "NAME") {
			lineparts := strings.SplitN(line, ":", 1)
			v.SourceName = lineparts[1]
		}
		if strings.HasPrefix(line, "NICKNAME") {
			//implement comment separated values
		}
		if strings.HasPrefix(line, "NOTE") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Note = lineparts[1]
		}
		if strings.HasPrefix(line, "ORG") {

		}
		if strings.HasPrefix(line, "PHOTO") {

		}
		if strings.HasPrefix(line, "PRODID") {
			lineparts := strings.SplitN(line, ":", 1)
			v.ProdID = lineparts[1]
		}
		if strings.HasPrefix(line, "PROFILE") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Profile = lineparts[1]
		}
		if strings.HasPrefix(line, "RELATED") {

		}
		if strings.HasPrefix(line, "REV") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Revision = lineparts[1]
		}
		if strings.HasPrefix(line, "ROLE") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Role = lineparts[1]
		}
		if strings.HasPrefix(line, "SOUND") {

		}
		if strings.HasPrefix(line, "SOURCE") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Source = lineparts[1]
		}
		if strings.HasPrefix(line, "TEL") {

		}
		if strings.HasPrefix(line, "TITLE") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Title = lineparts[1]
		}
		if strings.HasPrefix(line, "TZ") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Tz = lineparts[1] //TODO: look at implementing a conversion function for pre 4.0 Tz's
		}
		if strings.HasPrefix(line, "UID") {

		}
		if strings.HasPrefix(line, "URL") {
			lineparts := strings.SplitN(line, ":", 1)
			v.URL = lineparts[1]
		}
		if strings.HasPrefix(line, "VERSION") {
			lineparts := strings.SplitN(line, ":", 1)
			v.Version = lineparts[1]
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
