package sap_api_output_formatter

type DistributionChannelReads struct {
	ConnectionKey 		string `json:"connection_key"`
	Result        		bool   `json:"result"`
	RedisKey      		string `json:"redis_key"`
	Filepath      		string `json:"filepath"`
	Product       		string `json:"Product"`
	APISchema     		string `json:"api_schema"`
	DistributionChannel string `json:"distribution_channel"`
	Deleted       		string `json:"deleted"`
}

type DistributionChannel struct {
	DistributionChannel string `json:"DistributionChannel"`
	ToText       		string `json:"to_Text"`
}

type Text struct {
	DistributionChannel     string `json:"DistributionChannel"`
	Language         		string `json:"Language"`
	DistributionChannelName string `json:"DistributionChannelName"`
}
