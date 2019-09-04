---
Title: "virtual_machine_interface_refs"
type: "md"
weight: 3
Description: " "
---

## Example Usage

### Basic Creation

```hcl
resource "contrail_virtual_machine_interface" "port_test1" {
 	 name  = "port_test1"
	 display_name = "test_display_name1"
  	 parent_uuid = "${var.vm_id}"
     virtual_network_refs {
    	to = contrail_virtual_network.network_test.id     # <-- required field
     }
} 

resource "contrail_virtual_network" "network_test" {
  name = "test_name"
  display_name = "test_display_name" 
  parent_uuid = "${var.project_id}"
}

resource "contrail_virtual_machine_interface_refs" "health_check_refs_test" {
   uuid = contrail_virtual_machine_interface.port_test1.id
   service_health_check_refs  {
       to = "${var.health_check_id}"
   }
}
```

For more information about structure of objects see [documentation](http://www.opencontrail.org/documentation/api/r4.1/contrail_openapi.html)
## Argument Reference

The following arguments are supported:

* `security_logging_object_refs` - *optional* List of security-logging-object references
* `qos_config_refs` - *optional*	List of qos-config references	
* `security_group_refs` - *optional*	List of security-group references	
* `virtual_machine_interface_refs` - *optional*	List of virtual-machine-interface references	
* `virtual_machine_refs` - *optional*	List of virtual-machine references	
* `routing_instance_refs` - *optional*	List of routing-instance references
* `port_tuple_refs` - *optional*	List of port-tuple references
* `service_health_check_refs` - *optional*	List of service-health-check references
* `interface_route_table_refs` - *optional*	List of interface-route-table references
* `physical_interface_refs` - *optional*	List of physical-interface references
* `bridge_domain_refs` - *optional*	List of bridge-domain references
* `service_endpoint_refs` - *optional*	List of service-endpoint references	
* `tag_refs` - *optional*	List of tag references
