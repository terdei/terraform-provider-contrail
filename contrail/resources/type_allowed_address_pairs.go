//
// Automatically generated. DO NOT EDIT.
//

package resources

type AllowedAddressPairs struct {
	AllowedAddressPair []AllowedAddressPair `json:"allowed_address_pair,omitempty"`
}

func (obj *AllowedAddressPairs) AddAllowedAddressPair(value *AllowedAddressPair) {
	obj.AllowedAddressPair = append(obj.AllowedAddressPair, *value)
}
