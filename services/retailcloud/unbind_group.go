package retailcloud

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

// UnbindGroup invokes the retailcloud.UnbindGroup API synchronously
func (client *Client) UnbindGroup(request *UnbindGroupRequest) (response *UnbindGroupResponse, err error) {
	response = CreateUnbindGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindGroupWithChan invokes the retailcloud.UnbindGroup API asynchronously
func (client *Client) UnbindGroupWithChan(request *UnbindGroupRequest) (<-chan *UnbindGroupResponse, <-chan error) {
	responseChan := make(chan *UnbindGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindGroup(request)
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

// UnbindGroupWithCallback invokes the retailcloud.UnbindGroup API asynchronously
func (client *Client) UnbindGroupWithCallback(request *UnbindGroupRequest, callback func(response *UnbindGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindGroupResponse
		var err error
		defer close(result)
		response, err = client.UnbindGroup(request)
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

// UnbindGroupRequest is the request struct for api UnbindGroup
type UnbindGroupRequest struct {
	*requests.RpcRequest
	BizCode string           `position:"Query" name:"BizCode"`
	AppId   requests.Integer `position:"Query" name:"AppId"`
	Name    string           `position:"Query" name:"Name"`
}

// UnbindGroupResponse is the response struct for api UnbindGroup
type UnbindGroupResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUnbindGroupRequest creates a request to invoke UnbindGroup API
func CreateUnbindGroupRequest() (request *UnbindGroupRequest) {
	request = &UnbindGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "UnbindGroup", "retailcloud", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbindGroupResponse creates a response to parse from UnbindGroup response
func CreateUnbindGroupResponse() (response *UnbindGroupResponse) {
	response = &UnbindGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
