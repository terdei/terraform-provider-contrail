//
// Automatically generated. DO NOT EDIT.
//

package resources

type UveKeysType struct {
	UveKey []string `json:"uve_key,omitempty"`
}

func (obj *UveKeysType) AddUveKey(value string) {
	obj.UveKey = append(obj.UveKey, value)
}
