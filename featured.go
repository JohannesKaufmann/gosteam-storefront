package storefront

import (
	"encoding/json"
	"net/http"
)

type App struct {
	ID                      json.Number `json:"id"`
	Type                    int         `json:"type"`
	Name                    string      `json:"name"`
	Discounted              bool        `json:"discounted"`
	DiscountPercent         int         `json:"discount_percent"`
	OriginalPrice           interface{} `json:"original_price"`
	FinalPrice              int         `json:"final_price"`
	Currency                string      `json:"currency"`
	LargeCapsuleImage       string      `json:"large_capsule_image"`
	SmallCapsuleImage       string      `json:"small_capsule_image"`
	WindowsAvailable        bool        `json:"windows_available"`
	MacAvailable            bool        `json:"mac_available"`
	LinuxAvailable          bool        `json:"linux_available"`
	StreamingvideoAvailable bool        `json:"streamingvideo_available"`
	DiscountExpiration      int         `json:"discount_expiration,omitempty"`
	HeaderImage             string      `json:"header_image"`
	Headline                string      `json:"headline"`
	ControllerSupport       string      `json:"controller_support,omitempty"`
	PurchasePackage         int         `json:"purchase_package,omitempty"`
	Body                    string      `json:"body,omitempty"`
	URL                     string      `json:"url,omitempty"`
}

func Featured() (map[string][]App, error) {
	url := "https://store.steampowered.com/api/featured/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var res map[string]*json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]App)
	for key, val := range res {
		// we dont care about these
		if key == "layout" || key == "status" {
			continue
		}

		var apps []App
		err = json.Unmarshal(*val, &apps)
		if err != nil {
			return nil, err
		}
		result[key] = apps
	}

	return result, nil
}
