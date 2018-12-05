package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

var devMessage typesMultilineStringResp

func printErrors(w http.ResponseWriter, r *http.Request, status int, devMessage typesMultilineStringResp, userMessages []typesMultilineStringWithOptionalLangResp) {

	erro := typesErrorResponse{
		XMLnsXS:     XMLnsXS,
		XMLnsXML:    XMLnsXML,
		XMLns:       XMLnsCommon,
		DevMessage:  devMessage,
		UserMessage: userMessages,
	}

	x, err := xml.MarshalIndent(erro, "", "  ")
	if err != nil {
		status = http.StatusInternalServerError
		fmt.Println(err)
		writeResponse(r, w, status, []byte{})
	} else {
		writeResponse(r, w, status, x)
	}
}

// Invalid parameters
func getError400(w http.ResponseWriter, r *http.Request, reason string) {
	devMessage = "The request cannot be fulfilled."
	userMessages := []typesMultilineStringWithOptionalLangResp{
		typesMultilineStringWithOptionalLangResp{Value: "The parameters passed do not correspond to the server specifications."},
		typesMultilineStringWithOptionalLangResp{Value: "Please correct your request."},
		typesMultilineStringWithOptionalLangResp{Value: reason},
	}

	printErrors(w, r, http.StatusBadRequest, devMessage, userMessages)
}

// Invalid parameters
func getError401(w http.ResponseWriter, r *http.Request, reason string) {
	devMessage = "The request cannot be fulfilled."
	userMessages := []typesMultilineStringWithOptionalLangResp{
		typesMultilineStringWithOptionalLangResp{Value: "Forbidden."},
		typesMultilineStringWithOptionalLangResp{Value: "Please correct your request."},
		typesMultilineStringWithOptionalLangResp{Value: reason},
	}

	printErrors(w, r, http.StatusUnauthorized, devMessage, userMessages)
}

// Forbidden Access
func getError403(w http.ResponseWriter, r *http.Request, reason string) {
	devMessage = "The request was a legal request, but the server is refusing to respond to it."
	userMessages := []typesMultilineStringWithOptionalLangResp{
		typesMultilineStringWithOptionalLangResp{Value: "You don't have permissions to access this data."},
		typesMultilineStringWithOptionalLangResp{Value: "Reason: " + reason},
	}
	printErrors(w, r, http.StatusForbidden, devMessage, userMessages)
}

// Inappropriate method
func getError405(w http.ResponseWriter, r *http.Request) {
	devMessage = "A request was made of a page using a request method not supported by that page."
	userMessages := []typesMultilineStringWithOptionalLangResp{
		typesMultilineStringWithOptionalLangResp{Value: "The method used to request the data is not accepted."},
		typesMultilineStringWithOptionalLangResp{Value: "Please change your request method."},
	}
	printErrors(w, r, http.StatusMethodNotAllowed, devMessage, userMessages)
}

// Unsupported
func getError415(w http.ResponseWriter, r *http.Request) {
	devMessage = "The server will not accept the request, because the media type is not supported."
	userMessages := []typesMultilineStringWithOptionalLangResp{
		typesMultilineStringWithOptionalLangResp{Value: "The type of your request is not supported."},
		typesMultilineStringWithOptionalLangResp{Value: "Please review your request in order to match the correct type."},
	}
	printErrors(w, r, http.StatusUnsupportedMediaType, devMessage, userMessages)
}

// Server Error
func getError500(w http.ResponseWriter, r *http.Request) {
	devMessage = "A generic error message, given when no more specific message is suitable."
	userMessages := []typesMultilineStringWithOptionalLangResp{
		typesMultilineStringWithOptionalLangResp{Value: "Something went wrong on our side. We apologize for any inconvenience."},
		typesMultilineStringWithOptionalLangResp{Value: "Please check if your request is legal and if the problem persists contact the administrators."},
	}
	printErrors(w, r, http.StatusInternalServerError, devMessage, userMessages)
}

// Server Error
func getError503(w http.ResponseWriter, r *http.Request) {
	devMessage = "The server is currently unavailable (overloaded or down)."
	userMessages := []typesMultilineStringWithOptionalLangResp{
		typesMultilineStringWithOptionalLangResp{Value: "We are currently making some changes in order to make our service better."},
		typesMultilineStringWithOptionalLangResp{Value: "We apologize for any inconvenience."},
	}
	printErrors(w, r, http.StatusServiceUnavailable, devMessage, userMessages)
}
