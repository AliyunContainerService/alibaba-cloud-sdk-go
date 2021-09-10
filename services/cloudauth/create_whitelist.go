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

// CreateWhitelist invokes the cloudauth.CreateWhitelist API synchronously
func (client *Client) CreateWhitelist(request *CreateWhitelistRequest) (response *CreateWhitelistResponse, err error) {
	response = CreateCreateWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// CreateWhitelistWithChan invokes the cloudauth.CreateWhitelist API asynchronously
func (client *Client) CreateWhitelistWithChan(request *CreateWhitelistRequest) (<-chan *CreateWhitelistResponse, <-chan error) {
	responseChan := make(chan *CreateWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateWhitelist(request)
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

// CreateWhitelistWithCallback invokes the cloudauth.CreateWhitelist API asynchronously
func (client *Client) CreateWhitelistWithCallback(request *CreateWhitelistRequest, callback func(response *CreateWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateWhitelistResponse
		var err error
		defer close(result)
		response, err = client.CreateWhitelist(request)
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

// CreateWhitelistRequest is the request struct for api CreateWhitelist
type CreateWhitelistRequest struct {
	*requests.RpcRequest
	ValidDay  string `position:"Query" name:"ValidDay"`
	BizType   string `position:"Query" name:"BizType"`
	IdCardNum string `position:"Query" name:"IdCardNum"`
	SourceIp  string `position:"Query" name:"SourceIp"`
	BizId     string `position:"Query" name:"BizId"`
	Lang      string `position:"Query" name:"Lang"`
}

// CreateWhitelistResponse is the response struct for api CreateWhitelist
type CreateWhitelistResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateWhitelistRequest creates a request to invoke CreateWhitelist API
func CreateCreateWhitelistRequest() (request *CreateWhitelistRequest) {
	request = &CreateWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "CreateWhitelist", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateWhitelistResponse creates a response to parse from CreateWhitelist response
func CreateCreateWhitelistResponse() (response *CreateWhitelistResponse) {
	response = &CreateWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
