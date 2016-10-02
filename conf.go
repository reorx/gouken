package gouken

import (
	"fmt"
	"sort"

	"github.com/spf13/viper"
)

// ConfString format all configs to string
func ConfString() string {
	s := ""
	items := viper.AllSettings()
	keys := make([]string, len(items))

	// Sort keys
	i := 0
	for k := range items {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, k := range keys {
		s += fmt.Sprintf("%v: %v\n", k, items[k])
	}
	return s
}
