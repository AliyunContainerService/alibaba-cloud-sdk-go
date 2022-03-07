package voicenavigator

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

// ListSandBoxWhileList invokes the voicenavigator.ListSandBoxWhileList API synchronously
func (client *Client) ListSandBoxWhileList(request *ListSandBoxWhileListRequest) (response *ListSandBoxWhileListResponse, err error) {
	response = CreateListSandBoxWhileListResponse()
	err = client.DoAction(request, response)
	return
}

// ListSandBoxWhileListWithChan invokes the voicenavigator.ListSandBoxWhileList API asynchronously
func (client *Client) ListSandBoxWhileListWithChan(request *ListSandBoxWhileListRequest) (<-chan *ListSandBoxWhileListResponse, <-chan error) {
	responseChan := make(chan *ListSandBoxWhileListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSandBoxWhileList(request)
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

// ListSandBoxWhileListWithCallback invokes the voicenavigator.ListSandBoxWhileList API asynchronously
func (client *Client) ListSandBoxWhileListWithCallback(request *ListSandBoxWhileListRequest, callback func(response *ListSandBoxWhileListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSandBoxWhileListResponse
		var err error
		defer close(result)
		response, err = client.ListSandBoxWhileList(request)
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

// ListSandBoxWhileListRequest is the request struct for api ListSandBoxWhileList
type ListSandBoxWhileListRequest struct {
	*requests.RpcRequest
	PhoneNumber string           `position:"Query" name:"PhoneNumber"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	Name        string           `position:"Query" name:"Name"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// ListSandBoxWhileListResponse is the response struct for api ListSandBoxWhileList
type ListSandBoxWhileListResponse struct {
	*responses.BaseResponse
	PageSize   int             `json:"PageSize" xml:"PageSize"`
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	PageNumber int             `json:"PageNumber" xml:"PageNumber"`
	TotalCount int64           `json:"TotalCount" xml:"TotalCount"`
	WhiteList  []WhiteListItem `json:"WhiteList" xml:"WhiteList"`
}

// CreateListSandBoxWhileListRequest creates a request to invoke ListSandBoxWhileList API
func CreateListSandBoxWhileListRequest() (request *ListSandBoxWhileListRequest) {
	request = &ListSandBoxWhileListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "ListSandBoxWhileList", "voicebot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSandBoxWhileListResponse creates a response to parse from ListSandBoxWhileList response
func CreateListSandBoxWhileListResponse() (response *ListSandBoxWhileListResponse) {
	response = &ListSandBoxWhileListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
