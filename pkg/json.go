package pkg

import (
	"encoding/json"
)

func JsonFormat(txt string) (result string, err error) {
	mapStruct := make(map[string]interface{})

	result = txt
	if err = json.Unmarshal([]byte(txt), &mapStruct); err == nil {
		var data []byte
		data, err = json.MarshalIndent(mapStruct, "", "  ")
		result = string(data)
	}
	return
}
