package apiClients

import (
	"encoding/json"
	"fmt"
	"footbal-api-parser/models"
	"os"

	"strconv"

	"github.com/olekukonko/tablewriter"
)

// func Team() {
// 	url := "https://v3.football.api-sports.io/teams"
// 	fetchData := FetchData(url)
// 	var response models.TeamResponse
// 	err := json.Unmarshal(fetchData, &response)
// 	if err != nil {
// 		fmt.Printf("Error unmarshaling JSON: %v\n", err)
// 		return
// 	}

// 	if len(response.Response) == 0 {
// 		fmt.Println("No teams found")
// 		return
// 	}

// 	for _, team := range response.Response {
// 		fmt.Printf("Team Name: %s\n", team.Name)
// 		fmt.Printf("Team ID: %d\n", team.ID)
// 		fmt.Printf("Venue Name: %s\n", team.Venue.Name)
// 		fmt.Printf("Capacity: %d\n", team.Venue.Capacity)
// 		fmt.Println("---")
// 	}
// }

func Get() {
	response := models.APIResponse{}
	url := "https://v3.football.api-sports.io/leagues"
	fetchData := FetchData(url)
	err := json.Unmarshal(fetchData, &response)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}

	uniqueCountries := make(map[string]bool)
	countries := []string{}

	for _, respCountry := range response.Response {
		countryName := respCountry.Country.Name
		if _, exists := uniqueCountries[countryName]; !exists {
			uniqueCountries[countryName] = true
			countries = append(countries, countryName)
		}
	}

	Countries(countries)

	fmt.Print("\nEnter the ID of the chosen country 🌍 (or 'q' to quit): ")
	var choice string
	fmt.Scanln(&choice)

	if choice == "" || choice == "q" || choice == "quit" {
		return
	}

	choiceInt, err := strconv.Atoi(choice)
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		return
	}

	if choiceInt <= 0 || choiceInt > len(countries) {
		fmt.Println("Invalid country number.")
		return
	}

	selectedCountry := countries[choiceInt-1]
	Leagues(&response, selectedCountry)

	fmt.Print("\nEnter the ID of the chosen league 🏆: ")
	var choiceLeague string
	fmt.Scanln(&choiceLeague)

	if choiceLeague == "" || choiceLeague == "q" || choiceLeague == "quit" {
		return
	}

	choiceInt, err = strconv.Atoi(choiceLeague)
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		return
	}

	if choiceInt <= 0 || choiceInt > len(response.Response) {
		fmt.Println("Invalid league number.")
		return
	}
	selectedLeague := &response.Response[choiceInt-1]
	fmt.Printf("\n%s Years: \n", selectedLeague.League.Name)

	if selectedLeague.Seasons != nil {
		Table := tablewriter.NewWriter(os.Stdout)
		Table.SetHeader([]string{"ID", "Year"})
		Table.SetAlignment(tablewriter.ALIGN_LEFT)
		for i, season := range selectedLeague.Seasons {
			Table.Append([]string{strconv.Itoa(i + 1), strconv.Itoa(season.Year)})
		}
		Table.Render()

	} else {
		fmt.Println("No seasons found for this league.")
	}

	fmt.Print("\nEnter the ID of the chosen year 🗓️: ")
	var choiceYear string
	fmt.Scanln(&choiceYear)

	if choiceYear == "" || choiceYear == "q" || choiceYear == "quit" {
		return
	}

	choiceInt, err = strconv.Atoi(choiceYear)
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		return
	}

	if choiceInt <= 0 || choiceInt > len(selectedLeague.Seasons) {
		fmt.Println("Invalid year number.")
		return
	}

	selectedYear := &selectedLeague.Seasons[choiceInt-1]

	fmt.Printf("\nYear: %d\n", selectedYear.Year)

	SeasonDetails(selectedYear)
}
