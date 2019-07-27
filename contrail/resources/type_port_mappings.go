//
// Automatically generated. DO NOT EDIT.
//

package resources

type PortMap struct {
	SrcPort int `json:"src_port,omitempty"`
	DstPort int `json:"dst_port,omitempty"`
}

type PortMappings struct {
	PortMappings []PortMap `json:"port_mappings,omitempty"`
}

func (obj *PortMappings) AddPortMappings(value *PortMap) {
	obj.PortMappings = append(obj.PortMappings, *value)
}
