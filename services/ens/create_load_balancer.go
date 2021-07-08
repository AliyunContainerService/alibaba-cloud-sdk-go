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

// CreateLoadBalancer invokes the ens.CreateLoadBalancer API synchronously
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	response = CreateCreateLoadBalancerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLoadBalancerWithChan invokes the ens.CreateLoadBalancer API asynchronously
func (client *Client) CreateLoadBalancerWithChan(request *CreateLoadBalancerRequest) (<-chan *CreateLoadBalancerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancer(request)
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

// CreateLoadBalancerWithCallback invokes the ens.CreateLoadBalancer API asynchronously
func (client *Client) CreateLoadBalancerWithCallback(request *CreateLoadBalancerRequest, callback func(response *CreateLoadBalancerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancer(request)
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

// CreateLoadBalancerRequest is the request struct for api CreateLoadBalancer
type CreateLoadBalancerRequest struct {
	*requests.RpcRequest
	LoadBalancerName string `position:"Query" name:"LoadBalancerName"`
	EnsRegionId      string `position:"Query" name:"EnsRegionId"`
	LoadBalancerSpec string `position:"Query" name:"LoadBalancerSpec"`
	VSwitchId        string `position:"Query" name:"VSwitchId"`
	NetworkId        string `position:"Query" name:"NetworkId"`
	PayType          string `position:"Query" name:"PayType"`
}

// CreateLoadBalancerResponse is the response struct for api CreateLoadBalancer
type CreateLoadBalancerResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	LoadBalancerId   string `json:"LoadBalancerId" xml:"LoadBalancerId"`
	VSwitchId        string `json:"VSwitchId" xml:"VSwitchId"`
	LoadBalancerName string `json:"LoadBalancerName" xml:"LoadBalancerName"`
	NetworkId        string `json:"NetworkId" xml:"NetworkId"`
}

// CreateCreateLoadBalancerRequest creates a request to invoke CreateLoadBalancer API
func CreateCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "CreateLoadBalancer", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLoadBalancerResponse creates a response to parse from CreateLoadBalancer response
func CreateCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
