package configparser

import "fmt"

type ConfigMap map[string]string

func Parse(input []byte) (ConfigMap, error) {
	s := string(input)

	rows := splitAndCleanRows(s)

	configMap := make(ConfigMap, len(rows))
	for _, row := range rows {
		key, val, err := parseRow(row)
		if err != nil {
			return ConfigMap{}, fmt.Errorf(
				"unable to parse config: %w",
				err,
			)
		}

		configMap[key] = val
	}

	return configMap, nil
}
