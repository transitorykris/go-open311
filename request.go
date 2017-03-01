// Provides a client for Open311. May be San Francisco specific for now.
// http://wiki.open311.org/GeoReport_v2/

package open311

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

// Location is the location for a service request. Either:
// Lat/Long or Address or AddressID is required.
type Location struct {
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
	Address   string  `json:"address_string"`
	AddressID string  `json:"address_id"`
}

// ServiceRequest is used to create a new service request
type ServiceRequest struct {
	Lat       float64 `json:"lat" url:"lat,omitempty"`
	Long      float64 `json:"long" url:"long,omitempty"`
	Address   string  `json:"address_string" url:"address_string,omitempty"`
	AddressID string  `json:"address_id" url:"address_id,omitempty"`

	Email       string `json:"email" url:"email"`
	DeviceID    string `json:"device_id" url:"device_id"`
	AccountID   string `json:"account_id" url:"account_id"`
	FirstName   string `json:"first_name" url:"first_name"`
	LastName    string `json:"last_name" url:"last_name"`
	Phone       string `json:"phone" url:"phone"`
	Description string `json:"description" url:"description"`
	MediaURL    string `json:"media_url" url:"media_url"`
}

// ServiceRequestCreation contains the response for a service request
type ServiceRequestCreation struct {
	ID            string `json:"service_request_id"`
	Token         string `json:"token"`
	ServiceNotice string `json:"service_notice"`
	AccountID     string `json:"account_id"`
}

// ReqType is a required attribute for creating a service request
const (
	Poop = iota
	Needles
	Garbage
)

// loc is lat/long or address_string or address_id
// attr is an array of key/values
func (c *Client) postServiceRequest(code string, reqType int, req ServiceRequest) (ServiceRequestCreation, error) {
	var r []ServiceRequestCreation

	v, err := query.Values(req)
	if err != nil {
		return ServiceRequestCreation{}, err
	}
	v.Add("api_key", c.key)
	v.Add("jurisdiction_id", c.jurisdiction)
	v.Add("service_code", code)
	var form string
	switch reqType {
	case Poop:
		form = v.Encode() + "&attribute[request_type]=Human_waste_or_urine"
	case Needles:
		form = v.Encode() + "&attribute[request_type]=Needles_less_than_20"
	case Garbage:
		form = v.Encode() + "&attribute[request_type]=Other_loose_garbage_debris_yard_waste"
	}
	formBody := strings.NewReader(form)
	resp, err := http.Post(c.url+"/requests.json", "application/x-www-form-urlencoded", formBody)
	if err != nil {
		return ServiceRequestCreation{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ServiceRequestCreation{}, err
	}
	if err = json.Unmarshal(body, &r); err != nil {
		return ServiceRequestCreation{}, err
	}

	return r[0], nil
}

// ServiceRequestID is returned when looking up with a token
type ServiceRequestID struct {
	ID    string `json:"service_request_id"`
	Token string `json:"token"`
}

// XXX SF311 uses the token as the request ID
// XXX this endpoint does not return the correct format
// http://wiki.open311.org/GeoReport_v2/#get-service_request_id-from-a-token
func (c *Client) getServiceRequestID(token string) (string, error) {
	var id ServiceRequestID
	getURL := fmt.Sprintf("%s/tokens/%s.json", c.url, token)
	err := c.get(getURL, &id)
	return id.ID, err
}

// ServiceRequestOpts contains optional arguments when getting service requests
type ServiceRequestOpts struct {
	ServiceRequestID string    `json:"service_request_id"`
	ServiceCode      string    `json:"service_code"`
	StartDate        time.Time `json:"start_date"`
	EndData          time.Time `json:"end_date"`
	Status           string    `json:"status"`
}

// ServiceRequestResponse is returned after opening a service request
type ServiceRequestResponse struct {
	ServiceRequestID  string    `json:"service_request_id"`
	Status            string    `json:"status"`
	StatusNotes       string    `json:"status_notes"`
	ServiceName       string    `json:"service_name"`
	ServiceCode       string    `json:"service_code"`
	Description       string    `json:"description"`
	AgencyResponsible string    `json:"agency_responsible"`
	ServiceNotice     string    `json:"service_notice"`
	RequestedTime     time.Time `json:"requested_datetime"`
	UpdateTime        time.Time `json:"updated_datetime"`
	ExpectedTime      time.Time `json:"expected_datetime"`
	Zipcode           string    `json:"zipcode"`
	Location
}

func (c *Client) getServiceRequests(opts ServiceRequestOpts) ([]ServiceRequestResponse, error) {
	return []ServiceRequestResponse{}, nil
}

func (c *Client) getServiceRequest(id string) (ServiceRequestResponse, error) {
	var s []ServiceRequestResponse
	if err := c.get(fmt.Sprintf("/requests/%s.json", id), &s); err != nil {
		return ServiceRequestResponse{}, err
	}
	return s[0], nil
}
