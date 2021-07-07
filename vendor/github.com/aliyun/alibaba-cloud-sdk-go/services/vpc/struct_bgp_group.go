package vpc

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

// BgpGroup is a nested struct in vpc response
type BgpGroup struct {
	Name        string `json:"Name" xml:"Name"`
	Description string `json:"Description" xml:"Description"`
	BgpGroupId  string `json:"BgpGroupId" xml:"BgpGroupId"`
	PeerAsn     string `json:"PeerAsn" xml:"PeerAsn"`
	AuthKey     string `json:"AuthKey" xml:"AuthKey"`
	RouterId    string `json:"RouterId" xml:"RouterId"`
	Status      string `json:"Status" xml:"Status"`
	Keepalive   string `json:"Keepalive" xml:"Keepalive"`
	LocalAsn    string `json:"LocalAsn" xml:"LocalAsn"`
	Hold        string `json:"Hold" xml:"Hold"`
	IsFake      string `json:"IsFake" xml:"IsFake"`
	RouteLimit  string `json:"RouteLimit" xml:"RouteLimit"`
	RegionId    string `json:"RegionId" xml:"RegionId"`
	IpVersion   string `json:"IpVersion" xml:"IpVersion"`
}
