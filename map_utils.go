package go_commons

func GetListKeyFromMap(input map[string]interface{}) []string {
	results := make([]string, 0)
	for k := range input {
		results = append(results, k)
	}
	return results
}

func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

func Append(original map[string]float64, input map[string]float64) map[string]float64 {
	for k, v := range input {
		value := v
		if val, ok := original[k]; ok {
			value += val
		}
		original[k] = value
	}
	return original
}
