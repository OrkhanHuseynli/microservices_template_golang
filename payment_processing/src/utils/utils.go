package utils

import (
	"github.com/lithammer/shortuuid"
)

func GenShortUUID() string {
	return shortuuid.New()
}
