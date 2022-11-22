package rds

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

// Data is a nested struct in rds response
type Data struct {
	ConnectionString     string               `json:"ConnectionString" xml:"ConnectionString"`
	PageNumber           int                  `json:"PageNumber" xml:"PageNumber"`
	DBInstanceName       string               `json:"DBInstanceName" xml:"DBInstanceName"`
	IpVersion            string               `json:"IpVersion" xml:"IpVersion"`
	PageSize             int                  `json:"PageSize" xml:"PageSize"`
	TotalRecordCount     int                  `json:"TotalRecordCount" xml:"TotalRecordCount"`
	DBInstanceEndpointId string               `json:"DBInstanceEndpointId" xml:"DBInstanceEndpointId"`
	Nodes                []Node               `json:"Nodes" xml:"Nodes"`
	Connections          []Connection         `json:"Connections" xml:"Connections"`
	NotifyItemList       []NotifyItemListItem `json:"NotifyItemList" xml:"NotifyItemList"`
	DBInstanceEndpoints  DBInstanceEndpoints  `json:"DBInstanceEndpoints" xml:"DBInstanceEndpoints"`
}
