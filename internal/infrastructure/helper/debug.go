package helper

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"
)

func PrintStruct(data interface{}) {
	logData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {

		return
	}
	log.Debug(string(logData))
}
