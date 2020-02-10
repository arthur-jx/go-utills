package xjson

import (
	"bytes"
	"encoding/json"
	"log"
	"reflect"
	"strings"
)

//json map对象操作工具函数

func GetValueType(a interface{}) string {
	return reflect.TypeOf(a).String()
}

//将Json字符串转换为map对象
func JsonToObject(str string) (data map[string]interface{}, err error) {
	//	buf := &bytes.Buffer{}
	//	buf.WriteString(str)
	//	decoder := json.NewDecoder(buf)
	decoder := json.NewDecoder(strings.NewReader(str))
	err = decoder.Decode(&data)
	return
}

//将map对象转为Json字符串
func JsonFromObject(obj *map[string]interface{}) (str string, err error) {
	buf := &bytes.Buffer{}

	encoder := json.NewEncoder(buf)
	err = encoder.Encode(*obj)
	if err == nil {
		str = buf.String()
	}
	return
}

func JsonGetObject(path []string, data *map[string]interface{}) interface{} {
	var tag interface{}
	mapObj := *data

	for idx := range path {
		name := string(path[idx])
		// log.Println("key:", name, idx)
		val, ok := mapObj[name]
		if ok {
			if len(path)-1 == idx {
				// log.Println("find value", name, val)
				tag = val
				break
			} else {
				v, isCorrectType := val.(map[string]interface{})
				if isCorrectType {
					mapObj = v
				}
			}
		}
	}
	return tag
}

func JsonGetValueString(path []string, def string, data *map[string]interface{}) string {
	var tag = JsonGetObject(path, data)
	if tag != nil {
		v, ok := tag.(string)
		if ok {
			return v
		} else {
			log.Println("Get value string:", GetValueType(tag))
		}
	}

	return def
}

func JsonGetValueInt(path []string, def int, data *map[string]interface{}) int {
	var tag = JsonGetObject(path, data)
	if tag != nil {
		v, ok := tag.(float64)
		if ok {
			return int(v)
		} else {
			log.Println("Value Int:", GetValueType(tag))
		}
	}
	return def
}

func JsonGetValueInt64(path []string, def int64, data *map[string]interface{}) int64 {
	var tag = JsonGetObject(path, data)
	if tag != nil {
		v, ok := tag.(float64)
		if ok {
			return int64(v)
		} else {
			log.Println("Value Int:", GetValueType(tag))
		}
	}
	return def
}

func JsonAddValue(key string, value interface{}, data *map[string]interface{}) {
	(*data)[key] = value
}
