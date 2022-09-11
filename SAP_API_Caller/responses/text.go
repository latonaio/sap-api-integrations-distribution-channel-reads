package responses

type Text struct {
	D struct {
		Results []struct {
			DistributionChannel     string `json:"DistributionChannel"`
			Language         		string `json:"Language"`
			DistributionChannelName string `json:"DistributionChannelName"`
		} `json:"results"`
	} `json:"d"`
}
