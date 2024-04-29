package live

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

// DescribeLiveHttpsDomainList invokes the live.DescribeLiveHttpsDomainList API synchronously
func (client *Client) DescribeLiveHttpsDomainList(request *DescribeLiveHttpsDomainListRequest) (response *DescribeLiveHttpsDomainListResponse, err error) {
	response = CreateDescribeLiveHttpsDomainListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveHttpsDomainListWithChan invokes the live.DescribeLiveHttpsDomainList API asynchronously
func (client *Client) DescribeLiveHttpsDomainListWithChan(request *DescribeLiveHttpsDomainListRequest) (<-chan *DescribeLiveHttpsDomainListResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveHttpsDomainListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveHttpsDomainList(request)
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

// DescribeLiveHttpsDomainListWithCallback invokes the live.DescribeLiveHttpsDomainList API asynchronously
func (client *Client) DescribeLiveHttpsDomainListWithCallback(request *DescribeLiveHttpsDomainListRequest, callback func(response *DescribeLiveHttpsDomainListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveHttpsDomainListResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveHttpsDomainList(request)
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

// DescribeLiveHttpsDomainListRequest is the request struct for api DescribeLiveHttpsDomainList
type DescribeLiveHttpsDomainListRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Keyword    string           `position:"Query" name:"Keyword"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveHttpsDomainListResponse is the response struct for api DescribeLiveHttpsDomainList
type DescribeLiveHttpsDomainListResponse struct {
	*responses.BaseResponse
	RequestId  string                                 `json:"RequestId" xml:"RequestId"`
	TotalCount int                                    `json:"TotalCount" xml:"TotalCount"`
	CertInfos  CertInfosInDescribeLiveHttpsDomainList `json:"CertInfos" xml:"CertInfos"`
}

// CreateDescribeLiveHttpsDomainListRequest creates a request to invoke DescribeLiveHttpsDomainList API
func CreateDescribeLiveHttpsDomainListRequest() (request *DescribeLiveHttpsDomainListRequest) {
	request = &DescribeLiveHttpsDomainListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveHttpsDomainList", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveHttpsDomainListResponse creates a response to parse from DescribeLiveHttpsDomainList response
func CreateDescribeLiveHttpsDomainListResponse() (response *DescribeLiveHttpsDomainListResponse) {
	response = &DescribeLiveHttpsDomainListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
