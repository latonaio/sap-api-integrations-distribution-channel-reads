package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-distribution-channel-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetDistributionChannel(distributionChannel, language, distributionChannelName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "DistributionChannel":
			func() {
				c.DistributionChannel(distributionChannel)
				wg.Done()
			}()
		case "Text":
			func() {
				c.Text(language, distributionChannelName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) DistributionChannel(distributionChannel string) {
	distributionChannelData, err := c.callDistributionChannelSrvAPIRequirementDistributionChannel("A_DistributionChannel", distributionChannel)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(distributionChannelData)

	textData, err := c.callToText(distributionChannelData[0].ToText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(textData)
}

func (c *SAPAPICaller) callDistributionChannelSrvAPIRequirementDistributionChannel(api, distributionChannel string) ([]sap_api_output_formatter.DistributionChannel, error) {
	url := strings.Join([]string{c.baseURL, "API_DISTRIBUTIONCHANNEL_SRV", api}, "/")
	param := c.getQueryWithDistributionChannel(map[string]string{}, distributionChannel)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToDistributionChannel(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToText(url string) ([]sap_api_output_formatter.Text, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Text(language, distributionChannelName string) {
	data, err := c.callDistributionChannelSrvAPIRequirementText("A_DistributionChannelText", language, distributionChannelName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callDistributionChannelSrvAPIRequirementText(api, language, distributionChannelName string) ([]sap_api_output_formatter.Text, error) {
	url := strings.Join([]string{c.baseURL, "API_DISTRIBUTIONCHANNEL_SRV", api}, "/")

	param := c.getQueryWithText(map[string]string{}, language, distributionChannelName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithDistributionChannel(params map[string]string, distributionChannel string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("DistributionChannel eq '%s'", distributionChannel)
	return params
}

func (c *SAPAPICaller) getQueryWithText(params map[string]string, language, distributionChannelName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Language eq '%s' and substringof('%s', DistributionChannelName)", language, distributionChannelName)
	return params
}
