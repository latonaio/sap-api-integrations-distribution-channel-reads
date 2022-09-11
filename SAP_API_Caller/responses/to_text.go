package responses

type ToText struct {
	D struct {
		Results []struct {
			DistributionChannel     string `json:"DistributionChannel"`
			Language         		string `json:"Language"`
			DistributionChannelName string `json:"DistributionChannelName"`
		} `json:"results"`
	} `json:"d"`
}
