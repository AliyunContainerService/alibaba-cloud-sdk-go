package apds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSurveyJobCount invokes the apds.DescribeSurveyJobCount API synchronously
func (client *Client) DescribeSurveyJobCount(request *DescribeSurveyJobCountRequest) (response *DescribeSurveyJobCountResponse, err error) {
	response = CreateDescribeSurveyJobCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSurveyJobCountWithChan invokes the apds.DescribeSurveyJobCount API asynchronously
func (client *Client) DescribeSurveyJobCountWithChan(request *DescribeSurveyJobCountRequest) (<-chan *DescribeSurveyJobCountResponse, <-chan error) {
	responseChan := make(chan *DescribeSurveyJobCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSurveyJobCount(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeSurveyJobCountWithCallback invokes the apds.DescribeSurveyJobCount API asynchronously
func (client *Client) DescribeSurveyJobCountWithCallback(request *DescribeSurveyJobCountRequest, callback func(response *DescribeSurveyJobCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSurveyJobCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeSurveyJobCount(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeSurveyJobCountRequest is the request struct for api DescribeSurveyJobCount
type DescribeSurveyJobCountRequest struct {
	*requests.RoaRequest
}

// DescribeSurveyJobCountResponse is the response struct for api DescribeSurveyJobCount
type DescribeSurveyJobCountResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"error" xml:"error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateDescribeSurveyJobCountRequest creates a request to invoke DescribeSurveyJobCount API
func CreateDescribeSurveyJobCountRequest() (request *DescribeSurveyJobCountRequest) {
	request = &DescribeSurveyJobCountRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "DescribeSurveyJobCount", "/okss-services/winback/count-survey-job", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSurveyJobCountResponse creates a response to parse from DescribeSurveyJobCount response
func CreateDescribeSurveyJobCountResponse() (response *DescribeSurveyJobCountResponse) {
	response = &DescribeSurveyJobCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
