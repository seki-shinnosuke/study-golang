package todo

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	code "github.com/seki-shinnosuke/study-golang/code"
	model "github.com/seki-shinnosuke/study-golang/model/db"
	response "github.com/seki-shinnosuke/study-golang/model/rest/response/todo"
	"github.com/seki-shinnosuke/study-golang/testhelper"
	todouc "github.com/seki-shinnosuke/study-golang/usecase/todo"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestGetTasks(t *testing.T) {

	type fields struct {
		todoUsecase todouc.TodoUsecase
	}
	type params struct {
		requestHeaders map[string]string
	}

	tests := []struct {
		name     string
		fields   fields
		params   params
		wantCode int
		wantBody map[string]any
	}{
		{
			name: "正常: タスク一覧取得",
			fields: func() fields {
				return fields{
					todoUsecase: *todouc.NewTodoUsecase(),
				}
			}(),
			params: func() params {
				return params{
					requestHeaders: nil,
				}
			}(),
			wantCode: http.StatusOK,
			wantBody: func() map[string]any {
				data := response.TodoResponse{
					Tasks: []response.Task{
						{
							TaskId:       1,
							PersonName:   "テスト",
							TaskName:     "テスト太郎",
							DeadlineDate: null.NewTime(time.Date(2022, 12, 31, 0, 0, 0, 0, time.Local), false),
							TaskStatus:   code.NO_PROCESSING.Name(),
						},
					},
				}
				b, _ := json.Marshal(data)
				body := make(map[string]any)
				json.Unmarshal(b, &body)
				return body
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// DB Setup
			testhelper.SetupTestDB(t)
			testhelper.CleanUpDB(t)
			registerTask := model.TaskManagement{
				PersonName:   "テスト",
				TaskName:     "テスト太郎",
				DeadlineDate: null.NewTime(time.Date(2022, 12, 31, 0, 0, 0, 0, time.Local), false),
				TaskStatus:   code.NO_PROCESSING.Name(),
			}
			ctx := context.Background()
			registerTask.InsertG(ctx, boil.Infer())

			// Controller Setup
			_, r, w := testhelper.NewHttp(t)
			controller := NewTodoController(&tt.fields.todoUsecase)
			r.GET("/api/v1/tasks", controller.GetTasks)
			req, _ := http.NewRequest(
				http.MethodGet,
				"/api/v1/tasks",
				nil,
			)

			// Test Start
			r.ServeHTTP(w, req)
			// Assert
			if w.Code != tt.wantCode {
				t.Errorf("gotCode = %v, wantCode = %v", w.Code, tt.wantCode)
			}
			gotBody := make(map[string]any)
			json.Unmarshal(w.Body.Bytes(), &gotBody)
			if diff := cmp.Diff(gotBody, tt.wantBody); diff != "" {
				t.Errorf("mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
