package cloud_siem

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

// ListCloudSiemPredefinedRules invokes the cloud_siem.ListCloudSiemPredefinedRules API synchronously
func (client *Client) ListCloudSiemPredefinedRules(request *ListCloudSiemPredefinedRulesRequest) (response *ListCloudSiemPredefinedRulesResponse, err error) {
	response = CreateListCloudSiemPredefinedRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListCloudSiemPredefinedRulesWithChan invokes the cloud_siem.ListCloudSiemPredefinedRules API asynchronously
func (client *Client) ListCloudSiemPredefinedRulesWithChan(request *ListCloudSiemPredefinedRulesRequest) (<-chan *ListCloudSiemPredefinedRulesResponse, <-chan error) {
	responseChan := make(chan *ListCloudSiemPredefinedRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCloudSiemPredefinedRules(request)
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

// ListCloudSiemPredefinedRulesWithCallback invokes the cloud_siem.ListCloudSiemPredefinedRules API asynchronously
func (client *Client) ListCloudSiemPredefinedRulesWithCallback(request *ListCloudSiemPredefinedRulesRequest, callback func(response *ListCloudSiemPredefinedRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCloudSiemPredefinedRulesResponse
		var err error
		defer close(result)
		response, err = client.ListCloudSiemPredefinedRules(request)
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

// ListCloudSiemPredefinedRulesRequest is the request struct for api ListCloudSiemPredefinedRules
type ListCloudSiemPredefinedRulesRequest struct {
	*requests.RpcRequest
	RuleName    string           `position:"Body" name:"RuleName"`
	StartTime   requests.Integer `position:"Body" name:"StartTime"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	Id          string           `position:"Body" name:"Id"`
	RuleType    string           `position:"Body" name:"RuleType"`
	EndTime     requests.Integer `position:"Body" name:"EndTime"`
	CurrentPage requests.Integer `position:"Body" name:"CurrentPage"`
	AlertType   string           `position:"Body" name:"AlertType"`
	ThreatLevel *[]string        `position:"Body" name:"ThreatLevel"  type:"Repeated"`
	Status      requests.Integer `position:"Body" name:"Status"`
}

// ListCloudSiemPredefinedRulesResponse is the response struct for api ListCloudSiemPredefinedRules
type ListCloudSiemPredefinedRulesResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListCloudSiemPredefinedRulesRequest creates a request to invoke ListCloudSiemPredefinedRules API
func CreateListCloudSiemPredefinedRulesRequest() (request *ListCloudSiemPredefinedRulesRequest) {
	request = &ListCloudSiemPredefinedRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListCloudSiemPredefinedRules", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListCloudSiemPredefinedRulesResponse creates a response to parse from ListCloudSiemPredefinedRules response
func CreateListCloudSiemPredefinedRulesResponse() (response *ListCloudSiemPredefinedRulesResponse) {
	response = &ListCloudSiemPredefinedRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
