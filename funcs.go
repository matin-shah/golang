package apiClients

import (
	"fmt"
	"footbal-api-parser/models"
	"os"
	"sort"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func Countries(countries []string) {
	fmt.Println("List countries 🌍:")
	sort.Strings(countries)
	for i, country := range countries {
		fmt.Printf("%d - %s\n", i+1, country)
	}
}

func Leagues(response *models.APIResponse, selectedCountry string) {
	fmt.Printf("\n%s leagues:\n", selectedCountry)
	printTable := tablewriter.NewWriter(os.Stdout)
	printTable.SetHeader([]string{"ID", "League"})
	printTable.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, respLeague := range response.Response {
		if respLeague.Country.Name == selectedCountry {
			printTable.Append([]string{
				strconv.Itoa(respLeague.League.ID),
				respLeague.League.Name,
			})
		}
	}
	printTable.Render()
	fmt.Println()
}

func SeasonDetails(selectedYear *models.LeagueSeason) {
	printTable := tablewriter.NewWriter(os.Stdout)
	printTable.SetHeader([]string{"Information", "Value"})

	printTable.Append([]string{"Start", selectedYear.Start})
	printTable.Append([]string{"End", selectedYear.End})

	coverage := []struct {
		name  string
		value interface{}
	}{
		{"Events", selectedYear.Coverage.Fixtures.Events},
		{"Lineups", selectedYear.Coverage.Fixtures.Lineups},
		{"Statistics Fixtures", selectedYear.Coverage.Fixtures.StatisticsFixtures},
		{"Statistics Players", selectedYear.Coverage.Fixtures.StatisticsPlayers},
		{"Standings", selectedYear.Coverage.Fixtures.Standings},
		{"Players", selectedYear.Coverage.Fixtures.Players},
		{"Top Scorers", selectedYear.Coverage.Fixtures.TopScorers},
		{"Top Assists", selectedYear.Coverage.Fixtures.TopAssists},
		{"Top Cards", selectedYear.Coverage.Fixtures.TopCards},
		{"Injuries", selectedYear.Coverage.Fixtures.Injuries},
		{"Predictions", selectedYear.Coverage.Fixtures.Predictions},
		{"Odds", selectedYear.Coverage.Fixtures.Odds},
	}

	for _, info := range coverage {
		printTable.Append([]string{info.name, fmt.Sprintf("%t", info.value)})
	}

	printTable.Render()
}
