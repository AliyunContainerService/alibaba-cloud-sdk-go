package swas_open

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

// DescribeCommands invokes the swas_open.DescribeCommands API synchronously
func (client *Client) DescribeCommands(request *DescribeCommandsRequest) (response *DescribeCommandsResponse, err error) {
	response = CreateDescribeCommandsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCommandsWithChan invokes the swas_open.DescribeCommands API asynchronously
func (client *Client) DescribeCommandsWithChan(request *DescribeCommandsRequest) (<-chan *DescribeCommandsResponse, <-chan error) {
	responseChan := make(chan *DescribeCommandsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCommands(request)
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

// DescribeCommandsWithCallback invokes the swas_open.DescribeCommands API asynchronously
func (client *Client) DescribeCommandsWithCallback(request *DescribeCommandsRequest, callback func(response *DescribeCommandsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCommandsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCommands(request)
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

// DescribeCommandsRequest is the request struct for api DescribeCommands
type DescribeCommandsRequest struct {
	*requests.RpcRequest
	Type       string                 `position:"Query" name:"Type"`
	CommandId  string                 `position:"Query" name:"CommandId"`
	PageNumber string                 `position:"Query" name:"PageNumber"`
	Provider   string                 `position:"Query" name:"Provider"`
	PageSize   string                 `position:"Query" name:"PageSize"`
	Tag        *[]DescribeCommandsTag `position:"Query" name:"Tag"  type:"Repeated"`
	Name       string                 `position:"Query" name:"Name"`
}

// DescribeCommandsTag is a repeated param struct in DescribeCommandsRequest
type DescribeCommandsTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeCommandsResponse is the response struct for api DescribeCommands
type DescribeCommandsResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	Commands   []Command `json:"Commands" xml:"Commands"`
}

// CreateDescribeCommandsRequest creates a request to invoke DescribeCommands API
func CreateDescribeCommandsRequest() (request *DescribeCommandsRequest) {
	request = &DescribeCommandsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DescribeCommands", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCommandsResponse creates a response to parse from DescribeCommands response
func CreateDescribeCommandsResponse() (response *DescribeCommandsResponse) {
	response = &DescribeCommandsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
