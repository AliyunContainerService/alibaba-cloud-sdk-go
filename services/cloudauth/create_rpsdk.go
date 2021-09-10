package cloudauth

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

// CreateRPSDK invokes the cloudauth.CreateRPSDK API synchronously
func (client *Client) CreateRPSDK(request *CreateRPSDKRequest) (response *CreateRPSDKResponse, err error) {
	response = CreateCreateRPSDKResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRPSDKWithChan invokes the cloudauth.CreateRPSDK API asynchronously
func (client *Client) CreateRPSDKWithChan(request *CreateRPSDKRequest) (<-chan *CreateRPSDKResponse, <-chan error) {
	responseChan := make(chan *CreateRPSDKResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRPSDK(request)
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

// CreateRPSDKWithCallback invokes the cloudauth.CreateRPSDK API asynchronously
func (client *Client) CreateRPSDKWithCallback(request *CreateRPSDKRequest, callback func(response *CreateRPSDKResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRPSDKResponse
		var err error
		defer close(result)
		response, err = client.CreateRPSDK(request)
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

// CreateRPSDKRequest is the request struct for api CreateRPSDK
type CreateRPSDKRequest struct {
	*requests.RpcRequest
	AppUrl   string `position:"Query" name:"AppUrl"`
	Platform string `position:"Query" name:"Platform"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
}

// CreateRPSDKResponse is the response struct for api CreateRPSDK
type CreateRPSDKResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateRPSDKRequest creates a request to invoke CreateRPSDK API
func CreateCreateRPSDKRequest() (request *CreateRPSDKRequest) {
	request = &CreateRPSDKRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "CreateRPSDK", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateRPSDKResponse creates a response to parse from CreateRPSDK response
func CreateCreateRPSDKResponse() (response *CreateRPSDKResponse) {
	response = &CreateRPSDKResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
