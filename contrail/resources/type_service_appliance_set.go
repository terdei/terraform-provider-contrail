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
	service_appliance_set_service_appliance_set_properties int = iota
	service_appliance_set_service_appliance_driver
	service_appliance_set_service_appliance_ha_mode
	service_appliance_set_id_perms
	service_appliance_set_perms2
	service_appliance_set_annotations
	service_appliance_set_display_name
	service_appliance_set_service_appliances
	service_appliance_set_tag_refs
	service_appliance_set_service_template_back_refs
	service_appliance_set_loadbalancer_pool_back_refs
	service_appliance_set_loadbalancer_back_refs
)

type ServiceApplianceSet struct {
	contrail.ObjectBase
	service_appliance_set_properties KeyValuePairs
	service_appliance_driver         string
	service_appliance_ha_mode        string
	id_perms                         IdPermsType
	perms2                           PermType2
	annotations                      KeyValuePairs
	display_name                     string
	service_appliances               contrail.ReferenceList
	tag_refs                         contrail.ReferenceList
	service_template_back_refs       contrail.ReferenceList
	loadbalancer_pool_back_refs      contrail.ReferenceList
	loadbalancer_back_refs           contrail.ReferenceList
	valid                            big.Int
	modified                         big.Int
	baseMap                          map[string]contrail.ReferenceList
}

func (obj *ServiceApplianceSet) GetType() string {
	return "service-appliance-set"
}

func (obj *ServiceApplianceSet) GetDefaultParent() []string {
	name := []string{"default-global-system-config"}
	return name
}

func (obj *ServiceApplianceSet) GetDefaultParentType() string {
	return "global-system-config"
}

func (obj *ServiceApplianceSet) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *ServiceApplianceSet) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *ServiceApplianceSet) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *ServiceApplianceSet) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *ServiceApplianceSet) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *ServiceApplianceSet) GetServiceApplianceSetProperties() KeyValuePairs {
	return obj.service_appliance_set_properties
}

func (obj *ServiceApplianceSet) SetServiceApplianceSetProperties(value *KeyValuePairs) {
	obj.service_appliance_set_properties = *value
	obj.modified.SetBit(&obj.modified, service_appliance_set_service_appliance_set_properties, 1)
}

func (obj *ServiceApplianceSet) GetServiceApplianceDriver() string {
	return obj.service_appliance_driver
}

func (obj *ServiceApplianceSet) SetServiceApplianceDriver(value string) {
	obj.service_appliance_driver = value
	obj.modified.SetBit(&obj.modified, service_appliance_set_service_appliance_driver, 1)
}

func (obj *ServiceApplianceSet) GetServiceApplianceHaMode() string {
	return obj.service_appliance_ha_mode
}

func (obj *ServiceApplianceSet) SetServiceApplianceHaMode(value string) {
	obj.service_appliance_ha_mode = value
	obj.modified.SetBit(&obj.modified, service_appliance_set_service_appliance_ha_mode, 1)
}

func (obj *ServiceApplianceSet) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *ServiceApplianceSet) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, service_appliance_set_id_perms, 1)
}

func (obj *ServiceApplianceSet) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *ServiceApplianceSet) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, service_appliance_set_perms2, 1)
}

func (obj *ServiceApplianceSet) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *ServiceApplianceSet) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, service_appliance_set_annotations, 1)
}

func (obj *ServiceApplianceSet) GetDisplayName() string {
	return obj.display_name
}

func (obj *ServiceApplianceSet) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, service_appliance_set_display_name, 1)
}

func (obj *ServiceApplianceSet) readServiceAppliances() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_appliance_set_service_appliances) == 0) {
		err := obj.GetField(obj, "service_appliances")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceApplianceSet) GetServiceAppliances() (
	contrail.ReferenceList, error) {
	err := obj.readServiceAppliances()
	if err != nil {
		return nil, err
	}
	return obj.service_appliances, nil
}

func (obj *ServiceApplianceSet) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_appliance_set_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceApplianceSet) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *ServiceApplianceSet) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_appliance_set_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, service_appliance_set_tag_refs, 1)
	return nil
}

func (obj *ServiceApplianceSet) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(service_appliance_set_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, service_appliance_set_tag_refs, 1)
	return nil
}

func (obj *ServiceApplianceSet) ClearTag() {
	if (obj.valid.Bit(service_appliance_set_tag_refs) != 0) &&
		(obj.modified.Bit(service_appliance_set_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, service_appliance_set_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, service_appliance_set_tag_refs, 1)
}

func (obj *ServiceApplianceSet) SetTagList(
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

func (obj *ServiceApplianceSet) readServiceTemplateBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_appliance_set_service_template_back_refs) == 0) {
		err := obj.GetField(obj, "service_template_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceApplianceSet) GetServiceTemplateBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readServiceTemplateBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.service_template_back_refs, nil
}

func (obj *ServiceApplianceSet) readLoadbalancerPoolBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_appliance_set_loadbalancer_pool_back_refs) == 0) {
		err := obj.GetField(obj, "loadbalancer_pool_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceApplianceSet) GetLoadbalancerPoolBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readLoadbalancerPoolBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.loadbalancer_pool_back_refs, nil
}

func (obj *ServiceApplianceSet) readLoadbalancerBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(service_appliance_set_loadbalancer_back_refs) == 0) {
		err := obj.GetField(obj, "loadbalancer_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceApplianceSet) GetLoadbalancerBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readLoadbalancerBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.loadbalancer_back_refs, nil
}

func (obj *ServiceApplianceSet) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(service_appliance_set_service_appliance_set_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_set_properties)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_set_properties"] = &value
	}

	if obj.modified.Bit(service_appliance_set_service_appliance_driver) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_driver)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_driver"] = &value
	}

	if obj.modified.Bit(service_appliance_set_service_appliance_ha_mode) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_ha_mode)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_ha_mode"] = &value
	}

	if obj.modified.Bit(service_appliance_set_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(service_appliance_set_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(service_appliance_set_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(service_appliance_set_display_name) != 0 {
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

func (obj *ServiceApplianceSet) UnmarshalJSON(body []byte) error {
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
		case "service_appliance_set_properties":
			err = json.Unmarshal(value, &obj.service_appliance_set_properties)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_service_appliance_set_properties, 1)
			}
			break
		case "service_appliance_driver":
			err = json.Unmarshal(value, &obj.service_appliance_driver)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_service_appliance_driver, 1)
			}
			break
		case "service_appliance_ha_mode":
			err = json.Unmarshal(value, &obj.service_appliance_ha_mode)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_service_appliance_ha_mode, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_display_name, 1)
			}
			break
		case "service_appliances":
			err = json.Unmarshal(value, &obj.service_appliances)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_service_appliances, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_tag_refs, 1)
			}
			break
		case "service_template_back_refs":
			err = json.Unmarshal(value, &obj.service_template_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_service_template_back_refs, 1)
			}
			break
		case "loadbalancer_pool_back_refs":
			err = json.Unmarshal(value, &obj.loadbalancer_pool_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_loadbalancer_pool_back_refs, 1)
			}
			break
		case "loadbalancer_back_refs":
			err = json.Unmarshal(value, &obj.loadbalancer_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, service_appliance_set_loadbalancer_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *ServiceApplianceSet) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(service_appliance_set_service_appliance_set_properties) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_set_properties)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_set_properties"] = &value
	}

	if obj.modified.Bit(service_appliance_set_service_appliance_driver) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_driver)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_driver"] = &value
	}

	if obj.modified.Bit(service_appliance_set_service_appliance_ha_mode) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.service_appliance_ha_mode)
		if err != nil {
			return nil, err
		}
		msg["service_appliance_ha_mode"] = &value
	}

	if obj.modified.Bit(service_appliance_set_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(service_appliance_set_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(service_appliance_set_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(service_appliance_set_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(service_appliance_set_tag_refs) != 0 {
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

func (obj *ServiceApplianceSet) UpdateReferences() error {

	if (obj.modified.Bit(service_appliance_set_tag_refs) != 0) &&
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

func ServiceApplianceSetByName(c contrail.ApiClient, fqn string) (*ServiceApplianceSet, error) {
	obj, err := c.FindByName("service-appliance-set", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*ServiceApplianceSet), nil
}

func ServiceApplianceSetByUuid(c contrail.ApiClient, uuid string) (*ServiceApplianceSet, error) {
	obj, err := c.FindByUuid("service-appliance-set", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*ServiceApplianceSet), nil
}
