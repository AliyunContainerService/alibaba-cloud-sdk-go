package domain

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

// SaveSingleTaskForQueryingTransferAuthorizationCode invokes the domain.SaveSingleTaskForQueryingTransferAuthorizationCode API synchronously
func (client *Client) SaveSingleTaskForQueryingTransferAuthorizationCode(request *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) (response *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse, err error) {
	response = CreateSaveSingleTaskForQueryingTransferAuthorizationCodeResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForQueryingTransferAuthorizationCodeWithChan invokes the domain.SaveSingleTaskForQueryingTransferAuthorizationCode API asynchronously
func (client *Client) SaveSingleTaskForQueryingTransferAuthorizationCodeWithChan(request *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) (<-chan *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForQueryingTransferAuthorizationCode(request)
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

// SaveSingleTaskForQueryingTransferAuthorizationCodeWithCallback invokes the domain.SaveSingleTaskForQueryingTransferAuthorizationCode API asynchronously
func (client *Client) SaveSingleTaskForQueryingTransferAuthorizationCodeWithCallback(request *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest, callback func(response *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForQueryingTransferAuthorizationCode(request)
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

// SaveSingleTaskForQueryingTransferAuthorizationCodeRequest is the request struct for api SaveSingleTaskForQueryingTransferAuthorizationCode
type SaveSingleTaskForQueryingTransferAuthorizationCodeRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// SaveSingleTaskForQueryingTransferAuthorizationCodeResponse is the response struct for api SaveSingleTaskForQueryingTransferAuthorizationCode
type SaveSingleTaskForQueryingTransferAuthorizationCodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForQueryingTransferAuthorizationCodeRequest creates a request to invoke SaveSingleTaskForQueryingTransferAuthorizationCode API
func CreateSaveSingleTaskForQueryingTransferAuthorizationCodeRequest() (request *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) {
	request = &SaveSingleTaskForQueryingTransferAuthorizationCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForQueryingTransferAuthorizationCode", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveSingleTaskForQueryingTransferAuthorizationCodeResponse creates a response to parse from SaveSingleTaskForQueryingTransferAuthorizationCode response
func CreateSaveSingleTaskForQueryingTransferAuthorizationCodeResponse() (response *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse) {
	response = &SaveSingleTaskForQueryingTransferAuthorizationCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
