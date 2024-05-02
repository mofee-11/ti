package conf

import (
	"os"
	"strconv"
	"strings"
)

type config struct {
	DBPath string
}

func Parse(path string) (config, error) {
	var ret config
	var err error
	var data []byte
	data, err = os.ReadFile(path)
	if err != nil {
		return ret, err
	}
	doc := unmarshal(string(data))
	ret.DBPath = doc["db"].(string)

	return ret, err
}

func unmarshal(text string) map[string]interface{} {
	ret := make(map[string]interface{})

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value, err := strconv.Unquote(strings.TrimSpace(parts[1]))
		if err != nil {
			panic(err)
		}

		ret[key] = value
	}

	return ret
}
