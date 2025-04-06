package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type CityResults struct {
	Name string
	Min  float64
	Mean float64
	Max  float64
}

func main() {
	cities := map[string][]float64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ";")
		temp, err := strconv.ParseFloat(text[1], 64)
		if err != nil {
			fmt.Printf("ERROR converting float: %v", err)
			os.Exit(1)
		}
		cities[text[0]] = append(cities[text[0]], temp)
	}
	cityResults := []CityResults{}
	for city, temps := range cities {
		var min, sum, max float64 = 0.0, 0.0, 0.0
		for i, temp := range temps {
			if i == 0 {
				min = temp
			} else if temp < min {
				min = temp
			}
			if temp > max {
				max = temp
			}
			sum += temp
		}
		cityResults = append(cityResults, CityResults{
			Name: city,
			Min:  min,
			Mean: sum / float64(len(temps)),
			Max:  max,
		})
	}

	slices.SortStableFunc(cityResults, func(i, j CityResults) int {
		return strings.Compare(i.Name, j.Name)
	})
	fmt.Print("{")
	for i, v := range cityResults {
		//Expected output: Abha=-23.0/18.0/59.2
		fmt.Printf("%s=%.1f/%.1f/%.1f", v.Name, v.Min, v.Mean, v.Max)
		if i != len(cityResults)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}\n")
}
