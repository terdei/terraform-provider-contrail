//
// Automatically generated. DO NOT EDIT.
//

package resources

type DhcpOptionsListType struct {
	DhcpOption []DhcpOptionType `json:"dhcp_option,omitempty"`
}

func (obj *DhcpOptionsListType) AddDhcpOption(value *DhcpOptionType) {
	obj.DhcpOption = append(obj.DhcpOption, *value)
}
