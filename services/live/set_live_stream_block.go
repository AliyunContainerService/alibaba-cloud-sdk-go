package live

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

// SetLiveStreamBlock invokes the live.SetLiveStreamBlock API synchronously
func (client *Client) SetLiveStreamBlock(request *SetLiveStreamBlockRequest) (response *SetLiveStreamBlockResponse, err error) {
	response = CreateSetLiveStreamBlockResponse()
	err = client.DoAction(request, response)
	return
}

// SetLiveStreamBlockWithChan invokes the live.SetLiveStreamBlock API asynchronously
func (client *Client) SetLiveStreamBlockWithChan(request *SetLiveStreamBlockRequest) (<-chan *SetLiveStreamBlockResponse, <-chan error) {
	responseChan := make(chan *SetLiveStreamBlockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLiveStreamBlock(request)
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

// SetLiveStreamBlockWithCallback invokes the live.SetLiveStreamBlock API asynchronously
func (client *Client) SetLiveStreamBlockWithCallback(request *SetLiveStreamBlockRequest, callback func(response *SetLiveStreamBlockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLiveStreamBlockResponse
		var err error
		defer close(result)
		response, err = client.SetLiveStreamBlock(request)
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

// SetLiveStreamBlockRequest is the request struct for api SetLiveStreamBlock
type SetLiveStreamBlockRequest struct {
	*requests.RpcRequest
	LocationList string           `position:"Query" name:"LocationList"`
	BlockType    string           `position:"Query" name:"BlockType"`
	ReleaseTime  string           `position:"Query" name:"ReleaseTime"`
	AppName      string           `position:"Query" name:"AppName"`
	StreamName   string           `position:"Query" name:"StreamName"`
	DomainName   string           `position:"Query" name:"DomainName"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
}

// SetLiveStreamBlockResponse is the response struct for api SetLiveStreamBlock
type SetLiveStreamBlockResponse struct {
	*responses.BaseResponse
	Description string `json:"Description" xml:"Description"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Status      string `json:"Status" xml:"Status"`
}

// CreateSetLiveStreamBlockRequest creates a request to invoke SetLiveStreamBlock API
func CreateSetLiveStreamBlockRequest() (request *SetLiveStreamBlockRequest) {
	request = &SetLiveStreamBlockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "SetLiveStreamBlock", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetLiveStreamBlockResponse creates a response to parse from SetLiveStreamBlock response
func CreateSetLiveStreamBlockResponse() (response *SetLiveStreamBlockResponse) {
	response = &SetLiveStreamBlockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
