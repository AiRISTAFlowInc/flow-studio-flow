package util

import (
	"github.com/AiRISTAFlowInc/flow-studio-core/data/coerce"
	"github.com/mohae/deepcopy"
)

func DeepCopy(data interface{}) interface{} {
	return deepcopy.Copy(data)
}

func DeepCopyMap(data map[string]interface{}) map[string]interface{} {
	copiedData := deepcopy.Copy(data)
	copiedMap, _ := coerce.ToObject(copiedData)
	return copiedMap
}
