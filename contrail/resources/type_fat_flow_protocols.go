//
// Automatically generated. DO NOT EDIT.
//

package resources

type ProtocolType struct {
	Protocol string `json:"protocol,omitempty"`
	Port     int    `json:"port,omitempty"`
}

type FatFlowProtocols struct {
	FatFlowProtocol []ProtocolType `json:"fat_flow_protocol,omitempty"`
}

func (obj *FatFlowProtocols) AddFatFlowProtocol(value *ProtocolType) {
	obj.FatFlowProtocol = append(obj.FatFlowProtocol, *value)
}
