//
// Automatically generated. DO NOT EDIT.
//

package resources

type MacAddressesType struct {
	MacAddress []string `json:"mac_address,omitempty"`
}

func (obj *MacAddressesType) AddMacAddress(value string) {
	obj.MacAddress = append(obj.MacAddress, value)
}
