package Test

import (
	controller "github.com/go-workflow/go-workflow/workflow-controller"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSaveProcdefByToken(t *testing.T) {
	bodyReader := strings.NewReader(`{"name":"请假","resource":{"name":"发起人","type":"start","nodeId":"sid-startevent"}}`)

	type t2 struct {
		name           string
		r              *http.Request
		w              *httptest.ResponseRecorder
		expectedStatus int
	}

	test := t2{
		name:           "good",
		r:              httptest.NewRequest("POST", "/api/v1/workflow/procdef/saveByToken", bodyReader),
		w:              httptest.NewRecorder(),
		expectedStatus: http.StatusOK,
	}

	test.r.Header.Set("Authorization", "ksaklaasdfasdf")
	controller.SaveProcdefByToken(test.w, test.r)

}


func TestStartByToken(t *testing.T) {
	bodyReader := strings.NewReader(`{"procName":"请假","title":"请假-张三","userId":"11025","department":"技术中心","company":"A公司","var":{"DDHolidayField-J2BWEN12__duration":"8","DDHolidayField-J2BWEN12__options":"年假"}}`)

	type t2 struct {
		name           string
		r              *http.Request
		w              *httptest.ResponseRecorder
		expectedStatus int
	}

	test := t2{
		name:           "good",
		r:              httptest.NewRequest("POST", "/api/v1/workflow/process/startByToken", bodyReader),
		w:              httptest.NewRecorder(),
		expectedStatus: http.StatusOK,
	}

	test.r.Header.Set("Authorization", "ksaklaasdfasdf")
	controller.StartProcessInstanceByToken(test.w, test.r)

}