//
// Automatically generated. DO NOT EDIT.
//

package resources

type EncapsulationPrioritiesType struct {
	Encapsulation []string `json:"encapsulation,omitempty"`
}

func (obj *EncapsulationPrioritiesType) AddEncapsulation(value string) {
	obj.Encapsulation = append(obj.Encapsulation, value)
}
