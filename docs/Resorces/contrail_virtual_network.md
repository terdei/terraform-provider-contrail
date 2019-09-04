---
Title: "contrail_virtual_network"
type: "md"
weight: 3
Description: " "
---

## Example Usage

### Basic Creation

```hcl
resource "contrail_virtual_network" "network_test" {
  name = "test_name"
  display_name = "test_display_name" 
  parent_uuid = "${var.project_id}"
}
```

For more information about structure of objects see [documentation](http://www.opencontrail.org/documentation/api/r4.1/contrail_openapi.html)
## Argument Reference

The following arguments are supported:

* `fq_name` - *required*	Fully Qualified Name of resource.
* `parent_type` - *required*	Parent resource type.
* `ecmp_hashing_include_fields` - *optional*	ECMP hashing config at global level.
* `virtual_network_properties` - *optional*	Virtual network miscellaneous configurations.
* `provider_properties` - *optional*	Virtual network is provider network.
* `virtual_network_network_id` - *optional*	System assigned unique 32 bit ID for every virtual network.	
* `port_security_enabled` - *optional*	Port security status on the network.
* `route_target_list` - *optional*	List of route targets that are used as both import and export for this virtual network.
* `import_route_target_list` - *optional*	List of route targets that are used as import for this virtual network.
* `export_route_target_list` - *optional*	List of route targets that are used as export for this virtual network.
* `router_external` - *optional*	When true, this virtual network is openstack router external network.
* `is_shared` - *optional*	When true, this virtual network is shared with all tenants.
* `external_ipam` - *optional*	IP address assignment to VM is done statically, outside of (external to) Contrail Ipam. vCenter only feature.
* `flood_unknown_unicast` - *optional*	When true, packets with unknown unicast MAC address are flooded within the network. Default they are dropped.
* `multi_policy_service_chains_enabled` - *optional*	Allow multiple service chains within same two networks based on network policy. Current limitation is that both networks must reside within cluster, except when right most service is NAT.
* `address_allocation_mode` - *optional*	Address allocation mode for virtual network.
* `mac_learning_enabled` - *optional*	Enable MAC learning on the network.
* `mac_limit_control` - *optional*	MAC limit control on the network.
* `mac_move_control` - *optional*	MAC move control on the network.
* `mac_aging_time` - *optional*	MAC aging time on the network.
* `pbb_evpn_enable` - *optional*	Enable/Disable PBB EVPN tunneling on the network.
* `pbb_etree_enable` - *optional*	Enable/Disable PBB ETREE mode on the network.
* `layer2_control_word` - *optional*	Enable/Disable adding control word to the Layer 2 encapsulation.
* `annotations` - *optional*	Dictionary of arbitrary (key, value) on a resource.	
* `display_name` - *optional*	Display name user configured string(name) that can be updated any time. Used as openstack name.
