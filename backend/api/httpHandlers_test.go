package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nimitsarup/rep/handlers"
	hMock "github.com/nimitsarup/rep/handlers/mock"
	"github.com/stretchr/testify/assert"
)

func TestAPI_GetPacksForItems(t *testing.T) {
	type fields struct {
		Handlers handlers.HandlersInterface
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedStatus int
	}{
		{
			name: "WhenNoQueryParam_ThenError",
			args: args{r: httptest.NewRequest(http.MethodGet, "/packs", nil),
				w: httptest.NewRecorder()},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "WhenBadQueryParam_ThenError",
			args: args{r: httptest.NewRequest(http.MethodGet, "/packs?items=xx", nil),
				w: httptest.NewRecorder()},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Success",
			args: args{r: httptest.NewRequest(http.MethodGet, "/packs?items=123", nil),
				w: httptest.NewRecorder()},
			expectedStatus: http.StatusOK,
			fields: fields{Handlers: &hMock.HandlersInterfaceMock{
				GetPacksForItemsFunc: func(items int) map[int]int { return map[int]int{} },
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{
				Handlers: tt.fields.Handlers,
			}
			a.GetPacksForItems(tt.args.w, tt.args.r)
			assert.Equal(t, tt.expectedStatus, tt.args.w.Code)
		})
	}
}
