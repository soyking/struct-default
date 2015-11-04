package structdefault

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Default(data interface{}) error {
	original := reflect.ValueOf(data)
	return convertToDefault(original)
}

// 转换成默认类型
func convertToDefault(oValue reflect.Value) error {
	var oType reflect.Type

	switch oValue.Kind() {
	case reflect.Ptr:
		oValue = oValue.Elem()
	case reflect.Struct:
	default:
		return nil
	}
	oType = oValue.Type()
	for i := 0; i < oValue.NumField(); i++ {
		f := oValue.Field(i)
		var defaultValue string
		if defaultValue = oType.Field(i).Tag.Get("default"); defaultValue == "" {
			if f.Kind() != reflect.Ptr && f.Kind() != reflect.Struct {
				continue
			}
		}
		switch f.Kind() {
		case reflect.String:
			if defaultValue == "$uuid" {
				f.SetString(genUUID())
			}
			if f.String() == "" {
				f.SetString(defaultValue)
			}
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
			if strings.Contains(defaultValue, "$range") {
				if x, err := rangeIt(defaultValue); err != nil {
					return err
				} else {
					f.SetInt(x)
					continue
				}
			}
			if defaultValue == "$timeNowUnix" {
				f.SetInt(time.Now().Unix())
				continue
			}
			if defaultValue == "$timeNowUnixNano" {
				f.SetInt(time.Now().UnixNano())
				continue
			}
			if x, err := strconv.ParseInt(defaultValue, 10, 0); err == nil && f.Int() == 0 {
				f.SetInt(x)
			} else {
				return err
			}
		case reflect.Float32, reflect.Float64:
			if x, err := strconv.ParseFloat(defaultValue, 10); err == nil && f.Float() == 0 {
				f.SetFloat(x)
			} else {
				return err
			}
		case reflect.Bool:
			if x, err := strconv.ParseBool(defaultValue); err == nil && f.Bool() == false {
				f.SetBool(x)
			} else {
				return err
			}
		case reflect.Struct:
			if err := convertToDefault(f.Addr()); err != nil {
				return err
			}
		case reflect.Ptr:
			if err := convertToDefault(f); err != nil {
				return err
			}
		}
	}
	return nil
}
