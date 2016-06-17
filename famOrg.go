package contactDatabase

type family struct {
	orginization //family is an orginization. 'inheritance'

}

type orginization struct {
	mainAddress address //the main mailing address for the orginization, Optional
	mainEmail   string  //main email address for the orginization, Optional

	members []individual //slice of members

}
