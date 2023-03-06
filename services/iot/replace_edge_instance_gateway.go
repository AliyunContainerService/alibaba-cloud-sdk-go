package iot

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

// ReplaceEdgeInstanceGateway invokes the iot.ReplaceEdgeInstanceGateway API synchronously
func (client *Client) ReplaceEdgeInstanceGateway(request *ReplaceEdgeInstanceGatewayRequest) (response *ReplaceEdgeInstanceGatewayResponse, err error) {
	response = CreateReplaceEdgeInstanceGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// ReplaceEdgeInstanceGatewayWithChan invokes the iot.ReplaceEdgeInstanceGateway API asynchronously
func (client *Client) ReplaceEdgeInstanceGatewayWithChan(request *ReplaceEdgeInstanceGatewayRequest) (<-chan *ReplaceEdgeInstanceGatewayResponse, <-chan error) {
	responseChan := make(chan *ReplaceEdgeInstanceGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReplaceEdgeInstanceGateway(request)
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

// ReplaceEdgeInstanceGatewayWithCallback invokes the iot.ReplaceEdgeInstanceGateway API asynchronously
func (client *Client) ReplaceEdgeInstanceGatewayWithCallback(request *ReplaceEdgeInstanceGatewayRequest, callback func(response *ReplaceEdgeInstanceGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReplaceEdgeInstanceGatewayResponse
		var err error
		defer close(result)
		response, err = client.ReplaceEdgeInstanceGateway(request)
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

// ReplaceEdgeInstanceGatewayRequest is the request struct for api ReplaceEdgeInstanceGateway
type ReplaceEdgeInstanceGatewayRequest struct {
	*requests.RpcRequest
	NewGatewayId     string `position:"Query" name:"NewGatewayId"`
	IotInstanceId    string `position:"Query" name:"IotInstanceId"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	CurrentGatewayId string `position:"Query" name:"CurrentGatewayId"`
	ApiProduct       string `position:"Body" name:"ApiProduct"`
	ApiRevision      string `position:"Body" name:"ApiRevision"`
}

// ReplaceEdgeInstanceGatewayResponse is the response struct for api ReplaceEdgeInstanceGateway
type ReplaceEdgeInstanceGatewayResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateReplaceEdgeInstanceGatewayRequest creates a request to invoke ReplaceEdgeInstanceGateway API
func CreateReplaceEdgeInstanceGatewayRequest() (request *ReplaceEdgeInstanceGatewayRequest) {
	request = &ReplaceEdgeInstanceGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ReplaceEdgeInstanceGateway", "", "")
	request.Method = requests.POST
	return
}

// CreateReplaceEdgeInstanceGatewayResponse creates a response to parse from ReplaceEdgeInstanceGateway response
func CreateReplaceEdgeInstanceGatewayResponse() (response *ReplaceEdgeInstanceGatewayResponse) {
	response = &ReplaceEdgeInstanceGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
