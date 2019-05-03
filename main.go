package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type State struct {
	Label            string
	Name             string
	IntegrationValue string
	Country          string
}

func main() {
	f, err := os.Open("mx_states.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // this needs to be after the err check

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	allStates := []*State{}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := State{
			Label:            line[0],
			Name:             line[1],
			IntegrationValue: line[2],
			Country:          line[3],
		}
		allStates = append(allStates, &data)
	}

	for i, state := range allStates {
		if i == 0 { // skip first line of .csv (header)
			continue
		}
		fileString := `<?xml version="1.0" encoding="UTF-8"?>
<CustomMetadata xmlns="http://soap.sforce.com/2006/04/metadata" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
	<label>` + state.Label + `</label>
	<values>
		<field>Integration_Value__c</field>
		<value xsi:type="xsd:string">` + state.IntegrationValue + `</value>
	</values>
	<values>
		<field>Country__c</field>
		<value xsi:type="xsd:string">` + state.Country + `</value>
	</values>
</CustomMetadata>`
		fmt.Println(fileString)
		f, err := os.Create("mx state metadata/State_Form_Value." + state.Name + ".md")
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := f.WriteString(fileString)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l)
	}
}
