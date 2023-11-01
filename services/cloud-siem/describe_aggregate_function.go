package cloud_siem

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

// DescribeAggregateFunction invokes the cloud_siem.DescribeAggregateFunction API synchronously
func (client *Client) DescribeAggregateFunction(request *DescribeAggregateFunctionRequest) (response *DescribeAggregateFunctionResponse, err error) {
	response = CreateDescribeAggregateFunctionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAggregateFunctionWithChan invokes the cloud_siem.DescribeAggregateFunction API asynchronously
func (client *Client) DescribeAggregateFunctionWithChan(request *DescribeAggregateFunctionRequest) (<-chan *DescribeAggregateFunctionResponse, <-chan error) {
	responseChan := make(chan *DescribeAggregateFunctionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAggregateFunction(request)
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

// DescribeAggregateFunctionWithCallback invokes the cloud_siem.DescribeAggregateFunction API asynchronously
func (client *Client) DescribeAggregateFunctionWithCallback(request *DescribeAggregateFunctionRequest, callback func(response *DescribeAggregateFunctionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAggregateFunctionResponse
		var err error
		defer close(result)
		response, err = client.DescribeAggregateFunction(request)
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

// DescribeAggregateFunctionRequest is the request struct for api DescribeAggregateFunction
type DescribeAggregateFunctionRequest struct {
	*requests.RpcRequest
}

// DescribeAggregateFunctionResponse is the response struct for api DescribeAggregateFunction
type DescribeAggregateFunctionResponse struct {
	*responses.BaseResponse
	Success   bool       `json:"Success" xml:"Success"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeAggregateFunctionRequest creates a request to invoke DescribeAggregateFunction API
func CreateDescribeAggregateFunctionRequest() (request *DescribeAggregateFunctionRequest) {
	request = &DescribeAggregateFunctionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeAggregateFunction", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAggregateFunctionResponse creates a response to parse from DescribeAggregateFunction response
func CreateDescribeAggregateFunctionResponse() (response *DescribeAggregateFunctionResponse) {
	response = &DescribeAggregateFunctionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
