package maputils

import "encoding/json"

func MapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}

func OptionalString(m map[string]interface{}, name string) string {
	vstr := m[name]
	if vstr != nil {
		return vstr.(string)
	}
	return ""
}
