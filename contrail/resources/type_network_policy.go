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
	network_policy_network_policy_entries int = iota
	network_policy_id_perms
	network_policy_perms2
	network_policy_annotations
	network_policy_display_name
	network_policy_tag_refs
	network_policy_security_logging_object_back_refs
	network_policy_virtual_network_back_refs
)

type NetworkPolicy struct {
	contrail.ObjectBase
	network_policy_entries            PolicyEntriesType
	id_perms                          IdPermsType
	perms2                            PermType2
	annotations                       KeyValuePairs
	display_name                      string
	tag_refs                          contrail.ReferenceList
	security_logging_object_back_refs contrail.ReferenceList
	virtual_network_back_refs         contrail.ReferenceList
	valid                             big.Int
	modified                          big.Int
	baseMap                           map[string]contrail.ReferenceList
}

func (obj *NetworkPolicy) GetType() string {
	return "network-policy"
}

func (obj *NetworkPolicy) GetDefaultParent() []string {
	name := []string{"default-domain", "default-project"}
	return name
}

func (obj *NetworkPolicy) GetDefaultParentType() string {
	return "project"
}

func (obj *NetworkPolicy) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *NetworkPolicy) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *NetworkPolicy) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *NetworkPolicy) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *NetworkPolicy) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *NetworkPolicy) GetNetworkPolicyEntries() PolicyEntriesType {
	return obj.network_policy_entries
}

func (obj *NetworkPolicy) SetNetworkPolicyEntries(value *PolicyEntriesType) {
	obj.network_policy_entries = *value
	obj.modified.SetBit(&obj.modified, network_policy_network_policy_entries, 1)
}

func (obj *NetworkPolicy) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *NetworkPolicy) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, network_policy_id_perms, 1)
}

func (obj *NetworkPolicy) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *NetworkPolicy) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, network_policy_perms2, 1)
}

func (obj *NetworkPolicy) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *NetworkPolicy) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, network_policy_annotations, 1)
}

func (obj *NetworkPolicy) GetDisplayName() string {
	return obj.display_name
}

func (obj *NetworkPolicy) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, network_policy_display_name, 1)
}

func (obj *NetworkPolicy) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(network_policy_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *NetworkPolicy) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *NetworkPolicy) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(network_policy_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, network_policy_tag_refs, 1)
	return nil
}

func (obj *NetworkPolicy) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(network_policy_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, network_policy_tag_refs, 1)
	return nil
}

func (obj *NetworkPolicy) ClearTag() {
	if (obj.valid.Bit(network_policy_tag_refs) != 0) &&
		(obj.modified.Bit(network_policy_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, network_policy_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, network_policy_tag_refs, 1)
}

func (obj *NetworkPolicy) SetTagList(
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

func (obj *NetworkPolicy) readSecurityLoggingObjectBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(network_policy_security_logging_object_back_refs) == 0) {
		err := obj.GetField(obj, "security_logging_object_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *NetworkPolicy) GetSecurityLoggingObjectBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readSecurityLoggingObjectBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.security_logging_object_back_refs, nil
}

func (obj *NetworkPolicy) readVirtualNetworkBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(network_policy_virtual_network_back_refs) == 0) {
		err := obj.GetField(obj, "virtual_network_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *NetworkPolicy) GetVirtualNetworkBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualNetworkBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_network_back_refs, nil
}

func (obj *NetworkPolicy) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(network_policy_network_policy_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.network_policy_entries)
		if err != nil {
			return nil, err
		}
		msg["network_policy_entries"] = &value
	}

	if obj.modified.Bit(network_policy_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(network_policy_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(network_policy_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(network_policy_display_name) != 0 {
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

func (obj *NetworkPolicy) UnmarshalJSON(body []byte) error {
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
		case "network_policy_entries":
			err = json.Unmarshal(value, &obj.network_policy_entries)
			if err == nil {
				obj.valid.SetBit(&obj.valid, network_policy_network_policy_entries, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, network_policy_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, network_policy_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, network_policy_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, network_policy_display_name, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, network_policy_tag_refs, 1)
			}
			break
		case "security_logging_object_back_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr SecurityLoggingObjectRuleListType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, network_policy_security_logging_object_back_refs, 1)
				obj.security_logging_object_back_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.security_logging_object_back_refs = append(obj.security_logging_object_back_refs, ref)
				}
				break
			}
		case "virtual_network_back_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr VirtualNetworkPolicyType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, network_policy_virtual_network_back_refs, 1)
				obj.virtual_network_back_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.virtual_network_back_refs = append(obj.virtual_network_back_refs, ref)
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

func (obj *NetworkPolicy) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(network_policy_network_policy_entries) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.network_policy_entries)
		if err != nil {
			return nil, err
		}
		msg["network_policy_entries"] = &value
	}

	if obj.modified.Bit(network_policy_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(network_policy_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(network_policy_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(network_policy_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(network_policy_tag_refs) != 0 {
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

func (obj *NetworkPolicy) UpdateReferences() error {

	if (obj.modified.Bit(network_policy_tag_refs) != 0) &&
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

func NetworkPolicyByName(c contrail.ApiClient, fqn string) (*NetworkPolicy, error) {
	obj, err := c.FindByName("network-policy", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*NetworkPolicy), nil
}

func NetworkPolicyByUuid(c contrail.ApiClient, uuid string) (*NetworkPolicy, error) {
	obj, err := c.FindByUuid("network-policy", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*NetworkPolicy), nil
}
