package voicenavigator

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

// Conversation is a nested struct in voicenavigator response
type Conversation struct {
	EffectiveAnswerCount int    `json:"EffectiveAnswerCount" xml:"EffectiveAnswerCount"`
	SkillGroup           string `json:"SkillGroup" xml:"SkillGroup"`
	TransferredToAgent   bool   `json:"TransferredToAgent" xml:"TransferredToAgent"`
	EndReason            int    `json:"EndReason" xml:"EndReason"`
	EndTime              int64  `json:"EndTime" xml:"EndTime"`
	StartTime            int64  `json:"StartTime" xml:"StartTime"`
	CallingNumber        string `json:"CallingNumber" xml:"CallingNumber"`
	BeginTime            int64  `json:"BeginTime" xml:"BeginTime"`
	SkillGroupId         string `json:"SkillGroupId" xml:"SkillGroupId"`
	UserUtteranceCount   int    `json:"UserUtteranceCount" xml:"UserUtteranceCount"`
	HasToAgent           bool   `json:"HasToAgent" xml:"HasToAgent"`
	ConversationId       string `json:"ConversationId" xml:"ConversationId"`
	Rounds               int    `json:"Rounds" xml:"Rounds"`
}
