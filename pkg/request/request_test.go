package request_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dewzzjr/ais/pkg/request"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	t.Run("ShouldReturnObject", func(t *testing.T) {
		t.Parallel()
		t.Run("WhenValidJsonBody", func(t *testing.T) {
			t.Parallel()
			t.Run("WhenContentTypeJson", func(t *testing.T) {
				t.Parallel()
				var obj struct {
					Test string `json:"test"`
				}
				err := request.Read(
					testRequestJson(t),
					&obj,
				)
				assert.NoError(t, err)
				assert.Equal(t, obj.Test, "hello")
			})
			t.Run("WhenNoHeader_WhenOptsJson", func(t *testing.T) {
				t.Parallel()
				var obj struct {
					Test string `json:"test"`
				}
				err := request.Read(
					testRequestJsonNoHeader(t),
					&obj,
					request.JSON,
				)
				assert.NoError(t, err)
				assert.Equal(t, obj.Test, "hello2")
			})

			t.Run("WhenOptsJsonAndQuery", func(t *testing.T) {
				t.Parallel()
				var obj struct {
					Test      string `json:"test" schema:"-"`
					TestQuery string `json:"-" schema:"testQuery"`
				}
				err := request.Read(
					testRequestJsonAndQuery(t),
					&obj,
					request.JSON,
					request.Query,
				)
				assert.NoError(t, err)
				assert.Equal(t, obj.Test, "hello1")
				assert.Equal(t, obj.TestQuery, "hello2")
			})
		})
		t.Run("WhenValidQuery_WhenOptsQuery", func(t *testing.T) {
			t.Parallel()
			var obj struct {
				Test string `schema:"test"`
			}
			err := request.Read(
				testRequestQuery(t),
				&obj,
				request.Query,
			)
			assert.NoError(t, err)
			assert.Equal(t, obj.Test, "hello3")
		})
		t.Run("WhenValidPostForm_WhenContentTypeFormUrlEncoded", func(t *testing.T) {
			t.Parallel()
			var obj struct {
				Test string `schema:"test"`
			}
			err := request.Read(
				testRequestPostForm(t),
				&obj,
			)
			assert.NoError(t, err)
			assert.Equal(t, obj.Test, "hello4")
		})
	})
}

func testRequestJson(t *testing.T) *http.Request {
	t.Helper()
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(`{"test":"hello"}`))
	r.Header.Set("content-type", "application/json")
	return r
}

func testRequestJsonNoHeader(t *testing.T) *http.Request {
	t.Helper()
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(`{"test":"hello2"}`))
	return r
}

func testRequestQuery(t *testing.T) *http.Request {
	t.Helper()
	r := httptest.NewRequest(http.MethodGet, "/test?test=hello3", nil)
	return r
}

func testRequestPostForm(t *testing.T) *http.Request {
	t.Helper()
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBufferString(`test=hello4`))
	r.Header.Set("content-type", "application/x-www-form-urlencoded")
	return r
}

func testRequestJsonAndQuery(t *testing.T) *http.Request {
	t.Helper()
	r := httptest.NewRequest(http.MethodPost, "/test?testQuery=hello2", bytes.NewBufferString(`{"test":"hello1"}`))
	r.Header.Set("content-type", "application/json")
	return r
}
