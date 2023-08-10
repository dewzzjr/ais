package api_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dewzzjr/ais/internal/config"
	"github.com/dewzzjr/ais/internal/delivery/api"
	"github.com/dewzzjr/ais/internal/model"
	"github.com/dewzzjr/ais/internal/usecase/mocks"
	"github.com/dewzzjr/ais/pkg/errs"
	"github.com/dewzzjr/ais/pkg/pointer"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*httptest.ResponseRecorder, *mocks.MockArticle) {
	t.Helper()
	u := mocks.NewMockArticle(gomock.NewController(t))
	w := httptest.NewRecorder()
	return w, u
}
func TestCreateArticles(t *testing.T) {
	t.Run("ShouldReturnCreated_WhenSuccessCreate", func(t *testing.T) {
		// setup
		t.Parallel()
		w, m := setup(t)
		r := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(`{"author":"author test","title":"title test","body":"body test"}`))
		r.Header.Set("content-type", "application/json")

		// mock
		m.EXPECT().Insert(gomock.Any(), model.Article{Author: "author test", Title: "title test", Body: "body test"}).Return(nil)

		// execute
		api.New(m, config.API{}).CreateArticles(w, r)
		got := w.Result()

		// assert
		assert.Equal(t, got.StatusCode, http.StatusCreated)
	})
	t.Run("ShouldReturnBadRequest_WhenNotValidJson", func(t *testing.T) {
		// setup
		t.Parallel()
		w, m := setup(t)
		r := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(`invalid json`))
		r.Header.Set("content-type", "application/json")

		// execute
		api.New(m, config.API{}).CreateArticles(w, r)
		got := w.Result()

		// assert
		assert.Equal(t, got.StatusCode, http.StatusBadRequest)
	})
	t.Run("ShouldReturnInternalServer_WhenFailedCreate", func(t *testing.T) {
		// setup
		t.Parallel()
		w, m := setup(t)
		r := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(`{"author":"author test","title":"title test","body":"body test"}`))
		r.Header.Set("content-type", "application/json")

		// mock
		m.EXPECT().Insert(gomock.Any(), model.Article{Author: "author test", Title: "title test", Body: "body test"}).Return(errors.New("failed"))

		// execute
		api.New(m, config.API{}).CreateArticles(w, r)
		got := w.Result()

		// assert
		assert.Equal(t, got.StatusCode, http.StatusInternalServerError)
	})
	t.Run("ShouldReturnUnauthorized_WhenReturnWrapErr", func(t *testing.T) {
		// setup
		t.Parallel()
		w, m := setup(t)
		r := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader(`{"author":"author test","title":"title test","body":"body test"}`))
		r.Header.Set("content-type", "application/json")

		// mock
		m.EXPECT().Insert(gomock.Any(), model.Article{Author: "author test", Title: "title test", Body: "body test"}).Return(errs.Wrap(http.StatusUnauthorized, errors.New("unauthorized")))

		// execute
		api.New(m, config.API{}).CreateArticles(w, r)
		got := w.Result()

		// assert
		assert.Equal(t, got.StatusCode, http.StatusUnauthorized)
	})
}
func TestFetchArticles(t *testing.T) {
	t.Run("ShouldReturnOK_WhenSuccessFetch", func(t *testing.T) {
		t.Parallel()
		t.Run("WhenQueryParamEmpty", func(t *testing.T) {
			// setup
			t.Parallel()
			w, m := setup(t)
			r := httptest.NewRequest(http.MethodGet, "/articles", nil)
			r.Header.Set("content-type", "application/json")

			// mock
			m.EXPECT().Fetch(gomock.Any(), model.Filter{}).Return([]model.Article{}, nil)

			// execute
			api.New(m, config.API{}).FetchArticles(w, r)
			got := w.Result()

			// assert
			assert.Equal(t, got.StatusCode, http.StatusOK)
		})
		t.Run("WhenQueryParamExist", func(t *testing.T) {
			// setup
			t.Parallel()
			w, m := setup(t)
			r := httptest.NewRequest(http.MethodGet, "/articles?query=test&author=me", nil)
			r.Header.Set("content-type", "application/json")

			// mock
			m.EXPECT().Fetch(gomock.Any(), model.Filter{Query: pointer.New("test"), Author: pointer.New("me")}).Return([]model.Article{}, nil)

			// execute
			api.New(m, config.API{}).FetchArticles(w, r)
			got := w.Result()

			// assert
			assert.Equal(t, got.StatusCode, http.StatusOK)
		})
	})
	t.Run("ShouldReturnInternalServer_WhenFailedFetch", func(t *testing.T) {
		// setup
		t.Parallel()
		w, m := setup(t)
		r := httptest.NewRequest(http.MethodGet, "/articles", nil)
		r.Header.Set("content-type", "application/json")

		// mock
		m.EXPECT().Fetch(gomock.Any(), model.Filter{}).Return(nil, errors.New("failed"))

		// execute
		api.New(m, config.API{}).FetchArticles(w, r)
		got := w.Result()

		// assert
		assert.Equal(t, got.StatusCode, http.StatusInternalServerError)
	})
	t.Run("ShouldReturnUnauthorized_WhenReturnWrapErr", func(t *testing.T) {
		// setup
		t.Parallel()
		w, m := setup(t)
		r := httptest.NewRequest(http.MethodGet, "/articles", nil)
		r.Header.Set("content-type", "application/json")

		// mock
		m.EXPECT().Fetch(gomock.Any(), model.Filter{}).Return(nil, errs.Wrap(http.StatusUnauthorized, errors.New("unauthorized")))

		// execute
		api.New(m, config.API{}).FetchArticles(w, r)
		got := w.Result()

		// assert
		assert.Equal(t, got.StatusCode, http.StatusUnauthorized)
	})
}
