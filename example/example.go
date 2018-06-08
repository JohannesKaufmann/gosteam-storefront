package main

import (
	"fmt"
	"log"

	storefront "github.com/JohannesKaufmann/gosteam-storefront"
)

func main() {
	// - - - - featured - - - - //
	featured, err := storefront.Featured()
	if err != nil {
		log.Fatal(err)
	}

	for key, apps := range featured {
		fmt.Printf("\n - - %s - - \n", key)
		for i, app := range apps {
			fmt.Printf("%d -> id:%s name:%s discount:%d headline:%s \n", i, app.ID, app.Name, app.DiscountPercent, app.Headline)
		}
	}

	fmt.Println("==================")

	// - - - - featured categories - - - - //
	categories, err := storefront.FeaturedCategories()
	if err != nil {
		log.Fatal(err)
	}

	for key, category := range categories {
		fmt.Printf("\n - - %s: %s - - \n", key, category.Name)

		for i, app := range category.Items {
			fmt.Printf("%d -> id:%s name:%s discount:%d headline:%s \n", i, app.ID, app.Name, app.DiscountPercent, app.Headline)
		}
	}
}
