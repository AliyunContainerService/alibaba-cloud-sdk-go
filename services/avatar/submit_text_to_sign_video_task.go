package avatar

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

// SubmitTextToSignVideoTask invokes the avatar.SubmitTextToSignVideoTask API synchronously
func (client *Client) SubmitTextToSignVideoTask(request *SubmitTextToSignVideoTaskRequest) (response *SubmitTextToSignVideoTaskResponse, err error) {
	response = CreateSubmitTextToSignVideoTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitTextToSignVideoTaskWithChan invokes the avatar.SubmitTextToSignVideoTask API asynchronously
func (client *Client) SubmitTextToSignVideoTaskWithChan(request *SubmitTextToSignVideoTaskRequest) (<-chan *SubmitTextToSignVideoTaskResponse, <-chan error) {
	responseChan := make(chan *SubmitTextToSignVideoTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitTextToSignVideoTask(request)
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

// SubmitTextToSignVideoTaskWithCallback invokes the avatar.SubmitTextToSignVideoTask API asynchronously
func (client *Client) SubmitTextToSignVideoTaskWithCallback(request *SubmitTextToSignVideoTaskRequest, callback func(response *SubmitTextToSignVideoTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitTextToSignVideoTaskResponse
		var err error
		defer close(result)
		response, err = client.SubmitTextToSignVideoTask(request)
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

// SubmitTextToSignVideoTaskRequest is the request struct for api SubmitTextToSignVideoTask
type SubmitTextToSignVideoTaskRequest struct {
	*requests.RpcRequest
	App       SubmitTextToSignVideoTaskApp       `position:"Query" name:"App"  type:"Struct"`
	VideoInfo SubmitTextToSignVideoTaskVideoInfo `position:"Query" name:"VideoInfo"  type:"Struct"`
	TenantId  requests.Integer                   `position:"Query" name:"TenantId"`
	Text      string                             `position:"Query" name:"Text"`
	Title     string                             `position:"Query" name:"Title"`
}

// SubmitTextToSignVideoTaskApp is a repeated param struct in SubmitTextToSignVideoTaskRequest
type SubmitTextToSignVideoTaskApp struct {
	AppId string `name:"AppId"`
}

// SubmitTextToSignVideoTaskVideoInfo is a repeated param struct in SubmitTextToSignVideoTaskRequest
type SubmitTextToSignVideoTaskVideoInfo struct {
	IsAlpha     string `name:"IsAlpha"`
	IsSubtitles string `name:"IsSubtitles"`
	Resolution  string `name:"Resolution"`
}

// SubmitTextToSignVideoTaskResponse is the response struct for api SubmitTextToSignVideoTask
type SubmitTextToSignVideoTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   string `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSubmitTextToSignVideoTaskRequest creates a request to invoke SubmitTextToSignVideoTask API
func CreateSubmitTextToSignVideoTaskRequest() (request *SubmitTextToSignVideoTaskRequest) {
	request = &SubmitTextToSignVideoTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "SubmitTextToSignVideoTask", "", "")
	request.Method = requests.POST
	return
}

// CreateSubmitTextToSignVideoTaskResponse creates a response to parse from SubmitTextToSignVideoTask response
func CreateSubmitTextToSignVideoTaskResponse() (response *SubmitTextToSignVideoTaskResponse) {
	response = &SubmitTextToSignVideoTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
