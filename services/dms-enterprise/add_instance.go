package dms_enterprise

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

// AddInstance invokes the dms_enterprise.AddInstance API synchronously
func (client *Client) AddInstance(request *AddInstanceRequest) (response *AddInstanceResponse, err error) {
	response = CreateAddInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// AddInstanceWithChan invokes the dms_enterprise.AddInstance API asynchronously
func (client *Client) AddInstanceWithChan(request *AddInstanceRequest) (<-chan *AddInstanceResponse, <-chan error) {
	responseChan := make(chan *AddInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddInstance(request)
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

// AddInstanceWithCallback invokes the dms_enterprise.AddInstance API asynchronously
func (client *Client) AddInstanceWithCallback(request *AddInstanceRequest, callback func(response *AddInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddInstanceResponse
		var err error
		defer close(result)
		response, err = client.AddInstance(request)
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

// AddInstanceRequest is the request struct for api AddInstance
type AddInstanceRequest struct {
	*requests.RpcRequest
	EcsRegion        string           `position:"Query" name:"EcsRegion"`
	NetworkType      string           `position:"Query" name:"NetworkType"`
	Tid              requests.Integer `position:"Query" name:"Tid"`
	UseSsl           requests.Integer `position:"Query" name:"UseSsl"`
	EnableSellCommon string           `position:"Query" name:"EnableSellCommon"`
	EnableSellSitd   string           `position:"Query" name:"EnableSellSitd"`
	InstanceSource   string           `position:"Query" name:"InstanceSource"`
	EnvType          string           `position:"Query" name:"EnvType"`
	Host             string           `position:"Query" name:"Host"`
	QueryTimeout     requests.Integer `position:"Query" name:"QueryTimeout"`
	EcsInstanceId    string           `position:"Query" name:"EcsInstanceId"`
	ExportTimeout    requests.Integer `position:"Query" name:"ExportTimeout"`
	TemplateId       requests.Integer `position:"Query" name:"TemplateId"`
	EnableSellTrust  string           `position:"Query" name:"EnableSellTrust"`
	Port             requests.Integer `position:"Query" name:"Port"`
	SafeRule         string           `position:"Query" name:"SafeRule"`
	DdlOnline        requests.Integer `position:"Query" name:"DdlOnline"`
	UseDsql          requests.Integer `position:"Query" name:"UseDsql"`
	EnableSellStable string           `position:"Query" name:"EnableSellStable"`
	Sid              string           `position:"Query" name:"Sid"`
	DbaId            requests.Integer `position:"Query" name:"DbaId"`
	DataLinkName     string           `position:"Query" name:"DataLinkName"`
	TemplateType     string           `position:"Query" name:"TemplateType"`
	InstanceType     string           `position:"Query" name:"InstanceType"`
	DatabasePassword string           `position:"Query" name:"DatabasePassword"`
	InstanceAlias    string           `position:"Query" name:"InstanceAlias"`
	DatabaseUser     string           `position:"Query" name:"DatabaseUser"`
	VpcId            string           `position:"Query" name:"VpcId"`
	SkipTest         requests.Boolean `position:"Query" name:"SkipTest"`
}

// AddInstanceResponse is the response struct for api AddInstance
type AddInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateAddInstanceRequest creates a request to invoke AddInstance API
func CreateAddInstanceRequest() (request *AddInstanceRequest) {
	request = &AddInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "AddInstance", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddInstanceResponse creates a response to parse from AddInstance response
func CreateAddInstanceResponse() (response *AddInstanceResponse) {
	response = &AddInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
