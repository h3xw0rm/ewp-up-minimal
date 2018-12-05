package main

import "encoding/xml"

//XMLnsXS Schema
var XMLnsXS = "http://www.w3.org/2001/XMLSchema"

//XMLnsXML Schema
var XMLnsXML = "http://www.w3.org/XML/1998/namespace"

/**************************************************************************************************************
* Manifest (Discovery API)                                                                       (12/03/2018) *
**************************************************************************************************************/

//XMLnsManifest Schema
var XMLnsManifest = "https://github.com/erasmus-without-paper/ewp-specs-api-discovery/tree/stable-v5"

//XMLnsRegistryAPI Schema
var XMLnsRegistryAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-registry/tree/stable-v1"

type typeManifestResp struct {
	XMLName    xml.Name               `xml:"manifest"`
	XMLns      string                 `xml:"xmlns,attr"`
	XMLnsEWP   string                 `xml:"xmlns:ewp,attr"`
	XMLnsR     string                 `xml:"xmlns:r,attr"`
	XMLnsSec   string                 `xml:"xmlns:sec,attr"`
	XMLnsSecA  string                 `xml:"xmlns:sec-cli-anon,attr"`
	XMLnsSecCT string                 `xml:"xmlns:sec-cli-tls,attr"`
	XMLnsSecCH string                 `xml:"xmlns:sec-cli-http,attr"`
	XMLnsSecST string                 `xml:"xmlns:sec-svr-tls,attr"`
	XMLnsSecSH string                 `xml:"xmlns:sec-svr-http,attr"`
	Hosts      []typeManifestHostResp `xml:"host"` //(0-inf)
}

type typeManifestHostResp struct {
	AdminEmail          []typesAdminEmailResp              `xml:"ewp:admin-email"`           //(0-inf)
	AdminNotes          []typesAdminNotesResp              `xml:"ewp:admin-notes"`           //(0-1)
	APIsImplemented     []typeCatalogueAPISImplementedResp `xml:"r:apis-implemented"`        //(0-1)
	InstitutionsCovered []typeManifestInstCovered          `xml:"institutions-covered"`      //(0-1)
	ClientCredInUse     []typeManifestCliCredInUse         `xml:"client-credentials-in-use"` //(0-1)
	ServerCredInUse     []typeManifestSvrCredInUse         `xml:"server-credentials-in-use"` //(0-1)
}

type typeManifestInstCovered struct {
	XMLName xml.Name               `xml:"institutions-covered"`
	Heis    []typeCatalogueHeiResp `xml:"r:hei"` //(0-inf)
}

type typeCatalogueAPISImplementedResp struct {
	XMLName     xml.Name                       `xml:"r:apis-implemented"`
	Discovery   []typesManifestAPIDiscovery    `xml:"discovery"`
	Echo        []typesManifestAPIEcho         `xml:"echo"`
	Institution []typesManifestAPIInstitutions `xml:"institutions"`
	OUnits      []typesManifestAPIOUnits       `xml:"organizational-units"`
	Courses     []typesManifestAPICourses      `xml:"courses"`
	CourseRep   []typesManifestAPICourseRep    `xml:"simple-course-replication"`
	IIAs        []typesManifestAPIIIAS         `xml:"iias"`
	IIAcnr      []typesManifestAPIIIASCNR      `xml:"iia-cnr"`
	OMob        []typesManifestAPIoMobility    `xml:"omobilities"`
	OMobCNR     []typesManifestAPIoMobilityCNR `xml:"omobility-cnr"`
	IMob        []typesManifestAPIiMobility    `xml:"imobilities"`
	IMobCNR     []typesManifestAPIiMobilityCNR `xml:"imobility-cnr"`
	IMobTOR     []typesManifestAPITORS         `xml:"imobility-tor"`
	IMobTORCNR  []typesManifestAPITORSCNR      `xml:"imobility-tor-cnr"`
}

type typeCatalogueHeiResp struct {
	XMLName xml.Name                          `xml:"r:hei"`
	OtherID []catalogueOtherHeiID             `xml:"r:other-id"` //(0-inf) (pic,erasmus,euc,erasmus-charter)
	Name    []typesStringWithOptionalLangResp `xml:"r:name"`     //(1-inf)
	ID      string                            `xml:"id,attr"`    //required
}

type typeManifestCliCredInUse struct {
	XMLName     xml.Name `xml:"client-credentials-in-use"`
	Certificate []string `xml:"certificate"`    //(0-inf) xs:base64binary
	RSAPubKey   []string `xml:"rsa-public-key"` //(0-inf) xs:base64binary
}

type typeManifestSvrCredInUse struct {
	XMLName   xml.Name `xml:"server-credentials-in-use"`
	RSAPubKey []string `xml:"rsa-public-key"` //(0-inf) xs:base64binary
}

// manifest APIs (12/03/2018)

//XMLnsManifestDiscovery schema
var XMLnsManifestDiscovery = "https://github.com/erasmus-without-paper/ewp-specs-api-discovery/blob/stable-v5/manifest-entry.xsd"

//VersionDiscovery API
var VersionDiscovery = "5.0.0"

type typesManifestAPIDiscovery struct {
	XMLName    xml.Name              `xml:"discovery"`
	XMLns      string                `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp `xml:"ewp:admin-email"` //(0-inf)
	AdminNotes []typesAdminNotesResp `xml:"ewp:admin-notes"` //(0-1)
	Version    string                `xml:"version,attr"`    //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	URL        typesHTTPS            `xml:"url"`             //(1)
}

//XMLnsManifestEcho schema
var XMLnsManifestEcho = "https://github.com/erasmus-without-paper/ewp-specs-api-echo/blob/stable-v2/manifest-entry.xsd"

//VersionEcho API
var VersionEcho = "2.0.1"

type typesManifestAPIEcho struct {
	XMLName    xml.Name                       `xml:"echo"`
	XMLns      string                         `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp          `xml:"ewp:admin-email"` //(0-inf)
	AdminNotes []typesAdminNotesResp          `xml:"ewp:admin-notes"` //(0-1)
	Version    string                         `xml:"version,attr"`    //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec        []typesHTTPSecurityOptionsResp `xml:"http-security"`   //(0-1)
	URL        typesHTTPS                     `xml:"url"`             //(1)
}

//XMLnsManifestInstitutions schema
var XMLnsManifestInstitutions = "https://github.com/erasmus-without-paper/ewp-specs-api-institutions/blob/stable-v2/manifest-entry.xsd"

//VersionInstitutions API
var VersionInstitutions = "2.1.0"

type typesManifestAPIInstitutions struct {
	XMLName    xml.Name                       `xml:"institutions"`
	XMLns      string                         `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp          `xml:"ewp:admin-email"` //(0-inf)
	AdminNotes []typesAdminNotesResp          `xml:"ewp:admin-notes"` //(0-1)
	Version    string                         `xml:"version,attr"`    //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec        []typesHTTPSecurityOptionsResp `xml:"http-security"`   //(0-1)
	URL        typesHTTPS                     `xml:"url"`             //(1)
	MaxHeiIDs  int                            `xml:"max-hei-ids"`     //(1) pattern:positive integer
}

//XMLnsManifestOUnits schema
var XMLnsManifestOUnits = "https://github.com/erasmus-without-paper/ewp-specs-api-ounits/blob/stable-v2/manifest-entry.xsd"

//VersionOUnits API
var VersionOUnits = "2.1.0"

type typesManifestAPIOUnits struct {
	XMLName    xml.Name                       `xml:"organizational-units"`
	XMLns      string                         `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp          `xml:"ewp:admin-email"` //(0-inf)
	AdminNotes []typesAdminNotesResp          `xml:"ewp:admin-notes"` //(0-1)
	Version    string                         `xml:"version,attr"`    //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec        []typesHTTPSecurityOptionsResp `xml:"http-security"`   //(0-1)
	URL        typesHTTPS                     `xml:"url"`             //(1)
	MaxOUID    int                            `xml:"max-ounit-ids"`   //(1) pattern:positive integer
	MaxOUCode  int                            `xml:"max-ounit-codes"` //(1) pattern:positive integer
}

//XMLnsManifestCourses schema
var XMLnsManifestCourses = "https://github.com/erasmus-without-paper/ewp-specs-api-courses/blob/stable-v1/manifest-entry.xsd"

//VersionCourses API
var VersionCourses = "0.7.1"

type typesManifestAPICourses struct {
	XMLName    xml.Name                       `xml:"courses"`
	XMLns      string                         `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp          `xml:"ewp:admin-email"` //(0-inf)
	AdminNotes []typesAdminNotesResp          `xml:"ewp:admin-notes"` //(0-1)
	Version    string                         `xml:"version,attr"`    //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec        []typesHTTPSecurityOptionsResp `xml:"http-security"`   //(0-1)
	URL        typesHTTPS                     `xml:"url"`             //(1)
	MaxLOSID   int                            `xml:"max-los-ids"`     //(1) pattern:positive integer
	MaxLOSCode int                            `xml:"max-los-codes"`   //(1) pattern:positive integer
}

//XMLnsManifestCourseRep schema
var XMLnsManifestCourseRep = "https://github.com/erasmus-without-paper/ewp-specs-api-course-replication/blob/stable-v1/manifest-entry.xsd"

//VersionCourseRep API
var VersionCourseRep = "1.0.0"

type typesManifestAPICourseRep struct {
	XMLName     xml.Name                       `xml:"simple-course-replication"`
	XMLns       string                         `xml:"xmlns,attr"`
	AdminEmail  []typesAdminEmailResp          `xml:"ewp:admin-email"`         //(0-inf)
	AdminNotes  []typesAdminNotesResp          `xml:"ewp:admin-notes"`         //(0-1)
	Version     string                         `xml:"version,attr"`            //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec         []typesHTTPSecurityOptionsResp `xml:"http-security"`           //(0-1)
	URL         typesHTTPS                     `xml:"url"`                     //(1)
	AllowsAnon  bool                           `xml:"allows-anonymous-access"` //(1)
	SupModSince bool                           `xml:"supports-modified-since"` //(1)
}

//XMLnsManifestIIAS schema
var XMLnsManifestIIAS = "https://github.com/erasmus-without-paper/ewp-specs-api-iias/blob/stable-v2/manifest-entry.xsd"

//VersionIIAS API
var VersionIIAS = "2.1.0"

type typesManifestAPIIIAS struct {
	XMLName       xml.Name                       `xml:"iias"`
	XMLns         string                         `xml:"xmlns,attr"`
	AdminEmail    []typesAdminEmailResp          `xml:"ewp:admin-email"`     //(0-inf)
	AdminNotes    []typesAdminNotesResp          `xml:"ewp:admin-notes"`     //(0-1)
	Version       string                         `xml:"version,attr"`        //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec           []typesHTTPSecurityOptionsResp `xml:"http-security"`       //(0-1)
	GetURL        typesHTTPS                     `xml:"get-url"`             //(1)
	MaxIIAIDs     int                            `xml:"max-iia-ids"`         //(1) positiveInteger
	MaxIIACodes   int                            `xml:"max-iia-codes"`       //(1) positiveInteger
	IndexURL      typesHTTPS                     `xml:"index-url"`           //(1)
	Notifications []typesEmptyResp               `xml:"sends-notifications"` //(0-1) if present
}

//XMLnsManifestIIASCNR schema
var XMLnsManifestIIASCNR = "https://github.com/erasmus-without-paper/ewp-specs-api-iia-cnr/blob/stable-v2/manifest-entry.xsd"

//VersionIIASCNR API
var VersionIIASCNR = "2.0.2"

type typesManifestAPIIIASCNR struct {
	XMLName    xml.Name                       `xml:"iia-cnr"`
	XMLns      string                         `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp          `xml:"ewp:admin-email"` //(0-inf)
	AdminNotes []typesAdminNotesResp          `xml:"ewp:admin-notes"` //(0-1)
	Version    string                         `xml:"version,attr"`    //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec        []typesHTTPSecurityOptionsResp `xml:"http-security"`   //(0-1)
	URL        typesHTTPS                     `xml:"url"`             //(1)
}

//XMLnsManifestoMobility schema
var XMLnsManifestoMobility = "https://github.com/erasmus-without-paper/ewp-specs-api-omobilities/blob/stable-v1/manifest-entry.xsd"

//VersionoMobility API
var VersionoMobility = "0.15.0"

type typesManifestAPIoMobility struct {
	XMLName       xml.Name                               `xml:"omobilities"`
	XMLns         string                                 `xml:"xmlns,attr"`
	AdminEmail    []typesAdminEmailResp                  `xml:"ewp:admin-email"`        //(0-inf)
	AdminNotes    []typesAdminNotesResp                  `xml:"ewp:admin-notes"`        //(0-1)
	Version       string                                 `xml:"version,attr"`           //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec           []typesHTTPSecurityOptionsResp         `xml:"http-security"`          //(0-1)
	GetURL        typesHTTPS                             `xml:"get-url"`                //(1)
	IndexURL      typesHTTPS                             `xml:"index-url"`              //(1)
	UpdateURL     []typesHTTPS                           `xml:"update-url"`             //(0-1)
	MaxMobIDs     int                                    `xml:"max-omobility-ids"`      //(1) positiveInteger
	Notifications []typesEmptyResp                       `xml:"sends-notifications"`    //(0-1) if present
	SupUpdTypes   []typesManifestAPIoMobilitySupUpdTypes `xml:"supported-update-types"` //(0-1)
}

type typesManifestAPIoMobilitySupUpdTypes struct {
	XMLName           xml.Name         `xml:"supported-update-types"`
	ApproveComponents []typesEmptyResp `xml:"approve-components-studied-draft-v1"` //(0-1)
	UpdateComponents  []typesEmptyResp `xml:"update-components-studied-v1"`        //(0-1)
	UpdateStatuses    []typesEmptyResp `xml:"update-statuses-v1"`                  //(0-1)
}

//XMLnsManifestoMobilityCNR schema
var XMLnsManifestoMobilityCNR = "https://github.com/erasmus-without-paper/ewp-specs-api-omobility-cnr/blob/stable-v1/manifest-entry.xsd"

//VersionoMobilityCNR API
var VersionoMobilityCNR = "0.4.1"

type typesManifestAPIoMobilityCNR struct {
	XMLName    xml.Name                       `xml:"omobility-cnr"`
	XMLns      string                         `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp          `xml:"ewp:admin-email"`   //(0-inf)
	AdminNotes []typesAdminNotesResp          `xml:"ewp:admin-notes"`   //(0-1)
	Version    string                         `xml:"version,attr"`      //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec        []typesHTTPSecurityOptionsResp `xml:"http-security"`     //(0-1)
	URL        typesHTTPS                     `xml:"url"`               //(1)
	MaxMobIDs  int                            `xml:"max-omobility-ids"` //(1) positiveInteger
}

//XMLnsManifestiMobility schema
var XMLnsManifestiMobility = "https://github.com/erasmus-without-paper/ewp-specs-api-imobilities/blob/stable-v1/manifest-entry.xsd"

//VersioniMobility API
var VersioniMobility = "0.2.0"

type typesManifestAPIiMobility struct {
	XMLName       xml.Name                       `xml:"imobilities"`
	XMLns         string                         `xml:"xmlns,attr"`
	AdminEmail    []typesAdminEmailResp          `xml:"ewp:admin-email"`     //(0-inf)
	AdminNotes    []typesAdminNotesResp          `xml:"ewp:admin-notes"`     //(0-1)
	Version       string                         `xml:"version,attr"`        //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec           []typesHTTPSecurityOptionsResp `xml:"http-security"`       //(0-1)
	GetURL        typesHTTPS                     `xml:"get-url"`             //(1)
	MaxMobIDs     int                            `xml:"max-omobility-ids"`   //(1) positiveInteger
	Notifications []typesEmptyResp               `xml:"sends-notifications"` //(0-1) if present
}

//XMLnsManifestiMobilityCNR schema
var XMLnsManifestiMobilityCNR = "https://github.com/erasmus-without-paper/ewp-specs-api-imobility-cnr/blob/stable-v1/manifest-entry.xsd"

//VersioniMobilityCNR API
var VersioniMobilityCNR = "0.1.1"

type typesManifestAPIiMobilityCNR struct {
	XMLName    xml.Name                       `xml:"imobility-cnr"`
	XMLns      string                         `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp          `xml:"ewp:admin-email"`   //(0-inf)
	AdminNotes []typesAdminNotesResp          `xml:"ewp:admin-notes"`   //(0-1)
	Version    string                         `xml:"version,attr"`      //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec        []typesHTTPSecurityOptionsResp `xml:"http-security"`     //(0-1)
	URL        typesHTTPS                     `xml:"url"`               //(1)
	MaxMobIDs  int                            `xml:"max-omobility-ids"` //(1) positiveInteger
}

//XMLnsManifestTORS schema
var XMLnsManifestTORS = "https://github.com/erasmus-without-paper/ewp-specs-api-imobility-tors/blob/stable-v1/manifest-entry.xsd"

//VersionTORS API
var VersionTORS = "0.7.0"

type typesManifestAPITORS struct {
	XMLName      xml.Name                       `xml:"imobility-tors"`
	XMLns        string                         `xml:"xmlns,attr"`
	AdminEmail   []typesAdminEmailResp          `xml:"ewp:admin-email"`     //(0-inf)
	AdminNotes   []typesAdminNotesResp          `xml:"ewp:admin-notes"`     //(0-1)
	Version      string                         `xml:"version,attr"`        //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec          []typesHTTPSecurityOptionsResp `xml:"http-security"`       //(0-1)
	GetURL       typesHTTPS                     `xml:"get-url"`             //(1)
	IndexURL     typesHTTPS                     `xml:"index-url"`           //(1)
	MaxTORIDs    int                            `xml:"max-omobility-ids"`   //(1) positiveInteger
	Notification []typesEmptyResp               `xml:"sends-notifications"` //0-1
}

//XMLnsManifestTORSCNR schema
var XMLnsManifestTORSCNR = "https://github.com/erasmus-without-paper/ewp-specs-api-imobility-tor-cnr/blob/stable-v1/manifest-entry.xsd"

//VersionTORSCNR API
var VersionTORSCNR = "0.1.1"

type typesManifestAPITORSCNR struct {
	XMLName    xml.Name                       `xml:"imobility-tor-cnr"`
	XMLns      string                         `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp          `xml:"ewp:admin-email"`   //(0-inf)
	AdminNotes []typesAdminNotesResp          `xml:"ewp:admin-notes"`   //(0-1)
	Version    string                         `xml:"version,attr"`      //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
	Sec        []typesHTTPSecurityOptionsResp `xml:"http-security"`     //(0-1)
	URL        typesHTTPS                     `xml:"url"`               //(1)
	MaxMobIDs  int                            `xml:"max-omobility-ids"` //(1) positiveInteger
}

/**************************************************************************************************************
* Common Types                                                                                   (12/03/2018) *
**************************************************************************************************************/

//XMLnsCommon schema
var XMLnsCommon = "https://github.com/erasmus-without-paper/ewp-specs-architecture/blob/stable-v1/common-types.xsd"

//ArchitectureVersion versão
var ArchitectureVersion = "1.10.0"

type typesStringWithOptionalLangResp struct {
	Lang  []string `xml:"xml:lang,attr"`
	Value string   `xml:",innerxml"`
}

type typesMultilineStringResp string //string limited by back quotes ` `

type typesMultilineStringWithOptionalLangResp struct {
	Lang  []string `xml:"xml:lang,attr"`
	Value string   `xml:",innerxml"` //string limited by back quotes ` `
}

type typesUUIDResp struct {
	Value string `xml:",innerxml"` //pattern: "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
}

type typesASCIIPrintableIdentifierResp struct {
	Value string `xml:",innerxml"` //pattern: "[&#x0021;-&#x007E;]{1,64}"
}

type typesEmailResp string //pattern: "[^@]+@[^\.]+\..+"

type typesSha256Hex string //pattern: "[0-9a-f]{64}"

type typesEmptyResp struct{}

type typesHTTPSResp struct {
	Value string `xml:",innerxml"` //pattern: "https://.+"
}

type typesHTTPResp string //pattern: "https?://.+"

type typesHTTPWithOptionalLangResp struct {
	Lang  []string `xml:"xml:lang,attr"`
	Value string   `xml:",innerxml"` //pattern:https?://.+
}

type typesManifestAPIEntryBaseResp struct {
	XMLns      string                `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmailResp `xml:"ewp:admin-email"` //(0-inf)
	AdminNotes []typesAdminNotesResp `xml:"ewp:admin-notes"` //(0-1)
	Version    string                `xml:"version,attr"`    //(1) pattern: "[0-9]+\.[0-9]+\.[0-9]+"
}

type typesCountryCodeResp struct {
	Value string `xml:",innerxml,regexp='[A-Z][A-Z]'"` //pattern:[A-Z][A-Z]
}

type typesEqfLevelResp struct {
	Value []byte `xml:",innerxml"` //xs:byte <xs:minInclusive value="1"/> <xs:maxInclusive value="8"/>
}

type typesCerfLevelResp struct {
	Value string `xml:",innerxml,regexp:'[ABC][12]'"` //pattern: [ABC][12]
}

type typesGenderResp struct {
	Gender int `xml:",innerxml"` //pattern: 0-not known , 1-male , 2-female , 9-not applicable
}

type typesAdminEmailResp typesEmailResp

type typesAdminNotesResp typesMultilineStringResp

type typesErrorResponseResp struct {
	XMLName     xml.Name                                 `cml:"error-response"`
	DevMessage  typesMultilineStringResp                 `xml:"developer-message"` //1
	UserMessage typesMultilineStringWithOptionalLangResp `xml:"user-message"`      //0-inf
}

type typesSuccessUserMessageResp typesMultilineStringWithOptionalLangResp

/**************************************************************************************************************
* Types Address                                                                                  (12/03/2018) *
**************************************************************************************************************/

//XMLnsAddress schema
var XMLnsAddress = "https://github.com/erasmus-without-paper/ewp-specs-types-address/tree/stable-v1"

//AddressVersion number
var AddressVersion = "1.0.1"

type typesFlexibleAddressResp struct {
	RepName           []string               `xml:"a:recipientName"`     //(0-inf)
	AddressLine       []string               `xml:"a:addressLine"`       //op 1 (0-4)
	BuildingNumber    []string               `xml:"a:buildingNumber"`    //op 2 (0-1)
	BuildingName      []string               `xml:"a:buildingName"`      //op 2 (0-1)
	StreetName        []string               `xml:"a:StreetName"`        //op 2 (0-1)
	Unit              []string               `xml:"a:unit"`              //op 2 (0-1)
	Floor             []string               `xml:"a:floor"`             //op 2 (0-1)
	PostOfficeBox     []string               `xml:"a:postOfficeBox"`     //op 2 (0-1)
	DeliveryPointCode []string               `xml:"a:deliveryPointCode"` //op 2 (0-inf)
	PostalCode        []string               `xml:"a:postalCode"`        //(0-1)
	Locality          []string               `xml:"a:locality"`          //(0-1)
	Region            []string               `xml:"a:region"`            //(0-1)
	Country           []typesCountryCodeResp `xml:"a:country"`           //type: ewp:CountryCode
}

type typesFlexibleAddressOp1Resp struct {
	AddressLine []string `xml:"addressLine"` //op 1 (0-4)
}

type typesFlexibleAddressOp2Resp struct {
	BuildingNumber    []string `xml:"a:buildingNumber"`    //op 2 (0-1)
	BuildingName      []string `xml:"a:buildingName"`      //op 2 (0-1)
	StreetName        []string `xml:"a:StreetName"`        //op 2 (0-1)
	Unit              []string `xml:"a:unit"`              //op 2 (0-1)
	Floor             []string `xml:"a:floor"`             //op 2 (0-1)
	PostOfficeBox     []string `xml:"a:postOfficeBox"`     //op 2 (0-1)
	DeliveryPointCode []string `xml:"a:deliveryPointCode"` //op 2 (0-inf)
}

//type typesFlexibleAddressResp typesFlexibleAddressResp

//type typesAddressMailingAddressResp typesFlexibleAddressResp

/**************************************************************************************************************
* Types Phone Number                                                                             (12/03/2018) *
**************************************************************************************************************/

//XMLnsPhoneNumber schema
var XMLnsPhoneNumber = "https://github.com/erasmus-without-paper/ewp-specs-types-phonenumber/tree/stable-v1"

//PhoneNumberVersion number
var PhoneNumberVersion = "1.0.1"

type typesPhoneNumberResp struct {
	E164        string   `xml:"pn:e164"`         //(0-1) (pattern:\+[0-9]{1,15})
	Ext         int      `xml:"pn:ext"`          //(0-1)
	OtherFormat []string `xml:"pn:other-format"` //(0-1)
}

type typesPhoneResp struct {
	XMLName xml.Name `xml:"pn:phone-number"`
	*typesPhoneNumberResp
}
type typesFaxResp struct {
	XMLName xml.Name `xml:"pn:fax-number"`
	*typesPhoneNumberResp
}

/**************************************************************************************************************
* Types Academic Term                                                                            (12/03/2018) *
**************************************************************************************************************/

//XMLnsAcademicTerm schema
var XMLnsAcademicTerm = "https://github.com/erasmus-without-paper/ewp-specs-types-academic-term/tree/stable-v1"

//AcademicTermVersion number
var AcademicTermVersion = "1.1.0"

type typesAcademicTermResp struct {
	XMLName xml.Name `xml:"trm:academic-term"`
	*typesAcademicTermCompResp
}

type typesAcademicTermCompResp struct {
	AcademicTermID    typesAcademicYearID               `xml:"trm:academic-year-id"` //(1)
	EwpAcademicTermID []typesEwpAcademicTermIDResp      `xml:"trm:ewp-id"`           //(0-1)
	DisplayName       []typesStringWithOptionalLangResp `xml:"trm:display-name"`     //(1-inf)
	StartDate         string                            `xml:"trm:start-date"`       //(1) date
	EndDate           string                            `xml:"trm:end-date"`         //(1) date
}

type typesAcademicYearID string //value="[0-9]{4}/[0-9]{4}"

type typesEwpAcademicTermIDResp string //pattern: "[0-9]{4}/[0-9]{4}-[1-9]/[1-9]"

/**************************************************************************************************************
* Types Contact                                                                                  (12/03/2018) *
**************************************************************************************************************/

//XMLnsContact schema
var XMLnsContact = "https://github.com/erasmus-without-paper/ewp-specs-types-contact/tree/stable-v1"

//ContactVersion number
var ContactVersion = "1.1.0"

type typesComplexContactResp struct {
	ContactName       []typesStringWithOptionalLangResp          `xml:"c:contact-name"`       //(1-inf) type: ewp:StringWithOptionalLang
	PersonGivenNames  []typesStringWithOptionalLangResp          `xml:"c:person-given-names"` //(0-inf) type: ewp:StringWithOptionalLang
	PersonFamilyNames []typesStringWithOptionalLangResp          `xml:"c:person-family-name"` //(0-inf) type: ewp:StringWithOptionalLang
	Gender            []typesGenderResp                          `xml:"ewp:person-gender"`    //(0-1)
	PhoneNumbers      []typesPhoneResp                           `xml:"pn:phone-number"`      //(0-inf)
	FaxNumbers        []typesFaxResp                             `xml:"pn:fax-number"`        //(0-inf)
	Emails            []typesEmailResp                           `xml:"c:email"`              //(0-inf) type: ewp:email
	StreetAddress     []typesFlexibleAddressResp                 `xml:"a:street-address"`     //(0-1)
	MailingAddress    []typesFlexibleAddressResp                 `xml:"a:mailing-address"`    //(0-1)
	RoleDescription   []typesMultilineStringWithOptionalLangResp `xml:"c:role-description"`   //(0-inf) type: ewp:MultilineStringWithOptionalLan
}

type typesContactResp struct {
	XMLName xml.Name `xml:"c:contact"`
	*typesComplexContactResp
}

/**************************************************************************************************************
* Types Person ID                                                                                (12/03/2018) *
**************************************************************************************************************/

//XMLnsPersonID schema
var XMLnsPersonID = "https://github.com/erasmus-without-paper/ewp-specs-types-person-id/tree/stable-v1"

//PersonIDVersion number
var PersonIDVersion = "0.1.1"

type typesUPIDResp struct {
	XMLName xml.Name `xml:"UPID"`
	Value   string   `xml:",innerxml"` //pattern: "[A-Z][A-Z]-[^ ].*[^ ]"
}

/**************************************************************************************************************
* Types Auth And Sec                                                                             (12/03/2018) *
**************************************************************************************************************/

//XMLnsAuthAndSec schema
var XMLnsAuthAndSec = "https://github.com/erasmus-without-paper/ewp-specs-sec-intro/tree/stable-v2"

//AuthAndSecVersion number
var AuthAndSecVersion = "2.0.2"

type typesHTTPSecurityOptionsResp struct {
	CliAuthMethods []typesHTTPSecurityOptionsCliResp `xml:"sec:client-auth-methods"`         //(0-1) <xs:any minOccurs="1" maxOccurs="unbounded" namespace="##other" processContents="lax"/>
	SrvAuthMethods []typesHTTPSecurityOptionsSrvResp `xml:"sec:server-auth-methods"`         //(0-1) <xs:any minOccurs="1" maxOccurs="unbounded" namespace="##other" processContents="lax"/>
	ReqEncMethods  []typesHTTPSecurityOptionsReqResp `xml:"sec:request-encryption-methods"`  //(0-1) <xs:any minOccurs="1" maxOccurs="unbounded" namespace="##other" processContents="lax"/>
	RespEncMethods []typesHTTPSecurityOptionsResResp `xml:"sec:response-encryption-methods"` //(0-1) <xs:any minOccurs="1" maxOccurs="unbounded" namespace="##other" processContents="lax"/>
}

type typesHTTPSecurityOptionsCliResp struct {
	Anon    []typesAuthCliAnonResp    `xml:"anonymous"`
	TLSCert []typesAuthCliTLSCertResp `xml:"tlscert"`
	HTTPSig []typesAuthCliHTTPSigResp `xml:"httpsig"`
}
type typesHTTPSecurityOptionsSrvResp struct {
	TLSCert []typesAuthSrvTLSCertResp `xml:"tlscert"`
	HTTPSig []typesAuthSrvHTTPSigResp `xml:"httpsig"`
}
type typesHTTPSecurityOptionsReqResp struct {
	TLS []typesAuthReqTLSResp `xml:"tls"`
	RSA []typesAuthReqRSAResp `xml:"ewp-rsa-aes128gcm"`
}
type typesHTTPSecurityOptionsResResp struct {
	TLS []typesAuthResTLSResp `xml:"tls"`
	RSA []typesAuthResRSAResp `xml:"ewp-rsa-aes128gcm"`
}

/**************************************************************************************************************
* Types Client Authentication Anonymous                                                          (12/03/2018) *
**************************************************************************************************************/

//XMLnsCliAuthAnon schema
var XMLnsCliAuthAnon = "https://github.com/erasmus-without-paper/ewp-specs-sec-cliauth-none/tree/stable-v1"

//CliAuthAnonVersion number
var CliAuthAnonVersion = "1.1.0"

type typesAuthCliAnonResp struct {
	XMLName xml.Name `xml:"anonymous"`
	XMLns   string   `xml:"xmlns,attr"`
	*typesEmptyResp
}

/**************************************************************************************************************
* Types Client Authentication TLSCert                                                            (12/03/2018) *
**************************************************************************************************************/

//XMLnsCliAuthTLSCert schema
var XMLnsCliAuthTLSCert = "https://github.com/erasmus-without-paper/ewp-specs-sec-cliauth-tlscert/tree/stable-v1"

//CliAuthTLSCertVersion number
var CliAuthTLSCertVersion = "1.1.0"

type typesAuthCliTLSCertResp struct {
	XMLName         xml.Name `xml:"tlscert"`
	XMLns           string   `xml:"xmlns,attr"`
	AllowSelfSigned bool     `xml:"allows-self-signed,attr"` //(1) boolean
	*typesEmptyResp
}

/**************************************************************************************************************
* Types Client Authentication HTTPSig                                                            (12/03/2018) *
**************************************************************************************************************/

//XMLnsCliAuthHTTPSig schema
var XMLnsCliAuthHTTPSig = "https://github.com/erasmus-without-paper/ewp-specs-sec-cliauth-httpsig/tree/stable-v1"

//CliAuthHTTPSigVersion number
var CliAuthHTTPSigVersion = "1.0.1"

type typesAuthCliHTTPSigResp struct {
	XMLName xml.Name `xml:"httpsig"`
	XMLns   string   `xml:"xmlns,attr"`
	*typesEmptyResp
}

/**************************************************************************************************************
* Types Server Authentication TLSCert                                                            (12/03/2018) *
**************************************************************************************************************/

//XMLnsSrvAuthTLSCert schema
var XMLnsSrvAuthTLSCert = "https://github.com/erasmus-without-paper/ewp-specs-sec-srvauth-tlscert/tree/stable-v1"

//SrvAuthTLSCertVersion number
var SrvAuthTLSCertVersion = "1.1.0"

type typesAuthSrvTLSCertResp struct {
	XMLName xml.Name `xml:"tlscert"`
	XMLns   string   `xml:"xmlns,attr"`
	*typesEmptyResp
}

/**************************************************************************************************************
* Types Server Authentication HTTPSig                                                            (12/03/2018) *
**************************************************************************************************************/

//XMLnsSrvAuthHTTPSig schema
var XMLnsSrvAuthHTTPSig = "https://github.com/erasmus-without-paper/ewp-specs-sec-srvauth-httpsig/tree/stable-v1"

//SrvAuthHTTPSigVersion number
var SrvAuthHTTPSigVersion = "1.0.1"

type typesAuthSrvHTTPSigResp struct {
	XMLName xml.Name `xml:"httpsig"`
	XMLns   string   `xml:"xmlns,attr"`
	*typesEmptyResp
}

/**************************************************************************************************************
* Types Request Encription TLS                                                                   (12/03/2018) *
**************************************************************************************************************/

//XMLnsReqAuthTLS schema
var XMLnsReqAuthTLS = "https://github.com/erasmus-without-paper/ewp-specs-sec-reqencr-tls/tree/stable-v1"

//ReqAuthTLSVersion number
var ReqAuthTLSVersion = "1.1.0"

type typesAuthReqTLSResp struct {
	XMLName xml.Name `xml:"tls"`
	XMLns   string   `xml:"xmlns,attr"`
	*typesEmptyResp
}

/**************************************************************************************************************
* Types Request Encryption RSA                                                                   (12/03/2018) *
**************************************************************************************************************/

//XMLnsReqAuthRSA schema
var XMLnsReqAuthRSA = "https://github.com/erasmus-without-paper/ewp-specs-sec-reqencr-rsa-aes128gcm/tree/stable-v1"

//ReqAuthRSAVersion number
var ReqAuthRSAVersion = "0.4.0"

type typesAuthReqRSAResp struct {
	XMLName xml.Name `xml:"ewp-rsa-aes128gcm"`
	XMLns   string   `xml:"xmlns,attr"`
	*typesEmptyResp
}

/**************************************************************************************************************
* Types Response Encription TLS                                                                  (12/03/2018) *
**************************************************************************************************************/

//XMLnsResAuthTLS schema
var XMLnsResAuthTLS = "https://github.com/erasmus-without-paper/ewp-specs-sec-resencr-tls/tree/stable-v1"

//ResAuthTLSVersion number
var ResAuthTLSVersion = "1.1.0"

type typesAuthResTLSResp struct {
	XMLName xml.Name `xml:"tls"`
	XMLns   string   `xml:"xmlns,attr"`
	*typesEmptyResp
}

/**************************************************************************************************************
* Types Response Encryption RSA                                                                  (12/03/2018) *
**************************************************************************************************************/

//XMLnsResAuthRSA schema
var XMLnsResAuthRSA = "https://github.com/erasmus-without-paper/ewp-specs-sec-resencr-rsa-aes128gcm/tree/stable-v1"

//ResAuthRSAVersion number
var ResAuthRSAVersion = "0.3.0"

type typesAuthResRSAResp struct {
	XMLName xml.Name `xml:"ewp-rsa-aes128gcm"`
	XMLns   string   `xml:"xmlns,attr"`
	*typesEmptyResp
}

/**************************************************************************************************************
* Echo API                                                                                      (15/03/2018) *
**************************************************************************************************************/

//XMLnsEchoAPI schema
var XMLnsEchoAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-echo/tree/stable-v2"

type echoResponse struct {
	XMLName  xml.Name `xml:"response"`
	XMLns    string   `xml:"xmlns,attr"` // https://github.com/erasmus-without-paper/ewp-specs-api-echo/blob/stable-v1/response.xsd
	XMLnsXML string   `xml:"xmlns:xml,attr"`
	XMLnsXS  string   `xml:"xmlns:xs,attr"`
	HeiID    []string `xml:"hei-id"` //(0-inf)
	Echo     []string `xml:"echo"`   //(0-inf)
}

/**************************************************************************************************************
* Institutions API                                                                               (14/03/2018) *
**************************************************************************************************************/

//XMLnsInstitutionsAPI schema
var XMLnsInstitutionsAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-institutions/tree/stable-v2"

type institutionsResponse struct {
	XMLName  xml.Name `xml:"institutions-response"`
	XMLns    string   `xml:"xmlns,attr"`
	XMLnsXML string   `xml:"xmlns:xml,attr"`
	XMLnsXS  string   `xml:"xmlns:xs,attr"`
	XMLnsC   string   `xml:"xmlns:c,attr"`
	XMLnsA   string   `xml:"xmlns:a,attr"`
	XMLnsEWP string   `xml:"xmlns:ewp,attr"`
	XMLnsP   string   `xml:"xmlns:pn,attr"`
	XMLnsR   string   `xml:"xmlns:r,attr"`
	//XMLnsTRM string                  `xml:"xml:trm,attr"`
	Hei []institutionsHeiIDResp `xml:"hei"` //(0-inf)
}

type institutionsHeiIDResp struct {
	XMLName              xml.Name                            `xml:"hei"`
	HeiID                string                              `xml:"hei-id"`                 //(1)
	OtherID              []catalogueOtherHeiID               `xml:"other-id"`               //(0-inf) type: r:OtherHeiId
	Name                 []typesStringWithOptionalLangResp   `xml:"name"`                   //(1-inf) type: ewp:StringWithOptionalLang
	Abbreviation         []string                            `xml:"abbreviation"`           //(0-1)
	StreetAddress        []typesFlexibleAddressResp          `xml:"a:street-address"`       //(0-1)
	MailingAddress       []typesFlexibleAddressResp          `xml:"a:mailing-address"`      //(0-1)
	WebsiteURL           []typesHTTPWithOptionalLangResp     `xml:"website-url"`            //(0-inf) type: ewp:HTTPWithOptionalLang
	LogoURL              []typesHTTPSResp                    `xml:"logo-url"`               //(0-1) type: ewp:HTTPS
	MobilityFactsheetURL []typesHTTPWithOptionalLangResp     `xml:"mobility-factsheet-url"` //(0-inf) type: ewp:HTTPWithOptionalLang
	Contacts             []typesContactResp                  `xml:"c:contact"`              //(0-inf)
	RootOunitID          []typesASCIIPrintableIdentifierResp `xml:"root-ounit-id"`          //(0-1) type: ewp:AsciiPrintableIdentifie
	OUnitID              []typesASCIIPrintableIdentifierResp `xml:"ounit-id"`               //(0-inf) type: ewp:AsciiPrintableIdentifier
}

/**************************************************************************************************************
* OUnits API                                                                                     (14/03/2018) *
**************************************************************************************************************/

//XMLnsOUnits schema
var XMLnsOUnits = "https://github.com/erasmus-without-paper/ewp-specs-api-ounits/tree/stable-v2"

type ounitsResponse struct {
	XMLName  xml.Name    `xml:"ounits-response"`
	XMLns    string      `xml:"xmlns,attr"` // https://github.com/erasmus-without-paper/ewp-specs-api-ounits/blob/master/response.xsd
	XMLnsXML string      `xml:"xmlns:xml,attr"`
	XMLnsXS  string      `xml:"xmlns:xs,attr"`
	XMLnsC   string      `xml:"xmlns:c,attr"`
	XMLnsA   string      `xml:"xmlns:a,attr"`
	XMLnsEWP string      `xml:"xmlns:ewp,attr"`
	XMLnsP   string      `xml:"xmlns:pn,attr"`
	XMLnsR   string      `xml:"xmlns:r,attr"`
	OUnits   []ounitResp `xml:"ounit"` //(0-inf)
}

type ounitResp struct {
	XMLName              xml.Name                            `xml:"ounit"`
	OUnitID              typesASCIIPrintableIdentifierResp   `xml:"ounit-id"`               //(1)
	OUnitCode            string                              `xml:"ounit-code"`             //(1)
	Name                 []typesStringWithOptionalLangResp   `xml:"name"`                   //(1-inf)
	Abbreviation         []string                            `xml:"abbreviation"`           //(0-1)
	ParentOUnitID        []typesASCIIPrintableIdentifierResp `xml:"parent-ounit-id"`        //(0-1)
	StreetAddress        []typesFlexibleAddressResp          `xml:"a:street-address"`       //(0-1)
	MailingAddress       []typesFlexibleAddressResp          `xml:"a:mailing-address"`      //(0-1)
	WebsiteURL           []typesHTTPWithOptionalLangResp     `xml:"website-url"`            //(0-inf) type: ewp:HTTPWithOptionalLang
	LogoURL              []typesHTTPSResp                    `xml:"logo-url"`               //(0-1) type: ewp:HTTPS
	MobilityFactsheetURL []typesHTTPWithOptionalLangResp     `xml:"mobility-factsheet-url"` //(0-inf) type: ewp:HTTPWithOptionalLang
	Contact              []typesContactResp                  `xml:"c:contact"`              //(0-inf) type: c:Contact
}

/**************************************************************************************************************
* Courses API                                                                                    (16/03/2018) *
**************************************************************************************************************/

//XMLnsCoursesAPI schema
var XMLnsCoursesAPI = "https://raw.githubusercontent.com/erasmus-without-paper/ewp-specs-api-courses/master/response.xsd"

type coursesResponse struct {
	XMLName  xml.Name `xml:"courses-response"`
	XMLns    string   `xml:"xmlns,attr"` // https://github.com/erasmus-without-paper/ewp-specs-api-ounits/blob/master/response.xsd
	XMLnsXML string   `xml:"xmlns:xml,attr"`
	XMLnsXS  string   `xml:"xmlns:xs,attr"`
	//XMLnsC   string                            `xml:"xmlns:c,attr"`
	//XMLnsA   string                            `xml:"xmlns:a,attr"`
	XMLnsEWP string `xml:"xmlns:ewp,attr"`
	//XMLnsP   string                            `xml:"xmlns:pn,attr"`
	//XMLnsR   string                            `xml:"xmlns:r,attr"`*/
	XMLnsTRM string                            `xml:"xmlns:trm,attr"`
	OUnits   []courseLearningOpportunitiesResp `xml:"learningOpportunitySpecification"` //(0-inf)
}

type courseLearningOpportunitiesResp struct {
	XMLName     xml.Name                                   `xml:"learningOpportunitySpecification"`
	LOSID       typesASCIIPrintableIdentifierResp          `xml:"los-id"`           //(1) pattern="(CR|CLS|MOD|DEP)/(.{1,40})"
	LOSCode     []string                                   `xml:"los-code"`         //(0-1)
	OUnitID     []typesASCIIPrintableIdentifierResp        `xml:"ounit-id"`         //(0-1)
	Title       []typesStringWithOptionalLangResp          `xml:"title"`            //(1-inf)
	Type        []string                                   `xml:"type"`             //(0-1) Degree Programme, Module, Course, Class
	SubjectArea []string                                   `xml:"subjectArea"`      //(0-1)
	IscedCode   []string                                   `xml:"iscedCode"`        //(0-1)
	EqfLevel    []typesEqfLevelResp                        `xml:"eqfLevelProvided"` //(0-1)
	URL         []typesHTTPResp                            `xml:"url"`              //(0-1)
	Description []typesMultilineStringWithOptionalLangResp `xml:"description"`      //(0-inf)
	Specifies   []courseSpecifiesResp                      `xml:"specifies"`        //(0-1)
	Contains    []courseContainsResp                       `xml:"contains"`         //(0-1)
}

type courseSpecifiesResp struct {
	XMLName xml.Name           `xml:"specifies"`
	LOInst  []courseLOInstResp `xml:"learningOpportunityInstance"` //(0-inf)
}

type courseLOInstResp struct {
	XMLName               xml.Name                          `xml:"learningOpportunityInstance"`
	LoiID                 typesASCIIPrintableIdentifierResp `xml:"loi-id"`                //(1) pattern="(CRI|CLSI|MODI|DEPI)/(.{1,40})"
	Start                 string                            `xml:"start"`                 //(1) date
	End                   string                            `xml:"end"`                   //(1) date
	AcademicTerm          []typesAcademicTermResp           `xml:"trm:academic-term"`     //(0-1)
	GrandingScheme        []courseGrandingSchemeResp        `xml:"grandingScheme"`        //(0-1)
	ResultDistribution    []courseResultDistributionResp    `xml:"resultDistribution"`    //(0-1)
	Credit                []coursesCreditResp               `xml:"credit"`                //(0-inf)
	LanguageOfInstruction []string                          `xml:"languageOfInstruction"` //(0-1) xs:language
	EngagementHours       []float32                         `xml:"engagementHours"`       //(0-1) xs:decimal
}

type courseGrandingSchemeResp struct {
	XMLName     xml.Name                                   `xml:"gradingScheme"`
	Label       []typesStringWithOptionalLangResp          `xml:"label"`       //(1-inf)
	Description []typesMultilineStringWithOptionalLangResp `xml:"description"` //(0-inf)
}

type courseResultDistributionResp struct {
	XMLName     xml.Name                                   `xml:"resultDistribution"`
	Category    []courseCategoryResp                       `xml:"category"`    //(1-inf)
	Description []typesMultilineStringWithOptionalLangResp `xml:"description"` //(0-inf)
}

type courseCategoryResp struct {
	XMLName xml.Name `xml:"category"`
	Label   string   `xml:"label,attr"` //required
	Count   int      `xml:"count,attr"` //xs:nonNegativeInteger required
}

type coursesCreditResp struct {
	XMLName xml.Name `xml:"credit"`
	Scheme  string   `xml:"scheme"` //(1)
	Level   []string `xml:"level"`  //(0-1) Bachelor, Master, PhD
	Value   float32  `xml:"value"`  //(1) xs:decimal
}

type courseContainsResp struct {
	XMLName xml.Name                            `xml:"contains"`
	LOSID   []typesASCIIPrintableIdentifierResp `xml:"los-id"` //(1) pattern="(CR|CLS|MOD|DEP)/(.{1,40})"
}

/**************************************************************************************************************
* Interinstitutional Agreements API                                                              (16/03/2018) *
**************************************************************************************************************/

//XMLnsIIAIndexAPI schema
var XMLnsIIAIndexAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-iias/blob/stable-v2/endpoints/index-response.xsd"

//XMLnsIIAGetAPI schema
var XMLnsIIAGetAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-iias/blob/stable-v2/endpoints/get-response.xsd"

type iiasIndexResp struct {
	XMLName  xml.Name                            `xml:"iias-index-response"`
	XMLns    string                              `xml:"xmlns,attr"`
	XMLnsXML string                              `xml:"xmlns:xml,attr"`
	XMLnsEWP string                              `xml:"xmlns:ewp,attr"`
	IIA      []typesASCIIPrintableIdentifierResp `xml:"iia-id"`
}

type iiasGetResp struct {
	XMLName  xml.Name  `xml:"iias-get-response"`
	XMLns    string    `xml:"xmlns,attr"`
	XMLnsXML string    `xml:"xmlns:xml,attr"`
	XMLnsC   string    `xml:"xmlns:c,attr"`
	XMLnsEWP string    `xml:"xmlns:ewp,attr"`
	XMLnsTRM string    `xml:"xmlns:trm,attr"`
	IIA      []iiaResp `xml:"iia"`
}

type iiaResp struct {
	XMLName  xml.Name         `xml:"iia"`
	Partner  []iiaPartnerResp `xml:"partner"` //2-inf
	InEffect bool             `xml:"in-effect"`
	CoopCond iiaCoopCondResp  `xml:"cooperation-conditions"`
}

type iiaPartnerResp struct {
	XMLName    xml.Name                            `xml:"partner"`
	HeiID      string                              `xml:"hei-id"`
	OUnitID    []typesASCIIPrintableIdentifierResp `xml:"ounit-id"`        //0-1
	IIAID      []typesASCIIPrintableIdentifierResp `xml:"iia-id"`          //0-1
	IIACode    []string                            `xml:"iia-code"`        //0-1
	SigContact []typesContactResp                  `xml:"signing-contact"` //0-1
	SigDate    []string                            `xml:"signing-date"`    //0-1 xs:date
	Contact    []typesContactResp                  `xml:"c:contact"`       //0-inf ref=c:contact
}

type iiaCoopCondResp struct {
	XMLName    xml.Name            `xml:"cooperation-conditions"`
	StudMob    []iiaStudMobSpec    `xml:"student-studies-mobility-spec"`    //0-inf
	StudTrain  []iiaStudTrainSpec  `xml:"student-trainsheep-mobility-spec"` //0-inf
	StaffTeach []iiaStaffTeachSpec `xml:"staff-teacher-mobility-spec"`      //0-inf
	StaffTrain []iiaStaffTrainSpec `xml:"staff-training-mobility-spec"`     //0-inf
}

type iiaStudMobSpec struct {
	*iiaMobStudSpec
	EqfLevel []typesEqfLevelResp `xml:"eqf-level"` //1-inf
}

type iiaStudTrainSpec struct {
	*iiaMobStudSpec
}

type iiaStaffTeachSpec struct {
	*iiaMobStaffSpec
}

type iiaStaffTrainSpec struct {
	*iiaMobStaffSpec
}

type iiaMobStudSpec struct {
	*iiaMobSpec
	AVGMon []int `xml:"avg-months"`   //um ou outro xs:decimal min=0 and max 2 decimal
	TOTMon []int `xml:"total-months"` //um ou outro xs:decimal min=0 and max 2 decimal
}

type iiaMobStaffSpec struct {
	*iiaMobSpec
	AVGDay []int `xml:"avg-days"`   //um ou outro xs:decimal min=0 and max 2 decimal
	TOTDay []int `xml:"total-days"` //um ou outro xs:decimal min=0 and max 2 decimal
}

type iiaMobSpec struct {
	SendHEI      string                              `xml:"sending-hei-id"`
	SendOUnit    []typesASCIIPrintableIdentifierResp `xml:"sending-ounit-id"` //0-inf
	SendContact  []typesContactResp                  `xml:"sending-contact"`  //0-inf
	RecHEI       string                              `xml:"receiving-hei-id"`
	RecOUnit     []typesASCIIPrintableIdentifierResp `xml:"receiving-ounit-id"`         //0-inf
	RecContact   []typesContactResp                  `xml:"receiving-contact"`          //0-inf
	RecAcadYear  []typesAcademicYearID               `xml:"receiving-academic-year-id"` //1-inf
	MobPerYear   int                                 `xml:"mobilities-per-year"`
	RecLangSkill []iiaRecLangSkill                   `xml:"recommended-language-skill"` //0-inf
	ISCEDCode    []string                            `xml:"isced-f-code"`               //0-1
}

type iiaRecLangSkill struct {
	Language  string               `xml:"language"`
	CEFRLevel []typesCerfLevelResp `xml:"cefr-level"` //[ABC][12] 0-1
}

/**************************************************************************************************************
* Outgoing Mobility API                                                                          (14/03/2018) *
**************************************************************************************************************/

//XMLnsOMobilityIndexAPI schema
var XMLnsOMobilityIndexAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-omobilities/blob/stable-v1/endpoints/index-response.xsd"

//XMLnsOMobilityGetAPI schema
var XMLnsOMobilityGetAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-omobilities/blob/stable-v1/endpoints/get-response.xsd"

//XMLnsMobilityUpdateAPI schema
//var XMLnsMobilityUpdateAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-omobilities/blob/stable-v1/endpoints/update-response.xsd"

type oMobilityIndexResp struct {
	XMLName  xml.Name                            `xml:"omobilities-index-response"`
	XMLns    string                              `xml:"xmlns,attr"`
	XMLnsXML string                              `xml:"xmlns:xml,attr"`
	XMLnsXS  string                              `xml:"xmlns:xs,attr"`
	XMLnsEWP string                              `xml:"xmlns:ewp,attr"`
	MobID    []typesASCIIPrintableIdentifierResp `xml:"omobility-id"` //(0-inf)
}

type oMobilityGetResp struct {
	XMLName    xml.Name                        `xml:"omobilities-get-response"`
	XMLns      string                          `xml:"xmlns,attr"`
	XMLnsXML   string                          `xml:"xmlns:xml,attr"`
	XMLnsXS    string                          `xml:"xmlns:xs,attr"`
	XMLnsEWP   string                          `xml:"xmlns:ewp,attr"`
	XMLnsA     string                          `xml:"xmlns:a,attr"`
	XMLnsP     string                          `xml:"xmlns:pn,attr"`
	XMLnsC     string                          `xml:"xmlns:c,attr"`
	SequenceOM []oMobilityGetMobForStudiesResp `xml:"student-mobility-for-studies"` //(0-inf)
}

type oMobilityGetMobForStudiesResp struct {
	XMLName          xml.Name                                    `xml:"student-mobility-for-studies"`
	MobID            typesASCIIPrintableIdentifierResp           `xml:"omobility-id"`                 //(1)
	SendHei          oMobilityGetMobForStudiesSendReceiveHeiResp `xml:"sending-hei"`                  //(1)
	RecHei           oMobilityGetMobForStudiesSendReceiveHeiResp `xml:"receiving-hei"`                //(1)
	ChoiceSendAcad   []typesEwpAcademicTermIDResp                `xml:"sending-academic-term-ewp-id"` //(este ou o abaixo)
	ChoiceNonStand   []typesEmptyResp                            `xml:"non-standard-mobility-period"` //(este ou o acima!)
	RecAcadYearID    typesAcademicYearID                         `xml:"receiving-academic-year-id"`   //(1)
	Student          oMobilityGetMobForStudiesStudentResp        `xml:"student"`                      //(1)
	Status           string                                      `xml:"status"`                       //(1) - nomination, live, recognized, cancelled
	PlannedArrival   string                                      `xml:"planned-arrival-date"`         //(1) xs:date
	PlannedDeparture string                                      `xml:"planned-departure-date"`       //(1) xs:date
	//ActualArrival        string                                      `xml:"actual-arrival-date"`             //(1) xs:date
	//ActualDeparture      string                                      `xml:"actual-departure-date"`           //(1) xs:date
	EQFInitial           []typesEqfLevelResp                     `xml:"eqf-level-studied-at-nomination"` //(0-1)
	EQFEnd               []typesEqfLevelResp                     `xml:"eqf-level-studied-at-departure"`  //(0-1)
	ISCEDCode            []string                                `xml:"nominee-isced-f-code"`            //(0-1)
	LanguageSkill        []oMobilityGetMobForStudiesLanguageResp `xml:"nominee-language-skill"`          //(0-inf)
	ComponentsStudied    oMobilityGetMobForStudiesCompStudResp   `xml:"components-studied"`              //(1)
	ComponentsRecognized []oMobilityGetMobForStudiesCompStudResp `xml:"components-recognized"`           //(0-1)

}

type oMobilityGetMobForStudiesSendReceiveHeiResp struct {
	HeiID   string                              `xml:"hei-id"`   //(1)
	OUnitID []typesASCIIPrintableIdentifierResp `xml:"ounit-id"` //(0-1)
	IIAID   []typesASCIIPrintableIdentifierResp `xml:"iia-id"`   //(0-1)
}

type oMobilityGetMobForStudiesStudentResp struct {
	XMLName        xml.Name                                    `xml:"student"`
	GivenNames     []typesStringWithOptionalLangResp           `xml:"given-names"`       //(1-inf)
	FamilyNames    []typesStringWithOptionalLangResp           `xml:"family-name"`       //(1-inf)
	BirthDate      []string                                    `xml:"birth-date"`        //(0-1) type: xs:date
	Citizenship    []typesCountryCodeResp                      `xml:"citizenship"`       //(0-1)
	Gender         []int                                       `xml:"gender"`            //(0-1) 0 (not known), 1 (male), 2 (female), 9 (not applicable)
	Email          []typesEmailResp                            `xml:"email"`             //(0-1)
	PhotoURL       []oMobilityGetMobForStudiesStudentPhotoResp `xml:"photo-url"`         //(0-inf)
	StreetAddress  []typesFlexibleAddressResp                  `xml:"a:street-address"`  //(0-1)
	MailingAddress []typesFlexibleAddressResp                  `xml:"a:mailing-address"` //(0-1)
	PhoneNumber    []typesPhoneNumberResp                      `xml:"p:phone-number"`    //(0-inf)
}

type oMobilityGetMobForStudiesStudentPhotoResp struct {
	XMLName xml.Name `xml:"photo-url"`
	Value   string   `xml:",innerxml"`    //pattern: "https?://.+"
	Size    []string `xml:"size-px,attr"` //optional value="[0-9]+x[0-9]+"
	Public  []bool   `xml:"public,attr"`  //optional type="xs:boolean" default="false"
	Date    []string `xml:"date,attr"`    //optional type="xs:date"
}

type oMobilityGetMobForStudiesLanguageResp struct {
	Language  string               `xml:"language"`   //(1) xs:language
	CERFLevel []typesCerfLevelResp `xml:"cefr-level"` //(0-1)
}

type oMobilityGetMobForStudiesCompStudResp struct {
	BeforeMobChanges   []oMobilityGetMobForStudiesCompStudLOCResp `xml:"before-mobility-changes"`  //(1) type="ListOfChangesTo_ComponentsStudied" The "before oMobility" and "latest approved" sections are optional.
	BeforeMobSnap      []oMobilityGetMobForStudiesCompStudSOResp  `xml:"before-mobility-snapshot"` //(1) type="SnapshotOf_ComponentsStudied"
	LatestAppChanges   []oMobilityGetMobForStudiesCompStudLOCResp `xml:"latest-approved-changes"`  //(1) type="ListOfChangesTo_ComponentsStudied"
	LatestAppSnap      []oMobilityGetMobForStudiesCompStudSOResp  `xml:"latest-approved-snapshot"` //(1) type="SnapshotOf_ComponentsStudied" However, if "before oMobility" section exists, then we also require "latest approved" section to exist
	LatestDraftChanges oMobilityGetMobForStudiesCompStudLOCResp   `xml:"latest-draft-changes"`     //(1) type="ListOfChangesTo_ComponentsStudied"
	LatestDraftSnap    oMobilityGetMobForStudiesCompStudSOResp    `xml:"latest-draft-snapshot"`    //(1) type="SnapshotOf_ComponentsStudied"
}

type oMobilityGetMobForStudiesCompStudLOCResp struct {
	Insert []oMobilityGetMobForStudiesCompStudLOCICSResp `xml:"insert-component-studied"` //este ou o outro
	Remove []oMobilityGetMobForStudiesCompStudLOCRCSResp `xml:"remove-component-studied"` //este ou o outro
}

type oMobilityGetMobForStudiesCompStudLOCICSResp struct {
	Reason []oMobilityGetMobForStudiesSCTALReasonResp `xml:"reason"` //(0-1)
	Index  int                                        `xml:"index"`  //(1) type="xs:nonNegativeInteger"
	CS     oMobilityGetMobForStudiesCompStudCSCResp   `xml:"component-studied"`
}

type oMobilityGetMobForStudiesCompStudLOCRCSResp struct {
	*oMobilityGetMobForStudiesSCTALDResp
}

type oMobilityGetMobForStudiesCompRecLOCResp struct {
	Insert []oMobilityGetMobForStudiesCompRecSOICRResp `xml:"insert-component-recognized"` //este ou o outro
	Remove []oMobilityGetMobForStudiesCompRecSORCRResp `xml:"remove-component-recognized"` //este ou o outro
}

type oMobilityGetMobForStudiesCompRecSOICRResp struct {
	XMLName xml.Name                                     `xml:"insert-component-recognized"`
	Insert  oMobilityGetMobForStudiesCompRecSOICRRespAux `xml:"insert-component-recognized-aux"`
}

type oMobilityGetMobForStudiesCompRecSOICRRespAux struct {
	Reason []oMobilityGetMobForStudiesSCTALReasonResp `xml:"reason"` //(0-1)
	Index  int                                        `xml:"index"`  //(1) type="xs:nonNegativeInteger"
	CS     oMobilityGetMobForStudiesCompStudCSCResp   `xml:"component-studied"`
}

type oMobilityGetMobForStudiesCompRecSORCRResp struct {
	XMLName xml.Name `xml:"remove-component-recognized"`
	*oMobilityGetMobForStudiesSCTALDResp
}

type oMobilityGetMobForStudiesSCTALIResp struct {
	Reason []oMobilityGetMobForStudiesSCTALReasonResp `xml:"reason"` //(0-1)
	Index  int                                        `xml:"index"`  //(1) type="xs:nonNegativeInteger"
}

type oMobilityGetMobForStudiesSCTALResp struct {
	Reason []oMobilityGetMobForStudiesSCTALReasonResp `xml:"reason"` //(0-1)
}

type oMobilityGetMobForStudiesSCTALReasonResp struct {
	DisplayText   typesMultilineStringResp `xml:"display-text"`    //(1)
	EWPReasonCode []string                 `xml:"ewp-reason-code"` //(0-1) not-available, language-mismatch, timetable-conflict, substituting-deleted, extended-mobility, properties-update
}

type oMobilityGetMobForStudiesSCTALDResp struct {
	Reason []oMobilityGetMobForStudiesSCTALReasonResp `xml:"reason"` //(0-1)
	Index  int                                        `xml:"index"`  //(1) type="xs:nonNegativeInteger"
}

type oMobilityGetMobForStudiesCompStudCSCResp struct {
	XMLName xml.Name                                         `xml:"component-studied"`
	LOSID   []typesASCIIPrintableIdentifierResp              `xml:"los-id"` //(0-1)
	LOSCode []string                                         `xml:"los-code"`
	Title   string                                           `xml:"title"`                      //(1)
	LOIID   []typesASCIIPrintableIdentifierResp              `xml:"loi-id"`                     //(0-1)
	ATDN    []string                                         `xml:"academic-term-display-name"` //(0-1)
	Credit  []oMobilityGetMobForStudiesCompStudCSCCreditResp `xml:"credit"`                     //(0-inf)
}

type oMobilityGetMobForStudiesCompStudCSCCreditResp struct {
	XMLName xml.Name `xml:"credit"`
	Scheme  string   `xml:"scheme"` //(1)
	Value   string   `xml:"value"`  //(1) type="decimal"
}

type oMobilityGetMobForStudiesCompStudSOResp struct {
	CS            []oMobilityGetMobForStudiesCompStudCSCResp                `xml:"component-studied"`         //(0-inf)
	Approval      []oMobilityGetMobForStudiesCompStudSOApprovalResp         `xml:"approval"`                  //(0-3)
	SNBAB         []oMobilityGetMobForStudiesCompStudSOApprovalAppPartyResp `xml:"should-now-be-approved-by"` //(0-3) type="ApprovingParty"
	InEffectSince []string                                                  `xml:"in-effect-since,attr"`      //(optional) type="xs:dateTime"
}

type oMobilityGetMobForStudiesCompStudSOApprovalResp struct {
	ByParty   oMobilityGetMobForStudiesCompStudSOApprovalAppPartyResp `xml:"by-party"`  //(1) type="ApprovingParty" //student, sending-hei, receiving-hei
	Timestamp []string                                                `xml:"timestamp"` //(0-1) type="xs:dateTime"
}

type oMobilityGetMobForStudiesCompStudSOApprovalAppPartyResp string //student, sending-hei, receiving-hei

type oMobilityGetMobForStudiesCompRecSOResp struct {
	CS []oMobilityGetMobForStudiesCompStudCSCResp `xml:"component-studied"` //(0-inf)
}

/**************************************************************************************************************
* Incoming Mobility API                                                                          (14/03/2018) *
**************************************************************************************************************/

//XMLnsiMobilityGetAPI schema
var XMLnsiMobilityGetAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-imobilities/blob/stable-v1/endpoints/get-response.xsd"

type imobilityGetResp struct {
	XMLName    xml.Name                        `xml:"imobilities-get-response"`
	XMLns      string                          `xml:"xmlns,attr"`
	XMLnsXML   string                          `xml:"xmlns:xml,attr"`
	XMLnsXS    string                          `xml:"xmlns:xs,attr"`
	XMLnsEWP   string                          `xml:"xmlns:ewp,attr"`
	XMLnsA     string                          `xml:"xmlns:a,attr"`
	XMLnsP     string                          `xml:"xmlns:pn,attr"`
	XMLnsC     string                          `xml:"xmlns:c,attr"`
	SequenceOM []imobilityGetMobForStudiesResp `xml:"student-mobility-for-studies"` //(0-inf)
}

type imobilityGetMobForStudiesResp struct {
	XMLName    xml.Name                          `xml:"student-mobility-for-studies"`
	MobID      typesASCIIPrintableIdentifierResp `xml:"omobility-id"`          //1
	Status     string                            `xml:"status"`                //1 (pending, verified, rejected)
	ActArrDate []string                          `xml:"actual-arrival-date"`   //0-1
	ActDepDate []string                          `xml:"actual-departure-date"` //0-1
}

/**************************************************************************************************************
* Transcript of Records API                                                                      (16/03/2018) *
**************************************************************************************************************/

//XMLEMREXAPI schema
var XMLEMREXAPI = "https://github.com/emrex-eu/elmo-schemas/tree/v1"

//XMLnsTORSGetAPI schema
var XMLnsTORSGetAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-imobility-tors/blob/stable-v1/endpoints/get-response.xsd"

//XMLnsTORSIndexAPI schema
var XMLnsTORSIndexAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-imobility-tors/blob/stable-v1/endpoints/index-response.xsd"

type torsIndexResp struct {
	XMLName  xml.Name `xml:"imobility-tors-index-response"`
	XMLns    string   `xml:"xmlns,attr"`
	XMLnsXML string   `xml:"xmlns:xml,attr"`
	XMLnsXSI string   `xml:"xmlns:xsi,attr"`
	//XMLEMREX string        `xml:"xmlns:emrex,attr"`
	TORS []typesASCIIPrintableIdentifierResp `xml:"omobility-id"` //(0-inf)
}

type torsGetResp struct {
	XMLName  xml.Name `xml:"imobility-tors-get-response"`
	XMLns    string   `xml:"xmlns,attr"`
	XMLnsXML string   `xml:"xmlns:xml,attr"`
	XMLnsXSI string   `xml:"xmlns:xsi,attr"`
	//XMLEMREX string        `xml:"xmlns:emrex,attr"`
	TOR []torsTORResp `xml:"tor"`
}

type torsTORResp struct {
	XMLName   xml.Name                          `xml:"tor"`
	MobID     typesASCIIPrintableIdentifierResp `xml:"omobility-id"`
	EMREXELMO emrexResp                         `xml:"elmo"`
}

type emrexResp struct {
	XMLName  xml.Name          `xml:"elmo"`
	XMLEMREX string            `xml:"xmlns,attr"`
	GenDate  string            `xml:"generatedDate"` //dateTime
	Learner  emrexLearnerResp  `xml:"learner"`
	Report   []emrexReportResp `xml:"report"`
	Attach   []emrexAttachResp `xml:"attachment"` //0-inf
	//Extension emrexExtensionResp `xml:"extension"`  //0-inf  CustomExtensionsContainer sequece <xs:any minOccurs="0" maxOccurs="unbounded" namespace="##other" processContents="lax"/>
	//Signature emrexSigResp
}

type emrexLearnerResp struct {
	Citizenship []typesCountryCodeResp `xml:"citizenship"` //0-inf
	Identifier  []emrexIdentResp       `xml:"identifier"`  //0-inf
	GivenNames  string                 `xml:"givenNames"`  //xs:token
	FamilyName  string                 `xml:"familyName"`  //xs:token
	BDay        []string               `xml:"bday"`        //xs:date
}

type emrexIdentResp struct {
	XMLName xml.Name `xml:"identifier"`
	Type    string   `xml:"type,attr"` //type=xs:token
	Value   string   `xml:",innerxml"`
}

type emrexReportResp struct {
	Issuer    emrexRepIssuerResp `xml:"issuer"`
	LOSpec    []emrexLOSpecResp  `xml:"learningOpportunitySpecification"` //0-inf learningOpportunitySpecification
	IssueDate string             `xml:"issueDate"`                        //xs:dateTime
	Attach    []emrexAttachResp  `xml:"attachment"`                       //0-inf
}

type emrexRepIssuerResp struct {
	Country    []typesCountryCodeResp            `xml:"country"`    //0-inf
	Identifier []emrexRepIssuerIdentResp         `xml:"identifier"` //1-inf  <xs:extension base="xs:token">  <xs:attribute name="type" type="xs:token" use="required">
	Title      []typesStringWithOptionalLangResp `xml:"title"`      //1-inf
	URL        typesHTTPSResp                    `xml:"url"`
}

type emrexRepIssuerIdentResp struct {
	XMLName xml.Name `xml:"identifier"`
	Type    string   `xml:"type,attr"` //pic, erasmus, schac
	Value   string   `xml:",innerxml"`
}

type emrexExtensionResp string

type emrexAttachResp struct {
	Title     []typesStringWithOptionalLangResp          `xml:"title"`       //0-inf
	Type      []string                                   `xml:"type"`        //0-1 Diploma Supplement, Transcript of Records, EMREX trasncript, Letter of Nomination, Certificate of Training, Learning Agreement
	Desc      []typesMultilineStringWithOptionalLangResp `xml:"description"` //0-1
	Content   string                                     `xml:"content"`     //base64
	Extension []emrexExtensionResp                       `xml:"extension"`   //0-1
}

type emrexLOSpecResp struct {
	Identifier  []emrexIdentResp                           `xml:"identifier"`  //0-inf token
	Title       []typesStringWithOptionalLangResp          `xml:"title"`       //1-inf
	Type        []string                                   `xml:"type"`        //0-inf Degree Programme, Module, Course, Class
	SubjectArea []string                                   `xml:"subjectArea"` //(0-1)
	IscedCode   []string                                   `xml:"iscedCode"`   //(0-1)
	URL         []typesHTTPResp                            `xml:"url"`         //(0-1)
	Description []typesMultilineStringWithOptionalLangResp `xml:"description"` //(0-inf)
	Specifies   []emrexLOSpecifiesResp                     `xml:"specifies"`   //(0-1)
	HasPart     []emrexLOSpecHPResp                        `xml:"contains"`    //(0-1)
}

type emrexLOSpecHPResp struct {
	LOS []emrexLOSpecResp `xml:"learningOpportunitySpecification"`
}

type emrexLOSpecifiesResp struct {
	Instance  emrexLOInstResp      `xml:"learningOpportunityInstance"`
	Extension []emrexExtensionResp `xml:"extension"` //0-inf
}

type emrexLOInstResp struct {
	Start        []string                   `xml:"start"`                 //0-inf xs:date
	Date         []string                   `xml:"date"`                  //0-inf xs:date
	AcademicTerm []emrexLOInstAcadTermResp  `xml:"academic-term"`         //0-inf xs:date
	ResultLabel  []string                   `xml:"resultLabel"`           //
	ShortGrad    []emrexLOInstShortGradResp `xml:"shortenedGrading"`      //0-
	ResultDist   []emrexLOIInstResDistResp  `xml:"resultDistribution"`    //0-
	Credit       []emrexLOIInstCredResp     `xml:"credit"`                //0-inf
	LangOI       []string                   `xml:"languageOfInstruction"` //0-
	EngageHours  []float32                  `xml:"engagementHours"`       //0-
	Extension    []emrexExtensionResp       `xml:"extension"`             //0-inf
}

type emrexLOInstAcadTermResp struct {
	Title []typesStringWithOptionalLangResp `xml:"title"` //xs:token inf
	Start string                            `xml:"start"` //0-inf xs:date
	End   string                            `xml:"date"`  //0-inf xs:date
}

type emrexLOInstShortGradResp struct {
	PercLow  float32 `xml:"percentageLower"`
	PercEq   float32 `xml:"percentageEqual"`
	PercHigh float32 `xml:"percentageHigher"`
}

type emrexLOIInstResDistResp struct {
	Category    emrexLOIInstResDistCatResp        `xml:"category"`    //0-inf
	Description []typesStringWithOptionalLangResp `xml:"description"` //0-
}

type emrexLOIInstResDistCatResp struct {
	XMLName xml.Name `xml:"category"`
	Label   string   `xml:"label,attr"`
	Count   int      `xml:"count,attr"`
}

type emrexLOIInstCredResp struct {
	Scheme string    `xml:"scheme"`
	Level  []string  `xml:"level"` //0- Bachelor, Master, PhD
	Value  []float32 `xml:"value"` //0-
}

type emrexSigResp struct {
	XMLnsNS        string `xml:"xmlns,attr"`
	SignedInfo     emrexSigInfoResp
	SignatureValue string `xml:"SignatureValue"`
	KeyInfo        emrexSigKeyInfoResp
}

type emrexSigInfoResp struct {
	CanonicalizationMethod emrexSigAttrAlgResp
	SignatureMethod        emrexSigAttrAlgResp
	Reference              emrexSigRefResp
}

type emrexSigRefResp struct {
	URI          string `xml:"URI,attr"`
	Transforms   emrexSigRefTransResp
	DigestMethod emrexSigAttrAlgResp
	DigestValue  string `xml:"DigestValue"`
}

type emrexSigRefTransResp struct {
	Transform emrexSigAttrAlgResp
}

type emrexSigKeyInfoResp struct {
	X509Data emrexSigKeyInfox509Resp
}

type emrexSigKeyInfox509Resp struct {
	X509SubjectName string `xml:"X509SubjectName"`
	X509Certificate string `xml:"X509Certificate"`
}

type emrexSigAttrAlgResp struct {
	Algorithm string `xml:"Algorithm,attr"`
}

/**************************************************************************************************************
* Outros                                                                                                      *
**************************************************************************************************************/

type empty string

type updatesStructList struct {
	XMLName xml.Name        `xml:"updates"`
	Updates []updatesStruct `xml:"update"`
}

type updatesStruct struct {
	XMLName  xml.Name `xml:"update"`
	UpdateID int      `xml:"id"`
	State    int      `xml:"state"`
}

type apistruct struct {
	API     string
	URL     string
	Version string
}

/**************************************************************************************************************
* CNR                                                                                                      *
**************************************************************************************************************/

//XMLnsIIACNRAPI schema
var XMLnsIIACNRAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-iia-cnr/tree/stable-v2"

type iiaCNRResponse struct {
	XMLName  xml.Name `xml:"iia-cnr-response"`
	XMLns    string   `xml:"xmlns,attr"`
	XMLnsXML string   `xml:"xmlns:xml,attr"`
	XMLnsXS  string   `xml:"xmlns:xs,attr"`
	XMLnsEWP string   `xml:"xmlns:ewp,attr"`
}

//XMLnsOMobCNRAPI schema
var XMLnsOMobCNRAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-omobility-cnr/tree/stable-v1"

type oMobCNRResponse struct {
	XMLName  xml.Name `xml:"omobility-cnr-response"`
	XMLns    string   `xml:"xmlns,attr"`
	XMLnsXML string   `xml:"xmlns:xml,attr"`
	XMLnsXS  string   `xml:"xmlns:xs,attr"`
	XMLnsEWP string   `xml:"xmlns:ewp,attr"`
}

//XMLnsIMobCNRAPI schema
var XMLnsIMobCNRAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-imobility-cnr/tree/stable-v1"

type iMobCNRResponse struct {
	XMLName  xml.Name `xml:"imobility-cnr-response"`
	XMLns    string   `xml:"xmlns,attr"`
	XMLnsXML string   `xml:"xmlns:xml,attr"`
	XMLnsXS  string   `xml:"xmlns:xs,attr"`
	XMLnsEWP string   `xml:"xmlns:ewp,attr"`
}

//XMLnsTORCNRAPI schema
var XMLnsTORCNRAPI = "https://github.com/erasmus-without-paper/ewp-specs-api-imobility-tor-cnr/tree/stable-v1"

type torCNRResponse struct {
	XMLName  xml.Name `xml:"imobility-tor-cnr-response"`
	XMLns    string   `xml:"xmlns,attr"`
	XMLnsXML string   `xml:"xmlns:xml,attr"`
	XMLnsXS  string   `xml:"xmlns:xs,attr"`
	XMLnsEWP string   `xml:"xmlns:ewp,attr"`
}
