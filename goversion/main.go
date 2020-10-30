package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Country struct {
	Name       string `json:name`
	Population int    `json:population`
}

type Response struct {
	PerPage    int       `json:per_page`
	Total      int       `json:total`
	TotalPages int       `json:total_pages`
	Data       []Country `json:data`
}

func readJSON(resp *http.Response) (Response, error) {
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	var data Response
	err = json.Unmarshal(res, &data)

	if err != nil {
		return Response{}, err
	}

	return data, nil

}

// This would be a helper function that I would have in a util package
func getEnvVariavble(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

var base string = getEnvVariavble("url")

// This would also be a util function that I would have abstracted somewhere else so it can be used with other api calls
func GetTotalPages(param string) (int, error) {
	url := base + "?name=" + param
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := readJSON(res)
	if err != nil {
		return -1, err
	}

	totalPages := data.TotalPages
	return totalPages, nil

}

func getCountries(s string, p int) (int, error) {
	totalPop := 0
	totalPages, err := GetTotalPages(s)
	if err != nil {
		return -1, err
	}

	for i := 0; i <= totalPages; i++ {
		url := fmt.Sprintf("%s?name=%s&page=%d", base, s, i)
		fmt.Println(url)
		res, err := http.Get(url)
		response, err := readJSON(res)
		if err != nil {
			return -1, err
		}
		tot := response.Total
		for j := 0; j < tot; j++ {
			pop := response.Data[j].Population
			if pop > p {
				totalPop++
			}
		}
	}
	return totalPop, nil
}

func main() {

	pop, err := getCountries("un", 100090)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pop)
}
