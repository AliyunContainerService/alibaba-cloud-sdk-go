package csas

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

// Application is a nested struct in csas response
type Application struct {
	Name            string      `json:"Name" xml:"Name"`
	ApplicationId   string      `json:"ApplicationId" xml:"ApplicationId"`
	Protocol        string      `json:"Protocol" xml:"Protocol"`
	ApplicationName string      `json:"ApplicationName" xml:"ApplicationName"`
	CreateTime      string      `json:"CreateTime" xml:"CreateTime"`
	Status          string      `json:"Status" xml:"Status"`
	Description     string      `json:"Description" xml:"Description"`
	Addresses       []string    `json:"Addresses" xml:"Addresses"`
	PolicyIds       []string    `json:"PolicyIds" xml:"PolicyIds"`
	TagIds          []string    `json:"TagIds" xml:"TagIds"`
	Tags            []Tag       `json:"Tags" xml:"Tags"`
	PortRanges      []PortRange `json:"PortRanges" xml:"PortRanges"`
	Policies        []Policy    `json:"Policies" xml:"Policies"`
}
