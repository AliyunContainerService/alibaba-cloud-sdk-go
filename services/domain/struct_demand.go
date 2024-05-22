package domain

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

// Demand is a nested struct in domain response
type Demand struct {
	Status              string  `json:"Status" xml:"Status"`
	DemandPrice         float64 `json:"DemandPrice" xml:"DemandPrice"`
	BizId               string  `json:"BizId" xml:"BizId"`
	BargainSellerMobile string  `json:"BargainSellerMobile" xml:"BargainSellerMobile"`
	PublishTime         int64   `json:"PublishTime" xml:"PublishTime"`
	ProduceType         int     `json:"ProduceType" xml:"ProduceType"`
	DemandDomain        string  `json:"DemandDomain" xml:"DemandDomain"`
	Description         string  `json:"Description" xml:"Description"`
	Mobile              string  `json:"Mobile" xml:"Mobile"`
	ServicePayPrice     float64 `json:"ServicePayPrice" xml:"ServicePayPrice"`
	PayPrice            float64 `json:"PayPrice" xml:"PayPrice"`
	PayTime             int64   `json:"PayTime" xml:"PayTime"`
	BargainSellerPrice  float64 `json:"BargainSellerPrice" xml:"BargainSellerPrice"`
	OrderType           int     `json:"OrderType" xml:"OrderType"`
	PayDomain           string  `json:"PayDomain" xml:"PayDomain"`
	AuditStatus         int     `json:"AuditStatus" xml:"AuditStatus"`
	PartnerDomain       string  `json:"PartnerDomain" xml:"PartnerDomain"`
	PurchaseStatus      int     `json:"PurchaseStatus" xml:"PurchaseStatus"`
	Email               string  `json:"Email" xml:"Email"`
	SettleBasePrice     float64 `json:"SettleBasePrice" xml:"SettleBasePrice"`
}
