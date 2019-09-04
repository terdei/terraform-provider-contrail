//
// Automatically generated. DO NOT EDIT.
//

package resources

type FirewallServiceType struct {
	Protocol   string    `json:"protocol,omitempty"`
	ProtocolId int       `json:"protocol_id,omitempty"`
	SrcPorts   *PortType `json:"src_ports,omitempty"`
	DstPorts   *PortType `json:"dst_ports,omitempty"`
}
