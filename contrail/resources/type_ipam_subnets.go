//
// Automatically generated. DO NOT EDIT.
//

package resources

type IpamSubnets struct {
	Subnets []IpamSubnetType `json:"subnets,omitempty"`
}

func (obj *IpamSubnets) AddSubnets(value *IpamSubnetType) {
	obj.Subnets = append(obj.Subnets, *value)
}
