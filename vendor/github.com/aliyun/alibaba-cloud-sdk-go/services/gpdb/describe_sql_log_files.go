package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSQLLogFiles invokes the gpdb.DescribeSQLLogFiles API synchronously
// api document: https://help.aliyun.com/api/gpdb/describesqllogfiles.html
func (client *Client) DescribeSQLLogFiles(request *DescribeSQLLogFilesRequest) (response *DescribeSQLLogFilesResponse, err error) {
	response = CreateDescribeSQLLogFilesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLLogFilesWithChan invokes the gpdb.DescribeSQLLogFiles API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describesqllogfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLLogFilesWithChan(request *DescribeSQLLogFilesRequest) (<-chan *DescribeSQLLogFilesResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLLogFilesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLLogFiles(request)
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

// DescribeSQLLogFilesWithCallback invokes the gpdb.DescribeSQLLogFiles API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describesqllogfiles.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLLogFilesWithCallback(request *DescribeSQLLogFilesRequest, callback func(response *DescribeSQLLogFilesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLLogFilesResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLLogFiles(request)
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

// DescribeSQLLogFilesRequest is the request struct for api DescribeSQLLogFiles
type DescribeSQLLogFilesRequest struct {
	*requests.RpcRequest
	FileName     string           `position:"Query" name:"FileName"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId string           `position:"Query" name:"DBInstanceId"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeSQLLogFilesResponse is the response struct for api DescribeSQLLogFiles
type DescribeSQLLogFilesResponse struct {
	*responses.BaseResponse
	RequestId        string                     `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int                        `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                        `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                        `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeSQLLogFiles `json:"Items" xml:"Items"`
}

// CreateDescribeSQLLogFilesRequest creates a request to invoke DescribeSQLLogFiles API
func CreateDescribeSQLLogFilesRequest() (request *DescribeSQLLogFilesRequest) {
	request = &DescribeSQLLogFilesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeSQLLogFiles", "gpdb", "openAPI")
	return
}

// CreateDescribeSQLLogFilesResponse creates a response to parse from DescribeSQLLogFiles response
func CreateDescribeSQLLogFilesResponse() (response *DescribeSQLLogFilesResponse) {
	response = &DescribeSQLLogFilesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
