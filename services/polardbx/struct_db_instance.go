package polardbx

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

// DBInstance is a nested struct in polardbx response
type DBInstance struct {
	DBNodeClass             string       `json:"DBNodeClass" xml:"DBNodeClass"`
	DBType                  string       `json:"DBType" xml:"DBType"`
	Engine                  string       `json:"Engine" xml:"Engine"`
	TertiaryZone            string       `json:"TertiaryZone" xml:"TertiaryZone"`
	RightsSeparationEnabled bool         `json:"RightsSeparationEnabled" xml:"RightsSeparationEnabled"`
	CnNodeCount             int          `json:"CnNodeCount" xml:"CnNodeCount"`
	DBInstanceType          string       `json:"DBInstanceType" xml:"DBInstanceType"`
	MinorVersion            string       `json:"MinorVersion" xml:"MinorVersion"`
	CommodityCode           string       `json:"CommodityCode" xml:"CommodityCode"`
	Id                      string       `json:"Id" xml:"Id"`
	Type                    string       `json:"Type" xml:"Type"`
	VPCId                   string       `json:"VPCId" xml:"VPCId"`
	PayType                 string       `json:"PayType" xml:"PayType"`
	MaintainStartTime       string       `json:"MaintainStartTime" xml:"MaintainStartTime"`
	ZoneId                  string       `json:"ZoneId" xml:"ZoneId"`
	LatestMinorVersion      string       `json:"LatestMinorVersion" xml:"LatestMinorVersion"`
	DnNodeCount             int          `json:"DnNodeCount" xml:"DnNodeCount"`
	CanNotCreateColumnar    bool         `json:"CanNotCreateColumnar" xml:"CanNotCreateColumnar"`
	Status                  string       `json:"Status" xml:"Status"`
	ColumnarInstanceName    string       `json:"ColumnarInstanceName" xml:"ColumnarInstanceName"`
	LockMode                string       `json:"LockMode" xml:"LockMode"`
	CnNodeClassCode         string       `json:"CnNodeClassCode" xml:"CnNodeClassCode"`
	CreateTime              string       `json:"CreateTime" xml:"CreateTime"`
	TopologyType            string       `json:"TopologyType" xml:"TopologyType"`
	DifferentDNSpec         bool         `json:"DifferentDNSpec" xml:"DifferentDNSpec"`
	DBNodeCount             int          `json:"DBNodeCount" xml:"DBNodeCount"`
	RightsSeparationStatus  string       `json:"RightsSeparationStatus" xml:"RightsSeparationStatus"`
	RegionId                string       `json:"RegionId" xml:"RegionId"`
	ResourceGroupId         string       `json:"ResourceGroupId" xml:"ResourceGroupId"`
	KindCode                int          `json:"KindCode" xml:"KindCode"`
	MaintainEndTime         string       `json:"MaintainEndTime" xml:"MaintainEndTime"`
	Network                 string       `json:"Network" xml:"Network"`
	PrimaryZone             string       `json:"PrimaryZone" xml:"PrimaryZone"`
	Expired                 string       `json:"Expired" xml:"Expired"`
	Description             string       `json:"Description" xml:"Description"`
	DBVersion               string       `json:"DBVersion" xml:"DBVersion"`
	SecondaryZone           string       `json:"SecondaryZone" xml:"SecondaryZone"`
	Port                    string       `json:"Port" xml:"Port"`
	ExpireDate              string       `json:"ExpireDate" xml:"ExpireDate"`
	VSwitchId               string       `json:"VSwitchId" xml:"VSwitchId"`
	StorageUsed             int64        `json:"StorageUsed" xml:"StorageUsed"`
	Series                  string       `json:"Series" xml:"Series"`
	ConnectionString        string       `json:"ConnectionString" xml:"ConnectionString"`
	DnNodeClassCode         string       `json:"DnNodeClassCode" xml:"DnNodeClassCode"`
	LTSVersions             []string     `json:"LTSVersions" xml:"LTSVersions"`
	ReadDBInstances         []string     `json:"ReadDBInstances" xml:"ReadDBInstances"`
	ColumnarReadDBInstances []string     `json:"ColumnarReadDBInstances" xml:"ColumnarReadDBInstances"`
	ConnAddrs               []ConnAddr   `json:"ConnAddrs" xml:"ConnAddrs"`
	DBNodes                 []DBNode     `json:"DBNodes" xml:"DBNodes"`
	TagSet                  []TagSetItem `json:"TagSet" xml:"TagSet"`
}
