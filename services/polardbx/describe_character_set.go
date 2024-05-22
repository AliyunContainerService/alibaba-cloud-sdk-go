package polardbx

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

// DescribeCharacterSet invokes the polardbx.DescribeCharacterSet API synchronously
func (client *Client) DescribeCharacterSet(request *DescribeCharacterSetRequest) (response *DescribeCharacterSetResponse, err error) {
	response = CreateDescribeCharacterSetResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCharacterSetWithChan invokes the polardbx.DescribeCharacterSet API asynchronously
func (client *Client) DescribeCharacterSetWithChan(request *DescribeCharacterSetRequest) (<-chan *DescribeCharacterSetResponse, <-chan error) {
	responseChan := make(chan *DescribeCharacterSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCharacterSet(request)
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

// DescribeCharacterSetWithCallback invokes the polardbx.DescribeCharacterSet API asynchronously
func (client *Client) DescribeCharacterSetWithCallback(request *DescribeCharacterSetRequest, callback func(response *DescribeCharacterSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCharacterSetResponse
		var err error
		defer close(result)
		response, err = client.DescribeCharacterSet(request)
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

// DescribeCharacterSetRequest is the request struct for api DescribeCharacterSet
type DescribeCharacterSetRequest struct {
	*requests.RpcRequest
	DBInstanceName string `position:"Query" name:"DBInstanceName"`
}

// DescribeCharacterSetResponse is the response struct for api DescribeCharacterSet
type DescribeCharacterSetResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeCharacterSetRequest creates a request to invoke DescribeCharacterSet API
func CreateDescribeCharacterSetRequest() (request *DescribeCharacterSetRequest) {
	request = &DescribeCharacterSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "DescribeCharacterSet", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCharacterSetResponse creates a response to parse from DescribeCharacterSet response
func CreateDescribeCharacterSetResponse() (response *DescribeCharacterSetResponse) {
	response = &DescribeCharacterSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
