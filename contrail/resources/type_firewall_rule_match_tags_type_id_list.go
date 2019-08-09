//
// Automatically generated. DO NOT EDIT.
//

package resources

type FirewallRuleMatchTagsTypeIdList struct {
	TagType []int `json:"tag_type,omitempty"`
}

func (obj *FirewallRuleMatchTagsTypeIdList) AddTagType(value int) {
	obj.TagType = append(obj.TagType, value)
}
