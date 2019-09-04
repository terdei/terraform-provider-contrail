//
// Automatically generated. DO NOT EDIT.
//

package resources

type FirewallRuleMatchTagsType struct {
	TagList []string `json:"tag_list,omitempty"`
}

func (obj *FirewallRuleMatchTagsType) AddTagList(value string) {
	obj.TagList = append(obj.TagList, value)
}
