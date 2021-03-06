package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDBInstances invokes the gpdb.DescribeDBInstances API synchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstances.html
func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (response *DescribeDBInstancesResponse, err error) {
	response = CreateDescribeDBInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstancesWithChan invokes the gpdb.DescribeDBInstances API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstancesWithChan(request *DescribeDBInstancesRequest) (<-chan *DescribeDBInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstances(request)
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

// DescribeDBInstancesWithCallback invokes the gpdb.DescribeDBInstances API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describedbinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstancesWithCallback(request *DescribeDBInstancesRequest, callback func(response *DescribeDBInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstances(request)
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

// DescribeDBInstancesRequest is the request struct for api DescribeDBInstances
type DescribeDBInstancesRequest struct {
	*requests.RpcRequest
	DBInstanceIds         string                    `position:"Query" name:"DBInstanceIds"`
	PageSize              requests.Integer          `position:"Query" name:"PageSize"`
	DBInstanceDescription string                    `position:"Query" name:"DBInstanceDescription"`
	Tag                   *[]DescribeDBInstancesTag `position:"Query" name:"Tag"  type:"Repeated"`
	OwnerId               requests.Integer          `position:"Query" name:"OwnerId"`
	InstanceNetworkType   string                    `position:"Query" name:"InstanceNetworkType"`
	PageNumber            requests.Integer          `position:"Query" name:"PageNumber"`
}

// DescribeDBInstancesTag is a repeated param struct in DescribeDBInstancesRequest
type DescribeDBInstancesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeDBInstancesResponse is the response struct for api DescribeDBInstances
type DescribeDBInstancesResponse struct {
	*responses.BaseResponse
	RequestId        string                     `json:"RequestId" xml:"RequestId"`
	PageNumber       int                        `json:"PageNumber" xml:"PageNumber"`
	TotalRecordCount int                        `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int                        `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeDBInstances `json:"Items" xml:"Items"`
}

// CreateDescribeDBInstancesRequest creates a request to invoke DescribeDBInstances API
func CreateDescribeDBInstancesRequest() (request *DescribeDBInstancesRequest) {
	request = &DescribeDBInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDBInstances", "gpdb", "openAPI")
	return
}

// CreateDescribeDBInstancesResponse creates a response to parse from DescribeDBInstances response
func CreateDescribeDBInstancesResponse() (response *DescribeDBInstancesResponse) {
	response = &DescribeDBInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
