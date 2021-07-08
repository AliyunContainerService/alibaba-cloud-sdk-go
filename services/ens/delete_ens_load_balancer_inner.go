package ens

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

// DeleteEnsLoadBalancerInner invokes the ens.DeleteEnsLoadBalancerInner API synchronously
func (client *Client) DeleteEnsLoadBalancerInner(request *DeleteEnsLoadBalancerInnerRequest) (response *DeleteEnsLoadBalancerInnerResponse, err error) {
	response = CreateDeleteEnsLoadBalancerInnerResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteEnsLoadBalancerInnerWithChan invokes the ens.DeleteEnsLoadBalancerInner API asynchronously
func (client *Client) DeleteEnsLoadBalancerInnerWithChan(request *DeleteEnsLoadBalancerInnerRequest) (<-chan *DeleteEnsLoadBalancerInnerResponse, <-chan error) {
	responseChan := make(chan *DeleteEnsLoadBalancerInnerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteEnsLoadBalancerInner(request)
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

// DeleteEnsLoadBalancerInnerWithCallback invokes the ens.DeleteEnsLoadBalancerInner API asynchronously
func (client *Client) DeleteEnsLoadBalancerInnerWithCallback(request *DeleteEnsLoadBalancerInnerRequest, callback func(response *DeleteEnsLoadBalancerInnerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteEnsLoadBalancerInnerResponse
		var err error
		defer close(result)
		response, err = client.DeleteEnsLoadBalancerInner(request)
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

// DeleteEnsLoadBalancerInnerRequest is the request struct for api DeleteEnsLoadBalancerInner
type DeleteEnsLoadBalancerInnerRequest struct {
	*requests.RpcRequest
	LoadBalancerId string `position:"Query" name:"LoadBalancerId"`
}

// DeleteEnsLoadBalancerInnerResponse is the response struct for api DeleteEnsLoadBalancerInner
type DeleteEnsLoadBalancerInnerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteEnsLoadBalancerInnerRequest creates a request to invoke DeleteEnsLoadBalancerInner API
func CreateDeleteEnsLoadBalancerInnerRequest() (request *DeleteEnsLoadBalancerInnerRequest) {
	request = &DeleteEnsLoadBalancerInnerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DeleteEnsLoadBalancerInner", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteEnsLoadBalancerInnerResponse creates a response to parse from DeleteEnsLoadBalancerInner response
func CreateDeleteEnsLoadBalancerInnerResponse() (response *DeleteEnsLoadBalancerInnerResponse) {
	response = &DeleteEnsLoadBalancerInnerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
