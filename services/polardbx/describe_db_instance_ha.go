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

// DescribeDBInstanceHA invokes the polardbx.DescribeDBInstanceHA API synchronously
func (client *Client) DescribeDBInstanceHA(request *DescribeDBInstanceHARequest) (response *DescribeDBInstanceHAResponse, err error) {
	response = CreateDescribeDBInstanceHAResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceHAWithChan invokes the polardbx.DescribeDBInstanceHA API asynchronously
func (client *Client) DescribeDBInstanceHAWithChan(request *DescribeDBInstanceHARequest) (<-chan *DescribeDBInstanceHAResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceHAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceHA(request)
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

// DescribeDBInstanceHAWithCallback invokes the polardbx.DescribeDBInstanceHA API asynchronously
func (client *Client) DescribeDBInstanceHAWithCallback(request *DescribeDBInstanceHARequest, callback func(response *DescribeDBInstanceHAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceHAResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceHA(request)
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

// DescribeDBInstanceHARequest is the request struct for api DescribeDBInstanceHA
type DescribeDBInstanceHARequest struct {
	*requests.RpcRequest
	DBInstanceName string `position:"Query" name:"DBInstanceName"`
}

// DescribeDBInstanceHAResponse is the response struct for api DescribeDBInstanceHA
type DescribeDBInstanceHAResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeDBInstanceHARequest creates a request to invoke DescribeDBInstanceHA API
func CreateDescribeDBInstanceHARequest() (request *DescribeDBInstanceHARequest) {
	request = &DescribeDBInstanceHARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "DescribeDBInstanceHA", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceHAResponse creates a response to parse from DescribeDBInstanceHA response
func CreateDescribeDBInstanceHAResponse() (response *DescribeDBInstanceHAResponse) {
	response = &DescribeDBInstanceHAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
