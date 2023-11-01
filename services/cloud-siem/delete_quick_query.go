package cloud_siem

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

// DeleteQuickQuery invokes the cloud_siem.DeleteQuickQuery API synchronously
func (client *Client) DeleteQuickQuery(request *DeleteQuickQueryRequest) (response *DeleteQuickQueryResponse, err error) {
	response = CreateDeleteQuickQueryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteQuickQueryWithChan invokes the cloud_siem.DeleteQuickQuery API asynchronously
func (client *Client) DeleteQuickQueryWithChan(request *DeleteQuickQueryRequest) (<-chan *DeleteQuickQueryResponse, <-chan error) {
	responseChan := make(chan *DeleteQuickQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteQuickQuery(request)
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

// DeleteQuickQueryWithCallback invokes the cloud_siem.DeleteQuickQuery API asynchronously
func (client *Client) DeleteQuickQueryWithCallback(request *DeleteQuickQueryRequest, callback func(response *DeleteQuickQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteQuickQueryResponse
		var err error
		defer close(result)
		response, err = client.DeleteQuickQuery(request)
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

// DeleteQuickQueryRequest is the request struct for api DeleteQuickQuery
type DeleteQuickQueryRequest struct {
	*requests.RpcRequest
	SearchName string `position:"Body" name:"SearchName"`
}

// DeleteQuickQueryResponse is the response struct for api DeleteQuickQuery
type DeleteQuickQueryResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	ErrCode   string `json:"ErrCode" xml:"ErrCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	DyCode    string `json:"DyCode" xml:"DyCode"`
	DyMessage string `json:"DyMessage" xml:"DyMessage"`
}

// CreateDeleteQuickQueryRequest creates a request to invoke DeleteQuickQuery API
func CreateDeleteQuickQueryRequest() (request *DeleteQuickQueryRequest) {
	request = &DeleteQuickQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DeleteQuickQuery", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteQuickQueryResponse creates a response to parse from DeleteQuickQuery response
func CreateDeleteQuickQueryResponse() (response *DeleteQuickQueryResponse) {
	response = &DeleteQuickQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
