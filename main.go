package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Remove name, add NumRecords
type CityRecords struct {
	NumRecords float64
	Min        float64
	Mean       float64
	Max        float64
}

func updateCity(m map[string]CityRecords, name string, temp float64) {
	if records, ok := m[name]; ok {
		if temp < records.Min {
			records.Min = temp
		}
		if temp > records.Max {
			records.Max = temp
		}
		records.Mean += temp
		records.NumRecords += 1.0
	} else {
		m[name] = CityRecords{
			NumRecords: 1.0,
			Min:        temp,
			Mean:       temp,
			Max:        temp,
		}
	}

}

func main() {
	cityTemperatures := map[string]CityRecords{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ";")
		name := text[0]
		temp, err := strconv.ParseFloat(text[1], 64)
		if err != nil {
			fmt.Printf("ERROR converting float: %v", err)
			os.Exit(1)
		}
		//Race condition for concurrent writes to map
		updateCity(cityTemperatures, name, temp)
	}

	cityNames := make([]string, 0, len(cityTemperatures))
	for name, records := range cityTemperatures {
		cityNames = append(cityNames, name)
		records.Mean = records.Mean / records.NumRecords
	}

	slices.Sort(cityNames)
	fmt.Print("{")
	for i, name := range cityNames {
		records := cityTemperatures[name]
		//Expected output: Abha=-23.0/18.0/59.2
		fmt.Printf("%s=%.1f/%.1f/%.1f", name, records.Min, records.Mean, records.Max)
		if i < len(cityNames)-1 {
			fmt.Print(", ")
		}

	}
	fmt.Print("}\n")
}
