package ddoscoo

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

// ModifyElasticBizQps invokes the ddoscoo.ModifyElasticBizQps API synchronously
func (client *Client) ModifyElasticBizQps(request *ModifyElasticBizQpsRequest) (response *ModifyElasticBizQpsResponse, err error) {
	response = CreateModifyElasticBizQpsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyElasticBizQpsWithChan invokes the ddoscoo.ModifyElasticBizQps API asynchronously
func (client *Client) ModifyElasticBizQpsWithChan(request *ModifyElasticBizQpsRequest) (<-chan *ModifyElasticBizQpsResponse, <-chan error) {
	responseChan := make(chan *ModifyElasticBizQpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyElasticBizQps(request)
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

// ModifyElasticBizQpsWithCallback invokes the ddoscoo.ModifyElasticBizQps API asynchronously
func (client *Client) ModifyElasticBizQpsWithCallback(request *ModifyElasticBizQpsRequest, callback func(response *ModifyElasticBizQpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyElasticBizQpsResponse
		var err error
		defer close(result)
		response, err = client.ModifyElasticBizQps(request)
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

// ModifyElasticBizQpsRequest is the request struct for api ModifyElasticBizQps
type ModifyElasticBizQpsRequest struct {
	*requests.RpcRequest
	Mode          string           `position:"Query" name:"Mode"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	SourceIp      string           `position:"Query" name:"SourceIp"`
	OpsElasticQps requests.Integer `position:"Query" name:"OpsElasticQps"`
}

// ModifyElasticBizQpsResponse is the response struct for api ModifyElasticBizQps
type ModifyElasticBizQpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyElasticBizQpsRequest creates a request to invoke ModifyElasticBizQps API
func CreateModifyElasticBizQpsRequest() (request *ModifyElasticBizQpsRequest) {
	request = &ModifyElasticBizQpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ModifyElasticBizQps", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyElasticBizQpsResponse creates a response to parse from ModifyElasticBizQps response
func CreateModifyElasticBizQpsResponse() (response *ModifyElasticBizQpsResponse) {
	response = &ModifyElasticBizQpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
