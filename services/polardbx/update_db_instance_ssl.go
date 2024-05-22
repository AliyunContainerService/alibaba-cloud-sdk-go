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

// UpdateDBInstanceSSL invokes the polardbx.UpdateDBInstanceSSL API synchronously
func (client *Client) UpdateDBInstanceSSL(request *UpdateDBInstanceSSLRequest) (response *UpdateDBInstanceSSLResponse, err error) {
	response = CreateUpdateDBInstanceSSLResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDBInstanceSSLWithChan invokes the polardbx.UpdateDBInstanceSSL API asynchronously
func (client *Client) UpdateDBInstanceSSLWithChan(request *UpdateDBInstanceSSLRequest) (<-chan *UpdateDBInstanceSSLResponse, <-chan error) {
	responseChan := make(chan *UpdateDBInstanceSSLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDBInstanceSSL(request)
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

// UpdateDBInstanceSSLWithCallback invokes the polardbx.UpdateDBInstanceSSL API asynchronously
func (client *Client) UpdateDBInstanceSSLWithCallback(request *UpdateDBInstanceSSLRequest, callback func(response *UpdateDBInstanceSSLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDBInstanceSSLResponse
		var err error
		defer close(result)
		response, err = client.UpdateDBInstanceSSL(request)
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

// UpdateDBInstanceSSLRequest is the request struct for api UpdateDBInstanceSSL
type UpdateDBInstanceSSLRequest struct {
	*requests.RpcRequest
	DBInstanceName string           `position:"Query" name:"DBInstanceName"`
	CertCommonName string           `position:"Query" name:"CertCommonName"`
	EnableSSL      requests.Boolean `position:"Query" name:"EnableSSL"`
	Force          requests.Boolean `position:"Query" name:"Force"`
}

// UpdateDBInstanceSSLResponse is the response struct for api UpdateDBInstanceSSL
type UpdateDBInstanceSSLResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateDBInstanceSSLRequest creates a request to invoke UpdateDBInstanceSSL API
func CreateUpdateDBInstanceSSLRequest() (request *UpdateDBInstanceSSLRequest) {
	request = &UpdateDBInstanceSSLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "UpdateDBInstanceSSL", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDBInstanceSSLResponse creates a response to parse from UpdateDBInstanceSSL response
func CreateUpdateDBInstanceSSLResponse() (response *UpdateDBInstanceSSLResponse) {
	response = &UpdateDBInstanceSSLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
