package open311

import "fmt"

// Service are categories that requests can be submitted to
type Service struct {
	ServiceCode string   `json:"service_code"`
	ServiceName string   `json:"service_name"`
	Description string   `json:"description"`
	Metadata    bool     `json:"metadata"`
	Type        string   `json:"type"`
	Keywords    []string `json:"keywords"`
	Group       string   `json:"group"`
}

// Provide a list of acceptable 311 service request types and their associated service codes.
// These request types can be unique to the city/jurisdiction.
func (c *Client) getServiceList() ([]Service, error) {
	var s []Service
	if err := c.get("/services.json", &s); err != nil {
		return s, err
	}
	return s, nil
}

// ServiceDefinition describes a service's metadata
type ServiceDefinition struct {
	ServiceCode string       `json:"service_code"`
	Definition  []Attributes `json:"attributes"`
}

// Attributes are the metadata of a service
type Attributes struct {
	Variable            bool       `json:"variable"`
	Code                string     `json:"code"`
	DataType            string     `json:"datatype"`
	Required            bool       `json:"required"`
	DataTypeDescription string     `json:"string"`
	Order               int        `json:"order"`
	Description         string     `json:"description"`
	Values              []KeyValue `json:"values"`
}

// KeyValue is used for arbitrary key/value parameters
type KeyValue struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

// GetServiceDefinition returns details about a type of service request
// This call is only necessary if the Service selected has metadata set as true from the GET Services
// response
func (c *Client) GetServiceDefinition(code string) (ServiceDefinition, error) {
	var d ServiceDefinition
	if err := c.get(fmt.Sprintf("/services/%s.json", code), &d); err != nil {
		return d, err
	}
	return d, nil
}
