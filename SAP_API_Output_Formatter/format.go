package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-distribution-channel-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToDistributionChannel(raw []byte, l *logger.Logger) ([]DistributionChannel, error) {
	pm := &responses.DistributionChannel{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to DistributionChannel. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	distributionChannel := make([]DistributionChannel, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		distributionChannel = append(distributionChannel, DistributionChannel{
			DistributionChannel:         data.DistributionChannel,
			ToText:                      data.ToText.Deferred.URI,
		})
	}

	return distributionChannel, nil
}

func ConvertToText(raw []byte, l *logger.Logger) ([]Text, error) {
	pm := &responses.Text{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Text. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	text := make([]Text, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		text = append(text, Text{
			DistributionChannel:    	data.DistributionChannel,
			Language:           		data.Language,
			DistributionChannelName:	data.DistributionChannelName,
		})
	}

	return text, nil
}