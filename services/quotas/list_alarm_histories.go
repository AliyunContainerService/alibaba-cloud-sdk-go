package quotas

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

// ListAlarmHistories invokes the quotas.ListAlarmHistories API synchronously
func (client *Client) ListAlarmHistories(request *ListAlarmHistoriesRequest) (response *ListAlarmHistoriesResponse, err error) {
	response = CreateListAlarmHistoriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAlarmHistoriesWithChan invokes the quotas.ListAlarmHistories API asynchronously
func (client *Client) ListAlarmHistoriesWithChan(request *ListAlarmHistoriesRequest) (<-chan *ListAlarmHistoriesResponse, <-chan error) {
	responseChan := make(chan *ListAlarmHistoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAlarmHistories(request)
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

// ListAlarmHistoriesWithCallback invokes the quotas.ListAlarmHistories API asynchronously
func (client *Client) ListAlarmHistoriesWithCallback(request *ListAlarmHistoriesRequest, callback func(response *ListAlarmHistoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAlarmHistoriesResponse
		var err error
		defer close(result)
		response, err = client.ListAlarmHistories(request)
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

// ListAlarmHistoriesRequest is the request struct for api ListAlarmHistories
type ListAlarmHistoriesRequest struct {
	*requests.RpcRequest
	ProductCode     string           `position:"Body" name:"ProductCode"`
	StartTime       requests.Integer `position:"Body" name:"StartTime"`
	NextToken       string           `position:"Body" name:"NextToken"`
	Keyword         string           `position:"Body" name:"Keyword"`
	OriginalContext string           `position:"Body" name:"OriginalContext"`
	EndTime         requests.Integer `position:"Body" name:"EndTime"`
	AlarmId         string           `position:"Body" name:"AlarmId"`
	MaxResults      requests.Integer `position:"Body" name:"MaxResults"`
}

// ListAlarmHistoriesResponse is the response struct for api ListAlarmHistories
type ListAlarmHistoriesResponse struct {
	*responses.BaseResponse
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	NextToken      string         `json:"NextToken" xml:"NextToken"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	MaxResults     int            `json:"MaxResults" xml:"MaxResults"`
	AlarmHistories []AlarmHistory `json:"AlarmHistories" xml:"AlarmHistories"`
}

// CreateListAlarmHistoriesRequest creates a request to invoke ListAlarmHistories API
func CreateListAlarmHistoriesRequest() (request *ListAlarmHistoriesRequest) {
	request = &ListAlarmHistoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quotas", "2020-05-10", "ListAlarmHistories", "quotas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAlarmHistoriesResponse creates a response to parse from ListAlarmHistories response
func CreateListAlarmHistoriesResponse() (response *ListAlarmHistoriesResponse) {
	response = &ListAlarmHistoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
