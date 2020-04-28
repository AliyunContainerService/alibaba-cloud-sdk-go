package cloudauth

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

// CompareFaceVerify invokes the cloudauth.CompareFaceVerify API synchronously
// api document: https://help.aliyun.com/api/cloudauth/comparefaceverify.html
func (client *Client) CompareFaceVerify(request *CompareFaceVerifyRequest) (response *CompareFaceVerifyResponse, err error) {
	response = CreateCompareFaceVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// CompareFaceVerifyWithChan invokes the cloudauth.CompareFaceVerify API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/comparefaceverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CompareFaceVerifyWithChan(request *CompareFaceVerifyRequest) (<-chan *CompareFaceVerifyResponse, <-chan error) {
	responseChan := make(chan *CompareFaceVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CompareFaceVerify(request)
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

// CompareFaceVerifyWithCallback invokes the cloudauth.CompareFaceVerify API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/comparefaceverify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CompareFaceVerifyWithCallback(request *CompareFaceVerifyRequest, callback func(response *CompareFaceVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CompareFaceVerifyResponse
		var err error
		defer close(result)
		response, err = client.CompareFaceVerify(request)
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

// CompareFaceVerifyRequest is the request struct for api CompareFaceVerify
type CompareFaceVerifyRequest struct {
	*requests.RpcRequest
	TargetFaceContrastPictureUrl string           `position:"Body" name:"TargetFaceContrastPictureUrl"`
	ProductCode                  string           `position:"Body" name:"ProductCode"`
	TargetCertifyId              string           `position:"Body" name:"TargetCertifyId"`
	SourceOssObjectName          string           `position:"Body" name:"SourceOssObjectName"`
	TargetFaceContrastPicture    string           `position:"Body" name:"TargetFaceContrastPicture"`
	TargetOssBucketName          string           `position:"Body" name:"TargetOssBucketName"`
	SourceOssBucketName          string           `position:"Body" name:"SourceOssBucketName"`
	OuterOrderNo                 string           `position:"Body" name:"OuterOrderNo"`
	TargetOssObjectName          string           `position:"Body" name:"TargetOssObjectName"`
	SourceFaceContrastPicture    string           `position:"Body" name:"SourceFaceContrastPicture"`
	SceneId                      requests.Integer `position:"Body" name:"SceneId"`
	SourceFaceContrastPictureUrl string           `position:"Body" name:"SourceFaceContrastPictureUrl"`
	SourceCertifyId              string           `position:"Body" name:"SourceCertifyId"`
}

// CompareFaceVerifyResponse is the response struct for api CompareFaceVerify
type CompareFaceVerifyResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Message      string       `json:"Message" xml:"Message"`
	Code         string       `json:"Code" xml:"Code"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateCompareFaceVerifyRequest creates a request to invoke CompareFaceVerify API
func CreateCompareFaceVerifyRequest() (request *CompareFaceVerifyRequest) {
	request = &CompareFaceVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "CompareFaceVerify", "cloudauth", "openAPI")
	return
}

// CreateCompareFaceVerifyResponse creates a response to parse from CompareFaceVerify response
func CreateCompareFaceVerifyResponse() (response *CompareFaceVerifyResponse) {
	response = &CompareFaceVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
