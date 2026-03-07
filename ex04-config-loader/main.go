package main

import (
	"fmt"
	"os"
	"io"
	"encoding/json"
	"strconv"
	"reflect"
)

type Config struct {
	Port int `json:"port" validate:"1-65535"`
	TimeoutMs int `json:"timeout_ms" validate:"pos"`
	LogLevel string `json:"log_level" validate:"notempty"`
	AppPort int `json:"APP_PORT""`
}

func AutoValidate(s interface{}) error {
	value := reflect.ValueOf(s)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	typ := value.Type()
	for i:=0;i<typ.NumField();i++ {
		field := typ.Field(i)
		v := value.Field(i)
		tag := field.Tag.Get("validate")
		if tag == ""{
			continue
		}
		switch v.Kind() {
		case reflect.Int:
			val := v.Int()
			if tag == "pos" && val <= 0 {
				return fmt.Errorf("[%s] must be a positive", field.Name)
			}
			if tag == "1-65535" && (val < 1 || val > 65535) {
				return fmt.Errorf("[%s] must be between 1 and 65535", field.Name)
			}
		case reflect.String:
			if tag == "notempty" && v.String() == "" {
				return fmt.Errorf("[%s] cannnot be empty", field.Name)
			}
		}
	}
	return nil
}

func main() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Cannot open config file: ", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Cannot read json data: ", err)
		return
	}
	var config Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		fmt.Println("Cannot parse json data: ", err)
		return
	}

	if envPort, ok := os.LookupEnv("APP_PORT"); ok {
		if p, err := strconv.Atoi(envPort); err == nil {
			config.Port = p
		}
	}
	err = AutoValidate(&config)
	if err != nil {
		fmt.Println("Config validation error: ", err)
		return
	}
}