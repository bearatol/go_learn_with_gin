package domain

import (
	"strconv"

	"github.com/bearatol/go_learn_with_gin/cmd/functions"
)

func GlobalGinParams() map[string]interface{} {
	globalGinParams := map[string]interface{}{
		"year":  strconv.Itoa(functions.GetYear()),
		"users": GetAllUsers(),
	}
	return globalGinParams
}
