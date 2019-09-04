//
// Automatically generated. DO NOT EDIT.
//

package resources

type JunosServicePorts struct {
	ServicePort []string `json:"service_port,omitempty"`
}

func (obj *JunosServicePorts) AddServicePort(value string) {
	obj.ServicePort = append(obj.ServicePort, value)
}
