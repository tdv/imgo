package service

import (
	"strings"
)

func configPath(items ...string) string {
	return strings.Join(items, ".")
}
