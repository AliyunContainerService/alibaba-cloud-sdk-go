package dypnsapi

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

// QueryGateVerifyStatisticPublic invokes the dypnsapi.QueryGateVerifyStatisticPublic API synchronously
func (client *Client) QueryGateVerifyStatisticPublic(request *QueryGateVerifyStatisticPublicRequest) (response *QueryGateVerifyStatisticPublicResponse, err error) {
	response = CreateQueryGateVerifyStatisticPublicResponse()
	err = client.DoAction(request, response)
	return
}

// QueryGateVerifyStatisticPublicWithChan invokes the dypnsapi.QueryGateVerifyStatisticPublic API asynchronously
func (client *Client) QueryGateVerifyStatisticPublicWithChan(request *QueryGateVerifyStatisticPublicRequest) (<-chan *QueryGateVerifyStatisticPublicResponse, <-chan error) {
	responseChan := make(chan *QueryGateVerifyStatisticPublicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryGateVerifyStatisticPublic(request)
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

// QueryGateVerifyStatisticPublicWithCallback invokes the dypnsapi.QueryGateVerifyStatisticPublic API asynchronously
func (client *Client) QueryGateVerifyStatisticPublicWithCallback(request *QueryGateVerifyStatisticPublicRequest, callback func(response *QueryGateVerifyStatisticPublicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryGateVerifyStatisticPublicResponse
		var err error
		defer close(result)
		response, err = client.QueryGateVerifyStatisticPublic(request)
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

// QueryGateVerifyStatisticPublicRequest is the request struct for api QueryGateVerifyStatisticPublic
type QueryGateVerifyStatisticPublicRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AuthenticationType   requests.Integer `position:"Query" name:"AuthenticationType"`
	StartDate            string           `position:"Query" name:"StartDate"`
	SceneCode            string           `position:"Query" name:"SceneCode"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ProdCode             string           `position:"Query" name:"ProdCode"`
	OsType               string           `position:"Query" name:"OsType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	EndDate              string           `position:"Query" name:"EndDate"`
}

// QueryGateVerifyStatisticPublicResponse is the response struct for api QueryGateVerifyStatisticPublic
type QueryGateVerifyStatisticPublicResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryGateVerifyStatisticPublicRequest creates a request to invoke QueryGateVerifyStatisticPublic API
func CreateQueryGateVerifyStatisticPublicRequest() (request *QueryGateVerifyStatisticPublicRequest) {
	request = &QueryGateVerifyStatisticPublicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "QueryGateVerifyStatisticPublic", "dypnsapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryGateVerifyStatisticPublicResponse creates a response to parse from QueryGateVerifyStatisticPublic response
func CreateQueryGateVerifyStatisticPublicResponse() (response *QueryGateVerifyStatisticPublicResponse) {
	response = &QueryGateVerifyStatisticPublicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
