package avatar

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

// Data is a nested struct in avatar response
type Data struct {
	GrabType   string     `json:"GrabType" xml:"GrabType"`
	Token      string     `json:"Token" xml:"Token"`
	TaskUuid   string     `json:"TaskUuid" xml:"TaskUuid"`
	TenantId   int64      `json:"TenantId" xml:"TenantId"`
	IsCancel   bool       `json:"IsCancel" xml:"IsCancel"`
	Process    string     `json:"Process" xml:"Process"`
	SessionId  string     `json:"SessionId" xml:"SessionId"`
	FailReason string     `json:"FailReason" xml:"FailReason"`
	ActionType string     `json:"ActionType" xml:"ActionType"`
	InstanceId string     `json:"InstanceId" xml:"InstanceId"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	OutputText string     `json:"OutputText" xml:"OutputText"`
	Status     string     `json:"Status" xml:"Status"`
	Type       string     `json:"Type" xml:"Type"`
	TaskResult TaskResult `json:"TaskResult" xml:"TaskResult"`
	Channel    Channel    `json:"Channel" xml:"Channel"`
}
