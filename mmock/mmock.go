package mmock

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type Mmock struct {
	mmockUrl string
}

func NewMmock(mmockUrl string) *Mmock {
	return &Mmock{mmockUrl: mmockUrl}
}

// resetMmock resets the mmock storage of invoked endpoints.
func (m *Mmock) ResetMmock(t *testing.T) {
	_, err := http.Get(m.mmockUrl + "/api/request/reset")
	assert.NoError(t, err)
}

// checkHttpMockWasInvoked queries the HTTP API of mmock that returns what mocked endpoints have been called.
// It will fail the test if the mock does not contain at least one request matching the expectedPath.
func (m *Mmock) CheckHttpMockWasInvoked(t *testing.T, expectedPath string) {
	matchedReqs := m.queryMatchedRequests(t)

	found := m.searchForRequestedPath(matchedReqs, expectedPath)
	if !found {
		t.Errorf("did not find expected call")
	}
}

// checkHttpMockWasNotInvoked fails the test if the mock contains any requests matching unexpectedPath
func (m *Mmock) CheckHttpMockWasNotInvoked(t *testing.T, unexpectedPath string) {
	matchedReqs := m.queryMatchedRequests(t)

	found := m.searchForRequestedPath(matchedReqs, unexpectedPath)
	if found {
		t.Errorf("found call to %v we did not expect to find", unexpectedPath)
	}
}

func (m *Mmock) searchForRequestedPath(reqs []interface{}, path string) bool {
	found := false
	// Fugly, find the existence of the expected call in the response from mmock
	for _, a := range reqs {
		m1 := a.(map[string]interface{})
		for k, v := range m1 {
			// find the request property
			if k == "request" {
				actualPath := v.(map[string]interface{})["path"].(string)
				if actualPath == path {
					found = true
					break
				}
			}
		}
	}
	return found
}

func (m *Mmock) queryMatchedRequests(t *testing.T) []interface{} {
	resp, err := http.Get(m.mmockUrl + "/api/request/matched")
	assert.NoError(t, err)

	if resp.StatusCode != 200 {
		t.Fatalf("unable to check matched requests in mock, message: %v", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	respMap := make([]interface{}, 0)
	err = json.Unmarshal(bytes, &respMap)
	assert.NoError(t, err)

	return respMap
}
