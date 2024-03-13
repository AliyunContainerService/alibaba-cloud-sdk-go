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

// DescribeCloudAssistantAttributes invokes the swas_open.DescribeCloudAssistantAttributes API synchronously
func (client *Client) DescribeCloudAssistantAttributes(request *DescribeCloudAssistantAttributesRequest) (response *DescribeCloudAssistantAttributesResponse, err error) {
	response = CreateDescribeCloudAssistantAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudAssistantAttributesWithChan invokes the swas_open.DescribeCloudAssistantAttributes API asynchronously
func (client *Client) DescribeCloudAssistantAttributesWithChan(request *DescribeCloudAssistantAttributesRequest) (<-chan *DescribeCloudAssistantAttributesResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudAssistantAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudAssistantAttributes(request)
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

// DescribeCloudAssistantAttributesWithCallback invokes the swas_open.DescribeCloudAssistantAttributes API asynchronously
func (client *Client) DescribeCloudAssistantAttributesWithCallback(request *DescribeCloudAssistantAttributesRequest, callback func(response *DescribeCloudAssistantAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudAssistantAttributesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudAssistantAttributes(request)
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

// DescribeCloudAssistantAttributesRequest is the request struct for api DescribeCloudAssistantAttributes
type DescribeCloudAssistantAttributesRequest struct {
	*requests.RpcRequest
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	InstanceIds *[]string        `position:"Query" name:"InstanceIds"  type:"Json"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeCloudAssistantAttributesResponse is the response struct for api DescribeCloudAssistantAttributes
type DescribeCloudAssistantAttributesResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	TotalCount     int      `json:"TotalCount" xml:"TotalCount"`
	PageSize       int      `json:"PageSize" xml:"PageSize"`
	PageNumber     int      `json:"PageNumber" xml:"PageNumber"`
	CloudAssistant []Status `json:"CloudAssistant" xml:"CloudAssistant"`
}

// CreateDescribeCloudAssistantAttributesRequest creates a request to invoke DescribeCloudAssistantAttributes API
func CreateDescribeCloudAssistantAttributesRequest() (request *DescribeCloudAssistantAttributesRequest) {
	request = &DescribeCloudAssistantAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DescribeCloudAssistantAttributes", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudAssistantAttributesResponse creates a response to parse from DescribeCloudAssistantAttributes response
func CreateDescribeCloudAssistantAttributesResponse() (response *DescribeCloudAssistantAttributesResponse) {
	response = &DescribeCloudAssistantAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
