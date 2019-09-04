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
	floating_ip_pool_floating_ip_pool_subnets int = iota
	floating_ip_pool_id_perms
	floating_ip_pool_perms2
	floating_ip_pool_annotations
	floating_ip_pool_display_name
	floating_ip_pool_floating_ips
	floating_ip_pool_tag_refs
	floating_ip_pool_project_back_refs
)

type FloatingIpPool struct {
	contrail.ObjectBase
	floating_ip_pool_subnets FloatingIpPoolSubnetType
	id_perms                 IdPermsType
	perms2                   PermType2
	annotations              KeyValuePairs
	display_name             string
	floating_ips             contrail.ReferenceList
	tag_refs                 contrail.ReferenceList
	project_back_refs        contrail.ReferenceList
	valid                    big.Int
	modified                 big.Int
	baseMap                  map[string]contrail.ReferenceList
}

func (obj *FloatingIpPool) GetType() string {
	return "floating-ip-pool"
}

func (obj *FloatingIpPool) GetDefaultParent() []string {
	name := []string{"default-domain", "default-project", "default-virtual-network"}
	return name
}

func (obj *FloatingIpPool) GetDefaultParentType() string {
	return "virtual-network"
}

func (obj *FloatingIpPool) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *FloatingIpPool) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *FloatingIpPool) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *FloatingIpPool) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *FloatingIpPool) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *FloatingIpPool) GetFloatingIpPoolSubnets() FloatingIpPoolSubnetType {
	return obj.floating_ip_pool_subnets
}

func (obj *FloatingIpPool) SetFloatingIpPoolSubnets(value *FloatingIpPoolSubnetType) {
	obj.floating_ip_pool_subnets = *value
	obj.modified.SetBit(&obj.modified, floating_ip_pool_floating_ip_pool_subnets, 1)
}

func (obj *FloatingIpPool) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *FloatingIpPool) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, floating_ip_pool_id_perms, 1)
}

func (obj *FloatingIpPool) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *FloatingIpPool) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, floating_ip_pool_perms2, 1)
}

func (obj *FloatingIpPool) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *FloatingIpPool) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, floating_ip_pool_annotations, 1)
}

func (obj *FloatingIpPool) GetDisplayName() string {
	return obj.display_name
}

func (obj *FloatingIpPool) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, floating_ip_pool_display_name, 1)
}

func (obj *FloatingIpPool) readFloatingIps() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(floating_ip_pool_floating_ips) == 0) {
		err := obj.GetField(obj, "floating_ips")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FloatingIpPool) GetFloatingIps() (
	contrail.ReferenceList, error) {
	err := obj.readFloatingIps()
	if err != nil {
		return nil, err
	}
	return obj.floating_ips, nil
}

func (obj *FloatingIpPool) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(floating_ip_pool_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FloatingIpPool) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *FloatingIpPool) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(floating_ip_pool_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, floating_ip_pool_tag_refs, 1)
	return nil
}

func (obj *FloatingIpPool) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(floating_ip_pool_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, floating_ip_pool_tag_refs, 1)
	return nil
}

func (obj *FloatingIpPool) ClearTag() {
	if (obj.valid.Bit(floating_ip_pool_tag_refs) != 0) &&
		(obj.modified.Bit(floating_ip_pool_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, floating_ip_pool_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, floating_ip_pool_tag_refs, 1)
}

func (obj *FloatingIpPool) SetTagList(
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

func (obj *FloatingIpPool) readProjectBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(floating_ip_pool_project_back_refs) == 0) {
		err := obj.GetField(obj, "project_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FloatingIpPool) GetProjectBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readProjectBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.project_back_refs, nil
}

func (obj *FloatingIpPool) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(floating_ip_pool_floating_ip_pool_subnets) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.floating_ip_pool_subnets)
		if err != nil {
			return nil, err
		}
		msg["floating_ip_pool_subnets"] = &value
	}

	if obj.modified.Bit(floating_ip_pool_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(floating_ip_pool_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(floating_ip_pool_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(floating_ip_pool_display_name) != 0 {
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

func (obj *FloatingIpPool) UnmarshalJSON(body []byte) error {
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
		case "floating_ip_pool_subnets":
			err = json.Unmarshal(value, &obj.floating_ip_pool_subnets)
			if err == nil {
				obj.valid.SetBit(&obj.valid, floating_ip_pool_floating_ip_pool_subnets, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, floating_ip_pool_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, floating_ip_pool_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, floating_ip_pool_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, floating_ip_pool_display_name, 1)
			}
			break
		case "floating_ips":
			err = json.Unmarshal(value, &obj.floating_ips)
			if err == nil {
				obj.valid.SetBit(&obj.valid, floating_ip_pool_floating_ips, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, floating_ip_pool_tag_refs, 1)
			}
			break
		case "project_back_refs":
			err = json.Unmarshal(value, &obj.project_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, floating_ip_pool_project_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *FloatingIpPool) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(floating_ip_pool_floating_ip_pool_subnets) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.floating_ip_pool_subnets)
		if err != nil {
			return nil, err
		}
		msg["floating_ip_pool_subnets"] = &value
	}

	if obj.modified.Bit(floating_ip_pool_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(floating_ip_pool_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(floating_ip_pool_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(floating_ip_pool_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(floating_ip_pool_tag_refs) != 0 {
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

func (obj *FloatingIpPool) UpdateReferences() error {

	if (obj.modified.Bit(floating_ip_pool_tag_refs) != 0) &&
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

func FloatingIpPoolByName(c contrail.ApiClient, fqn string) (*FloatingIpPool, error) {
	obj, err := c.FindByName("floating-ip-pool", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*FloatingIpPool), nil
}

func FloatingIpPoolByUuid(c contrail.ApiClient, uuid string) (*FloatingIpPool, error) {
	obj, err := c.FindByUuid("floating-ip-pool", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*FloatingIpPool), nil
}
