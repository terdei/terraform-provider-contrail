//
// Automatically generated. DO NOT EDIT.
//

package resources

type SubnetListType struct {
	Subnet []SubnetType `json:"subnet,omitempty"`
}

func (obj *SubnetListType) AddSubnet(value *SubnetType) {
	obj.Subnet = append(obj.Subnet, *value)
}
