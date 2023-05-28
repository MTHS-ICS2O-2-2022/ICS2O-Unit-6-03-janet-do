// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file lets you check the weather

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func weather2(urlAddress string) {
	response, err := http.Get(urlAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var jsonData WeatherData
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	temperature1 := jsonData.Main.Temp
	temperature2 := temperature1 - 273.15

	fmt.Printf("The temperature outside is %.2f °C!\n", temperature2)
	fmt.Println("\nDone.")
}

func main() {
	urlAddress := "https://api.openweathermap.org/data/2.5/weather?lat=45.4211435&lon=-75.6900574&appid=fe1d80e1e103cff8c6afd190cad23fa5"
	weather2(urlAddress)
}


