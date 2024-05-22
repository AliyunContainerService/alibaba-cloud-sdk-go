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

// CreateFixedPriceDemandOrder invokes the domain.CreateFixedPriceDemandOrder API synchronously
func (client *Client) CreateFixedPriceDemandOrder(request *CreateFixedPriceDemandOrderRequest) (response *CreateFixedPriceDemandOrderResponse, err error) {
	response = CreateCreateFixedPriceDemandOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFixedPriceDemandOrderWithChan invokes the domain.CreateFixedPriceDemandOrder API asynchronously
func (client *Client) CreateFixedPriceDemandOrderWithChan(request *CreateFixedPriceDemandOrderRequest) (<-chan *CreateFixedPriceDemandOrderResponse, <-chan error) {
	responseChan := make(chan *CreateFixedPriceDemandOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFixedPriceDemandOrder(request)
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

// CreateFixedPriceDemandOrderWithCallback invokes the domain.CreateFixedPriceDemandOrder API asynchronously
func (client *Client) CreateFixedPriceDemandOrderWithCallback(request *CreateFixedPriceDemandOrderRequest, callback func(response *CreateFixedPriceDemandOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFixedPriceDemandOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateFixedPriceDemandOrder(request)
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

// CreateFixedPriceDemandOrderRequest is the request struct for api CreateFixedPriceDemandOrder
type CreateFixedPriceDemandOrderRequest struct {
	*requests.RpcRequest
	Code      string `position:"Query" name:"Code"`
	ContactId string `position:"Query" name:"ContactId"`
	Source    string `position:"Query" name:"Source"`
	Domain    string `position:"Query" name:"Domain"`
}

// CreateFixedPriceDemandOrderResponse is the response struct for api CreateFixedPriceDemandOrder
type CreateFixedPriceDemandOrderResponse struct {
	*responses.BaseResponse
	RequestId      string                              `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int                                 `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorCode      string                              `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool                                `json:"Success" xml:"Success"`
	Module         ModuleInCreateFixedPriceDemandOrder `json:"Module" xml:"Module"`
}

// CreateCreateFixedPriceDemandOrderRequest creates a request to invoke CreateFixedPriceDemandOrder API
func CreateCreateFixedPriceDemandOrderRequest() (request *CreateFixedPriceDemandOrderRequest) {
	request = &CreateFixedPriceDemandOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "CreateFixedPriceDemandOrder", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateFixedPriceDemandOrderResponse creates a response to parse from CreateFixedPriceDemandOrder response
func CreateCreateFixedPriceDemandOrderResponse() (response *CreateFixedPriceDemandOrderResponse) {
	response = &CreateFixedPriceDemandOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
