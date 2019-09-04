//
// Automatically generated. DO NOT EDIT.
//

package resources

import (
	"encoding/json"
	"math/big"

	"github.com/Juniper/contrail-go-api"
)

const (
	global_vrouter_config_ecmp_hashing_include_fields int = iota
	global_vrouter_config_linklocal_services
	global_vrouter_config_encapsulation_priorities
	global_vrouter_config_vxlan_network_identifier_mode
	global_vrouter_config_flow_export_rate
	global_vrouter_config_flow_aging_timeout_list
	global_vrouter_config_enable_security_logging
	global_vrouter_config_forwarding_mode
	global_vrouter_config_id_perms
	global_vrouter_config_perms2
	global_vrouter_config_annotations
	global_vrouter_config_display_name
	global_vrouter_config_security_logging_objects
	global_vrouter_config_tag_refs
	global_vrouter_config_application_policy_set_back_refs
)

type GlobalVrouterConfig struct {
	contrail.ObjectBase
	ecmp_hashing_include_fields      EcmpHashingIncludeFields
	linklocal_services               LinklocalServicesTypes
	encapsulation_priorities         EncapsulationPrioritiesType
	vxlan_network_identifier_mode    string
	flow_export_rate                 int
	flow_aging_timeout_list          FlowAgingTimeoutList
	enable_security_logging          bool
	forwarding_mode                  string
	id_perms                         IdPermsType
	perms2                           PermType2
	annotations                      KeyValuePairs
	display_name                     string
	security_logging_objects         contrail.ReferenceList
	tag_refs                         contrail.ReferenceList
	application_policy_set_back_refs contrail.ReferenceList
	valid                            big.Int
	modified                         big.Int
	baseMap                          map[string]contrail.ReferenceList
}

func (obj *GlobalVrouterConfig) GetType() string {
	return "global-vrouter-config"
}

func (obj *GlobalVrouterConfig) GetDefaultParent() []string {
	name := []string{"default-global-system-config"}
	return name
}

func (obj *GlobalVrouterConfig) GetDefaultParentType() string {
	return "global-system-config"
}

func (obj *GlobalVrouterConfig) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *GlobalVrouterConfig) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *GlobalVrouterConfig) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *GlobalVrouterConfig) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *GlobalVrouterConfig) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *GlobalVrouterConfig) GetEcmpHashingIncludeFields() EcmpHashingIncludeFields {
	return obj.ecmp_hashing_include_fields
}

func (obj *GlobalVrouterConfig) SetEcmpHashingIncludeFields(value *EcmpHashingIncludeFields) {
	obj.ecmp_hashing_include_fields = *value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_ecmp_hashing_include_fields, 1)
}

func (obj *GlobalVrouterConfig) GetLinklocalServices() LinklocalServicesTypes {
	return obj.linklocal_services
}

func (obj *GlobalVrouterConfig) SetLinklocalServices(value *LinklocalServicesTypes) {
	obj.linklocal_services = *value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_linklocal_services, 1)
}

func (obj *GlobalVrouterConfig) GetEncapsulationPriorities() EncapsulationPrioritiesType {
	return obj.encapsulation_priorities
}

func (obj *GlobalVrouterConfig) SetEncapsulationPriorities(value *EncapsulationPrioritiesType) {
	obj.encapsulation_priorities = *value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_encapsulation_priorities, 1)
}

func (obj *GlobalVrouterConfig) GetVxlanNetworkIdentifierMode() string {
	return obj.vxlan_network_identifier_mode
}

func (obj *GlobalVrouterConfig) SetVxlanNetworkIdentifierMode(value string) {
	obj.vxlan_network_identifier_mode = value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_vxlan_network_identifier_mode, 1)
}

func (obj *GlobalVrouterConfig) GetFlowExportRate() int {
	return obj.flow_export_rate
}

func (obj *GlobalVrouterConfig) SetFlowExportRate(value int) {
	obj.flow_export_rate = value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_flow_export_rate, 1)
}

func (obj *GlobalVrouterConfig) GetFlowAgingTimeoutList() FlowAgingTimeoutList {
	return obj.flow_aging_timeout_list
}

func (obj *GlobalVrouterConfig) SetFlowAgingTimeoutList(value *FlowAgingTimeoutList) {
	obj.flow_aging_timeout_list = *value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_flow_aging_timeout_list, 1)
}

func (obj *GlobalVrouterConfig) GetEnableSecurityLogging() bool {
	return obj.enable_security_logging
}

func (obj *GlobalVrouterConfig) SetEnableSecurityLogging(value bool) {
	obj.enable_security_logging = value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_enable_security_logging, 1)
}

func (obj *GlobalVrouterConfig) GetForwardingMode() string {
	return obj.forwarding_mode
}

func (obj *GlobalVrouterConfig) SetForwardingMode(value string) {
	obj.forwarding_mode = value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_forwarding_mode, 1)
}

func (obj *GlobalVrouterConfig) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *GlobalVrouterConfig) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_id_perms, 1)
}

func (obj *GlobalVrouterConfig) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *GlobalVrouterConfig) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_perms2, 1)
}

func (obj *GlobalVrouterConfig) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *GlobalVrouterConfig) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_annotations, 1)
}

func (obj *GlobalVrouterConfig) GetDisplayName() string {
	return obj.display_name
}

func (obj *GlobalVrouterConfig) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, global_vrouter_config_display_name, 1)
}

func (obj *GlobalVrouterConfig) readSecurityLoggingObjects() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(global_vrouter_config_security_logging_objects) == 0) {
		err := obj.GetField(obj, "security_logging_objects")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *GlobalVrouterConfig) GetSecurityLoggingObjects() (
	contrail.ReferenceList, error) {
	err := obj.readSecurityLoggingObjects()
	if err != nil {
		return nil, err
	}
	return obj.security_logging_objects, nil
}

func (obj *GlobalVrouterConfig) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(global_vrouter_config_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *GlobalVrouterConfig) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *GlobalVrouterConfig) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(global_vrouter_config_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, global_vrouter_config_tag_refs, 1)
	return nil
}

func (obj *GlobalVrouterConfig) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(global_vrouter_config_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	for i, ref := range obj.tag_refs {
		if ref.Uuid == uuid {
			obj.tag_refs = append(
				obj.tag_refs[:i],
				obj.tag_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, global_vrouter_config_tag_refs, 1)
	return nil
}

func (obj *GlobalVrouterConfig) ClearTag() {
	if (obj.valid.Bit(global_vrouter_config_tag_refs) != 0) &&
		(obj.modified.Bit(global_vrouter_config_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, global_vrouter_config_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, global_vrouter_config_tag_refs, 1)
}

func (obj *GlobalVrouterConfig) SetTagList(
	refList []contrail.ReferencePair) {
	obj.ClearTag()
	obj.tag_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.tag_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *GlobalVrouterConfig) readApplicationPolicySetBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(global_vrouter_config_application_policy_set_back_refs) == 0) {
		err := obj.GetField(obj, "application_policy_set_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *GlobalVrouterConfig) GetApplicationPolicySetBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readApplicationPolicySetBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.application_policy_set_back_refs, nil
}

func (obj *GlobalVrouterConfig) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(global_vrouter_config_ecmp_hashing_include_fields) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.ecmp_hashing_include_fields)
		if err != nil {
			return nil, err
		}
		msg["ecmp_hashing_include_fields"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_linklocal_services) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.linklocal_services)
		if err != nil {
			return nil, err
		}
		msg["linklocal_services"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_encapsulation_priorities) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.encapsulation_priorities)
		if err != nil {
			return nil, err
		}
		msg["encapsulation_priorities"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_vxlan_network_identifier_mode) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.vxlan_network_identifier_mode)
		if err != nil {
			return nil, err
		}
		msg["vxlan_network_identifier_mode"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_flow_export_rate) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.flow_export_rate)
		if err != nil {
			return nil, err
		}
		msg["flow_export_rate"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_flow_aging_timeout_list) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.flow_aging_timeout_list)
		if err != nil {
			return nil, err
		}
		msg["flow_aging_timeout_list"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_enable_security_logging) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.enable_security_logging)
		if err != nil {
			return nil, err
		}
		msg["enable_security_logging"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_forwarding_mode) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.forwarding_mode)
		if err != nil {
			return nil, err
		}
		msg["forwarding_mode"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.tag_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.tag_refs)
		if err != nil {
			return nil, err
		}
		msg["tag_refs"] = &value
	}

	return json.Marshal(msg)
}

func (obj *GlobalVrouterConfig) UnmarshalJSON(body []byte) error {
	var m map[string]json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	err = obj.UnmarshalCommon(m)
	if err != nil {
		return err
	}
	for key, value := range m {
		switch key {
		case "ecmp_hashing_include_fields":
			err = json.Unmarshal(value, &obj.ecmp_hashing_include_fields)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_ecmp_hashing_include_fields, 1)
			}
			break
		case "linklocal_services":
			err = json.Unmarshal(value, &obj.linklocal_services)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_linklocal_services, 1)
			}
			break
		case "encapsulation_priorities":
			err = json.Unmarshal(value, &obj.encapsulation_priorities)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_encapsulation_priorities, 1)
			}
			break
		case "vxlan_network_identifier_mode":
			err = json.Unmarshal(value, &obj.vxlan_network_identifier_mode)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_vxlan_network_identifier_mode, 1)
			}
			break
		case "flow_export_rate":
			err = json.Unmarshal(value, &obj.flow_export_rate)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_flow_export_rate, 1)
			}
			break
		case "flow_aging_timeout_list":
			err = json.Unmarshal(value, &obj.flow_aging_timeout_list)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_flow_aging_timeout_list, 1)
			}
			break
		case "enable_security_logging":
			err = json.Unmarshal(value, &obj.enable_security_logging)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_enable_security_logging, 1)
			}
			break
		case "forwarding_mode":
			err = json.Unmarshal(value, &obj.forwarding_mode)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_forwarding_mode, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_display_name, 1)
			}
			break
		case "security_logging_objects":
			err = json.Unmarshal(value, &obj.security_logging_objects)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_security_logging_objects, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_tag_refs, 1)
			}
			break
		case "application_policy_set_back_refs":
			err = json.Unmarshal(value, &obj.application_policy_set_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, global_vrouter_config_application_policy_set_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *GlobalVrouterConfig) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(global_vrouter_config_ecmp_hashing_include_fields) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.ecmp_hashing_include_fields)
		if err != nil {
			return nil, err
		}
		msg["ecmp_hashing_include_fields"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_linklocal_services) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.linklocal_services)
		if err != nil {
			return nil, err
		}
		msg["linklocal_services"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_encapsulation_priorities) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.encapsulation_priorities)
		if err != nil {
			return nil, err
		}
		msg["encapsulation_priorities"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_vxlan_network_identifier_mode) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.vxlan_network_identifier_mode)
		if err != nil {
			return nil, err
		}
		msg["vxlan_network_identifier_mode"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_flow_export_rate) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.flow_export_rate)
		if err != nil {
			return nil, err
		}
		msg["flow_export_rate"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_flow_aging_timeout_list) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.flow_aging_timeout_list)
		if err != nil {
			return nil, err
		}
		msg["flow_aging_timeout_list"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_enable_security_logging) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.enable_security_logging)
		if err != nil {
			return nil, err
		}
		msg["enable_security_logging"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_forwarding_mode) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.forwarding_mode)
		if err != nil {
			return nil, err
		}
		msg["forwarding_mode"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(global_vrouter_config_tag_refs) != 0 {
		if len(obj.tag_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["tag_refs"] = &value
		} else if !obj.hasReferenceBase("tag") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.tag_refs)
			if err != nil {
				return nil, err
			}
			msg["tag_refs"] = &value
		}
	}

	return json.Marshal(msg)
}

func (obj *GlobalVrouterConfig) UpdateReferences() error {

	if (obj.modified.Bit(global_vrouter_config_tag_refs) != 0) &&
		len(obj.tag_refs) > 0 &&
		obj.hasReferenceBase("tag") {
		err := obj.UpdateReference(
			obj, "tag",
			obj.tag_refs,
			obj.baseMap["tag"])
		if err != nil {
			return err
		}
	}

	return nil
}

func GlobalVrouterConfigByName(c contrail.ApiClient, fqn string) (*GlobalVrouterConfig, error) {
	obj, err := c.FindByName("global-vrouter-config", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*GlobalVrouterConfig), nil
}

func GlobalVrouterConfigByUuid(c contrail.ApiClient, uuid string) (*GlobalVrouterConfig, error) {
	obj, err := c.FindByUuid("global-vrouter-config", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*GlobalVrouterConfig), nil
}
