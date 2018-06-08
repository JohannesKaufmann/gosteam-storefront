package storefront

import (
	"encoding/json"
	"net/http"
)

type Category struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Items []App  `json:"items"`
}

func FeaturedCategories() (map[string]Category, error) {
	url := "https://store.steampowered.com/api/featuredcategories/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var res map[string]*json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	result := make(map[string]Category)
	for key, val := range res {
		if key == "genres" || key == "trailerslideshow" || key == "status" {
			continue
		}

		var r Category
		err = json.Unmarshal(*val, &r)
		if err != nil {
			return nil, err
		}

		result[key] = r
	}

	return result, nil
}
