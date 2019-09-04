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
	policy_management_id_perms int = iota
	policy_management_perms2
	policy_management_annotations
	policy_management_display_name
	policy_management_service_groups
	policy_management_address_groups
	policy_management_firewall_rules
	policy_management_firewall_policys
	policy_management_application_policy_sets
	policy_management_tag_refs
)

type PolicyManagement struct {
	contrail.ObjectBase
	id_perms                IdPermsType
	perms2                  PermType2
	annotations             KeyValuePairs
	display_name            string
	service_groups          contrail.ReferenceList
	address_groups          contrail.ReferenceList
	firewall_rules          contrail.ReferenceList
	firewall_policys        contrail.ReferenceList
	application_policy_sets contrail.ReferenceList
	tag_refs                contrail.ReferenceList
	valid                   big.Int
	modified                big.Int
	baseMap                 map[string]contrail.ReferenceList
}

func (obj *PolicyManagement) GetType() string {
	return "policy-management"
}

func (obj *PolicyManagement) GetDefaultParent() []string {
	name := []string{}
	return name
}

func (obj *PolicyManagement) GetDefaultParentType() string {
	return ""
}

func (obj *PolicyManagement) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *PolicyManagement) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *PolicyManagement) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *PolicyManagement) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *PolicyManagement) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *PolicyManagement) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *PolicyManagement) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, policy_management_id_perms, 1)
}

func (obj *PolicyManagement) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *PolicyManagement) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, policy_management_perms2, 1)
}

func (obj *PolicyManagement) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *PolicyManagement) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, policy_management_annotations, 1)
}

func (obj *PolicyManagement) GetDisplayName() string {
	return obj.display_name
}

func (obj *PolicyManagement) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, policy_management_display_name, 1)
}

func (obj *PolicyManagement) readServiceGroups() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(policy_management_service_groups) == 0) {
		err := obj.GetField(obj, "service_groups")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *PolicyManagement) GetServiceGroups() (
	contrail.ReferenceList, error) {
	err := obj.readServiceGroups()
	if err != nil {
		return nil, err
	}
	return obj.service_groups, nil
}

func (obj *PolicyManagement) readAddressGroups() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(policy_management_address_groups) == 0) {
		err := obj.GetField(obj, "address_groups")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *PolicyManagement) GetAddressGroups() (
	contrail.ReferenceList, error) {
	err := obj.readAddressGroups()
	if err != nil {
		return nil, err
	}
	return obj.address_groups, nil
}

func (obj *PolicyManagement) readFirewallRules() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(policy_management_firewall_rules) == 0) {
		err := obj.GetField(obj, "firewall_rules")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *PolicyManagement) GetFirewallRules() (
	contrail.ReferenceList, error) {
	err := obj.readFirewallRules()
	if err != nil {
		return nil, err
	}
	return obj.firewall_rules, nil
}

func (obj *PolicyManagement) readFirewallPolicys() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(policy_management_firewall_policys) == 0) {
		err := obj.GetField(obj, "firewall_policys")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *PolicyManagement) GetFirewallPolicys() (
	contrail.ReferenceList, error) {
	err := obj.readFirewallPolicys()
	if err != nil {
		return nil, err
	}
	return obj.firewall_policys, nil
}

func (obj *PolicyManagement) readApplicationPolicySets() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(policy_management_application_policy_sets) == 0) {
		err := obj.GetField(obj, "application_policy_sets")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *PolicyManagement) GetApplicationPolicySets() (
	contrail.ReferenceList, error) {
	err := obj.readApplicationPolicySets()
	if err != nil {
		return nil, err
	}
	return obj.application_policy_sets, nil
}

func (obj *PolicyManagement) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(policy_management_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *PolicyManagement) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *PolicyManagement) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(policy_management_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, policy_management_tag_refs, 1)
	return nil
}

func (obj *PolicyManagement) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(policy_management_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, policy_management_tag_refs, 1)
	return nil
}

func (obj *PolicyManagement) ClearTag() {
	if (obj.valid.Bit(policy_management_tag_refs) != 0) &&
		(obj.modified.Bit(policy_management_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, policy_management_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, policy_management_tag_refs, 1)
}

func (obj *PolicyManagement) SetTagList(
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

func (obj *PolicyManagement) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(policy_management_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(policy_management_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(policy_management_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(policy_management_display_name) != 0 {
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

func (obj *PolicyManagement) UnmarshalJSON(body []byte) error {
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
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_display_name, 1)
			}
			break
		case "service_groups":
			err = json.Unmarshal(value, &obj.service_groups)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_service_groups, 1)
			}
			break
		case "address_groups":
			err = json.Unmarshal(value, &obj.address_groups)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_address_groups, 1)
			}
			break
		case "firewall_rules":
			err = json.Unmarshal(value, &obj.firewall_rules)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_firewall_rules, 1)
			}
			break
		case "firewall_policys":
			err = json.Unmarshal(value, &obj.firewall_policys)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_firewall_policys, 1)
			}
			break
		case "application_policy_sets":
			err = json.Unmarshal(value, &obj.application_policy_sets)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_application_policy_sets, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, policy_management_tag_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *PolicyManagement) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(policy_management_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(policy_management_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(policy_management_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(policy_management_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(policy_management_tag_refs) != 0 {
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

func (obj *PolicyManagement) UpdateReferences() error {

	if (obj.modified.Bit(policy_management_tag_refs) != 0) &&
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

func PolicyManagementByName(c contrail.ApiClient, fqn string) (*PolicyManagement, error) {
	obj, err := c.FindByName("policy-management", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*PolicyManagement), nil
}

func PolicyManagementByUuid(c contrail.ApiClient, uuid string) (*PolicyManagement, error) {
	obj, err := c.FindByUuid("policy-management", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*PolicyManagement), nil
}
