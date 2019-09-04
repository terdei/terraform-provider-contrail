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
	application_policy_set_all_applications int = iota
	application_policy_set_id_perms
	application_policy_set_perms2
	application_policy_set_annotations
	application_policy_set_display_name
	application_policy_set_firewall_policy_refs
	application_policy_set_global_vrouter_config_refs
	application_policy_set_tag_refs
	application_policy_set_project_back_refs
)

type ApplicationPolicySet struct {
	contrail.ObjectBase
	all_applications           bool
	id_perms                   IdPermsType
	perms2                     PermType2
	annotations                KeyValuePairs
	display_name               string
	firewall_policy_refs       contrail.ReferenceList
	global_vrouter_config_refs contrail.ReferenceList
	tag_refs                   contrail.ReferenceList
	project_back_refs          contrail.ReferenceList
	valid                      big.Int
	modified                   big.Int
	baseMap                    map[string]contrail.ReferenceList
}

func (obj *ApplicationPolicySet) GetType() string {
	return "application-policy-set"
}

func (obj *ApplicationPolicySet) GetDefaultParent() []string {
	name := []string{"default-policy-management"}
	return name
}

func (obj *ApplicationPolicySet) GetDefaultParentType() string {
	return "policy-management"
}

func (obj *ApplicationPolicySet) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *ApplicationPolicySet) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *ApplicationPolicySet) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *ApplicationPolicySet) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *ApplicationPolicySet) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *ApplicationPolicySet) GetAllApplications() bool {
	return obj.all_applications
}

func (obj *ApplicationPolicySet) SetAllApplications(value bool) {
	obj.all_applications = value
	obj.modified.SetBit(&obj.modified, application_policy_set_all_applications, 1)
}

func (obj *ApplicationPolicySet) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *ApplicationPolicySet) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, application_policy_set_id_perms, 1)
}

func (obj *ApplicationPolicySet) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *ApplicationPolicySet) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, application_policy_set_perms2, 1)
}

func (obj *ApplicationPolicySet) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *ApplicationPolicySet) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, application_policy_set_annotations, 1)
}

func (obj *ApplicationPolicySet) GetDisplayName() string {
	return obj.display_name
}

func (obj *ApplicationPolicySet) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, application_policy_set_display_name, 1)
}

func (obj *ApplicationPolicySet) readFirewallPolicyRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(application_policy_set_firewall_policy_refs) == 0) {
		err := obj.GetField(obj, "firewall_policy_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ApplicationPolicySet) GetFirewallPolicyRefs() (
	contrail.ReferenceList, error) {
	err := obj.readFirewallPolicyRefs()
	if err != nil {
		return nil, err
	}
	return obj.firewall_policy_refs, nil
}

func (obj *ApplicationPolicySet) AddFirewallPolicy(
	rhs *FirewallPolicy, data FirewallSequence) error {
	err := obj.readFirewallPolicyRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(application_policy_set_firewall_policy_refs) == 0 {
		obj.storeReferenceBase("firewall-policy", obj.firewall_policy_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
	obj.firewall_policy_refs = append(obj.firewall_policy_refs, ref)
	obj.modified.SetBit(&obj.modified, application_policy_set_firewall_policy_refs, 1)
	return nil
}

func (obj *ApplicationPolicySet) DeleteFirewallPolicy(uuid string) error {
	err := obj.readFirewallPolicyRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(application_policy_set_firewall_policy_refs) == 0 {
		obj.storeReferenceBase("firewall-policy", obj.firewall_policy_refs)
	}

	for i, ref := range obj.firewall_policy_refs {
		if ref.Uuid == uuid {
			obj.firewall_policy_refs = append(
				obj.firewall_policy_refs[:i],
				obj.firewall_policy_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, application_policy_set_firewall_policy_refs, 1)
	return nil
}

func (obj *ApplicationPolicySet) ClearFirewallPolicy() {
	if (obj.valid.Bit(application_policy_set_firewall_policy_refs) != 0) &&
		(obj.modified.Bit(application_policy_set_firewall_policy_refs) == 0) {
		obj.storeReferenceBase("firewall-policy", obj.firewall_policy_refs)
	}
	obj.firewall_policy_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, application_policy_set_firewall_policy_refs, 1)
	obj.modified.SetBit(&obj.modified, application_policy_set_firewall_policy_refs, 1)
}

func (obj *ApplicationPolicySet) SetFirewallPolicyList(
	refList []contrail.ReferencePair) {
	obj.ClearFirewallPolicy()
	obj.firewall_policy_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.firewall_policy_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *ApplicationPolicySet) readGlobalVrouterConfigRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(application_policy_set_global_vrouter_config_refs) == 0) {
		err := obj.GetField(obj, "global_vrouter_config_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ApplicationPolicySet) GetGlobalVrouterConfigRefs() (
	contrail.ReferenceList, error) {
	err := obj.readGlobalVrouterConfigRefs()
	if err != nil {
		return nil, err
	}
	return obj.global_vrouter_config_refs, nil
}

func (obj *ApplicationPolicySet) AddGlobalVrouterConfig(
	rhs *GlobalVrouterConfig) error {
	err := obj.readGlobalVrouterConfigRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(application_policy_set_global_vrouter_config_refs) == 0 {
		obj.storeReferenceBase("global-vrouter-config", obj.global_vrouter_config_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.global_vrouter_config_refs = append(obj.global_vrouter_config_refs, ref)
	obj.modified.SetBit(&obj.modified, application_policy_set_global_vrouter_config_refs, 1)
	return nil
}

func (obj *ApplicationPolicySet) DeleteGlobalVrouterConfig(uuid string) error {
	err := obj.readGlobalVrouterConfigRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(application_policy_set_global_vrouter_config_refs) == 0 {
		obj.storeReferenceBase("global-vrouter-config", obj.global_vrouter_config_refs)
	}

	for i, ref := range obj.global_vrouter_config_refs {
		if ref.Uuid == uuid {
			obj.global_vrouter_config_refs = append(
				obj.global_vrouter_config_refs[:i],
				obj.global_vrouter_config_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, application_policy_set_global_vrouter_config_refs, 1)
	return nil
}

func (obj *ApplicationPolicySet) ClearGlobalVrouterConfig() {
	if (obj.valid.Bit(application_policy_set_global_vrouter_config_refs) != 0) &&
		(obj.modified.Bit(application_policy_set_global_vrouter_config_refs) == 0) {
		obj.storeReferenceBase("global-vrouter-config", obj.global_vrouter_config_refs)
	}
	obj.global_vrouter_config_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, application_policy_set_global_vrouter_config_refs, 1)
	obj.modified.SetBit(&obj.modified, application_policy_set_global_vrouter_config_refs, 1)
}

func (obj *ApplicationPolicySet) SetGlobalVrouterConfigList(
	refList []contrail.ReferencePair) {
	obj.ClearGlobalVrouterConfig()
	obj.global_vrouter_config_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.global_vrouter_config_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *ApplicationPolicySet) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(application_policy_set_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ApplicationPolicySet) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *ApplicationPolicySet) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(application_policy_set_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, application_policy_set_tag_refs, 1)
	return nil
}

func (obj *ApplicationPolicySet) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(application_policy_set_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, application_policy_set_tag_refs, 1)
	return nil
}

func (obj *ApplicationPolicySet) ClearTag() {
	if (obj.valid.Bit(application_policy_set_tag_refs) != 0) &&
		(obj.modified.Bit(application_policy_set_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, application_policy_set_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, application_policy_set_tag_refs, 1)
}

func (obj *ApplicationPolicySet) SetTagList(
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

func (obj *ApplicationPolicySet) readProjectBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(application_policy_set_project_back_refs) == 0) {
		err := obj.GetField(obj, "project_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ApplicationPolicySet) GetProjectBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readProjectBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.project_back_refs, nil
}

func (obj *ApplicationPolicySet) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(application_policy_set_all_applications) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.all_applications)
		if err != nil {
			return nil, err
		}
		msg["all_applications"] = &value
	}

	if obj.modified.Bit(application_policy_set_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(application_policy_set_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(application_policy_set_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(application_policy_set_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.firewall_policy_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.firewall_policy_refs)
		if err != nil {
			return nil, err
		}
		msg["firewall_policy_refs"] = &value
	}

	if len(obj.global_vrouter_config_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.global_vrouter_config_refs)
		if err != nil {
			return nil, err
		}
		msg["global_vrouter_config_refs"] = &value
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

func (obj *ApplicationPolicySet) UnmarshalJSON(body []byte) error {
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
		case "all_applications":
			err = json.Unmarshal(value, &obj.all_applications)
			if err == nil {
				obj.valid.SetBit(&obj.valid, application_policy_set_all_applications, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, application_policy_set_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, application_policy_set_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, application_policy_set_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, application_policy_set_display_name, 1)
			}
			break
		case "global_vrouter_config_refs":
			err = json.Unmarshal(value, &obj.global_vrouter_config_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, application_policy_set_global_vrouter_config_refs, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, application_policy_set_tag_refs, 1)
			}
			break
		case "project_back_refs":
			err = json.Unmarshal(value, &obj.project_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, application_policy_set_project_back_refs, 1)
			}
			break
		case "firewall_policy_refs":
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
				obj.valid.SetBit(&obj.valid, application_policy_set_firewall_policy_refs, 1)
				obj.firewall_policy_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.firewall_policy_refs = append(obj.firewall_policy_refs, ref)
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

func (obj *ApplicationPolicySet) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(application_policy_set_all_applications) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.all_applications)
		if err != nil {
			return nil, err
		}
		msg["all_applications"] = &value
	}

	if obj.modified.Bit(application_policy_set_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(application_policy_set_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(application_policy_set_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(application_policy_set_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(application_policy_set_firewall_policy_refs) != 0 {
		if len(obj.firewall_policy_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["firewall_policy_refs"] = &value
		} else if !obj.hasReferenceBase("firewall-policy") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.firewall_policy_refs)
			if err != nil {
				return nil, err
			}
			msg["firewall_policy_refs"] = &value
		}
	}

	if obj.modified.Bit(application_policy_set_global_vrouter_config_refs) != 0 {
		if len(obj.global_vrouter_config_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["global_vrouter_config_refs"] = &value
		} else if !obj.hasReferenceBase("global-vrouter-config") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.global_vrouter_config_refs)
			if err != nil {
				return nil, err
			}
			msg["global_vrouter_config_refs"] = &value
		}
	}

	if obj.modified.Bit(application_policy_set_tag_refs) != 0 {
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

func (obj *ApplicationPolicySet) UpdateReferences() error {

	if (obj.modified.Bit(application_policy_set_firewall_policy_refs) != 0) &&
		len(obj.firewall_policy_refs) > 0 &&
		obj.hasReferenceBase("firewall-policy") {
		err := obj.UpdateReference(
			obj, "firewall-policy",
			obj.firewall_policy_refs,
			obj.baseMap["firewall-policy"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(application_policy_set_global_vrouter_config_refs) != 0) &&
		len(obj.global_vrouter_config_refs) > 0 &&
		obj.hasReferenceBase("global-vrouter-config") {
		err := obj.UpdateReference(
			obj, "global-vrouter-config",
			obj.global_vrouter_config_refs,
			obj.baseMap["global-vrouter-config"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(application_policy_set_tag_refs) != 0) &&
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

func ApplicationPolicySetByName(c contrail.ApiClient, fqn string) (*ApplicationPolicySet, error) {
	obj, err := c.FindByName("application-policy-set", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*ApplicationPolicySet), nil
}

func ApplicationPolicySetByUuid(c contrail.ApiClient, uuid string) (*ApplicationPolicySet, error) {
	obj, err := c.FindByUuid("application-policy-set", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*ApplicationPolicySet), nil
}
