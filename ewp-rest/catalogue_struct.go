package main

import "encoding/xml"

/*****************************************
* Catalogue                (23/Fev/2017) *
*****************************************/

//XMLnsRegistry schema
var XMLnsRegistry = "https://github.com/erasmus-without-paper/ewp-specs-api-registry/blob/stable-v1/manifest-entry.xsd"

//XMLnsAPIRegistry schema
var XMLnsAPIRegistry = "https://github.com/erasmus-without-paper/ewp-specs-api-registry/tree/stable-v1"

//XMLnsCatalogue schema
var XMLnsCatalogue = "https://github.com/erasmus-without-paper/ewp-specs-api-registry/blob/stable-v1/catalogue.xsd"

type catalogue struct {
	XMLName xml.Name `xml:"catalogue"`
	//XMLns        string                `xml:"xmlns,attr"`   // https://github.com/erasmus-without-paper/ewp-specs-api-registry/blob/stable-v1/catalogue.xsd
	Hosts        []catalogueHosts      `xml:"host"`         //(1-inf)
	Institutions catalogueInstitutions `xml:"institutions"` //(1)
	Binaries     []catalogueBinaries   `xml:"binaries"`     //(0-1)
}

type catalogueHosts struct {
	//XMLns               string                   `xml:"xmlns,attr"`
	AdminEmail []typesAdminEmail `xml:"admin-email"` //(0-inf)
	AdminNotes []typesAdminNotes `xml:"admin-notes"` //(0-1)
	//APISImplemented []catalogueAPISImplemented `xml:"apis-implemented"` //(0-1) ref
	InstitutionsCovered []catalogueInstCovered `xml:"institutions-covered"`      //(0-1)
	CliCredInUse        catalogueCliCredInUse  `xml:"client-credentials-in-use"` //(0-1)
	SvrCredInUse        catalogueSvrCredInUse  `xml:"server-credentials-in-use"` //(0-1)
}

type catalogueInstCovered struct {
	XMLName xml.Name `xml:"institutions-covered"`
	HeiID   []string `xml:"hei-id"` //(0-inf)
}

type catalogueCliCredInUse struct {
	XMLName     xml.Name               `xml:"client-credentials-in-use"`
	Certificate []catalogueCliCredCert `xml:"certificate"`    //(0-inf)
	RSACert     []catalogueCliCredRSA  `xml:"rsa-public-key"` //(0-inf)
}
type catalogueSvrCredInUse struct {
	XMLName xml.Name              `xml:"server-credentials-in-use"`
	RSACert []catalogueCliCredRSA `xml:"rsa-public-key"` //(0-inf)
}

type catalogueCliCredCert struct {
	XMLName xml.Name       `xml:"certificate"`
	SHA256  typesSha256Hex `xml:"sha-256,attr"` //required pattern: [0-9a-f]{64}
}
type catalogueCliCredRSA struct {
	XMLName xml.Name       `xml:"rsa-public-key"`
	SHA256  typesSha256Hex `xml:"sha-256,attr"` //(0-inf) pattern: [0-9a-f]{64}
	Value   string         `xml:",innerxml"`
}

type catalogueInstitutions struct {
	XMLName xml.Name       `xml:"institutions"`
	Heis    []catalogueHei `xml:"hei"` //(0-inf) ref hei
}

type catalogueAPISImplemented struct {
	XMLName         xml.Name                        `xml:"apis-implemented"`
	DiscoverAPI     []typesCatalogueAPIDiscovery    `xml:"discovery"`
	EchoAPI         []typesCatalogueAPIEcho         `xml:"echo"`
	InstitutionsAPI []typesCatalogueAPIInstitutions `xml:"institutions"`
	OUnitsAPI       []typesCatalogueAPIOUnits       `xml:"organizational-units"`
	CoursesAPI      []typesCatalogueAPICourses      `xml:"courses"`
	IIASAPIS        []typesCatalogueAPIIIAS         `xml:"iias"`
	OMobilityAPI    []typesCatalogueAPIoMobility    `xml:"omobilities"`
	IMobilityAPI    []typesCatalogueAPIiMobility    `xml:"imobilities"`
	TORSAPI         []typesCatalogueAPITORS         `xml:"tors"`
}

type catalogueHei struct {
	XMLName xml.Name                      `xml:"hei"`
	OtherID []catalogueOtherHeiID         `xml:"other-id"` //(0-inf)
	Name    []typesStringWithOptionalLang `xml:"name"`     //(1-inf)
	ID      string                        `xml:"id,attr"`  //required
}

type catalogueOtherHeiID struct {
	Type  string `xml:"type,attr"` //previous-schac, pic, erasmus, euc
	Value string `xml:",innerxml"`
}

type catalogueBinaries struct {
	RSACert []catalogueCliCredRSA `xml:"rsa-public-key"` //(0-inf)
}

// catalogue APIs (12/03/2018)

type typesCatalogueAPIDiscovery struct {
	XMLName xml.Name `xml:"discovery"`
	*typesManifestAPIEntryBaseResp
	URL typesHTTPS `xml:"url"` //(1)
}

type typesCatalogueAPIEcho struct {
	XMLName xml.Name `xml:"echo"`
	*typesManifestAPIEntryBaseResp
	Sec []catalogueHTTPSecurityOptionsResp `xml:"http-security"` //(0-1)
	URL typesHTTPS                         `xml:"url"`           //(1)
}

type typesCatalogueAPIInstitutions struct {
	XMLName xml.Name `xml:"institutions"`
	*typesManifestAPIEntryBaseResp
	Sec       []catalogueHTTPSecurityOptionsResp `xml:"http-security"` //(0-1)
	URL       typesHTTPS                         `xml:"url"`           //(1)
	MaxHeiIDs int                                `xml:"max-hei-ids"`   //(1) pattern:positive integer
}

type typesCatalogueAPIOUnits struct {
	XMLName xml.Name `xml:"organizational-units"`
	*typesManifestAPIEntryBaseResp
	Sec       []catalogueHTTPSecurityOptionsResp `xml:"http-security"`   //(0-1)
	URL       typesHTTPS                         `xml:"url"`             //(1)
	MaxOUID   int                                `xml:"max-ounit-ids"`   //(1) pattern:positive integer
	MaxOUCode int                                `xml:"max-ounit-codes"` //(1) pattern:positive integer
}

type typesCatalogueAPICourses struct {
	XMLName xml.Name `xml:"courses"`
	*typesManifestAPIEntryBaseResp
	Sec        []catalogueHTTPSecurityOptionsResp `xml:"http-security"` //(0-1)
	URL        typesHTTPS                         `xml:"url"`           //(1)
	MaxLOSID   int                                `xml:"max-los-ids"`   //(1) pattern:positive integer
	MaxLOSCode int                                `xml:"max-los-codes"` //(1) pattern:positive integer
}

type typesCatalogueAPICourseRep struct {
	XMLName xml.Name `xml:"simple-course-replication"`
	*typesManifestAPIEntryBaseResp
	Sec         []catalogueHTTPSecurityOptionsResp `xml:"http-security"`           //(0-1)
	URL         typesHTTPS                         `xml:"url"`                     //(1)
	AllowsAnon  bool                               `xml:"allows-anonymous-access"` //(1)
	SupModSince bool                               `xml:"supports-modified-since"` //(1)
}

type typesCatalogueAPIIIAS struct {
	XMLName xml.Name `xml:"iias"`
	*typesManifestAPIEntryBaseResp
	Sec           []catalogueHTTPSecurityOptionsResp `xml:"http-security"`       //(0-1)
	GetURL        typesHTTPS                         `xml:"get-url"`             //(1)
	MaxIIAIDs     int                                `xml:"max-iia-ids"`         //(1) positiveInteger
	MaxIIACodes   int                                `xml:"max-iia-codes"`       //(1) positiveInteger
	IndexURL      typesHTTPS                         `xml:"index-url"`           //(1)
	Notifications []typesEmptyResp                   `xml:"sends-notifications"` //(0-1) if present
}

type typesCatalogueAPIIIASCNR struct {
	XMLName xml.Name `xml:"iia-cnr"`
	*typesManifestAPIEntryBaseResp
	Sec []catalogueHTTPSecurityOptionsResp `xml:"http-security"` //(0-1)
	URL typesHTTPS                         `xml:"url"`           //(1)
}

type typesCatalogueAPIoMobility struct {
	XMLName xml.Name `xml:"omobilities"`
	*typesManifestAPIEntryBaseResp
	Sec           []catalogueHTTPSecurityOptionsResp `xml:"http-security"`       //(0-1)
	GetURL        typesHTTPS                         `xml:"get-url"`             //(1)
	IndexURL      typesHTTPS                         `xml:"index-url"`           //(1)
	UpdateURL     []typesHTTPS                       `xml:"update-url"`          //(0-1)
	MaxMobIDs     int                                `xml:"max-omobility-ids"`   //(1) positiveInteger
	Notifications []typesEmptyResp                   `xml:"sends-notifications"` //(0-1) if present
	*typesManifestAPIoMobilitySupUpdTypes
}

type typesCatalogueAPIoMobilityCNR struct {
	XMLName xml.Name `xml:"omobility-cnr"`
	*typesManifestAPIEntryBaseResp
	Sec       []catalogueHTTPSecurityOptionsResp `xml:"http-security"`     //(0-1)
	URL       typesHTTPS                         `xml:"url"`               //(1)
	MaxMobIDs int                                `xml:"max-omobility-ids"` //(1) positiveInteger
}

type typesCatalogueAPIiMobility struct {
	XMLName xml.Name `xml:"imobilities"`
	*typesManifestAPIEntryBaseResp
	Sec           []catalogueHTTPSecurityOptionsResp `xml:"http-security"`       //(0-1)
	GetURL        typesHTTPS                         `xml:"get-url"`             //(1)
	MaxMobIDs     int                                `xml:"max-omobility-ids"`   //(1) positiveInteger
	Notifications []typesEmptyResp                   `xml:"sends-notifications"` //(0-1) if present
}

type typesCatalogueAPIiMobilityCNR struct {
	XMLName xml.Name `xml:"imobility-cnr"`
	*typesManifestAPIEntryBaseResp
	Sec       []catalogueHTTPSecurityOptionsResp `xml:"http-security"`     //(0-1)
	URL       typesHTTPS                         `xml:"url"`               //(1)
	MaxMobIDs int                                `xml:"max-omobility-ids"` //(1) positiveInteger
}

type typesCatalogueAPITORS struct {
	XMLName xml.Name `xml:"imobility-tors"`
	*typesManifestAPIEntryBaseResp
	Sec          []catalogueHTTPSecurityOptionsResp `xml:"http-security"`       //(0-1)
	GetURL       typesHTTPS                         `xml:"get-url"`             //(1)
	IndexURL     typesHTTPS                         `xml:"index-url"`           //(1)
	MaxTORIDs    int                                `xml:"max-omobility-ids"`   //(1) positiveInteger
	Notification []typesEmptyResp                   `xml:"sends-notifications"` //0-1
}

type typesCatalogueAPITORSCNR struct {
	XMLName xml.Name `xml:"imobility-tor-cnr"`
	*typesManifestAPIEntryBaseResp
	Sec       []catalogueHTTPSecurityOptionsResp `xml:"http-security"`     //(0-1)
	URL       typesHTTPS                         `xml:"url"`               //(1)
	MaxMobIDs int                                `xml:"max-omobility-ids"` //(1) positiveInteger
}

type catalogueHTTPSecurityOptionsResp struct {
	CliAuthMethods []typesHTTPSecurityOptionsCliResp `xml:"client-auth-methods"`         //(0-1) <xs:any minOccurs="1" maxOccurs="unbounded" namespace="##other" processContents="lax"/>
	SrvAuthMethods []typesHTTPSecurityOptionsSrvResp `xml:"server-auth-methods"`         //(0-1) <xs:any minOccurs="1" maxOccurs="unbounded" namespace="##other" processContents="lax"/>
	ReqEncMethods  []typesHTTPSecurityOptionsReqResp `xml:"request-encryption-methods"`  //(0-1) <xs:any minOccurs="1" maxOccurs="unbounded" namespace="##other" processContents="lax"/>
	RespEncMethods []typesHTTPSecurityOptionsResResp `xml:"response-encryption-methods"` //(0-1) <xs:any minOccurs="1" maxOccurs="unbounded" namespace="##other" processContents="lax"/>
}
