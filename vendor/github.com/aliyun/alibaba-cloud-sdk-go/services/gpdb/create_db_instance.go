package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateDBInstance invokes the gpdb.CreateDBInstance API synchronously
// api document: https://help.aliyun.com/api/gpdb/createdbinstance.html
func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
	response = CreateCreateDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBInstanceWithChan invokes the gpdb.CreateDBInstance API asynchronously
// api document: https://help.aliyun.com/api/gpdb/createdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBInstanceWithChan(request *CreateDBInstanceRequest) (<-chan *CreateDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBInstance(request)
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

// CreateDBInstanceWithCallback invokes the gpdb.CreateDBInstance API asynchronously
// api document: https://help.aliyun.com/api/gpdb/createdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDBInstanceWithCallback(request *CreateDBInstanceRequest, callback func(response *CreateDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateDBInstance(request)
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

// CreateDBInstanceRequest is the request struct for api CreateDBInstance
type CreateDBInstanceRequest struct {
	*requests.RpcRequest
	DBInstanceGroupCount  string           `position:"Query" name:"DBInstanceGroupCount"`
	Period                string           `position:"Query" name:"Period"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	EngineVersion         string           `position:"Query" name:"EngineVersion"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	UsedTime              string           `position:"Query" name:"UsedTime"`
	DBInstanceClass       string           `position:"Query" name:"DBInstanceClass"`
	SecurityIPList        string           `position:"Query" name:"SecurityIPList"`
	VSwitchId             string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string           `position:"Query" name:"PrivateIpAddress"`
	Engine                string           `position:"Query" name:"Engine"`
	VPCId                 string           `position:"Query" name:"VPCId"`
	ZoneId                string           `position:"Query" name:"ZoneId"`
	DBInstanceDescription string           `position:"Query" name:"DBInstanceDescription"`
	PayType               string           `position:"Query" name:"PayType"`
	InstanceNetworkType   string           `position:"Query" name:"InstanceNetworkType"`
}

// CreateDBInstanceResponse is the response struct for api CreateDBInstance
type CreateDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DBInstanceId     string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId          string `json:"OrderId" xml:"OrderId"`
	ConnectionString string `json:"ConnectionString" xml:"ConnectionString"`
	Port             string `json:"Port" xml:"Port"`
}

// CreateCreateDBInstanceRequest creates a request to invoke CreateDBInstance API
func CreateCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
	request = &CreateDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "CreateDBInstance", "gpdb", "openAPI")
	return
}

// CreateCreateDBInstanceResponse creates a response to parse from CreateDBInstance response
func CreateCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
	response = &CreateDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
