package sas

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

// ModifyWebLockDeleteConfig invokes the sas.ModifyWebLockDeleteConfig API synchronously
func (client *Client) ModifyWebLockDeleteConfig(request *ModifyWebLockDeleteConfigRequest) (response *ModifyWebLockDeleteConfigResponse, err error) {
	response = CreateModifyWebLockDeleteConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyWebLockDeleteConfigWithChan invokes the sas.ModifyWebLockDeleteConfig API asynchronously
func (client *Client) ModifyWebLockDeleteConfigWithChan(request *ModifyWebLockDeleteConfigRequest) (<-chan *ModifyWebLockDeleteConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyWebLockDeleteConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyWebLockDeleteConfig(request)
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

// ModifyWebLockDeleteConfigWithCallback invokes the sas.ModifyWebLockDeleteConfig API asynchronously
func (client *Client) ModifyWebLockDeleteConfigWithCallback(request *ModifyWebLockDeleteConfigRequest, callback func(response *ModifyWebLockDeleteConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyWebLockDeleteConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyWebLockDeleteConfig(request)
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

// ModifyWebLockDeleteConfigRequest is the request struct for api ModifyWebLockDeleteConfig
type ModifyWebLockDeleteConfigRequest struct {
	*requests.RpcRequest
	Uuid     string           `position:"Query" name:"Uuid"`
	SourceIp string           `position:"Query" name:"SourceIp"`
	Id       requests.Integer `position:"Query" name:"Id"`
	Lang     string           `position:"Query" name:"Lang"`
}

// ModifyWebLockDeleteConfigResponse is the response struct for api ModifyWebLockDeleteConfig
type ModifyWebLockDeleteConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyWebLockDeleteConfigRequest creates a request to invoke ModifyWebLockDeleteConfig API
func CreateModifyWebLockDeleteConfigRequest() (request *ModifyWebLockDeleteConfigRequest) {
	request = &ModifyWebLockDeleteConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyWebLockDeleteConfig", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyWebLockDeleteConfigResponse creates a response to parse from ModifyWebLockDeleteConfig response
func CreateModifyWebLockDeleteConfigResponse() (response *ModifyWebLockDeleteConfigResponse) {
	response = &ModifyWebLockDeleteConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
