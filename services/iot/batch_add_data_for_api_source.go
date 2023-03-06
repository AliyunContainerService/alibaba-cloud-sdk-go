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

// BatchAddDataForApiSource invokes the iot.BatchAddDataForApiSource API synchronously
func (client *Client) BatchAddDataForApiSource(request *BatchAddDataForApiSourceRequest) (response *BatchAddDataForApiSourceResponse, err error) {
	response = CreateBatchAddDataForApiSourceResponse()
	err = client.DoAction(request, response)
	return
}

// BatchAddDataForApiSourceWithChan invokes the iot.BatchAddDataForApiSource API asynchronously
func (client *Client) BatchAddDataForApiSourceWithChan(request *BatchAddDataForApiSourceRequest) (<-chan *BatchAddDataForApiSourceResponse, <-chan error) {
	responseChan := make(chan *BatchAddDataForApiSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchAddDataForApiSource(request)
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

// BatchAddDataForApiSourceWithCallback invokes the iot.BatchAddDataForApiSource API asynchronously
func (client *Client) BatchAddDataForApiSourceWithCallback(request *BatchAddDataForApiSourceRequest, callback func(response *BatchAddDataForApiSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchAddDataForApiSourceResponse
		var err error
		defer close(result)
		response, err = client.BatchAddDataForApiSource(request)
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

// BatchAddDataForApiSourceRequest is the request struct for api BatchAddDataForApiSource
type BatchAddDataForApiSourceRequest struct {
	*requests.RpcRequest
	ContentList   string `position:"Query" name:"ContentList"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// BatchAddDataForApiSourceResponse is the response struct for api BatchAddDataForApiSource
type BatchAddDataForApiSourceResponse struct {
	*responses.BaseResponse
	RequestId    string                 `json:"RequestId" xml:"RequestId"`
	Success      bool                   `json:"Success" xml:"Success"`
	Code         string                 `json:"Code" xml:"Code"`
	ErrorMessage string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateBatchAddDataForApiSourceRequest creates a request to invoke BatchAddDataForApiSource API
func CreateBatchAddDataForApiSourceRequest() (request *BatchAddDataForApiSourceRequest) {
	request = &BatchAddDataForApiSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchAddDataForApiSource", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchAddDataForApiSourceResponse creates a response to parse from BatchAddDataForApiSource response
func CreateBatchAddDataForApiSourceResponse() (response *BatchAddDataForApiSourceResponse) {
	response = &BatchAddDataForApiSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
