package csas

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

// CreatePrivateAccessApplication invokes the csas.CreatePrivateAccessApplication API synchronously
func (client *Client) CreatePrivateAccessApplication(request *CreatePrivateAccessApplicationRequest) (response *CreatePrivateAccessApplicationResponse, err error) {
	response = CreateCreatePrivateAccessApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePrivateAccessApplicationWithChan invokes the csas.CreatePrivateAccessApplication API asynchronously
func (client *Client) CreatePrivateAccessApplicationWithChan(request *CreatePrivateAccessApplicationRequest) (<-chan *CreatePrivateAccessApplicationResponse, <-chan error) {
	responseChan := make(chan *CreatePrivateAccessApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePrivateAccessApplication(request)
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

// CreatePrivateAccessApplicationWithCallback invokes the csas.CreatePrivateAccessApplication API asynchronously
func (client *Client) CreatePrivateAccessApplicationWithCallback(request *CreatePrivateAccessApplicationRequest, callback func(response *CreatePrivateAccessApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePrivateAccessApplicationResponse
		var err error
		defer close(result)
		response, err = client.CreatePrivateAccessApplication(request)
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

// CreatePrivateAccessApplicationRequest is the request struct for api CreatePrivateAccessApplication
type CreatePrivateAccessApplicationRequest struct {
	*requests.RpcRequest
	Addresses   *[]string                                   `position:"Body" name:"Addresses"  type:"Json"`
	Description string                                      `position:"Body" name:"Description"`
	Protocol    string                                      `position:"Body" name:"Protocol"`
	SourceIp    string                                      `position:"Query" name:"SourceIp"`
	TagIds      *[]string                                   `position:"Body" name:"TagIds"  type:"Json"`
	PortRanges  *[]CreatePrivateAccessApplicationPortRanges `position:"Body" name:"PortRanges"  type:"Json"`
	Name        string                                      `position:"Body" name:"Name"`
	Status      string                                      `position:"Body" name:"Status"`
}

// CreatePrivateAccessApplicationPortRanges is a repeated param struct in CreatePrivateAccessApplicationRequest
type CreatePrivateAccessApplicationPortRanges struct {
	End   string `name:"End"`
	Begin string `name:"Begin"`
}

// CreatePrivateAccessApplicationResponse is the response struct for api CreatePrivateAccessApplication
type CreatePrivateAccessApplicationResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ApplicationId string `json:"ApplicationId" xml:"ApplicationId"`
}

// CreateCreatePrivateAccessApplicationRequest creates a request to invoke CreatePrivateAccessApplication API
func CreateCreatePrivateAccessApplicationRequest() (request *CreatePrivateAccessApplicationRequest) {
	request = &CreatePrivateAccessApplicationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "CreatePrivateAccessApplication", "", "")
	request.Method = requests.POST
	return
}

// CreateCreatePrivateAccessApplicationResponse creates a response to parse from CreatePrivateAccessApplication response
func CreateCreatePrivateAccessApplicationResponse() (response *CreatePrivateAccessApplicationResponse) {
	response = &CreatePrivateAccessApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
