package options

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var CamelCaseRegexp, _ = regexp.Compile("[A-Z][a-z]+")

func SplitCamelCase(s string) (ret []string) {
	start := 0
	result := CamelCaseRegexp.FindAllStringIndex(s, -1)
	for _, loc := range result {
		if start != loc[0] {
			ret = append(ret, s[start:loc[0]])
		}
		ret = append(ret, s[loc[0]:loc[1]])
		start = loc[1]
	}
	return ret
}

var nilStruct struct{}

func Parse(v interface{}, parseEnv, parseFlag bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	typ := rv.Type()
	envMap := make(map[int]string)
	flagMap := make(map[int]*string)
	requiredMap := make(map[int]string)
	// typ := reflect.ValueOf(v).Type()
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		envKey := f.Tag.Get("env")
		flagKey := f.Tag.Get("flag")
		optionSet := make(map[string]struct{})
		for _, option := range strings.Split(f.Tag.Get("options"), " ") {
			optionSet[option] = nilStruct
		}
		if f.Tag.Get("required") != "" {
			optionSet["required"] = nilStruct
		}
		if _, auto := optionSet["auto"]; auto {
			if envKey == "" {
				envKey = strings.ToUpper(strings.Join(SplitCamelCase(f.Name), "_"))
			}
			if flagKey == "" {
				flagKey = strings.ToLower(strings.Join(SplitCamelCase(f.Name), "-"))
			}
		}
		if parseEnv && envKey != "" && os.Getenv(envKey) != "" {
			envMap[i] = os.Getenv(envKey)
		}
		if parseFlag && flagKey != "" {
			flagMap[i] = flag.String(flagKey, "", f.Tag.Get("usage"))
		}
		if _, required := optionSet["required"]; required {
			requiredMap[i] = f.Name
		}
	}
	if parseFlag {
		flag.Parse()
	}
	for i := 0; i < typ.NumField(); i++ {
		var val string
		if s, exists := envMap[i]; exists {
			val = s
		}
		if s, exists := flagMap[i]; exists && *s != "" {
			val = *s
		}
		if val != "" {
			set(rv.FieldByIndex([]int{i}), val)
		} else if fieldName, required := requiredMap[i]; required {
			if rv.Field(i).Kind() == reflect.String && rv.Field(i).String() == "" {
				panic(fmt.Sprintf("%s is required", fieldName))
			}
		}
	}
}

func set(v reflect.Value, s string) {
	switch v.Kind() {
	case reflect.String:
		v.SetString(s)
	case reflect.Bool:
		s = strings.ToLower(s)
		if s == "true" || s == "yes" || s == "1" {
			v.SetBool(true)
		} else {
			v.SetBool(false)
		}
	case reflect.Uint, reflect.Uint32:
		if i, err := strconv.ParseUint(s, 0, 32); err == nil {
			v.SetUint(i)
		}
	case reflect.Uint8:
		if i, err := strconv.ParseUint(s, 0, 8); err == nil {
			v.SetUint(i)
		}
	case reflect.Uint16:
		if i, err := strconv.ParseUint(s, 0, 16); err == nil {
			v.SetUint(i)
		}
	case reflect.Uint64:
		if i, err := strconv.ParseUint(s, 0, 64); err == nil {
			v.SetUint(i)
		}
	case reflect.Int, reflect.Int32:
		if i, err := strconv.ParseInt(s, 0, 32); err == nil {
			v.SetInt(i)
		}
	case reflect.Int8:
		if i, err := strconv.ParseInt(s, 0, 8); err == nil {
			v.SetInt(i)
		}
	case reflect.Int16:
		if i, err := strconv.ParseInt(s, 0, 16); err == nil {
			v.SetInt(i)
		}
	case reflect.Int64:
		if i, err := strconv.ParseInt(s, 0, 64); err == nil {
			v.SetInt(i)
		}
	case reflect.Float32:
		if i, err := strconv.ParseFloat(s, 32); err == nil {
			v.SetFloat(i)
		}
	case reflect.Float64:
		if i, err := strconv.ParseFloat(s, 64); err == nil {
			v.SetFloat(i)
		}
	}
}
