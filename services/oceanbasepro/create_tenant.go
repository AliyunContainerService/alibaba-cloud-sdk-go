package oceanbasepro

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

// CreateTenant invokes the oceanbasepro.CreateTenant API synchronously
func (client *Client) CreateTenant(request *CreateTenantRequest) (response *CreateTenantResponse, err error) {
	response = CreateCreateTenantResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTenantWithChan invokes the oceanbasepro.CreateTenant API asynchronously
func (client *Client) CreateTenantWithChan(request *CreateTenantRequest) (<-chan *CreateTenantResponse, <-chan error) {
	responseChan := make(chan *CreateTenantResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTenant(request)
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

// CreateTenantWithCallback invokes the oceanbasepro.CreateTenant API asynchronously
func (client *Client) CreateTenantWithCallback(request *CreateTenantRequest, callback func(response *CreateTenantResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTenantResponse
		var err error
		defer close(result)
		response, err = client.CreateTenant(request)
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

// CreateTenantRequest is the request struct for api CreateTenant
type CreateTenantRequest struct {
	*requests.RpcRequest
	Charset          string            `position:"Body" name:"Charset"`
	TenantMode       string            `position:"Body" name:"TenantMode"`
	Memory           requests.Integer  `position:"Body" name:"Memory"`
	LogDisk          requests.Integer  `position:"Body" name:"LogDisk"`
	TimeZone         string            `position:"Body" name:"TimeZone"`
	Description      string            `position:"Body" name:"Description"`
	UserVSwitchId    string            `position:"Body" name:"UserVSwitchId"`
	UserVpcId        string            `position:"Body" name:"UserVpcId"`
	Cpu              requests.Integer  `position:"Body" name:"Cpu"`
	UnitNum          requests.Integer  `position:"Body" name:"UnitNum"`
	InstanceId       string            `position:"Body" name:"InstanceId"`
	PrimaryZone      string            `position:"Body" name:"PrimaryZone"`
	UserVpcOwnerId   string            `position:"Body" name:"UserVpcOwnerId"`
	CreateParams     map[string]string `position:"Body" name:"CreateParams"  type:"Map"`
	TenantName       string            `position:"Body" name:"TenantName"`
	ReadOnlyZoneList string            `position:"Body" name:"ReadOnlyZoneList"`
}

// CreateTenantResponse is the response struct for api CreateTenant
type CreateTenantResponse struct {
	*responses.BaseResponse
	TenantId  string `json:"TenantId" xml:"TenantId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTenantRequest creates a request to invoke CreateTenant API
func CreateCreateTenantRequest() (request *CreateTenantRequest) {
	request = &CreateTenantRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "CreateTenant", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTenantResponse creates a response to parse from CreateTenant response
func CreateCreateTenantResponse() (response *CreateTenantResponse) {
	response = &CreateTenantResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
