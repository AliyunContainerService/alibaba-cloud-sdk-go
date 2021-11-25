package voicenavigator

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

// EndDialogue invokes the voicenavigator.EndDialogue API synchronously
func (client *Client) EndDialogue(request *EndDialogueRequest) (response *EndDialogueResponse, err error) {
	response = CreateEndDialogueResponse()
	err = client.DoAction(request, response)
	return
}

// EndDialogueWithChan invokes the voicenavigator.EndDialogue API asynchronously
func (client *Client) EndDialogueWithChan(request *EndDialogueRequest) (<-chan *EndDialogueResponse, <-chan error) {
	responseChan := make(chan *EndDialogueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EndDialogue(request)
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

// EndDialogueWithCallback invokes the voicenavigator.EndDialogue API asynchronously
func (client *Client) EndDialogueWithCallback(request *EndDialogueRequest, callback func(response *EndDialogueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EndDialogueResponse
		var err error
		defer close(result)
		response, err = client.EndDialogue(request)
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

// EndDialogueRequest is the request struct for api EndDialogue
type EndDialogueRequest struct {
	*requests.RpcRequest
	ConversationId  string           `position:"Query" name:"ConversationId"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	InstanceOwnerId requests.Integer `position:"Query" name:"InstanceOwnerId"`
}

// EndDialogueResponse is the response struct for api EndDialogue
type EndDialogueResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEndDialogueRequest creates a request to invoke EndDialogue API
func CreateEndDialogueRequest() (request *EndDialogueRequest) {
	request = &EndDialogueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "EndDialogue", "voicebot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEndDialogueResponse creates a response to parse from EndDialogue response
func CreateEndDialogueResponse() (response *EndDialogueResponse) {
	response = &EndDialogueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
