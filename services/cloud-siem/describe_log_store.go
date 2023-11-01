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

// DescribeLogStore invokes the cloud_siem.DescribeLogStore API synchronously
func (client *Client) DescribeLogStore(request *DescribeLogStoreRequest) (response *DescribeLogStoreResponse, err error) {
	response = CreateDescribeLogStoreResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLogStoreWithChan invokes the cloud_siem.DescribeLogStore API asynchronously
func (client *Client) DescribeLogStoreWithChan(request *DescribeLogStoreRequest) (<-chan *DescribeLogStoreResponse, <-chan error) {
	responseChan := make(chan *DescribeLogStoreResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLogStore(request)
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

// DescribeLogStoreWithCallback invokes the cloud_siem.DescribeLogStore API asynchronously
func (client *Client) DescribeLogStoreWithCallback(request *DescribeLogStoreRequest, callback func(response *DescribeLogStoreResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLogStoreResponse
		var err error
		defer close(result)
		response, err = client.DescribeLogStore(request)
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

// DescribeLogStoreRequest is the request struct for api DescribeLogStore
type DescribeLogStoreRequest struct {
	*requests.RpcRequest
}

// DescribeLogStoreResponse is the response struct for api DescribeLogStore
type DescribeLogStoreResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	ErrCode   string `json:"ErrCode" xml:"ErrCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	DyCode    string `json:"DyCode" xml:"DyCode"`
	DyMessage string `json:"DyMessage" xml:"DyMessage"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeLogStoreRequest creates a request to invoke DescribeLogStore API
func CreateDescribeLogStoreRequest() (request *DescribeLogStoreRequest) {
	request = &DescribeLogStoreRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeLogStore", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLogStoreResponse creates a response to parse from DescribeLogStore response
func CreateDescribeLogStoreResponse() (response *DescribeLogStoreResponse) {
	response = &DescribeLogStoreResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
