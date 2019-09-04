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
	firewall_rule_action_list int = iota
	firewall_rule_service
	firewall_rule_endpoint_1
	firewall_rule_endpoint_2
	firewall_rule_match_tags
	firewall_rule_match_tag_types
	firewall_rule_direction
	firewall_rule_id_perms
	firewall_rule_perms2
	firewall_rule_annotations
	firewall_rule_display_name
	firewall_rule_service_group_refs
	firewall_rule_address_group_refs
	firewall_rule_virtual_network_refs
	firewall_rule_security_logging_object_refs
	firewall_rule_tag_refs
	firewall_rule_firewall_policy_back_refs
)

type FirewallRule struct {
	contrail.ObjectBase
	action_list                  ActionListType
	service                      FirewallServiceType
	endpoint_1                   FirewallRuleEndpointType
	endpoint_2                   FirewallRuleEndpointType
	match_tags                   FirewallRuleMatchTagsType
	match_tag_types              FirewallRuleMatchTagsTypeIdList
	direction                    string
	id_perms                     IdPermsType
	perms2                       PermType2
	annotations                  KeyValuePairs
	display_name                 string
	service_group_refs           contrail.ReferenceList
	address_group_refs           contrail.ReferenceList
	virtual_network_refs         contrail.ReferenceList
	security_logging_object_refs contrail.ReferenceList
	tag_refs                     contrail.ReferenceList
	firewall_policy_back_refs    contrail.ReferenceList
	valid                        big.Int
	modified                     big.Int
	baseMap                      map[string]contrail.ReferenceList
}

func (obj *FirewallRule) GetType() string {
	return "firewall-rule"
}

func (obj *FirewallRule) GetDefaultParent() []string {
	name := []string{"default-policy-management"}
	return name
}

func (obj *FirewallRule) GetDefaultParentType() string {
	return "policy-management"
}

func (obj *FirewallRule) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *FirewallRule) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *FirewallRule) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *FirewallRule) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *FirewallRule) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *FirewallRule) GetActionList() ActionListType {
	return obj.action_list
}

func (obj *FirewallRule) SetActionList(value *ActionListType) {
	obj.action_list = *value
	obj.modified.SetBit(&obj.modified, firewall_rule_action_list, 1)
}

func (obj *FirewallRule) GetService() FirewallServiceType {
	return obj.service
}

func (obj *FirewallRule) SetService(value *FirewallServiceType) {
	obj.service = *value
	obj.modified.SetBit(&obj.modified, firewall_rule_service, 1)
}

func (obj *FirewallRule) GetEndpoint1() FirewallRuleEndpointType {
	return obj.endpoint_1
}

func (obj *FirewallRule) SetEndpoint1(value *FirewallRuleEndpointType) {
	obj.endpoint_1 = *value
	obj.modified.SetBit(&obj.modified, firewall_rule_endpoint_1, 1)
}

func (obj *FirewallRule) GetEndpoint2() FirewallRuleEndpointType {
	return obj.endpoint_2
}

func (obj *FirewallRule) SetEndpoint2(value *FirewallRuleEndpointType) {
	obj.endpoint_2 = *value
	obj.modified.SetBit(&obj.modified, firewall_rule_endpoint_2, 1)
}

func (obj *FirewallRule) GetMatchTags() FirewallRuleMatchTagsType {
	return obj.match_tags
}

func (obj *FirewallRule) SetMatchTags(value *FirewallRuleMatchTagsType) {
	obj.match_tags = *value
	obj.modified.SetBit(&obj.modified, firewall_rule_match_tags, 1)
}

func (obj *FirewallRule) GetMatchTagTypes() FirewallRuleMatchTagsTypeIdList {
	return obj.match_tag_types
}

func (obj *FirewallRule) SetMatchTagTypes(value *FirewallRuleMatchTagsTypeIdList) {
	obj.match_tag_types = *value
	obj.modified.SetBit(&obj.modified, firewall_rule_match_tag_types, 1)
}

func (obj *FirewallRule) GetDirection() string {
	return obj.direction
}

func (obj *FirewallRule) SetDirection(value string) {
	obj.direction = value
	obj.modified.SetBit(&obj.modified, firewall_rule_direction, 1)
}

func (obj *FirewallRule) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *FirewallRule) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, firewall_rule_id_perms, 1)
}

func (obj *FirewallRule) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *FirewallRule) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, firewall_rule_perms2, 1)
}

func (obj *FirewallRule) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *FirewallRule) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, firewall_rule_annotations, 1)
}

func (obj *FirewallRule) GetDisplayName() string {
	return obj.display_name
}

func (obj *FirewallRule) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, firewall_rule_display_name, 1)
}

func (obj *FirewallRule) readServiceGroupRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(firewall_rule_service_group_refs) == 0) {
		err := obj.GetField(obj, "service_group_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FirewallRule) GetServiceGroupRefs() (
	contrail.ReferenceList, error) {
	err := obj.readServiceGroupRefs()
	if err != nil {
		return nil, err
	}
	return obj.service_group_refs, nil
}

func (obj *FirewallRule) AddServiceGroup(
	rhs *ServiceGroup) error {
	err := obj.readServiceGroupRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_service_group_refs) == 0 {
		obj.storeReferenceBase("service-group", obj.service_group_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.service_group_refs = append(obj.service_group_refs, ref)
	obj.modified.SetBit(&obj.modified, firewall_rule_service_group_refs, 1)
	return nil
}

func (obj *FirewallRule) DeleteServiceGroup(uuid string) error {
	err := obj.readServiceGroupRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_service_group_refs) == 0 {
		obj.storeReferenceBase("service-group", obj.service_group_refs)
	}

	for i, ref := range obj.service_group_refs {
		if ref.Uuid == uuid {
			obj.service_group_refs = append(
				obj.service_group_refs[:i],
				obj.service_group_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, firewall_rule_service_group_refs, 1)
	return nil
}

func (obj *FirewallRule) ClearServiceGroup() {
	if (obj.valid.Bit(firewall_rule_service_group_refs) != 0) &&
		(obj.modified.Bit(firewall_rule_service_group_refs) == 0) {
		obj.storeReferenceBase("service-group", obj.service_group_refs)
	}
	obj.service_group_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, firewall_rule_service_group_refs, 1)
	obj.modified.SetBit(&obj.modified, firewall_rule_service_group_refs, 1)
}

func (obj *FirewallRule) SetServiceGroupList(
	refList []contrail.ReferencePair) {
	obj.ClearServiceGroup()
	obj.service_group_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.service_group_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *FirewallRule) readAddressGroupRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(firewall_rule_address_group_refs) == 0) {
		err := obj.GetField(obj, "address_group_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FirewallRule) GetAddressGroupRefs() (
	contrail.ReferenceList, error) {
	err := obj.readAddressGroupRefs()
	if err != nil {
		return nil, err
	}
	return obj.address_group_refs, nil
}

func (obj *FirewallRule) AddAddressGroup(
	rhs *AddressGroup) error {
	err := obj.readAddressGroupRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_address_group_refs) == 0 {
		obj.storeReferenceBase("address-group", obj.address_group_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.address_group_refs = append(obj.address_group_refs, ref)
	obj.modified.SetBit(&obj.modified, firewall_rule_address_group_refs, 1)
	return nil
}

func (obj *FirewallRule) DeleteAddressGroup(uuid string) error {
	err := obj.readAddressGroupRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_address_group_refs) == 0 {
		obj.storeReferenceBase("address-group", obj.address_group_refs)
	}

	for i, ref := range obj.address_group_refs {
		if ref.Uuid == uuid {
			obj.address_group_refs = append(
				obj.address_group_refs[:i],
				obj.address_group_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, firewall_rule_address_group_refs, 1)
	return nil
}

func (obj *FirewallRule) ClearAddressGroup() {
	if (obj.valid.Bit(firewall_rule_address_group_refs) != 0) &&
		(obj.modified.Bit(firewall_rule_address_group_refs) == 0) {
		obj.storeReferenceBase("address-group", obj.address_group_refs)
	}
	obj.address_group_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, firewall_rule_address_group_refs, 1)
	obj.modified.SetBit(&obj.modified, firewall_rule_address_group_refs, 1)
}

func (obj *FirewallRule) SetAddressGroupList(
	refList []contrail.ReferencePair) {
	obj.ClearAddressGroup()
	obj.address_group_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.address_group_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *FirewallRule) readVirtualNetworkRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(firewall_rule_virtual_network_refs) == 0) {
		err := obj.GetField(obj, "virtual_network_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FirewallRule) GetVirtualNetworkRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualNetworkRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_network_refs, nil
}

func (obj *FirewallRule) AddVirtualNetwork(
	rhs *VirtualNetwork) error {
	err := obj.readVirtualNetworkRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_virtual_network_refs) == 0 {
		obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.virtual_network_refs = append(obj.virtual_network_refs, ref)
	obj.modified.SetBit(&obj.modified, firewall_rule_virtual_network_refs, 1)
	return nil
}

func (obj *FirewallRule) DeleteVirtualNetwork(uuid string) error {
	err := obj.readVirtualNetworkRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_virtual_network_refs) == 0 {
		obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
	}

	for i, ref := range obj.virtual_network_refs {
		if ref.Uuid == uuid {
			obj.virtual_network_refs = append(
				obj.virtual_network_refs[:i],
				obj.virtual_network_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, firewall_rule_virtual_network_refs, 1)
	return nil
}

func (obj *FirewallRule) ClearVirtualNetwork() {
	if (obj.valid.Bit(firewall_rule_virtual_network_refs) != 0) &&
		(obj.modified.Bit(firewall_rule_virtual_network_refs) == 0) {
		obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
	}
	obj.virtual_network_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, firewall_rule_virtual_network_refs, 1)
	obj.modified.SetBit(&obj.modified, firewall_rule_virtual_network_refs, 1)
}

func (obj *FirewallRule) SetVirtualNetworkList(
	refList []contrail.ReferencePair) {
	obj.ClearVirtualNetwork()
	obj.virtual_network_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.virtual_network_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *FirewallRule) readSecurityLoggingObjectRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(firewall_rule_security_logging_object_refs) == 0) {
		err := obj.GetField(obj, "security_logging_object_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FirewallRule) GetSecurityLoggingObjectRefs() (
	contrail.ReferenceList, error) {
	err := obj.readSecurityLoggingObjectRefs()
	if err != nil {
		return nil, err
	}
	return obj.security_logging_object_refs, nil
}

func (obj *FirewallRule) AddSecurityLoggingObject(
	rhs *SecurityLoggingObject, data SloRateType) error {
	err := obj.readSecurityLoggingObjectRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_security_logging_object_refs) == 0 {
		obj.storeReferenceBase("security-logging-object", obj.security_logging_object_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
	obj.security_logging_object_refs = append(obj.security_logging_object_refs, ref)
	obj.modified.SetBit(&obj.modified, firewall_rule_security_logging_object_refs, 1)
	return nil
}

func (obj *FirewallRule) DeleteSecurityLoggingObject(uuid string) error {
	err := obj.readSecurityLoggingObjectRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_security_logging_object_refs) == 0 {
		obj.storeReferenceBase("security-logging-object", obj.security_logging_object_refs)
	}

	for i, ref := range obj.security_logging_object_refs {
		if ref.Uuid == uuid {
			obj.security_logging_object_refs = append(
				obj.security_logging_object_refs[:i],
				obj.security_logging_object_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, firewall_rule_security_logging_object_refs, 1)
	return nil
}

func (obj *FirewallRule) ClearSecurityLoggingObject() {
	if (obj.valid.Bit(firewall_rule_security_logging_object_refs) != 0) &&
		(obj.modified.Bit(firewall_rule_security_logging_object_refs) == 0) {
		obj.storeReferenceBase("security-logging-object", obj.security_logging_object_refs)
	}
	obj.security_logging_object_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, firewall_rule_security_logging_object_refs, 1)
	obj.modified.SetBit(&obj.modified, firewall_rule_security_logging_object_refs, 1)
}

func (obj *FirewallRule) SetSecurityLoggingObjectList(
	refList []contrail.ReferencePair) {
	obj.ClearSecurityLoggingObject()
	obj.security_logging_object_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.security_logging_object_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *FirewallRule) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(firewall_rule_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FirewallRule) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *FirewallRule) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, firewall_rule_tag_refs, 1)
	return nil
}

func (obj *FirewallRule) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(firewall_rule_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, firewall_rule_tag_refs, 1)
	return nil
}

func (obj *FirewallRule) ClearTag() {
	if (obj.valid.Bit(firewall_rule_tag_refs) != 0) &&
		(obj.modified.Bit(firewall_rule_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, firewall_rule_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, firewall_rule_tag_refs, 1)
}

func (obj *FirewallRule) SetTagList(
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

func (obj *FirewallRule) readFirewallPolicyBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(firewall_rule_firewall_policy_back_refs) == 0) {
		err := obj.GetField(obj, "firewall_policy_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FirewallRule) GetFirewallPolicyBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readFirewallPolicyBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.firewall_policy_back_refs, nil
}

func (obj *FirewallRule) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(firewall_rule_action_list) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.action_list)
		if err != nil {
			return nil, err
		}
		msg["action_list"] = &value
	}

	if obj.modified.Bit(firewall_rule_service) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service)
		if err != nil {
			return nil, err
		}
		msg["service"] = &value
	}

	if obj.modified.Bit(firewall_rule_endpoint_1) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.endpoint_1)
		if err != nil {
			return nil, err
		}
		msg["endpoint_1"] = &value
	}

	if obj.modified.Bit(firewall_rule_endpoint_2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.endpoint_2)
		if err != nil {
			return nil, err
		}
		msg["endpoint_2"] = &value
	}

	if obj.modified.Bit(firewall_rule_match_tags) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.match_tags)
		if err != nil {
			return nil, err
		}
		msg["match_tags"] = &value
	}

	if obj.modified.Bit(firewall_rule_match_tag_types) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.match_tag_types)
		if err != nil {
			return nil, err
		}
		msg["match_tag_types"] = &value
	}

	if obj.modified.Bit(firewall_rule_direction) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.direction)
		if err != nil {
			return nil, err
		}
		msg["direction"] = &value
	}

	if obj.modified.Bit(firewall_rule_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(firewall_rule_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(firewall_rule_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(firewall_rule_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.service_group_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_group_refs)
		if err != nil {
			return nil, err
		}
		msg["service_group_refs"] = &value
	}

	if len(obj.address_group_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.address_group_refs)
		if err != nil {
			return nil, err
		}
		msg["address_group_refs"] = &value
	}

	if len(obj.virtual_network_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.virtual_network_refs)
		if err != nil {
			return nil, err
		}
		msg["virtual_network_refs"] = &value
	}

	if len(obj.security_logging_object_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_logging_object_refs)
		if err != nil {
			return nil, err
		}
		msg["security_logging_object_refs"] = &value
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

func (obj *FirewallRule) UnmarshalJSON(body []byte) error {
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
		case "action_list":
			err = json.Unmarshal(value, &obj.action_list)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_action_list, 1)
			}
			break
		case "service":
			err = json.Unmarshal(value, &obj.service)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_service, 1)
			}
			break
		case "endpoint_1":
			err = json.Unmarshal(value, &obj.endpoint_1)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_endpoint_1, 1)
			}
			break
		case "endpoint_2":
			err = json.Unmarshal(value, &obj.endpoint_2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_endpoint_2, 1)
			}
			break
		case "match_tags":
			err = json.Unmarshal(value, &obj.match_tags)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_match_tags, 1)
			}
			break
		case "match_tag_types":
			err = json.Unmarshal(value, &obj.match_tag_types)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_match_tag_types, 1)
			}
			break
		case "direction":
			err = json.Unmarshal(value, &obj.direction)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_direction, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_display_name, 1)
			}
			break
		case "service_group_refs":
			err = json.Unmarshal(value, &obj.service_group_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_service_group_refs, 1)
			}
			break
		case "address_group_refs":
			err = json.Unmarshal(value, &obj.address_group_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_address_group_refs, 1)
			}
			break
		case "virtual_network_refs":
			err = json.Unmarshal(value, &obj.virtual_network_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_virtual_network_refs, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, firewall_rule_tag_refs, 1)
			}
			break
		case "security_logging_object_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr SloRateType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, firewall_rule_security_logging_object_refs, 1)
				obj.security_logging_object_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.security_logging_object_refs = append(obj.security_logging_object_refs, ref)
				}
				break
			}
		case "firewall_policy_back_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr FirewallSequence
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, firewall_rule_firewall_policy_back_refs, 1)
				obj.firewall_policy_back_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.firewall_policy_back_refs = append(obj.firewall_policy_back_refs, ref)
				}
				break
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FirewallRule) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(firewall_rule_action_list) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.action_list)
		if err != nil {
			return nil, err
		}
		msg["action_list"] = &value
	}

	if obj.modified.Bit(firewall_rule_service) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service)
		if err != nil {
			return nil, err
		}
		msg["service"] = &value
	}

	if obj.modified.Bit(firewall_rule_endpoint_1) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.endpoint_1)
		if err != nil {
			return nil, err
		}
		msg["endpoint_1"] = &value
	}

	if obj.modified.Bit(firewall_rule_endpoint_2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.endpoint_2)
		if err != nil {
			return nil, err
		}
		msg["endpoint_2"] = &value
	}

	if obj.modified.Bit(firewall_rule_match_tags) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.match_tags)
		if err != nil {
			return nil, err
		}
		msg["match_tags"] = &value
	}

	if obj.modified.Bit(firewall_rule_match_tag_types) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.match_tag_types)
		if err != nil {
			return nil, err
		}
		msg["match_tag_types"] = &value
	}

	if obj.modified.Bit(firewall_rule_direction) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.direction)
		if err != nil {
			return nil, err
		}
		msg["direction"] = &value
	}

	if obj.modified.Bit(firewall_rule_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(firewall_rule_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(firewall_rule_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(firewall_rule_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(firewall_rule_service_group_refs) != 0 {
		if len(obj.service_group_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["service_group_refs"] = &value
		} else if !obj.hasReferenceBase("service-group") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.service_group_refs)
			if err != nil {
				return nil, err
			}
			msg["service_group_refs"] = &value
		}
	}

	if obj.modified.Bit(firewall_rule_address_group_refs) != 0 {
		if len(obj.address_group_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["address_group_refs"] = &value
		} else if !obj.hasReferenceBase("address-group") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.address_group_refs)
			if err != nil {
				return nil, err
			}
			msg["address_group_refs"] = &value
		}
	}

	if obj.modified.Bit(firewall_rule_virtual_network_refs) != 0 {
		if len(obj.virtual_network_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["virtual_network_refs"] = &value
		} else if !obj.hasReferenceBase("virtual-network") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.virtual_network_refs)
			if err != nil {
				return nil, err
			}
			msg["virtual_network_refs"] = &value
		}
	}

	if obj.modified.Bit(firewall_rule_security_logging_object_refs) != 0 {
		if len(obj.security_logging_object_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["security_logging_object_refs"] = &value
		} else if !obj.hasReferenceBase("security-logging-object") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.security_logging_object_refs)
			if err != nil {
				return nil, err
			}
			msg["security_logging_object_refs"] = &value
		}
	}

	if obj.modified.Bit(firewall_rule_tag_refs) != 0 {
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

func (obj *FirewallRule) UpdateReferences() error {

	if (obj.modified.Bit(firewall_rule_service_group_refs) != 0) &&
		len(obj.service_group_refs) > 0 &&
		obj.hasReferenceBase("service-group") {
		err := obj.UpdateReference(
			obj, "service-group",
			obj.service_group_refs,
			obj.baseMap["service-group"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(firewall_rule_address_group_refs) != 0) &&
		len(obj.address_group_refs) > 0 &&
		obj.hasReferenceBase("address-group") {
		err := obj.UpdateReference(
			obj, "address-group",
			obj.address_group_refs,
			obj.baseMap["address-group"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(firewall_rule_virtual_network_refs) != 0) &&
		len(obj.virtual_network_refs) > 0 &&
		obj.hasReferenceBase("virtual-network") {
		err := obj.UpdateReference(
			obj, "virtual-network",
			obj.virtual_network_refs,
			obj.baseMap["virtual-network"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(firewall_rule_security_logging_object_refs) != 0) &&
		len(obj.security_logging_object_refs) > 0 &&
		obj.hasReferenceBase("security-logging-object") {
		err := obj.UpdateReference(
			obj, "security-logging-object",
			obj.security_logging_object_refs,
			obj.baseMap["security-logging-object"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(firewall_rule_tag_refs) != 0) &&
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

func FirewallRuleByName(c contrail.ApiClient, fqn string) (*FirewallRule, error) {
	obj, err := c.FindByName("firewall-rule", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*FirewallRule), nil
}

func FirewallRuleByUuid(c contrail.ApiClient, uuid string) (*FirewallRule, error) {
	obj, err := c.FindByUuid("firewall-rule", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*FirewallRule), nil
}
