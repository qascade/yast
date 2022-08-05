package core

import(
	"os"
	"encoding/json"

	"github.com/qascade/yast/tui"
	"github.com/qascade/yast/utils"
	"github.com/tidwall/sjson"
)

//This function will call interactive tui for taking input of user preferences. 
func FillConfigJSON(configFile *os.File) error {
	utils.LogUnimplementedFunc()
	return nil
}


type ConfigBuildSpec struct {
	Player string `json:"player"`
	TargetPreference string `json:"target-preference"`
	QueryHistory bool `json:"query-history"`
}
