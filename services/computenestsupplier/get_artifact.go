package computenestsupplier

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

// GetArtifact invokes the computenestsupplier.GetArtifact API synchronously
func (client *Client) GetArtifact(request *GetArtifactRequest) (response *GetArtifactResponse, err error) {
	response = CreateGetArtifactResponse()
	err = client.DoAction(request, response)
	return
}

// GetArtifactWithChan invokes the computenestsupplier.GetArtifact API asynchronously
func (client *Client) GetArtifactWithChan(request *GetArtifactRequest) (<-chan *GetArtifactResponse, <-chan error) {
	responseChan := make(chan *GetArtifactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetArtifact(request)
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

// GetArtifactWithCallback invokes the computenestsupplier.GetArtifact API asynchronously
func (client *Client) GetArtifactWithCallback(request *GetArtifactRequest, callback func(response *GetArtifactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetArtifactResponse
		var err error
		defer close(result)
		response, err = client.GetArtifact(request)
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

// GetArtifactRequest is the request struct for api GetArtifact
type GetArtifactRequest struct {
	*requests.RpcRequest
	ArtifactVersion string `position:"Query" name:"ArtifactVersion"`
	ArtifactName    string `position:"Query" name:"ArtifactName"`
	ArtifactId      string `position:"Query" name:"ArtifactId"`
}

// GetArtifactResponse is the response struct for api GetArtifact
type GetArtifactResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ArtifactId       string `json:"ArtifactId" xml:"ArtifactId"`
	ArtifactType     string `json:"ArtifactType" xml:"ArtifactType"`
	Name             string `json:"Name" xml:"Name"`
	VersionName      string `json:"VersionName" xml:"VersionName"`
	ArtifactVersion  string `json:"ArtifactVersion" xml:"ArtifactVersion"`
	Description      string `json:"Description" xml:"Description"`
	GmtModified      string `json:"GmtModified" xml:"GmtModified"`
	Status           string `json:"Status" xml:"Status"`
	MaxVersion       int64  `json:"MaxVersion" xml:"MaxVersion"`
	ArtifactProperty string `json:"ArtifactProperty" xml:"ArtifactProperty"`
	SupportRegionIds string `json:"SupportRegionIds" xml:"SupportRegionIds"`
	Progress         string `json:"Progress" xml:"Progress"`
	ResourceGroupId  string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tags             []Tag  `json:"Tags" xml:"Tags"`
}

// CreateGetArtifactRequest creates a request to invoke GetArtifact API
func CreateGetArtifactRequest() (request *GetArtifactRequest) {
	request = &GetArtifactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "GetArtifact", "", "")
	request.Method = requests.POST
	return
}

// CreateGetArtifactResponse creates a response to parse from GetArtifact response
func CreateGetArtifactResponse() (response *GetArtifactResponse) {
	response = &GetArtifactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
