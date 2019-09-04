---
Title: "contrail_tag"
type: "md"
weight: 3
Description: " "
---

## Example Usage

### Basic Creation

#### Example with tags
```hcl
resource "contrail_tag" "tag_test" {
    name = "test_tag"
    tag_value = "value"
    tag_type_name = "type"
	display_name = "type=value"
}
```

For more information about structure of objects see [documentation](http://www.opencontrail.org/documentation/api/r4.1/contrail_openapi.html)
## Argument Reference

The following arguments are supported:

* `fq_name` - *required*	Fully Qualified Name of resource
* `parent_type` - *required*	Parent resource type
* `tag_type_name` - *required*	Tag type string representation
* `tag_value` - *required*	Tag value string representation	
* `annotations` - *optional*	Dictionary of arbitrary (key, value) on a resource.	
* `display_name` - *optional*	Display name user configured string(name) that can be updated any time. 
