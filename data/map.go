package data

import (
	"gopkg.in/yaml.v2"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func updateMapSlice(ms *yaml.MapSlice, key string, updateFunc func(interface{}) interface{}) {
	for i, item := range *ms {
		if item.Key == key {
			(*ms)[i].Value = updateFunc(item.Value)
			break
		}
	}
}

func cloneMap(originalMap map[interface{}]interface{}) map[interface{}]interface{} {
	clonedMap := make(map[interface{}]interface{})
	for key, value := range originalMap {
		clonedMap[key] = value
	}
	return clonedMap
}

// StringMapToResourceList 将字符串类型的map转换为kubernetes的ResourceList
func StringMapToResourceList(m map[string]string) coreV1.ResourceList {
	resourceList := make(coreV1.ResourceList)
	for key, value := range m {
		quantity, err := resource.ParseQuantity(value)
		if err != nil {
			continue
		}
		resourceList[coreV1.ResourceName(key)] = quantity
	}
	return resourceList
}
