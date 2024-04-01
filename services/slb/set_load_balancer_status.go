package slb

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

// SetLoadBalancerStatus invokes the slb.SetLoadBalancerStatus API synchronously
func (client *Client) SetLoadBalancerStatus(request *SetLoadBalancerStatusRequest) (response *SetLoadBalancerStatusResponse, err error) {
	response = CreateSetLoadBalancerStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SetLoadBalancerStatusWithChan invokes the slb.SetLoadBalancerStatus API asynchronously
func (client *Client) SetLoadBalancerStatusWithChan(request *SetLoadBalancerStatusRequest) (<-chan *SetLoadBalancerStatusResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerStatus(request)
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

// SetLoadBalancerStatusWithCallback invokes the slb.SetLoadBalancerStatus API asynchronously
func (client *Client) SetLoadBalancerStatusWithCallback(request *SetLoadBalancerStatusRequest, callback func(response *SetLoadBalancerStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerStatusResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerStatus(request)
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

// SetLoadBalancerStatusRequest is the request struct for api SetLoadBalancerStatus
type SetLoadBalancerStatusRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	LoadBalancerStatus   string           `position:"Query" name:"LoadBalancerStatus"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// SetLoadBalancerStatusResponse is the response struct for api SetLoadBalancerStatus
type SetLoadBalancerStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetLoadBalancerStatusRequest creates a request to invoke SetLoadBalancerStatus API
func CreateSetLoadBalancerStatusRequest() (request *SetLoadBalancerStatusRequest) {
	request = &SetLoadBalancerStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2013-02-21", "SetLoadBalancerStatus", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetLoadBalancerStatusResponse creates a response to parse from SetLoadBalancerStatus response
func CreateSetLoadBalancerStatusResponse() (response *SetLoadBalancerStatusResponse) {
	response = &SetLoadBalancerStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
