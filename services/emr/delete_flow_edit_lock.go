package emr

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

// DeleteFlowEditLock invokes the emr.DeleteFlowEditLock API synchronously
func (client *Client) DeleteFlowEditLock(request *DeleteFlowEditLockRequest) (response *DeleteFlowEditLockResponse, err error) {
	response = CreateDeleteFlowEditLockResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFlowEditLockWithChan invokes the emr.DeleteFlowEditLock API asynchronously
func (client *Client) DeleteFlowEditLockWithChan(request *DeleteFlowEditLockRequest) (<-chan *DeleteFlowEditLockResponse, <-chan error) {
	responseChan := make(chan *DeleteFlowEditLockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFlowEditLock(request)
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

// DeleteFlowEditLockWithCallback invokes the emr.DeleteFlowEditLock API asynchronously
func (client *Client) DeleteFlowEditLockWithCallback(request *DeleteFlowEditLockRequest, callback func(response *DeleteFlowEditLockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFlowEditLockResponse
		var err error
		defer close(result)
		response, err = client.DeleteFlowEditLock(request)
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

// DeleteFlowEditLockRequest is the request struct for api DeleteFlowEditLock
type DeleteFlowEditLockRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EntityId        string           `position:"Query" name:"EntityId"`
}

// DeleteFlowEditLockResponse is the response struct for api DeleteFlowEditLock
type DeleteFlowEditLockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDeleteFlowEditLockRequest creates a request to invoke DeleteFlowEditLock API
func CreateDeleteFlowEditLockRequest() (request *DeleteFlowEditLockRequest) {
	request = &DeleteFlowEditLockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlowEditLock", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFlowEditLockResponse creates a response to parse from DeleteFlowEditLock response
func CreateDeleteFlowEditLockResponse() (response *DeleteFlowEditLockResponse) {
	response = &DeleteFlowEditLockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
