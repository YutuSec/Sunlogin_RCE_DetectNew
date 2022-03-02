package DataHandle

import (
	"encoding/json"
	"fmt"
)

type AutoGenerated struct {
	Code         int    `json:"__code"`
	Enabled      string `json:"enabled"`
	VerifyString string `json:"verify_string"`
}

func Jsonhandle(jsondata string) string {
	jsonstr := []byte(jsondata)
	var jsonResult AutoGenerated
	err := json.Unmarshal(jsonstr, &jsonResult)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return jsonResult.VerifyString
}