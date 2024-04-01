package slb

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

// DescribeLoadBalancerHTTPListenerAttribute invokes the slb.DescribeLoadBalancerHTTPListenerAttribute API synchronously
func (client *Client) DescribeLoadBalancerHTTPListenerAttribute(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (response *DescribeLoadBalancerHTTPListenerAttributeResponse, err error) {
	response = CreateDescribeLoadBalancerHTTPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoadBalancerHTTPListenerAttributeWithChan invokes the slb.DescribeLoadBalancerHTTPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithChan(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (<-chan *DescribeLoadBalancerHTTPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLoadBalancerHTTPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoadBalancerHTTPListenerAttribute(request)
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

// DescribeLoadBalancerHTTPListenerAttributeWithCallback invokes the slb.DescribeLoadBalancerHTTPListenerAttribute API asynchronously
func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithCallback(request *DescribeLoadBalancerHTTPListenerAttributeRequest, callback func(response *DescribeLoadBalancerHTTPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoadBalancerHTTPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoadBalancerHTTPListenerAttribute(request)
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

// DescribeLoadBalancerHTTPListenerAttributeRequest is the request struct for api DescribeLoadBalancerHTTPListenerAttribute
type DescribeLoadBalancerHTTPListenerAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         requests.Integer `position:"Query" name:"ListenerPort"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// DescribeLoadBalancerHTTPListenerAttributeResponse is the response struct for api DescribeLoadBalancerHTTPListenerAttribute
type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	*responses.BaseResponse
	Status                 string `json:"Status" xml:"Status"`
	VServerGroupId         string `json:"VServerGroupId" xml:"VServerGroupId"`
	Cookie                 string `json:"Cookie" xml:"Cookie"`
	Gzip                   string `json:"Gzip" xml:"Gzip"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
	HealthCheckConnectPort int    `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	Bandwidth              int    `json:"Bandwidth" xml:"Bandwidth"`
	HealthCheckTimeout     int    `json:"HealthCheckTimeout" xml:"HealthCheckTimeout"`
	BackendServerPort      int    `json:"BackendServerPort" xml:"BackendServerPort"`
	CookieTimeout          int    `json:"CookieTimeout" xml:"CookieTimeout"`
	URI                    string `json:"URI" xml:"URI"`
	UnhealthyThreshold     int    `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	XForwardedForSLBID     string `json:"XForwardedFor_SLBID" xml:"XForwardedFor_SLBID"`
	SecurityStatus         string `json:"SecurityStatus" xml:"SecurityStatus"`
	HealthCheckHttpCode    string `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
	Domain                 string `json:"Domain" xml:"Domain"`
	MaxConnection          int    `json:"MaxConnection" xml:"MaxConnection"`
	XForwardedFor          string `json:"XForwardedFor" xml:"XForwardedFor"`
	ListenerPort           int    `json:"ListenerPort" xml:"ListenerPort"`
	StickySessionType      string `json:"StickySessionType" xml:"StickySessionType"`
	Scheduler              string `json:"Scheduler" xml:"Scheduler"`
	Interval               int    `json:"Interval" xml:"Interval"`
	HealthyThreshold       int    `json:"HealthyThreshold" xml:"HealthyThreshold"`
	XForwardedForProto     string `json:"XForwardedFor_proto" xml:"XForwardedFor_proto"`
	XForwardedForSLBIP     string `json:"XForwardedFor_SLBIP" xml:"XForwardedFor_SLBIP"`
	StickySession          string `json:"StickySession" xml:"StickySession"`
	HealthCheck            string `json:"HealthCheck" xml:"HealthCheck"`
}

// CreateDescribeLoadBalancerHTTPListenerAttributeRequest creates a request to invoke DescribeLoadBalancerHTTPListenerAttribute API
func CreateDescribeLoadBalancerHTTPListenerAttributeRequest() (request *DescribeLoadBalancerHTTPListenerAttributeRequest) {
	request = &DescribeLoadBalancerHTTPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2013-02-21", "DescribeLoadBalancerHTTPListenerAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLoadBalancerHTTPListenerAttributeResponse creates a response to parse from DescribeLoadBalancerHTTPListenerAttribute response
func CreateDescribeLoadBalancerHTTPListenerAttributeResponse() (response *DescribeLoadBalancerHTTPListenerAttributeResponse) {
	response = &DescribeLoadBalancerHTTPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
