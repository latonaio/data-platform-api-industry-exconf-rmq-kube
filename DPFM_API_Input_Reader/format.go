package dpfm_api_input_reader

import (
	"data-platform-api-industry-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToIndustry() *requests.Industry {
	data := sdc.Industry
	return &requests.Industry{
		Industry: data.Industry,
	}
}
