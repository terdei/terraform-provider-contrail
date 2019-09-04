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
	route_table_routes int = iota
	route_table_id_perms
	route_table_perms2
	route_table_annotations
	route_table_display_name
	route_table_tag_refs
	route_table_virtual_network_back_refs
	route_table_logical_router_back_refs
)

type RouteTable struct {
	contrail.ObjectBase
	routes                    RouteTableType
	id_perms                  IdPermsType
	perms2                    PermType2
	annotations               KeyValuePairs
	display_name              string
	tag_refs                  contrail.ReferenceList
	virtual_network_back_refs contrail.ReferenceList
	logical_router_back_refs  contrail.ReferenceList
	valid                     big.Int
	modified                  big.Int
	baseMap                   map[string]contrail.ReferenceList
}

func (obj *RouteTable) GetType() string {
	return "route-table"
}

func (obj *RouteTable) GetDefaultParent() []string {
	name := []string{"default-domain", "default-project"}
	return name
}

func (obj *RouteTable) GetDefaultParentType() string {
	return "project"
}

func (obj *RouteTable) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *RouteTable) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *RouteTable) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *RouteTable) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *RouteTable) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *RouteTable) GetRoutes() RouteTableType {
	return obj.routes
}

func (obj *RouteTable) SetRoutes(value *RouteTableType) {
	obj.routes = *value
	obj.modified.SetBit(&obj.modified, route_table_routes, 1)
}

func (obj *RouteTable) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *RouteTable) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, route_table_id_perms, 1)
}

func (obj *RouteTable) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *RouteTable) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, route_table_perms2, 1)
}

func (obj *RouteTable) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *RouteTable) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, route_table_annotations, 1)
}

func (obj *RouteTable) GetDisplayName() string {
	return obj.display_name
}

func (obj *RouteTable) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, route_table_display_name, 1)
}

func (obj *RouteTable) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(route_table_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *RouteTable) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *RouteTable) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(route_table_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, route_table_tag_refs, 1)
	return nil
}

func (obj *RouteTable) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(route_table_tag_refs) == 0 {
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
	obj.modified.SetBit(&obj.modified, route_table_tag_refs, 1)
	return nil
}

func (obj *RouteTable) ClearTag() {
	if (obj.valid.Bit(route_table_tag_refs) != 0) &&
		(obj.modified.Bit(route_table_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, route_table_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, route_table_tag_refs, 1)
}

func (obj *RouteTable) SetTagList(
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

func (obj *RouteTable) readVirtualNetworkBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(route_table_virtual_network_back_refs) == 0) {
		err := obj.GetField(obj, "virtual_network_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *RouteTable) GetVirtualNetworkBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualNetworkBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_network_back_refs, nil
}

func (obj *RouteTable) readLogicalRouterBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(route_table_logical_router_back_refs) == 0) {
		err := obj.GetField(obj, "logical_router_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *RouteTable) GetLogicalRouterBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readLogicalRouterBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.logical_router_back_refs, nil
}

func (obj *RouteTable) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(route_table_routes) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.routes)
		if err != nil {
			return nil, err
		}
		msg["routes"] = &value
	}

	if obj.modified.Bit(route_table_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(route_table_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(route_table_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(route_table_display_name) != 0 {
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

func (obj *RouteTable) UnmarshalJSON(body []byte) error {
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
		case "routes":
			err = json.Unmarshal(value, &obj.routes)
			if err == nil {
				obj.valid.SetBit(&obj.valid, route_table_routes, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, route_table_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, route_table_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, route_table_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, route_table_display_name, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, route_table_tag_refs, 1)
			}
			break
		case "virtual_network_back_refs":
			err = json.Unmarshal(value, &obj.virtual_network_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, route_table_virtual_network_back_refs, 1)
			}
			break
		case "logical_router_back_refs":
			err = json.Unmarshal(value, &obj.logical_router_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, route_table_logical_router_back_refs, 1)
			}
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *RouteTable) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(route_table_routes) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.routes)
		if err != nil {
			return nil, err
		}
		msg["routes"] = &value
	}

	if obj.modified.Bit(route_table_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(route_table_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(route_table_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(route_table_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(route_table_tag_refs) != 0 {
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

func (obj *RouteTable) UpdateReferences() error {

	if (obj.modified.Bit(route_table_tag_refs) != 0) &&
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

func RouteTableByName(c contrail.ApiClient, fqn string) (*RouteTable, error) {
	obj, err := c.FindByName("route-table", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*RouteTable), nil
}

func RouteTableByUuid(c contrail.ApiClient, uuid string) (*RouteTable, error) {
	obj, err := c.FindByUuid("route-table", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*RouteTable), nil
}
