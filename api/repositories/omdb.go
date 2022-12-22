package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetOmdbMovies(title string, c chan MovieCollection) {
	omdb_key := os.Getenv("OMDB_API_KEY")
	res, _ := http.Get("https://www.omdbapi.com/?s=" + title + "&page=1&type=movie&apikey=" + omdb_key)
	var search SearchResponse
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body))
	json.NewDecoder(res.Body).Decode(&search)
	m := MovieCollection{search.Search}
	c <- m
}

type SearchResponse struct {
	Search []interface{}
}
