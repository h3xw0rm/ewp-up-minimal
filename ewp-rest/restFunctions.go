package main

import (
	"encoding/xml"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func getEWPManifest() *typeManifestResp {
	dat, err := ioutil.ReadFile(RawPubKey)
	if err != nil {
		fmt.Println(err)
	}
	certificate := strings.TrimSpace(string(dat))

	dat2, err := ioutil.ReadFile(RawRSAPubKey)
	if err != nil {
		fmt.Println(err)
	}
	rsaPublicKey := strings.TrimSpace(string(dat2))

	sec := []typesHTTPSecurityOptionsResp{
		typesHTTPSecurityOptionsResp{
			CliAuthMethods: []typesHTTPSecurityOptionsCliResp{
				typesHTTPSecurityOptionsCliResp{
					TLSCert: []typesAuthCliTLSCertResp{
						typesAuthCliTLSCertResp{
							XMLns:           XMLnsCliAuthTLSCert,
							AllowSelfSigned: true,
						},
					},
					HTTPSig: []typesAuthCliHTTPSigResp{
						typesAuthCliHTTPSigResp{
							XMLns: XMLnsCliAuthHTTPSig,
						},
					},
				},
			},
			SrvAuthMethods: []typesHTTPSecurityOptionsSrvResp{
				typesHTTPSecurityOptionsSrvResp{
					TLSCert: []typesAuthSrvTLSCertResp{
						typesAuthSrvTLSCertResp{
							XMLns: XMLnsSrvAuthTLSCert,
						},
					},
					HTTPSig: []typesAuthSrvHTTPSigResp{
						typesAuthSrvHTTPSigResp{
							XMLns: XMLnsSrvAuthHTTPSig,
						},
					},
				},
			},
			ReqEncMethods: []typesHTTPSecurityOptionsReqResp{
				typesHTTPSecurityOptionsReqResp{
					TLS: []typesAuthReqTLSResp{
						typesAuthReqTLSResp{
							XMLns: XMLnsReqAuthTLS,
						},
					},
					/*RSA: []typesAuthReqRSAResp{
						typesAuthReqRSAResp{
							XMLns: XMLnsReqAuthRSA,
						},
					},*/
				},
			},
			RespEncMethods: []typesHTTPSecurityOptionsResResp{
				typesHTTPSecurityOptionsResResp{
					TLS: []typesAuthResTLSResp{
						typesAuthResTLSResp{
							XMLns: XMLnsResAuthTLS,
						},
					},
					/*RSA: []typesAuthResRSAResp{
						typesAuthResRSAResp{
							XMLns: XMLnsResAuthRSA,
						},
					},*/
				},
			},
		},
	}

	hostURL := Hostname

	if strings.Compare(os.Getenv("EWP_REST_EXT_PORT"), "443") != 0 {
		hostURL = hostURL + ":" + os.Getenv("EWP_REST_EXT_PORT")
	}

	man := typeManifestResp{
		XMLns:      XMLnsManifest,
		XMLnsEWP:   XMLnsCommon,
		XMLnsR:     XMLnsAPIRegistry,
		XMLnsSec:   XMLnsAuthAndSec,
		XMLnsSecA:  XMLnsCliAuthAnon,
		XMLnsSecCT: XMLnsCliAuthTLSCert,
		XMLnsSecCH: XMLnsCliAuthHTTPSig,
		XMLnsSecST: XMLnsSrvAuthTLSCert,
		XMLnsSecSH: XMLnsSrvAuthHTTPSig,
		Hosts: []typeManifestHostResp{
			typeManifestHostResp{
				AdminEmail: []typesAdminEmailResp{
					typesAdminEmailResp(os.Getenv("EWP_DEV_EMAIL")),
				},
				AdminNotes: []typesAdminNotesResp{typesAdminNotesResp("Manifest File for University of " + os.Getenv("EWP_UNIV_NAME_EN") + ".")},
				APIsImplemented: []typeCatalogueAPISImplementedResp{
					typeCatalogueAPISImplementedResp{
						Discovery: []typesManifestAPIDiscovery{
							typesManifestAPIDiscovery{
								XMLns:   XMLnsManifestDiscovery,
								Version: VersionDiscovery,
								URL:     typesHTTPS{Value: "https://" + Hostname + "/rest/manifest"},
							},
						},
						Echo: []typesManifestAPIEcho{
							typesManifestAPIEcho{
								XMLns:   XMLnsManifestEcho,
								Version: VersionEcho,
								Sec:     sec,
								URL:     typesHTTPS{Value: "https://" + hostURL + "/rest/echo"},
							},
						},
						Institution: []typesManifestAPIInstitutions{
							typesManifestAPIInstitutions{
								XMLns:     XMLnsManifestInstitutions,
								Version:   VersionInstitutions,
								Sec:       sec,
								URL:       typesHTTPS{Value: "https://" + hostURL + "/rest/institutions"},
								MaxHeiIDs: MaxHeiIDs,
							},
						},
					},
				},
				InstitutionsCovered: []typeManifestInstCovered{
					typeManifestInstCovered{
						Heis: []typeCatalogueHeiResp{
							typeCatalogueHeiResp{
								ID: os.Getenv("EWP_SCHAC"),
								OtherID: []catalogueOtherHeiID{
									catalogueOtherHeiID{
										Type:  "pic",
										Value: os.Getenv("EWP_PIC"),
									},
									catalogueOtherHeiID{
										Type:  "erasmus",
										Value: os.Getenv("EWP_ERASMUS_1") + " " + os.Getenv("EWP_ERASMUS_2"),
									},
									catalogueOtherHeiID{
										Type:  "euc",
										Value: os.Getenv("EWP_EUC"),
									},
								},
								Name: []typesStringWithOptionalLangResp{
									typesStringWithOptionalLangResp{
										Lang:  []string{"en"},
										Value: "University of " + os.Getenv("EWP_UNIV_NAME_EN"),
									},
								},
							},
						},
					},
				},
				ClientCredInUse: []typeManifestCliCredInUse{
					typeManifestCliCredInUse{
						Certificate: []string{
							certificate,
						},
						RSAPubKey: []string{
							rsaPublicKey,
						},
					},
				},
				ServerCredInUse: []typeManifestSvrCredInUse{
					typeManifestSvrCredInUse{
						RSAPubKey: []string{
							rsaPublicKey,
						},
					},
				},
			},
		},
	}

	return &man
}

func retrieveEcho(w http.ResponseWriter, r *http.Request, hostsCovered []catalogueHosts, queryValues map[string][]string) {

	echosForm := queryValues["echo"]

	var echosResp []string

	for _, formEcho := range echosForm {
		echosResp = append(echosResp, html.EscapeString(formEcho))
	}

	var heiIDsResp []string

	for _, hostCov := range hostsCovered {
		for _, heiID := range hostCov.InstitutionsCovered[0].HeiID {
			heiIDsResp = append(heiIDsResp, heiID)
		}
	}

	response := echoResponse{
		XMLnsXS:  XMLnsXS,
		XMLnsXML: XMLnsXML,
		XMLns:    XMLnsEchoAPI,
		Echo:     echosResp,
		HeiID:    heiIDsResp,
	}

	x, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marsheling")
		fmt.Println(err)
		getError500(w, r)
		return
	}

	writeResponse(r, w, http.StatusOK, x)

}

func retrieveInstitutions(w http.ResponseWriter, r *http.Request, hostsCovered []catalogueHosts, queryValues map[string][]string) {

	heiIDsQueried := len(queryValues["hei_id"])

	if heiIDsQueried > MaxHeiIDs {
		getError400(w, r, "")
		return
	}

	instResp := institutionsResponse{
		XMLns:    XMLnsInstitutionsAPI,
		XMLnsXML: XMLnsXML,
		XMLnsXS:  XMLnsXS,
		XMLnsEWP: XMLnsCommon,
		XMLnsC:   XMLnsContact,
		XMLnsA:   XMLnsAddress,
		XMLnsP:   XMLnsPhoneNumber,
		XMLnsR:   XMLnsRegistryAPI,
		Hei:      []institutionsHeiIDResp{},
	}

	for _, validHeiID := range AcceptedInstitutions {
		for _, queriedHeiID := range queryValues["hei_id"] {
			if strings.Compare(validHeiID, queriedHeiID) == 0 {
				inst := institutionsHeiIDResp{
					HeiID: os.Getenv("EWP_SCHAC"),
					OtherID: []catalogueOtherHeiID{
						catalogueOtherHeiID{
							Type:  "pic",
							Value: os.Getenv("EWP_PIC"),
						},
						catalogueOtherHeiID{
							Type:  "erasmus",
							Value: os.Getenv("EWP_ERASMUS_1") + " " + os.Getenv("EWP_ERASMUS_2"),
						},
					},
					Name: []typesStringWithOptionalLangResp{
						typesStringWithOptionalLangResp{
							Lang:  []string{"en"},
							Value: "University of " + os.Getenv("EWP_UNIV_NAME_EN"),
						},
					},
					Abbreviation: []string{os.Getenv("EWP_UNIV_ABBREVIATION")},
					RootOunitID:  []typesASCIIPrintableIdentifierResp{typesASCIIPrintableIdentifierResp{Value: os.Getenv("EWP_SCHAC")}},
				}
				instResp.Hei = append(instResp.Hei, inst)
			}
		}
	}

	status := http.StatusOK

	x, err := xml.MarshalIndent(instResp, "", "  ")
	if err != nil {
		status = http.StatusInternalServerError
	}

	writeResponse(r, w, status, x)

}
