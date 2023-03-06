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

// BatchImportVehicleDevice invokes the iot.BatchImportVehicleDevice API synchronously
func (client *Client) BatchImportVehicleDevice(request *BatchImportVehicleDeviceRequest) (response *BatchImportVehicleDeviceResponse, err error) {
	response = CreateBatchImportVehicleDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// BatchImportVehicleDeviceWithChan invokes the iot.BatchImportVehicleDevice API asynchronously
func (client *Client) BatchImportVehicleDeviceWithChan(request *BatchImportVehicleDeviceRequest) (<-chan *BatchImportVehicleDeviceResponse, <-chan error) {
	responseChan := make(chan *BatchImportVehicleDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchImportVehicleDevice(request)
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

// BatchImportVehicleDeviceWithCallback invokes the iot.BatchImportVehicleDevice API asynchronously
func (client *Client) BatchImportVehicleDeviceWithCallback(request *BatchImportVehicleDeviceRequest, callback func(response *BatchImportVehicleDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchImportVehicleDeviceResponse
		var err error
		defer close(result)
		response, err = client.BatchImportVehicleDevice(request)
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

// BatchImportVehicleDeviceRequest is the request struct for api BatchImportVehicleDevice
type BatchImportVehicleDeviceRequest struct {
	*requests.RpcRequest
	IotInstanceId string                                `position:"Query" name:"IotInstanceId"`
	ProductKey    string                                `position:"Query" name:"ProductKey"`
	DeviceList    *[]BatchImportVehicleDeviceDeviceList `position:"Query" name:"DeviceList"  type:"Repeated"`
	ApiProduct    string                                `position:"Body" name:"ApiProduct"`
	ApiRevision   string                                `position:"Body" name:"ApiRevision"`
}

// BatchImportVehicleDeviceDeviceList is a repeated param struct in BatchImportVehicleDeviceRequest
type BatchImportVehicleDeviceDeviceList struct {
	DeviceId     string `name:"DeviceId"`
	Manufacturer string `name:"Manufacturer"`
	DeviceModel  string `name:"DeviceModel"`
}

// BatchImportVehicleDeviceResponse is the response struct for api BatchImportVehicleDevice
type BatchImportVehicleDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateBatchImportVehicleDeviceRequest creates a request to invoke BatchImportVehicleDevice API
func CreateBatchImportVehicleDeviceRequest() (request *BatchImportVehicleDeviceRequest) {
	request = &BatchImportVehicleDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchImportVehicleDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchImportVehicleDeviceResponse creates a response to parse from BatchImportVehicleDevice response
func CreateBatchImportVehicleDeviceResponse() (response *BatchImportVehicleDeviceResponse) {
	response = &BatchImportVehicleDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
