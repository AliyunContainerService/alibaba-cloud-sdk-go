package apds

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

// DeleteSurveyJob invokes the apds.DeleteSurveyJob API synchronously
func (client *Client) DeleteSurveyJob(request *DeleteSurveyJobRequest) (response *DeleteSurveyJobResponse, err error) {
	response = CreateDeleteSurveyJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSurveyJobWithChan invokes the apds.DeleteSurveyJob API asynchronously
func (client *Client) DeleteSurveyJobWithChan(request *DeleteSurveyJobRequest) (<-chan *DeleteSurveyJobResponse, <-chan error) {
	responseChan := make(chan *DeleteSurveyJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSurveyJob(request)
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

// DeleteSurveyJobWithCallback invokes the apds.DeleteSurveyJob API asynchronously
func (client *Client) DeleteSurveyJobWithCallback(request *DeleteSurveyJobRequest, callback func(response *DeleteSurveyJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSurveyJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteSurveyJob(request)
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

// DeleteSurveyJobRequest is the request struct for api DeleteSurveyJob
type DeleteSurveyJobRequest struct {
	*requests.RoaRequest
	Id string `position:"Query" name:"id"`
}

// DeleteSurveyJobResponse is the response struct for api DeleteSurveyJob
type DeleteSurveyJobResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"error" xml:"error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateDeleteSurveyJobRequest creates a request to invoke DeleteSurveyJob API
func CreateDeleteSurveyJobRequest() (request *DeleteSurveyJobRequest) {
	request = &DeleteSurveyJobRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "DeleteSurveyJob", "/okss-services/winback/delete-survey-job", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteSurveyJobResponse creates a response to parse from DeleteSurveyJob response
func CreateDeleteSurveyJobResponse() (response *DeleteSurveyJobResponse) {
	response = &DeleteSurveyJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
