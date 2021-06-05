package workorder

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

// CreateTicket invokes the workorder.CreateTicket API synchronously
func (client *Client) CreateTicket(request *CreateTicketRequest) (response *CreateTicketResponse, err error) {
	response = CreateCreateTicketResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTicketWithChan invokes the workorder.CreateTicket API asynchronously
func (client *Client) CreateTicketWithChan(request *CreateTicketRequest) (<-chan *CreateTicketResponse, <-chan error) {
	responseChan := make(chan *CreateTicketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTicket(request)
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

// CreateTicketWithCallback invokes the workorder.CreateTicket API asynchronously
func (client *Client) CreateTicketWithCallback(request *CreateTicketRequest, callback func(response *CreateTicketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTicketResponse
		var err error
		defer close(result)
		response, err = client.CreateTicket(request)
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

// CreateTicketRequest is the request struct for api CreateTicket
type CreateTicketRequest struct {
	*requests.RpcRequest
	Severity    requests.Integer `position:"Body" name:"Severity"`
	Description string           `position:"Body" name:"Description"`
	CategoryId  string           `position:"Body" name:"CategoryId"`
}

// CreateTicketResponse is the response struct for api CreateTicket
type CreateTicketResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateTicketRequest creates a request to invoke CreateTicket API
func CreateCreateTicketRequest() (request *CreateTicketRequest) {
	request = &CreateTicketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Workorder", "2021-05-10", "CreateTicket", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateTicketResponse creates a response to parse from CreateTicket response
func CreateCreateTicketResponse() (response *CreateTicketResponse) {
	response = &CreateTicketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
