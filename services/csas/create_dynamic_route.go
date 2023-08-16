package csas

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

// CreateDynamicRoute invokes the csas.CreateDynamicRoute API synchronously
func (client *Client) CreateDynamicRoute(request *CreateDynamicRouteRequest) (response *CreateDynamicRouteResponse, err error) {
	response = CreateCreateDynamicRouteResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDynamicRouteWithChan invokes the csas.CreateDynamicRoute API asynchronously
func (client *Client) CreateDynamicRouteWithChan(request *CreateDynamicRouteRequest) (<-chan *CreateDynamicRouteResponse, <-chan error) {
	responseChan := make(chan *CreateDynamicRouteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDynamicRoute(request)
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

// CreateDynamicRouteWithCallback invokes the csas.CreateDynamicRoute API asynchronously
func (client *Client) CreateDynamicRouteWithCallback(request *CreateDynamicRouteRequest, callback func(response *CreateDynamicRouteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDynamicRouteResponse
		var err error
		defer close(result)
		response, err = client.CreateDynamicRoute(request)
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

// CreateDynamicRouteRequest is the request struct for api CreateDynamicRoute
type CreateDynamicRouteRequest struct {
	*requests.RpcRequest
	Description      string           `position:"Body" name:"Description"`
	SourceIp         string           `position:"Query" name:"SourceIp"`
	DynamicRouteType string           `position:"Body" name:"DynamicRouteType"`
	TagIds           *[]string        `position:"Body" name:"TagIds"  type:"Repeated"`
	RegionIds        *[]string        `position:"Body" name:"RegionIds"  type:"Repeated"`
	Priority         requests.Integer `position:"Body" name:"Priority"`
	NextHop          string           `position:"Body" name:"NextHop"`
	ApplicationIds   *[]string        `position:"Body" name:"ApplicationIds"  type:"Repeated"`
	Name             string           `position:"Body" name:"Name"`
	ApplicationType  string           `position:"Body" name:"ApplicationType"`
	Status           string           `position:"Body" name:"Status"`
}

// CreateDynamicRouteResponse is the response struct for api CreateDynamicRoute
type CreateDynamicRouteResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DynamicRouteId string `json:"DynamicRouteId" xml:"DynamicRouteId"`
}

// CreateCreateDynamicRouteRequest creates a request to invoke CreateDynamicRoute API
func CreateCreateDynamicRouteRequest() (request *CreateDynamicRouteRequest) {
	request = &CreateDynamicRouteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "CreateDynamicRoute", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDynamicRouteResponse creates a response to parse from CreateDynamicRoute response
func CreateCreateDynamicRouteResponse() (response *CreateDynamicRouteResponse) {
	response = &CreateDynamicRouteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
