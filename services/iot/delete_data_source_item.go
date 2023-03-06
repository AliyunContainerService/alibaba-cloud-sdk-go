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

// DeleteDataSourceItem invokes the iot.DeleteDataSourceItem API synchronously
func (client *Client) DeleteDataSourceItem(request *DeleteDataSourceItemRequest) (response *DeleteDataSourceItemResponse, err error) {
	response = CreateDeleteDataSourceItemResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDataSourceItemWithChan invokes the iot.DeleteDataSourceItem API asynchronously
func (client *Client) DeleteDataSourceItemWithChan(request *DeleteDataSourceItemRequest) (<-chan *DeleteDataSourceItemResponse, <-chan error) {
	responseChan := make(chan *DeleteDataSourceItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDataSourceItem(request)
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

// DeleteDataSourceItemWithCallback invokes the iot.DeleteDataSourceItem API asynchronously
func (client *Client) DeleteDataSourceItemWithCallback(request *DeleteDataSourceItemRequest, callback func(response *DeleteDataSourceItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDataSourceItemResponse
		var err error
		defer close(result)
		response, err = client.DeleteDataSourceItem(request)
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

// DeleteDataSourceItemRequest is the request struct for api DeleteDataSourceItem
type DeleteDataSourceItemRequest struct {
	*requests.RpcRequest
	IotInstanceId    string           `position:"Query" name:"IotInstanceId"`
	DataSourceItemId requests.Integer `position:"Query" name:"DataSourceItemId"`
	ApiProduct       string           `position:"Body" name:"ApiProduct"`
	DataSourceId     requests.Integer `position:"Query" name:"DataSourceId"`
	ApiRevision      string           `position:"Body" name:"ApiRevision"`
}

// DeleteDataSourceItemResponse is the response struct for api DeleteDataSourceItem
type DeleteDataSourceItemResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateDeleteDataSourceItemRequest creates a request to invoke DeleteDataSourceItem API
func CreateDeleteDataSourceItemRequest() (request *DeleteDataSourceItemRequest) {
	request = &DeleteDataSourceItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteDataSourceItem", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDataSourceItemResponse creates a response to parse from DeleteDataSourceItem response
func CreateDeleteDataSourceItemResponse() (response *DeleteDataSourceItemResponse) {
	response = &DeleteDataSourceItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
