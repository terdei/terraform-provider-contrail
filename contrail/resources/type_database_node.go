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
	database_node_database_node_ip_address int = iota
	database_node_id_perms
	database_node_perms2
	database_node_annotations
	database_node_display_name
	database_node_tag_refs
)

type DatabaseNode struct {
	contrail.ObjectBase
	database_node_ip_address string
	id_perms                 IdPermsType
	perms2                   PermType2
	annotations              KeyValuePairs
	display_name             string
	tag_refs                 contrail.ReferenceList
	valid                    big.Int
	modified                 big.Int
	baseMap                  map[string]contrail.ReferenceList
}

func (obj *DatabaseNode) GetType() string {
	return "database-node"
}

func (obj *DatabaseNode) GetDefaultParent() []string {
	name := []string{"default-global-system-config"}
	return name
}

func (obj *DatabaseNode) GetDefaultParentType() string {
	return "global-system-config"
}

func (obj *DatabaseNode) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *DatabaseNode) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *DatabaseNode) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *DatabaseNode) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *DatabaseNode) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *DatabaseNode) GetDatabaseNodeIpAddress() string {
	return obj.database_node_ip_address
}

func (obj *DatabaseNode) SetDatabaseNodeIpAddress(value string) {
	obj.database_node_ip_address = value
	obj.modified.SetBit(&obj.modified, database_node_database_node_ip_address, 1)
}

func (obj *DatabaseNode) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *DatabaseNode) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, database_node_id_perms, 1)
}

func (obj *DatabaseNode) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *DatabaseNode) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, database_node_perms2, 1)
}

func (obj *DatabaseNode) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *DatabaseNode) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, database_node_annotations, 1)
}

func (obj *DatabaseNode) GetDisplayName() string {
	return obj.display_name
}

func (obj *DatabaseNode) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, database_node_display_name, 1)
}

func (obj *DatabaseNode) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(database_node_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *DatabaseNode) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *DatabaseNode) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(database_node_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, database_node_tag_refs, 1)
	return nil
}

func (obj *DatabaseNode) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(database_node_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, database_node_tag_refs, 1)
	return nil
}

func (obj *DatabaseNode) ClearTag() {
	if (obj.valid.Bit(database_node_tag_refs) != 0) &&
		(obj.modified.Bit(database_node_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, database_node_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, database_node_tag_refs, 1)
}

func (obj *DatabaseNode) SetTagList(
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

func (obj *DatabaseNode) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(database_node_database_node_ip_address) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.database_node_ip_address)
		if err != nil {
			return nil, err
		}
		msg["database_node_ip_address"] = &value
	}

	if obj.modified.Bit(database_node_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(database_node_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(database_node_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(database_node_display_name) != 0 {
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

func (obj *DatabaseNode) UnmarshalJSON(body []byte) error {
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
		case "database_node_ip_address":
			err = json.Unmarshal(value, &obj.database_node_ip_address)
			if err == nil {
				obj.valid.SetBit(&obj.valid, database_node_database_node_ip_address, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, database_node_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, database_node_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, database_node_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, database_node_display_name, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, database_node_tag_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *DatabaseNode) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(database_node_database_node_ip_address) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.database_node_ip_address)
		if err != nil {
			return nil, err
		}
		msg["database_node_ip_address"] = &value
	}

	if obj.modified.Bit(database_node_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(database_node_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(database_node_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(database_node_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(database_node_tag_refs) != 0 {
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

func (obj *DatabaseNode) UpdateReferences() error {

	if (obj.modified.Bit(database_node_tag_refs) != 0) &&
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

func DatabaseNodeByName(c contrail.ApiClient, fqn string) (*DatabaseNode, error) {
	obj, err := c.FindByName("database-node", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*DatabaseNode), nil
}

func DatabaseNodeByUuid(c contrail.ApiClient, uuid string) (*DatabaseNode, error) {
	obj, err := c.FindByUuid("database-node", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*DatabaseNode), nil
}
