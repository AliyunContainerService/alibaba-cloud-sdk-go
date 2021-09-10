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

// UpdateFaceConfig invokes the cloudauth.UpdateFaceConfig API synchronously
func (client *Client) UpdateFaceConfig(request *UpdateFaceConfigRequest) (response *UpdateFaceConfigResponse, err error) {
	response = CreateUpdateFaceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFaceConfigWithChan invokes the cloudauth.UpdateFaceConfig API asynchronously
func (client *Client) UpdateFaceConfigWithChan(request *UpdateFaceConfigRequest) (<-chan *UpdateFaceConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateFaceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFaceConfig(request)
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

// UpdateFaceConfigWithCallback invokes the cloudauth.UpdateFaceConfig API asynchronously
func (client *Client) UpdateFaceConfigWithCallback(request *UpdateFaceConfigRequest, callback func(response *UpdateFaceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFaceConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateFaceConfig(request)
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

// UpdateFaceConfigRequest is the request struct for api UpdateFaceConfig
type UpdateFaceConfigRequest struct {
	*requests.RpcRequest
	BizName  string `position:"Query" name:"BizName"`
	BizType  string `position:"Query" name:"BizType"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
}

// UpdateFaceConfigResponse is the response struct for api UpdateFaceConfig
type UpdateFaceConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateFaceConfigRequest creates a request to invoke UpdateFaceConfig API
func CreateUpdateFaceConfigRequest() (request *UpdateFaceConfigRequest) {
	request = &UpdateFaceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "UpdateFaceConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateFaceConfigResponse creates a response to parse from UpdateFaceConfig response
func CreateUpdateFaceConfigResponse() (response *UpdateFaceConfigResponse) {
	response = &UpdateFaceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
