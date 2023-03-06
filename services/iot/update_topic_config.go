package iot

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

// UpdateTopicConfig invokes the iot.UpdateTopicConfig API synchronously
func (client *Client) UpdateTopicConfig(request *UpdateTopicConfigRequest) (response *UpdateTopicConfigResponse, err error) {
	response = CreateUpdateTopicConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTopicConfigWithChan invokes the iot.UpdateTopicConfig API asynchronously
func (client *Client) UpdateTopicConfigWithChan(request *UpdateTopicConfigRequest) (<-chan *UpdateTopicConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateTopicConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTopicConfig(request)
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

// UpdateTopicConfigWithCallback invokes the iot.UpdateTopicConfig API asynchronously
func (client *Client) UpdateTopicConfigWithCallback(request *UpdateTopicConfigRequest, callback func(response *UpdateTopicConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTopicConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateTopicConfig(request)
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

// UpdateTopicConfigRequest is the request struct for api UpdateTopicConfig
type UpdateTopicConfigRequest struct {
	*requests.RpcRequest
	IotInstanceId   string           `position:"Query" name:"IotInstanceId"`
	TopicFullName   string           `position:"Query" name:"TopicFullName"`
	EnableBroadcast requests.Boolean `position:"Query" name:"EnableBroadcast"`
	ProductKey      string           `position:"Query" name:"ProductKey"`
	ApiProduct      string           `position:"Body" name:"ApiProduct"`
	ApiRevision     string           `position:"Body" name:"ApiRevision"`
}

// UpdateTopicConfigResponse is the response struct for api UpdateTopicConfig
type UpdateTopicConfigResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateTopicConfigRequest creates a request to invoke UpdateTopicConfig API
func CreateUpdateTopicConfigRequest() (request *UpdateTopicConfigRequest) {
	request = &UpdateTopicConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateTopicConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateTopicConfigResponse creates a response to parse from UpdateTopicConfig response
func CreateUpdateTopicConfigResponse() (response *UpdateTopicConfigResponse) {
	response = &UpdateTopicConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
