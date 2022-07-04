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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyNatGatewayAttribute invokes the vpc.ModifyNatGatewayAttribute API synchronously
func (client *Client) ModifyNatGatewayAttribute(request *ModifyNatGatewayAttributeRequest) (response *ModifyNatGatewayAttributeResponse, err error) {
	response = CreateModifyNatGatewayAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyNatGatewayAttributeWithChan invokes the vpc.ModifyNatGatewayAttribute API asynchronously
func (client *Client) ModifyNatGatewayAttributeWithChan(request *ModifyNatGatewayAttributeRequest) (<-chan *ModifyNatGatewayAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyNatGatewayAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyNatGatewayAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyNatGatewayAttributeWithCallback invokes the vpc.ModifyNatGatewayAttribute API asynchronously
func (client *Client) ModifyNatGatewayAttributeWithCallback(request *ModifyNatGatewayAttributeRequest, callback func(response *ModifyNatGatewayAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyNatGatewayAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyNatGatewayAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyNatGatewayAttributeRequest is the request struct for api ModifyNatGatewayAttribute
type ModifyNatGatewayAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	NatGatewayId         string           `position:"Query" name:"NatGatewayId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyNatGatewayAttributeResponse is the response struct for api ModifyNatGatewayAttribute
type ModifyNatGatewayAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyNatGatewayAttributeRequest creates a request to invoke ModifyNatGatewayAttribute API
func CreateModifyNatGatewayAttributeRequest() (request *ModifyNatGatewayAttributeRequest) {
	request = &ModifyNatGatewayAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNatGatewayAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyNatGatewayAttributeResponse creates a response to parse from ModifyNatGatewayAttribute response
func CreateModifyNatGatewayAttributeResponse() (response *ModifyNatGatewayAttributeResponse) {
	response = &ModifyNatGatewayAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
