package service

import (
	"strings"
)

func ConfigPath(items ...string) string {
	return strings.Join(items, ".")
}
