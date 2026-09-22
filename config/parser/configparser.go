package configparser

type ConfigMap map[string]string

func Parse(input []byte) (ConfigMap, error) {
	s := string(input)

	splitAndCleanRows(s)

	return ConfigMap{}, nil
}
