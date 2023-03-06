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

// ReBindLicenseDevice invokes the iot.ReBindLicenseDevice API synchronously
func (client *Client) ReBindLicenseDevice(request *ReBindLicenseDeviceRequest) (response *ReBindLicenseDeviceResponse, err error) {
	response = CreateReBindLicenseDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// ReBindLicenseDeviceWithChan invokes the iot.ReBindLicenseDevice API asynchronously
func (client *Client) ReBindLicenseDeviceWithChan(request *ReBindLicenseDeviceRequest) (<-chan *ReBindLicenseDeviceResponse, <-chan error) {
	responseChan := make(chan *ReBindLicenseDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReBindLicenseDevice(request)
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

// ReBindLicenseDeviceWithCallback invokes the iot.ReBindLicenseDevice API asynchronously
func (client *Client) ReBindLicenseDeviceWithCallback(request *ReBindLicenseDeviceRequest, callback func(response *ReBindLicenseDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReBindLicenseDeviceResponse
		var err error
		defer close(result)
		response, err = client.ReBindLicenseDevice(request)
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

// ReBindLicenseDeviceRequest is the request struct for api ReBindLicenseDevice
type ReBindLicenseDeviceRequest struct {
	*requests.RpcRequest
	DeviceNameList *[]string `position:"Body" name:"DeviceNameList"  type:"Repeated"`
	IotInstanceId  string    `position:"Query" name:"IotInstanceId"`
	ProductKey     string    `position:"Query" name:"ProductKey"`
	ApiProduct     string    `position:"Body" name:"ApiProduct"`
	ApiRevision    string    `position:"Body" name:"ApiRevision"`
	LicenseCode    string    `position:"Query" name:"LicenseCode"`
}

// ReBindLicenseDeviceResponse is the response struct for api ReBindLicenseDevice
type ReBindLicenseDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	Success      bool                      `json:"Success" xml:"Success"`
	Code         string                    `json:"Code" xml:"Code"`
	ErrorMessage string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInReBindLicenseDevice `json:"Data" xml:"Data"`
}

// CreateReBindLicenseDeviceRequest creates a request to invoke ReBindLicenseDevice API
func CreateReBindLicenseDeviceRequest() (request *ReBindLicenseDeviceRequest) {
	request = &ReBindLicenseDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ReBindLicenseDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateReBindLicenseDeviceResponse creates a response to parse from ReBindLicenseDevice response
func CreateReBindLicenseDeviceResponse() (response *ReBindLicenseDeviceResponse) {
	response = &ReBindLicenseDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
