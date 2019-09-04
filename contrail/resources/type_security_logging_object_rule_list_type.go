//
// Automatically generated. DO NOT EDIT.
//

package resources

type SecurityLoggingObjectRuleEntryType struct {
	RuleUuid string `json:"rule_uuid,omitempty"`
	Rate     int    `json:"rate,omitempty"`
}

type SecurityLoggingObjectRuleListType struct {
	Rule []SecurityLoggingObjectRuleEntryType `json:"rule,omitempty"`
}

func (obj *SecurityLoggingObjectRuleListType) AddRule(value *SecurityLoggingObjectRuleEntryType) {
	obj.Rule = append(obj.Rule, *value)
}
