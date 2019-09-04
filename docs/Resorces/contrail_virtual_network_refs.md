---
Title: "contrail_virtual_network_refs"
type: "md"
weight: 3
Description: " "
---

## Example Usage

### Basic Creation

#### Example with tags
```hcl
resource "contrail_virtual_network" "network_test" {
  name = "test_name"
  display_name = "test_display_name" 
  parent_uuid = "${var.project_id}"
}

resource "contrail_virtual_network_refs" "network_ref" {
	uuid = contrail_virtual_network.network_test.id
	tag_refs {
	  to = contrail_tag.tag_test.id
	}
	depends_on = [ contrail_tag.tag_test ]
}

resource "contrail_tag" "tag_test" {
    name = "test_tag"
    tag_value = "value"
    tag_type_name = "type"
	display_name = "type=value"
}
```

#### Example with ipam

```hcl

resource "contrail_virtual_network" "spocknet" {
	name = "spocknet"
	display_name = "spocknet"\
}

resource "contrail_virtual_network_refs" "spockrefs" {
	uuid = "${contrail_virtual_network.spocknet.id}"
	network_ipam_refs {
		to = "${contrail_network_ipam.spock_ipam.id}"
		attr {
			ipam_subnets {
				subnet_name = "spock_subnet"
				subnet {
					ip_prefix = "10.10.20.0"
					ip_prefix_len = 24
				}
			}
		}
	}
}

resource "contrail_network_ipam" "spock_ipam" {
	name = "spock_ipam"
	#dns_nameservers = ["1.1.1.1", "2.2.2.2", "3.3.3.3", "8.8.8.8"]
}
```

For more information about structure of objects see [documentation](http://www.opencontrail.org/documentation/api/r4.1/contrail_openapi.html)
## Argument Reference

The following arguments are supported:

* `security_logging_object_refs` - *optional*	List of security-logging-object references.
* `qos_config_refs` - *optional*	List of qos-config references.
* `network_ipam_refs` - *required*	List of network-ipam references.
* `network_policy_refs` - *optional*	List of network-policy references.
* `virtual_network_refs` - *optional*	List of virtual-network references.
* `route_table_refs` - *optional*	List of route-table references.
* `bgpvpn_refs` - *optional*	List of bgpvpn references.
* `tag_refs` - *optional*	List of tag references.
