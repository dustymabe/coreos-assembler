package ecs

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

// DedicatedHostType is a nested struct in ecs response
type DedicatedHostType struct {
	Cores                         int                                                       `json:"Cores" xml:"Cores"`
	LocalStorageCategory          string                                                    `json:"LocalStorageCategory" xml:"LocalStorageCategory"`
	GPUSpec                       string                                                    `json:"GPUSpec" xml:"GPUSpec"`
	TotalVcpus                    int                                                       `json:"TotalVcpus" xml:"TotalVcpus"`
	CpuOverCommitRatioRange       string                                                    `json:"CpuOverCommitRatioRange" xml:"CpuOverCommitRatioRange"`
	PhysicalGpus                  int                                                       `json:"PhysicalGpus" xml:"PhysicalGpus"`
	MemorySize                    float64                                                   `json:"MemorySize" xml:"MemorySize"`
	SupportCpuOverCommitRatio     bool                                                      `json:"SupportCpuOverCommitRatio" xml:"SupportCpuOverCommitRatio"`
	LocalStorageCapacity          int64                                                     `json:"LocalStorageCapacity" xml:"LocalStorageCapacity"`
	DedicatedHostType             string                                                    `json:"DedicatedHostType" xml:"DedicatedHostType"`
	LocalStorageAmount            int                                                       `json:"LocalStorageAmount" xml:"LocalStorageAmount"`
	TotalVgpus                    int                                                       `json:"TotalVgpus" xml:"TotalVgpus"`
	Sockets                       int                                                       `json:"Sockets" xml:"Sockets"`
	SupportedInstanceTypeFamilies SupportedInstanceTypeFamiliesInDescribeDedicatedHostTypes `json:"SupportedInstanceTypeFamilies" xml:"SupportedInstanceTypeFamilies"`
	SupportedInstanceTypesList    SupportedInstanceTypesListInDescribeDedicatedHostTypes    `json:"SupportedInstanceTypesList" xml:"SupportedInstanceTypesList"`
}
