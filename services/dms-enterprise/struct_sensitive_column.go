package dms_enterprise

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

// SensitiveColumn is a nested struct in dms_enterprise response
type SensitiveColumn struct {
	SchemaName                  string                      `json:"SchemaName" xml:"SchemaName"`
	ColumnCount                 int64                       `json:"ColumnCount" xml:"ColumnCount"`
	TableName                   string                      `json:"TableName" xml:"TableName"`
	SampleData                  string                      `json:"SampleData" xml:"SampleData"`
	UserSensitivityLevel        string                      `json:"UserSensitivityLevel" xml:"UserSensitivityLevel"`
	CategoryName                string                      `json:"CategoryName" xml:"CategoryName"`
	ColumnName                  string                      `json:"ColumnName" xml:"ColumnName"`
	SecurityLevel               string                      `json:"SecurityLevel" xml:"SecurityLevel"`
	InstanceId                  int                         `json:"InstanceId" xml:"InstanceId"`
	FunctionType                string                      `json:"FunctionType" xml:"FunctionType"`
	IsPlain                     bool                        `json:"IsPlain" xml:"IsPlain"`
	DefaultDesensitizationRule  DefaultDesensitizationRule  `json:"DefaultDesensitizationRule" xml:"DefaultDesensitizationRule"`
	SemiDesensitizationRuleList SemiDesensitizationRuleList `json:"SemiDesensitizationRuleList" xml:"SemiDesensitizationRuleList"`
}
