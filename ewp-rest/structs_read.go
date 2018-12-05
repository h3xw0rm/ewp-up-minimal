package main

import "encoding/xml"

/*********************************************************
* OMobilities Update Requests              (29/Aug/2018) *
*********************************************************/

type oMobUpdateReq struct {
	XMLName     xml.Name                      `xml:"omobilities-update-request"`
	SendHei     typesASCIIPrintableIdentifier `xml:"sending-hei-id"`
	AppCompStud oMobUpdateReqAppCompStd       `xml:"approve-components-studied-draft-v1"`
	UpdCompStud oMobUpdateReqUpdCompStd       `xml:"update-components-studied-v1"`
}

type oMobUpdateReqAppCompStd struct {
	OMobID   typesASCIIPrintableIdentifier                           `xml:"omobility-id"`
	AppParty oMobilityGetMobForStudiesCompStudSOApprovalAppPartyResp `xml:"approving-party"` //string - enumeration: student, sending-hei, receiving-hei
	//CurrLatestDraftSnap omobility:SnapshotOf_ComponentsStudied `xml:"current-latest-draft-snapshot"`
}

type oMobUpdateReqUpdCompStd struct {
	OMobID typesASCIIPrintableIdentifier `xml:"omobility-id"`
	//CurrLatestDraftSnap omobility:SnapshotOf_ComponentsStudied `xml:"current-latest-draft-snapshot"`
	//SuggestedChanges omobility:ListOfChangesTo_ComponentsStudied `xml:"suggested-changes"`
	//SnapWithChanges omobility:SnapshotOf_ComponentsStudied `xml:"snapshot-with-changes-applied"`
}

/*****************************************
* Types Common             (03/Mar/2017) *
*****************************************/

type typesStringWithOptionalLang struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",innerxml"`
}

type typesMultilineString string //string limited by back quotes ` `

type typesMultilineStringWithOptionalLang struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",innerxml"` //string limited by back quotes ` `
}

type typesUUID struct {
	Value string `xml:",innerxml"` //pattern:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}
}

type typesASCIIPrintableIdentifier struct {
	Value string `xml:",innerxml"` //pattern:[&#x0021;-&#x007E;]{1,64}
}

type email string //pattern:[^@]+@[^\.]+\..+

type typesEmpty struct {
}

type typesHTTPS struct {
	Value string `xml:",innerxml"` //pattern:https://.+
}

type typesHTTP struct {
	Value string `xml:",innerxml"` //pattern:https?://.+
}

type typesHTTPWithOptionalLang struct {
	Lang  string `xml:"lang,attr"`
	Value string `xml:",innerxml"` //pattern:https?://.+
}

type typesManifestAPIEntryBase struct {
	XMLns      string            `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmail `xml:"admin-email"`  //(0-inf)
	AdminNotes []typesAdminNotes `xml:"admin-notes"`  //(0-1)
	Version    string            `xml:"version,attr"` //(1) pattern:[0-9]+\.[0-9]+\.[0-9]+
}

type typesAdminEmail struct {
	XMLName xml.Name `xml:"admin-email"`
	XMLns   []string `xml:"xmlns,attr"`
	Email   string   `xml:",innerxml"`
}

type typesAdminNotes typesMultilineString

type typesVersion struct {
	XMLName xml.Name `xml:"version"`
	Value   string   `xml:",innerxml"`
}

type typesCountryCode struct {
	Value string `xml:",innerxml,regexp:'[A-Z][A-Z]'"` //pattern:[A-Z][A-Z]
}

type typesEqfLevel struct {
	Value byte `xml:",innerxml"` //xs:byte <xs:minInclusive value="1"/> <xs:maxInclusive value="8"/>
}

type typesCerfLevel struct {
	Value string `xml:",innerxml,regexp:'[ABC][12]'"` //pattern: [ABC][12]
}

type typesGender struct {
	Value int `xml:",innerxml"` //0 (not known), 1 (male), 2 (female), 9 (not applicable)
}

type typesErrorResponse struct {
	XMLName     xml.Name                                   `xml:"error-response"`
	XMLns       string                                     `xml:"xmlns,attr"` // https://github.com/erasmus-without-paper/ewp-specs-architecture/blob/stable-v1/common-types.xsd
	XMLnsXML    string                                     `xml:"xmlns:xml,attr"`
	XMLnsXS     string                                     `xml:"xmlns:xs,attr"`
	DevMessage  typesMultilineStringResp                   `xml:"developer-message"` //(1)
	UserMessage []typesMultilineStringWithOptionalLangResp `xml:"user-message"`      //(0-inf)
}

/*****************************************
* Types Contact            (22/Fev/2017) *
*****************************************/

type typesContact struct {
	XMLName           xml.Name                               `xml:"contact"`
	ContactName       []typesStringWithOptionalLang          `xml:"contact-name"`       //(1-inf) type: ewp:StringWithOptionalLang
	PersonGivenNames  []typesStringWithOptionalLang          `xml:"person-given-names"` //(0-inf) type: ewp:StringWithOptionalLang
	PersonFamilyNames []typesStringWithOptionalLang          `xml:"person-family-name"` //(0-inf) type: ewp:StringWithOptionalLang
	PhoneNumbers      []typesPhoneNumber                     `xml:"phone-number"`       //(0-inf)
	FaxNumbers        []typesFaxNumber                       `xml:"fax-number"`         //(0-inf)
	Emails            []email                                `xml:"email"`              //(0-inf) type: ewp:email
	StreetAddress     []typesAddressStreetAddress            `xml:"street-address"`     //(0-1)
	MailingAddress    []typesAddressMailingAddress           `xml:"mailing-address"`    //(0-1)
	RoleDescription   []typesMultilineStringWithOptionalLang `xml:"role-description"`   //(0-inf) type: ewp:MultilineStringWithOptionalLang
}

/*****************************************
* Types Phone Number       (21/Fev/2017) *
*****************************************/

type typesPhoneNumber struct {
	XMLName     xml.Name `xml:"phone-number"`
	E164        string   `xml:"e164"`         //(0-1) (pattern:\+[0-9]{1,15})
	Ext         int      `xml:"ext"`          //(0-1)
	OtherFormat string   `xml:"other-format"` //(0-1)
}

type typesFaxNumber struct {
	XMLName     xml.Name `xml:"fax-number"`
	E164        string   `xml:"e164"`         //(0-1) (pattern:\+[0-9]{1,15})
	Ext         int      `xml:"ext"`          //(0-1)
	OtherFormat string   `xml:"other-format"` //(0-1)
}

/*****************************************
* Types Address            (22/Fev/2017) *
*****************************************/

type typesAddressStreetAddress struct {
	XMLName           xml.Name         `xml:"street-address"`
	RepName           []string         `xml:"recipentName"`      //(0-inf)
	AddressLine       []string         `xml:"addressLine"`       //op 1 (0-4)
	BuildingNumber    string           `xml:"buildingNumber"`    //op 2 (0-1)
	BuildingName      string           `xml:"buildingName"`      //op 2 (0-1)
	StreetName        string           `xml:"StreetName"`        //op 2 (0-1)
	Unit              string           `xml:"unit"`              //op 2 (0-1)
	Floor             string           `xml:"floor"`             //op 2 (0-1)
	PostOfficeBox     string           `xml:"postOfficeBox"`     //op 2 (0-1)
	DeliveryPointCode []string         `xml:"deliveryPointCode"` //op 2 (0-inf)
	PostalCode        string           `xml:"postalCode"`        //(0-1)
	Locality          string           `xml:"locality"`          //(0-1)
	Region            string           `xml:"region"`            //(0-1)
	Country           typesCountryCode `xml:"country"`           //type: ewp:CountryCode
}

type typesAddressMailingAddress struct {
	XMLName           xml.Name         `xml:"mailing-address"`
	RepName           []string         `xml:"recipentName"`      //(0-inf)
	AddressLine       []string         `xml:"addressLine"`       //op 1 (0-4)
	BuildingNumber    string           `xml:"buildingNumber"`    //op 2 (0-1)
	BuildingName      string           `xml:"buildingName"`      //op 2 (0-1)
	StreetName        string           `xml:"StreetName"`        //op 2 (0-1)
	Unit              string           `xml:"unit"`              //op 2 (0-1)
	Floor             string           `xml:"floor"`             //op 2 (0-1)
	PostOfficeBox     string           `xml:"postOfficeBox"`     //op 2 (0-1)
	DeliveryPointCode []string         `xml:"deliveryPointCode"` //op 2 (0-inf)
	PostalCode        string           `xml:"postalCode"`        //(0-1)
	Locality          string           `xml:"locality"`          //(0-1)
	Region            string           `xml:"region"`            //(0-1)
	Country           typesCountryCode `xml:"country"`           //type: ewp:CountryCode
}

/*****************************************
* Types Academic Term      (21/Fev/2017) *
*****************************************/

type academicTerm struct {
	XMLName        xml.Name                          `xml:"academic-term"`
	AcademicTermID typesAcademicYearID               `xml:"academic-year-id"` //(1)
	DisplayName    []typesStringWithOptionalLangResp `xml:"display-name"`     //(1-inf)
	StartDate      string                            `xml:"start-date"`       //(1) date
	EndDate        string                            `xml:"end-date"`         //(1) date
}
