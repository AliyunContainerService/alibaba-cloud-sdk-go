package cdrs

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

// ListPersonTrack invokes the cdrs.ListPersonTrack API synchronously
func (client *Client) ListPersonTrack(request *ListPersonTrackRequest) (response *ListPersonTrackResponse, err error) {
	response = CreateListPersonTrackResponse()
	err = client.DoAction(request, response)
	return
}

// ListPersonTrackWithChan invokes the cdrs.ListPersonTrack API asynchronously
func (client *Client) ListPersonTrackWithChan(request *ListPersonTrackRequest) (<-chan *ListPersonTrackResponse, <-chan error) {
	responseChan := make(chan *ListPersonTrackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPersonTrack(request)
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

// ListPersonTrackWithCallback invokes the cdrs.ListPersonTrack API asynchronously
func (client *Client) ListPersonTrackWithCallback(request *ListPersonTrackRequest, callback func(response *ListPersonTrackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPersonTrackResponse
		var err error
		defer close(result)
		response, err = client.ListPersonTrack(request)
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

// ListPersonTrackRequest is the request struct for api ListPersonTrack
type ListPersonTrackRequest struct {
	*requests.RpcRequest
	Schema     string           `position:"Body" name:"Schema"`
	CorpId     string           `position:"Body" name:"CorpId"`
	EndTime    string           `position:"Body" name:"EndTime"`
	StartTime  string           `position:"Body" name:"StartTime"`
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
	PersonId   string           `position:"Body" name:"PersonId"`
}

// ListPersonTrackResponse is the response struct for api ListPersonTrack
type ListPersonTrackResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListPersonTrackRequest creates a request to invoke ListPersonTrack API
func CreateListPersonTrackRequest() (request *ListPersonTrackRequest) {
	request = &ListPersonTrackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListPersonTrack", "", "")
	request.Method = requests.POST
	return
}

// CreateListPersonTrackResponse creates a response to parse from ListPersonTrack response
func CreateListPersonTrackResponse() (response *ListPersonTrackResponse) {
	response = &ListPersonTrackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
