//
// Automatically generated. DO NOT EDIT.
//

package resources

type FirewallServiceGroupType struct {
	FirewallService []FirewallServiceType `json:"firewall_service,omitempty"`
}

func (obj *FirewallServiceGroupType) AddFirewallService(value *FirewallServiceType) {
	obj.FirewallService = append(obj.FirewallService, *value)
}
