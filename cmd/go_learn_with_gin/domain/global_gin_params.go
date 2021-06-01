package domain

import (
	"strconv"

	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/functions"
)

func GlobalGinParams() map[string]interface{} {
	return map[string]interface{}{
		"year":  strconv.Itoa(functions.GetYear()),
		"users": GetAllUsers(),
	}
}
