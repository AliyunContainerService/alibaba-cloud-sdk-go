package vpcpeer

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

// AcceptVpcPeerConnection invokes the vpcpeer.AcceptVpcPeerConnection API synchronously
func (client *Client) AcceptVpcPeerConnection(request *AcceptVpcPeerConnectionRequest) (response *AcceptVpcPeerConnectionResponse, err error) {
	response = CreateAcceptVpcPeerConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// AcceptVpcPeerConnectionWithChan invokes the vpcpeer.AcceptVpcPeerConnection API asynchronously
func (client *Client) AcceptVpcPeerConnectionWithChan(request *AcceptVpcPeerConnectionRequest) (<-chan *AcceptVpcPeerConnectionResponse, <-chan error) {
	responseChan := make(chan *AcceptVpcPeerConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AcceptVpcPeerConnection(request)
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

// AcceptVpcPeerConnectionWithCallback invokes the vpcpeer.AcceptVpcPeerConnection API asynchronously
func (client *Client) AcceptVpcPeerConnectionWithCallback(request *AcceptVpcPeerConnectionRequest, callback func(response *AcceptVpcPeerConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AcceptVpcPeerConnectionResponse
		var err error
		defer close(result)
		response, err = client.AcceptVpcPeerConnection(request)
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

// AcceptVpcPeerConnectionRequest is the request struct for api AcceptVpcPeerConnection
type AcceptVpcPeerConnectionRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Body" name:"ClientToken"`
	Channel              string           `position:"Body" name:"Channel"`
	RequestContent       string           `position:"Body" name:"RequestContent"`
	DryRun               requests.Boolean `position:"Body" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Body" name:"ResourceOwnerAccount"`
	InstanceId           string           `position:"Body" name:"InstanceId"`
}

// AcceptVpcPeerConnectionResponse is the response struct for api AcceptVpcPeerConnection
type AcceptVpcPeerConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAcceptVpcPeerConnectionRequest creates a request to invoke AcceptVpcPeerConnection API
func CreateAcceptVpcPeerConnectionRequest() (request *AcceptVpcPeerConnectionRequest) {
	request = &AcceptVpcPeerConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VpcPeer", "2022-01-01", "AcceptVpcPeerConnection", "vpcpeer", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAcceptVpcPeerConnectionResponse creates a response to parse from AcceptVpcPeerConnection response
func CreateAcceptVpcPeerConnectionResponse() (response *AcceptVpcPeerConnectionResponse) {
	response = &AcceptVpcPeerConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
