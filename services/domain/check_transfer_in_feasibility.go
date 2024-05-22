package domain

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

// CheckTransferInFeasibility invokes the domain.CheckTransferInFeasibility API synchronously
func (client *Client) CheckTransferInFeasibility(request *CheckTransferInFeasibilityRequest) (response *CheckTransferInFeasibilityResponse, err error) {
	response = CreateCheckTransferInFeasibilityResponse()
	err = client.DoAction(request, response)
	return
}

// CheckTransferInFeasibilityWithChan invokes the domain.CheckTransferInFeasibility API asynchronously
func (client *Client) CheckTransferInFeasibilityWithChan(request *CheckTransferInFeasibilityRequest) (<-chan *CheckTransferInFeasibilityResponse, <-chan error) {
	responseChan := make(chan *CheckTransferInFeasibilityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckTransferInFeasibility(request)
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

// CheckTransferInFeasibilityWithCallback invokes the domain.CheckTransferInFeasibility API asynchronously
func (client *Client) CheckTransferInFeasibilityWithCallback(request *CheckTransferInFeasibilityRequest, callback func(response *CheckTransferInFeasibilityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckTransferInFeasibilityResponse
		var err error
		defer close(result)
		response, err = client.CheckTransferInFeasibility(request)
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

// CheckTransferInFeasibilityRequest is the request struct for api CheckTransferInFeasibility
type CheckTransferInFeasibilityRequest struct {
	*requests.RpcRequest
	DomainName                string `position:"Query" name:"DomainName"`
	TransferAuthorizationCode string `position:"Query" name:"TransferAuthorizationCode"`
	UserClientIp              string `position:"Query" name:"UserClientIp"`
	Lang                      string `position:"Query" name:"Lang"`
}

// CheckTransferInFeasibilityResponse is the response struct for api CheckTransferInFeasibility
type CheckTransferInFeasibilityResponse struct {
	*responses.BaseResponse
	CanTransfer bool   `json:"CanTransfer" xml:"CanTransfer"`
	Message     string `json:"Message" xml:"Message"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	ProductId   string `json:"ProductId" xml:"ProductId"`
	Code        string `json:"Code" xml:"Code"`
}

// CreateCheckTransferInFeasibilityRequest creates a request to invoke CheckTransferInFeasibility API
func CreateCheckTransferInFeasibilityRequest() (request *CheckTransferInFeasibilityRequest) {
	request = &CheckTransferInFeasibilityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "CheckTransferInFeasibility", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckTransferInFeasibilityResponse creates a response to parse from CheckTransferInFeasibility response
func CreateCheckTransferInFeasibilityResponse() (response *CheckTransferInFeasibilityResponse) {
	response = &CheckTransferInFeasibilityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
