/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long: `This command fetches a random dad joke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		jokeTerm, _ := cmd.Flags().GetString("term")
		if jokeTerm != "" {
			getRandomJokeWithTerm(jokeTerm)
		} else {
			getRandomJoke()
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.PersistentFlags().String("term", "", "A search term for a dad joke.")
}

type Joke struct {
	ID string `json:"id"`
	Joke string `json:"joke"`
	Status int `json:"status"`
}

type SearchResult struct {
	Results json.RawMessage `json:"results"`
	SearchTerm string `json:"search_term"`
	Status int `json:"status"`
	TotalJokes int `json:"total_jokes"`
}

func getRandomJokeWithTerm(jokeTerm string) {
	total, results := getJokesDataWithTerm(jokeTerm)
	randomiseJokeList(total, results)
}

func randomiseJokeList(length int, jokelist []Joke) {
	rand.Seed(time.Now().Unix())

	min := 0
	max := length - 1

	if length <= 0 {
		err := fmt.Errorf("no jokes found with this term")
		fmt.Println(err.Error())
	} else {
		randomNum := min + rand.Intn(max - min)
		fmt.Println(jokelist[randomNum].Joke)
	}
}

func getJokesDataWithTerm(jokeTerm string) (int, []Joke) {
	url := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", jokeTerm)
	responseBytes := getJokeData(url)

	var jokeListRaw SearchResult
	if err := json.Unmarshal(responseBytes, &jokeListRaw); err != nil {
		log.Printf("Could not unmarshal resonseBytes - %v", err)
	}

	var jokes []Joke
	if err := json.Unmarshal(jokeListRaw.Results, &jokes); err != nil {
		log.Printf("Could not unmarshal jokeListRaw results - %v", err)
	}
	return len(jokes), jokes
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	var joke Joke
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	fmt.Println(joke.Joke)
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(http.MethodGet, baseAPI, nil)
	if err != nil {
		log.Printf("Could not request a dadjoke - %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/example.dadjoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request - %v", err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body - %v", err)
	}

	return responseBody
}