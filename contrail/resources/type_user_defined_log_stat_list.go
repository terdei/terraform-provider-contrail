//
// Automatically generated. DO NOT EDIT.
//

package resources

type UserDefinedLogStat struct {
	Name    string `json:"name,omitempty"`
	Pattern string `json:"pattern,omitempty"`
}

type UserDefinedLogStatList struct {
	Statlist []UserDefinedLogStat `json:"statlist,omitempty"`
}

func (obj *UserDefinedLogStatList) AddStatlist(value *UserDefinedLogStat) {
	obj.Statlist = append(obj.Statlist, *value)
}
