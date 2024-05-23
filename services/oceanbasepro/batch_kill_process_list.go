package oceanbasepro

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

// BatchKillProcessList invokes the oceanbasepro.BatchKillProcessList API synchronously
func (client *Client) BatchKillProcessList(request *BatchKillProcessListRequest) (response *BatchKillProcessListResponse, err error) {
	response = CreateBatchKillProcessListResponse()
	err = client.DoAction(request, response)
	return
}

// BatchKillProcessListWithChan invokes the oceanbasepro.BatchKillProcessList API asynchronously
func (client *Client) BatchKillProcessListWithChan(request *BatchKillProcessListRequest) (<-chan *BatchKillProcessListResponse, <-chan error) {
	responseChan := make(chan *BatchKillProcessListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchKillProcessList(request)
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

// BatchKillProcessListWithCallback invokes the oceanbasepro.BatchKillProcessList API asynchronously
func (client *Client) BatchKillProcessListWithCallback(request *BatchKillProcessListRequest, callback func(response *BatchKillProcessListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchKillProcessListResponse
		var err error
		defer close(result)
		response, err = client.BatchKillProcessList(request)
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

// BatchKillProcessListRequest is the request struct for api BatchKillProcessList
type BatchKillProcessListRequest struct {
	*requests.RpcRequest
	SessionList string `position:"Body" name:"SessionList"`
	InstanceId  string `position:"Body" name:"InstanceId"`
	TenantId    string `position:"Body" name:"TenantId"`
}

// BatchKillProcessListResponse is the response struct for api BatchKillProcessList
type BatchKillProcessListResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateBatchKillProcessListRequest creates a request to invoke BatchKillProcessList API
func CreateBatchKillProcessListRequest() (request *BatchKillProcessListRequest) {
	request = &BatchKillProcessListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "BatchKillProcessList", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchKillProcessListResponse creates a response to parse from BatchKillProcessList response
func CreateBatchKillProcessListResponse() (response *BatchKillProcessListResponse) {
	response = &BatchKillProcessListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
