//
// Automatically generated. DO NOT EDIT.
//

package resources

type FloatingIpPoolType struct {
	Subnet []SubnetType `json:"subnet,omitempty"`
}

func (obj *FloatingIpPoolType) AddSubnet(value *SubnetType) {
	obj.Subnet = append(obj.Subnet, *value)
}
