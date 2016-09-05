// Automatically generated by MockGen. DO NOT EDIT!
// Source: vendor/github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface/interface.go

package mocks

import (
	request "github.com/aws/aws-sdk-go/aws/request"
	cloudwatchlogs "github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

// Mock of CloudWatchLogsAPI interface
type MockCloudWatchLogsAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockCloudWatchLogsAPIRecorder
}

// Recorder for MockCloudWatchLogsAPI (not exported)
type _MockCloudWatchLogsAPIRecorder struct {
	mock *MockCloudWatchLogsAPI
}

func NewMockCloudWatchLogsAPI(ctrl *gomock.Controller) *MockCloudWatchLogsAPI {
	mock := &MockCloudWatchLogsAPI{ctrl: ctrl}
	mock.recorder = &_MockCloudWatchLogsAPIRecorder{mock}
	return mock
}

func (_m *MockCloudWatchLogsAPI) EXPECT() *_MockCloudWatchLogsAPIRecorder {
	return _m.recorder
}

func (_m *MockCloudWatchLogsAPI) CancelExportTaskRequest(_param0 *cloudwatchlogs.CancelExportTaskInput) (*request.Request, *cloudwatchlogs.CancelExportTaskOutput) {
	ret := _m.ctrl.Call(_m, "CancelExportTaskRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.CancelExportTaskOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) CancelExportTaskRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CancelExportTaskRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) CancelExportTask(_param0 *cloudwatchlogs.CancelExportTaskInput) (*cloudwatchlogs.CancelExportTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "CancelExportTask", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.CancelExportTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) CancelExportTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CancelExportTask", arg0)
}

func (_m *MockCloudWatchLogsAPI) CreateExportTaskRequest(_param0 *cloudwatchlogs.CreateExportTaskInput) (*request.Request, *cloudwatchlogs.CreateExportTaskOutput) {
	ret := _m.ctrl.Call(_m, "CreateExportTaskRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.CreateExportTaskOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) CreateExportTaskRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateExportTaskRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) CreateExportTask(_param0 *cloudwatchlogs.CreateExportTaskInput) (*cloudwatchlogs.CreateExportTaskOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateExportTask", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.CreateExportTaskOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) CreateExportTask(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateExportTask", arg0)
}

func (_m *MockCloudWatchLogsAPI) CreateLogGroupRequest(_param0 *cloudwatchlogs.CreateLogGroupInput) (*request.Request, *cloudwatchlogs.CreateLogGroupOutput) {
	ret := _m.ctrl.Call(_m, "CreateLogGroupRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.CreateLogGroupOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) CreateLogGroupRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateLogGroupRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) CreateLogGroup(_param0 *cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateLogGroup", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.CreateLogGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) CreateLogGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateLogGroup", arg0)
}

func (_m *MockCloudWatchLogsAPI) CreateLogStreamRequest(_param0 *cloudwatchlogs.CreateLogStreamInput) (*request.Request, *cloudwatchlogs.CreateLogStreamOutput) {
	ret := _m.ctrl.Call(_m, "CreateLogStreamRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.CreateLogStreamOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) CreateLogStreamRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateLogStreamRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) CreateLogStream(_param0 *cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error) {
	ret := _m.ctrl.Call(_m, "CreateLogStream", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.CreateLogStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) CreateLogStream(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateLogStream", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteDestinationRequest(_param0 *cloudwatchlogs.DeleteDestinationInput) (*request.Request, *cloudwatchlogs.DeleteDestinationOutput) {
	ret := _m.ctrl.Call(_m, "DeleteDestinationRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DeleteDestinationOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteDestinationRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteDestinationRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteDestination(_param0 *cloudwatchlogs.DeleteDestinationInput) (*cloudwatchlogs.DeleteDestinationOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteDestination", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DeleteDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteDestination(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteDestination", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteLogGroupRequest(_param0 *cloudwatchlogs.DeleteLogGroupInput) (*request.Request, *cloudwatchlogs.DeleteLogGroupOutput) {
	ret := _m.ctrl.Call(_m, "DeleteLogGroupRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DeleteLogGroupOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteLogGroupRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteLogGroupRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteLogGroup(_param0 *cloudwatchlogs.DeleteLogGroupInput) (*cloudwatchlogs.DeleteLogGroupOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteLogGroup", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DeleteLogGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteLogGroup(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteLogGroup", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteLogStreamRequest(_param0 *cloudwatchlogs.DeleteLogStreamInput) (*request.Request, *cloudwatchlogs.DeleteLogStreamOutput) {
	ret := _m.ctrl.Call(_m, "DeleteLogStreamRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DeleteLogStreamOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteLogStreamRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteLogStreamRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteLogStream(_param0 *cloudwatchlogs.DeleteLogStreamInput) (*cloudwatchlogs.DeleteLogStreamOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteLogStream", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DeleteLogStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteLogStream(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteLogStream", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteMetricFilterRequest(_param0 *cloudwatchlogs.DeleteMetricFilterInput) (*request.Request, *cloudwatchlogs.DeleteMetricFilterOutput) {
	ret := _m.ctrl.Call(_m, "DeleteMetricFilterRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DeleteMetricFilterOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteMetricFilterRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteMetricFilterRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteMetricFilter(_param0 *cloudwatchlogs.DeleteMetricFilterInput) (*cloudwatchlogs.DeleteMetricFilterOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteMetricFilter", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DeleteMetricFilterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteMetricFilter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteMetricFilter", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteRetentionPolicyRequest(_param0 *cloudwatchlogs.DeleteRetentionPolicyInput) (*request.Request, *cloudwatchlogs.DeleteRetentionPolicyOutput) {
	ret := _m.ctrl.Call(_m, "DeleteRetentionPolicyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DeleteRetentionPolicyOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteRetentionPolicyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRetentionPolicyRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteRetentionPolicy(_param0 *cloudwatchlogs.DeleteRetentionPolicyInput) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteRetentionPolicy", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DeleteRetentionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteRetentionPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRetentionPolicy", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteSubscriptionFilterRequest(_param0 *cloudwatchlogs.DeleteSubscriptionFilterInput) (*request.Request, *cloudwatchlogs.DeleteSubscriptionFilterOutput) {
	ret := _m.ctrl.Call(_m, "DeleteSubscriptionFilterRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DeleteSubscriptionFilterOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteSubscriptionFilterRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSubscriptionFilterRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DeleteSubscriptionFilter(_param0 *cloudwatchlogs.DeleteSubscriptionFilterInput) (*cloudwatchlogs.DeleteSubscriptionFilterOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteSubscriptionFilter", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DeleteSubscriptionFilterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DeleteSubscriptionFilter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSubscriptionFilter", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeDestinationsRequest(_param0 *cloudwatchlogs.DescribeDestinationsInput) (*request.Request, *cloudwatchlogs.DescribeDestinationsOutput) {
	ret := _m.ctrl.Call(_m, "DescribeDestinationsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DescribeDestinationsOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeDestinationsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeDestinationsRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeDestinations(_param0 *cloudwatchlogs.DescribeDestinationsInput) (*cloudwatchlogs.DescribeDestinationsOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeDestinations", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeDestinations(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeDestinations", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeDestinationsPages(_param0 *cloudwatchlogs.DescribeDestinationsInput, _param1 func(*cloudwatchlogs.DescribeDestinationsOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "DescribeDestinationsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeDestinationsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeDestinationsPages", arg0, arg1)
}

func (_m *MockCloudWatchLogsAPI) DescribeExportTasksRequest(_param0 *cloudwatchlogs.DescribeExportTasksInput) (*request.Request, *cloudwatchlogs.DescribeExportTasksOutput) {
	ret := _m.ctrl.Call(_m, "DescribeExportTasksRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DescribeExportTasksOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeExportTasksRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeExportTasksRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeExportTasks(_param0 *cloudwatchlogs.DescribeExportTasksInput) (*cloudwatchlogs.DescribeExportTasksOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeExportTasks", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeExportTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeExportTasks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeExportTasks", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeLogGroupsRequest(_param0 *cloudwatchlogs.DescribeLogGroupsInput) (*request.Request, *cloudwatchlogs.DescribeLogGroupsOutput) {
	ret := _m.ctrl.Call(_m, "DescribeLogGroupsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DescribeLogGroupsOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeLogGroupsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeLogGroupsRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeLogGroups(_param0 *cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeLogGroups", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeLogGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeLogGroups(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeLogGroups", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeLogGroupsPages(_param0 *cloudwatchlogs.DescribeLogGroupsInput, _param1 func(*cloudwatchlogs.DescribeLogGroupsOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "DescribeLogGroupsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeLogGroupsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeLogGroupsPages", arg0, arg1)
}

func (_m *MockCloudWatchLogsAPI) DescribeLogStreamsRequest(_param0 *cloudwatchlogs.DescribeLogStreamsInput) (*request.Request, *cloudwatchlogs.DescribeLogStreamsOutput) {
	ret := _m.ctrl.Call(_m, "DescribeLogStreamsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DescribeLogStreamsOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeLogStreamsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeLogStreamsRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeLogStreams(_param0 *cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeLogStreams", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeLogStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeLogStreams(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeLogStreams", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeLogStreamsPages(_param0 *cloudwatchlogs.DescribeLogStreamsInput, _param1 func(*cloudwatchlogs.DescribeLogStreamsOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "DescribeLogStreamsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeLogStreamsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeLogStreamsPages", arg0, arg1)
}

func (_m *MockCloudWatchLogsAPI) DescribeMetricFiltersRequest(_param0 *cloudwatchlogs.DescribeMetricFiltersInput) (*request.Request, *cloudwatchlogs.DescribeMetricFiltersOutput) {
	ret := _m.ctrl.Call(_m, "DescribeMetricFiltersRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DescribeMetricFiltersOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeMetricFiltersRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeMetricFiltersRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeMetricFilters(_param0 *cloudwatchlogs.DescribeMetricFiltersInput) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeMetricFilters", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeMetricFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeMetricFilters(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeMetricFilters", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeMetricFiltersPages(_param0 *cloudwatchlogs.DescribeMetricFiltersInput, _param1 func(*cloudwatchlogs.DescribeMetricFiltersOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "DescribeMetricFiltersPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeMetricFiltersPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeMetricFiltersPages", arg0, arg1)
}

func (_m *MockCloudWatchLogsAPI) DescribeSubscriptionFiltersRequest(_param0 *cloudwatchlogs.DescribeSubscriptionFiltersInput) (*request.Request, *cloudwatchlogs.DescribeSubscriptionFiltersOutput) {
	ret := _m.ctrl.Call(_m, "DescribeSubscriptionFiltersRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.DescribeSubscriptionFiltersOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeSubscriptionFiltersRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeSubscriptionFiltersRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeSubscriptionFilters(_param0 *cloudwatchlogs.DescribeSubscriptionFiltersInput) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {
	ret := _m.ctrl.Call(_m, "DescribeSubscriptionFilters", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeSubscriptionFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeSubscriptionFilters(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeSubscriptionFilters", arg0)
}

func (_m *MockCloudWatchLogsAPI) DescribeSubscriptionFiltersPages(_param0 *cloudwatchlogs.DescribeSubscriptionFiltersInput, _param1 func(*cloudwatchlogs.DescribeSubscriptionFiltersOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "DescribeSubscriptionFiltersPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCloudWatchLogsAPIRecorder) DescribeSubscriptionFiltersPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DescribeSubscriptionFiltersPages", arg0, arg1)
}

func (_m *MockCloudWatchLogsAPI) FilterLogEventsRequest(_param0 *cloudwatchlogs.FilterLogEventsInput) (*request.Request, *cloudwatchlogs.FilterLogEventsOutput) {
	ret := _m.ctrl.Call(_m, "FilterLogEventsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.FilterLogEventsOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) FilterLogEventsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilterLogEventsRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) FilterLogEvents(_param0 *cloudwatchlogs.FilterLogEventsInput) (*cloudwatchlogs.FilterLogEventsOutput, error) {
	ret := _m.ctrl.Call(_m, "FilterLogEvents", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.FilterLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) FilterLogEvents(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilterLogEvents", arg0)
}

func (_m *MockCloudWatchLogsAPI) FilterLogEventsPages(_param0 *cloudwatchlogs.FilterLogEventsInput, _param1 func(*cloudwatchlogs.FilterLogEventsOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "FilterLogEventsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCloudWatchLogsAPIRecorder) FilterLogEventsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FilterLogEventsPages", arg0, arg1)
}

func (_m *MockCloudWatchLogsAPI) GetLogEventsRequest(_param0 *cloudwatchlogs.GetLogEventsInput) (*request.Request, *cloudwatchlogs.GetLogEventsOutput) {
	ret := _m.ctrl.Call(_m, "GetLogEventsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.GetLogEventsOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) GetLogEventsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLogEventsRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) GetLogEvents(_param0 *cloudwatchlogs.GetLogEventsInput) (*cloudwatchlogs.GetLogEventsOutput, error) {
	ret := _m.ctrl.Call(_m, "GetLogEvents", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.GetLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) GetLogEvents(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLogEvents", arg0)
}

func (_m *MockCloudWatchLogsAPI) GetLogEventsPages(_param0 *cloudwatchlogs.GetLogEventsInput, _param1 func(*cloudwatchlogs.GetLogEventsOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "GetLogEventsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCloudWatchLogsAPIRecorder) GetLogEventsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLogEventsPages", arg0, arg1)
}

func (_m *MockCloudWatchLogsAPI) PutDestinationRequest(_param0 *cloudwatchlogs.PutDestinationInput) (*request.Request, *cloudwatchlogs.PutDestinationOutput) {
	ret := _m.ctrl.Call(_m, "PutDestinationRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.PutDestinationOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutDestinationRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutDestinationRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutDestination(_param0 *cloudwatchlogs.PutDestinationInput) (*cloudwatchlogs.PutDestinationOutput, error) {
	ret := _m.ctrl.Call(_m, "PutDestination", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.PutDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutDestination(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutDestination", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutDestinationPolicyRequest(_param0 *cloudwatchlogs.PutDestinationPolicyInput) (*request.Request, *cloudwatchlogs.PutDestinationPolicyOutput) {
	ret := _m.ctrl.Call(_m, "PutDestinationPolicyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.PutDestinationPolicyOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutDestinationPolicyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutDestinationPolicyRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutDestinationPolicy(_param0 *cloudwatchlogs.PutDestinationPolicyInput) (*cloudwatchlogs.PutDestinationPolicyOutput, error) {
	ret := _m.ctrl.Call(_m, "PutDestinationPolicy", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.PutDestinationPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutDestinationPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutDestinationPolicy", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutLogEventsRequest(_param0 *cloudwatchlogs.PutLogEventsInput) (*request.Request, *cloudwatchlogs.PutLogEventsOutput) {
	ret := _m.ctrl.Call(_m, "PutLogEventsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.PutLogEventsOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutLogEventsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutLogEventsRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutLogEvents(_param0 *cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error) {
	ret := _m.ctrl.Call(_m, "PutLogEvents", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.PutLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutLogEvents(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutLogEvents", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutMetricFilterRequest(_param0 *cloudwatchlogs.PutMetricFilterInput) (*request.Request, *cloudwatchlogs.PutMetricFilterOutput) {
	ret := _m.ctrl.Call(_m, "PutMetricFilterRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.PutMetricFilterOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutMetricFilterRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutMetricFilterRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutMetricFilter(_param0 *cloudwatchlogs.PutMetricFilterInput) (*cloudwatchlogs.PutMetricFilterOutput, error) {
	ret := _m.ctrl.Call(_m, "PutMetricFilter", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.PutMetricFilterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutMetricFilter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutMetricFilter", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutRetentionPolicyRequest(_param0 *cloudwatchlogs.PutRetentionPolicyInput) (*request.Request, *cloudwatchlogs.PutRetentionPolicyOutput) {
	ret := _m.ctrl.Call(_m, "PutRetentionPolicyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.PutRetentionPolicyOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutRetentionPolicyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutRetentionPolicyRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutRetentionPolicy(_param0 *cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, error) {
	ret := _m.ctrl.Call(_m, "PutRetentionPolicy", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.PutRetentionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutRetentionPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutRetentionPolicy", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutSubscriptionFilterRequest(_param0 *cloudwatchlogs.PutSubscriptionFilterInput) (*request.Request, *cloudwatchlogs.PutSubscriptionFilterOutput) {
	ret := _m.ctrl.Call(_m, "PutSubscriptionFilterRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.PutSubscriptionFilterOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutSubscriptionFilterRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutSubscriptionFilterRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) PutSubscriptionFilter(_param0 *cloudwatchlogs.PutSubscriptionFilterInput) (*cloudwatchlogs.PutSubscriptionFilterOutput, error) {
	ret := _m.ctrl.Call(_m, "PutSubscriptionFilter", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.PutSubscriptionFilterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) PutSubscriptionFilter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PutSubscriptionFilter", arg0)
}

func (_m *MockCloudWatchLogsAPI) TestMetricFilterRequest(_param0 *cloudwatchlogs.TestMetricFilterInput) (*request.Request, *cloudwatchlogs.TestMetricFilterOutput) {
	ret := _m.ctrl.Call(_m, "TestMetricFilterRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*cloudwatchlogs.TestMetricFilterOutput)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) TestMetricFilterRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TestMetricFilterRequest", arg0)
}

func (_m *MockCloudWatchLogsAPI) TestMetricFilter(_param0 *cloudwatchlogs.TestMetricFilterInput) (*cloudwatchlogs.TestMetricFilterOutput, error) {
	ret := _m.ctrl.Call(_m, "TestMetricFilter", _param0)
	ret0, _ := ret[0].(*cloudwatchlogs.TestMetricFilterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCloudWatchLogsAPIRecorder) TestMetricFilter(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TestMetricFilter", arg0)
}
