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
	loadbalancer_healthmonitor_loadbalancer_healthmonitor_properties int = iota
	loadbalancer_healthmonitor_id_perms
	loadbalancer_healthmonitor_perms2
	loadbalancer_healthmonitor_annotations
	loadbalancer_healthmonitor_display_name
	loadbalancer_healthmonitor_tag_refs
	loadbalancer_healthmonitor_loadbalancer_pool_back_refs
)

type LoadbalancerHealthmonitor struct {
	contrail.ObjectBase
	loadbalancer_healthmonitor_properties LoadbalancerHealthmonitorType
	id_perms                              IdPermsType
	perms2                                PermType2
	annotations                           KeyValuePairs
	display_name                          string
	tag_refs                              contrail.ReferenceList
	loadbalancer_pool_back_refs           contrail.ReferenceList
	valid                                 big.Int
	modified                              big.Int
	baseMap                               map[string]contrail.ReferenceList
}

func (obj *LoadbalancerHealthmonitor) GetType() string {
	return "loadbalancer-healthmonitor"
}

func (obj *LoadbalancerHealthmonitor) GetDefaultParent() []string {
	name := []string{"default-domain", "default-project"}
	return name
}

func (obj *LoadbalancerHealthmonitor) GetDefaultParentType() string {
	return "project"
}

func (obj *LoadbalancerHealthmonitor) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *LoadbalancerHealthmonitor) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *LoadbalancerHealthmonitor) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *LoadbalancerHealthmonitor) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *LoadbalancerHealthmonitor) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *LoadbalancerHealthmonitor) GetLoadbalancerHealthmonitorProperties() LoadbalancerHealthmonitorType {
	return obj.loadbalancer_healthmonitor_properties
}

func (obj *LoadbalancerHealthmonitor) SetLoadbalancerHealthmonitorProperties(value *LoadbalancerHealthmonitorType) {
	obj.loadbalancer_healthmonitor_properties = *value
	obj.modified.SetBit(&obj.modified, loadbalancer_healthmonitor_loadbalancer_healthmonitor_properties, 1)
}

func (obj *LoadbalancerHealthmonitor) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *LoadbalancerHealthmonitor) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, loadbalancer_healthmonitor_id_perms, 1)
}

func (obj *LoadbalancerHealthmonitor) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *LoadbalancerHealthmonitor) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, loadbalancer_healthmonitor_perms2, 1)
}

func (obj *LoadbalancerHealthmonitor) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *LoadbalancerHealthmonitor) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, loadbalancer_healthmonitor_annotations, 1)
}

func (obj *LoadbalancerHealthmonitor) GetDisplayName() string {
	return obj.display_name
}

func (obj *LoadbalancerHealthmonitor) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, loadbalancer_healthmonitor_display_name, 1)
}

func (obj *LoadbalancerHealthmonitor) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(loadbalancer_healthmonitor_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *LoadbalancerHealthmonitor) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *LoadbalancerHealthmonitor) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, loadbalancer_healthmonitor_tag_refs, 1)
	return nil
}

func (obj *LoadbalancerHealthmonitor) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, loadbalancer_healthmonitor_tag_refs, 1)
	return nil
}

func (obj *LoadbalancerHealthmonitor) ClearTag() {
	if (obj.valid.Bit(loadbalancer_healthmonitor_tag_refs) != 0) &&
		(obj.modified.Bit(loadbalancer_healthmonitor_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, loadbalancer_healthmonitor_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, loadbalancer_healthmonitor_tag_refs, 1)
}

func (obj *LoadbalancerHealthmonitor) SetTagList(
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

func (obj *LoadbalancerHealthmonitor) readLoadbalancerPoolBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(loadbalancer_healthmonitor_loadbalancer_pool_back_refs) == 0) {
		err := obj.GetField(obj, "loadbalancer_pool_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *LoadbalancerHealthmonitor) GetLoadbalancerPoolBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readLoadbalancerPoolBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.loadbalancer_pool_back_refs, nil
}

func (obj *LoadbalancerHealthmonitor) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_loadbalancer_healthmonitor_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.loadbalancer_healthmonitor_properties)
		if err != nil {
			return nil, err
		}
		msg["loadbalancer_healthmonitor_properties"] = &value
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_display_name) != 0 {
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

func (obj *LoadbalancerHealthmonitor) UnmarshalJSON(body []byte) error {
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
		case "loadbalancer_healthmonitor_properties":
			err = json.Unmarshal(value, &obj.loadbalancer_healthmonitor_properties)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_healthmonitor_loadbalancer_healthmonitor_properties, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_healthmonitor_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_healthmonitor_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_healthmonitor_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_healthmonitor_display_name, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_healthmonitor_tag_refs, 1)
			}
			break
		case "loadbalancer_pool_back_refs":
			err = json.Unmarshal(value, &obj.loadbalancer_pool_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, loadbalancer_healthmonitor_loadbalancer_pool_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *LoadbalancerHealthmonitor) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_loadbalancer_healthmonitor_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.loadbalancer_healthmonitor_properties)
		if err != nil {
			return nil, err
		}
		msg["loadbalancer_healthmonitor_properties"] = &value
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(loadbalancer_healthmonitor_tag_refs) != 0 {
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

func (obj *LoadbalancerHealthmonitor) UpdateReferences() error {

	if (obj.modified.Bit(loadbalancer_healthmonitor_tag_refs) != 0) &&
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

func LoadbalancerHealthmonitorByName(c contrail.ApiClient, fqn string) (*LoadbalancerHealthmonitor, error) {
	obj, err := c.FindByName("loadbalancer-healthmonitor", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*LoadbalancerHealthmonitor), nil
}

func LoadbalancerHealthmonitorByUuid(c contrail.ApiClient, uuid string) (*LoadbalancerHealthmonitor, error) {
	obj, err := c.FindByUuid("loadbalancer-healthmonitor", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*LoadbalancerHealthmonitor), nil
}
