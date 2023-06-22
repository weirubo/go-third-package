package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	//v.SetConfigFile("./configs/test.yaml")
	v.SetConfigFile("./configs/default.yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	//err = v.Unmarshal(&userData)
	err = v.UnmarshalKey("user_data", &userData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("userData=%+v\n", userData)
}

type UserData struct {
	Uid       int        `json:"uid"`
	Uname     string     `json:"uname"`
	OtherInfo OtherInfo  `json:"other_info" mapstructure:"other_info"`
	Language  []Language `json:"language" mapstructure:"language"`
}

type OtherInfo struct {
	Email   string
	Address string
}

type Language struct {
	Name  string
	Score int
}

var userData UserData
