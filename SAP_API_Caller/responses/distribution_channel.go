package responses

type DistributionChannel struct {
	D struct {
		Count   string `json:"__count"`
		Results []struct {
			DistributionChannel string `json:"DistributionChannel"`
			ToText       struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Text"`
		} `json:"results"`
	} `json:"d"`
}
