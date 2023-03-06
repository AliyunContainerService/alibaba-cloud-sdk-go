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

// QueryEdgeInstanceChannel invokes the iot.QueryEdgeInstanceChannel API synchronously
func (client *Client) QueryEdgeInstanceChannel(request *QueryEdgeInstanceChannelRequest) (response *QueryEdgeInstanceChannelResponse, err error) {
	response = CreateQueryEdgeInstanceChannelResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEdgeInstanceChannelWithChan invokes the iot.QueryEdgeInstanceChannel API asynchronously
func (client *Client) QueryEdgeInstanceChannelWithChan(request *QueryEdgeInstanceChannelRequest) (<-chan *QueryEdgeInstanceChannelResponse, <-chan error) {
	responseChan := make(chan *QueryEdgeInstanceChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEdgeInstanceChannel(request)
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

// QueryEdgeInstanceChannelWithCallback invokes the iot.QueryEdgeInstanceChannel API asynchronously
func (client *Client) QueryEdgeInstanceChannelWithCallback(request *QueryEdgeInstanceChannelRequest, callback func(response *QueryEdgeInstanceChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEdgeInstanceChannelResponse
		var err error
		defer close(result)
		response, err = client.QueryEdgeInstanceChannel(request)
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

// QueryEdgeInstanceChannelRequest is the request struct for api QueryEdgeInstanceChannel
type QueryEdgeInstanceChannelRequest struct {
	*requests.RpcRequest
	DriverId      string           `position:"Query" name:"DriverId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ChannelName   string           `position:"Query" name:"ChannelName"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryEdgeInstanceChannelResponse is the response struct for api QueryEdgeInstanceChannel
type QueryEdgeInstanceChannelResponse struct {
	*responses.BaseResponse
	RequestId    string                         `json:"RequestId" xml:"RequestId"`
	Success      bool                           `json:"Success" xml:"Success"`
	Code         string                         `json:"Code" xml:"Code"`
	ErrorMessage string                         `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryEdgeInstanceChannel `json:"Data" xml:"Data"`
}

// CreateQueryEdgeInstanceChannelRequest creates a request to invoke QueryEdgeInstanceChannel API
func CreateQueryEdgeInstanceChannelRequest() (request *QueryEdgeInstanceChannelRequest) {
	request = &QueryEdgeInstanceChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryEdgeInstanceChannel", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryEdgeInstanceChannelResponse creates a response to parse from QueryEdgeInstanceChannel response
func CreateQueryEdgeInstanceChannelResponse() (response *QueryEdgeInstanceChannelResponse) {
	response = &QueryEdgeInstanceChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
