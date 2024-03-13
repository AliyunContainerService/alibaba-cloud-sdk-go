package swas_open

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

// CustomImage is a nested struct in swas_open response
type CustomImage struct {
	CreationTime       string `json:"CreationTime" xml:"CreationTime"`
	DataSnapshotName   string `json:"DataSnapshotName" xml:"DataSnapshotName"`
	SystemSnapshotId   string `json:"SystemSnapshotId" xml:"SystemSnapshotId"`
	InShare            bool   `json:"InShare" xml:"InShare"`
	InstanceId         string `json:"InstanceId" xml:"InstanceId"`
	DataSnapshotId     string `json:"DataSnapshotId" xml:"DataSnapshotId"`
	RegionId           string `json:"RegionId" xml:"RegionId"`
	SystemSnapshotName string `json:"SystemSnapshotName" xml:"SystemSnapshotName"`
	Description        string `json:"Description" xml:"Description"`
	Name               string `json:"Name" xml:"Name"`
	ImageId            string `json:"ImageId" xml:"ImageId"`
	Status             string `json:"Status" xml:"Status"`
	InstanceName       string `json:"InstanceName" xml:"InstanceName"`
	ResourceGroupId    string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tags               []Tag  `json:"Tags" xml:"Tags"`
}
