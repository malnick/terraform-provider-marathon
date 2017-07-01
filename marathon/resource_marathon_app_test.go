package marathon

//func TestCreateApp(t *testing.T) {
//	mockResp, err := ioutil.ReadFile("fixtures/mock_marathon_create_response.json")
//	if err != nil {
//		t.Error(err)
//	}
//
//	m, err := New()
//	if err != nil {
//		t.Error(err)
//	}
//
//	httpmock.ActivateNonDefault(m.client)
//	defer httpmock.DeactivateAndReset()
//	httpmock.RegisterResponder("POST", "http://localhost/v2/apps",
//		httpmock.NewStringResponder(201, string(mockResp)))
//
//	m.JSON = "fixtures/mock_marathon_app.json"
//	m.URL = "http://localhost"
//
//	if err := m.create(); err != nil {
//		t.Error(err)
//	}
//}
//
//func TestUpdateApp(t *testing.T) {
//	mockResp, err := ioutil.ReadFile("fixtures/mock_marathon_create_response.json")
//	if err != nil {
//		t.Error(err)
//	}
//
//	m, err := New()
//	if err != nil {
//		t.Error(err)
//	}
//
//	httpmock.ActivateNonDefault(m.client)
//	defer httpmock.DeactivateAndReset()
//	httpmock.RegisterResponder("PUT", "http://localhost/v2/apps/mock",
//		httpmock.NewStringResponder(201, string(mockResp)))
//
//	m.JSON = "fixtures/mock_marathon_app.json"
//	m.URL = "http://localhost"
//
//	if err := m.update(); err != nil {
//		t.Error(err)
//	}
//}
