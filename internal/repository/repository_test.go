package repository

import (
	"errors"
	"slices"
	"task-manager/internal/model"
	"testing"
)

type TestsObject struct {
	testName     string
	data         *[]model.UpdateFields
	reqId        int
	correctQuery string
	correctArgs  []any
	wantError    error
}

func TestUpdateUserData(t *testing.T) {

	tests := []TestsObject{
		{
			testName: "all fields",
			data: &[]model.UpdateFields{
				{
					FieldName: "login",
					Data:      "newLogin",
				},
				{
					FieldName: "password_hash",
					Data:      "newPassword",
				},
				{
					FieldName: "role",
					Data:      "newRole",
				},
			},
			reqId:        1,
			correctQuery: "UPDATE users SET login=$1, password_hash=$2, role=$3, updated_at = NOW() WHERE id=$4 RETURNING id, login, role, updated_at, created_at",
			correctArgs:  []any{"newLogin", "newPassword", "newRole", 1},
			wantError:    nil,
		},
		{
			testName:     "no fields",
			data:         &[]model.UpdateFields{},
			reqId:        2,
			correctQuery: "",
			correctArgs:  nil,
			wantError:    errors.New("no fields to update"),
		},
		{
			testName: "only role",
			data: &[]model.UpdateFields{
				{
					FieldName: "role",
					Data:      "newRole",
				},
			},
			reqId:        3,
			correctQuery: "UPDATE users SET role=$1, updated_at = NOW() WHERE id=$2 RETURNING id, login, role, updated_at, created_at",
			correctArgs:  []any{"newRole", 3},
			wantError:    nil,
		},
		{
			testName: "only password",
			data: &[]model.UpdateFields{
				{
					FieldName: "password_hash",
					Data:      "newPassword",
				},
			},
			reqId:        3,
			correctQuery: "UPDATE users SET password_hash=$1, updated_at = NOW() WHERE id=$2 RETURNING id, login, role, updated_at, created_at",
			correctArgs:  []any{"newPassword", 3},
			wantError:    nil,
		},
		{
			testName: "only login",
			data: &[]model.UpdateFields{
				{
					FieldName: "login",
					Data:      "newLogin",
				},
			},
			reqId:        3,
			correctQuery: "UPDATE users SET login=$1, updated_at = NOW() WHERE id=$2 RETURNING id, login, role, updated_at, created_at",
			correctArgs:  []any{"newLogin", 3},
			wantError:    nil,
		},
		{
			testName: "login and password",
			data: &[]model.UpdateFields{
				{
					FieldName: "login",
					Data:      "newLogin",
				},
				{
					FieldName: "password_hash",
					Data:      "newPassword",
				},
			},
			reqId:        3,
			correctQuery: "UPDATE users SET login=$1, password_hash=$2, updated_at = NOW() WHERE id=$3 RETURNING id, login, role, updated_at, created_at",
			correctArgs:  []any{"newLogin", "newPassword", 3},
			wantError:    nil,
		},
	}

	for _, test := range tests {

		testCp := test

		t.Run(testCp.testName, func(t *testing.T) {
			t.Parallel()

			query, args, err := UpdateQueryForm(testCp.data, testCp.reqId)

			if query != testCp.correctQuery {
				t.Errorf("incorrect query want %s, get %s", testCp.correctQuery, query)
			}
			if !slices.Equal(args, testCp.correctArgs) {
				t.Errorf("incorrect args want %s, get %s", testCp.correctArgs, args)
			}
			if testCp.wantError != nil {
				if err != nil && testCp.wantError.Error() != err.Error() {
					t.Errorf("unexpected error: want %s, get %s", testCp.wantError.Error(), err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("no error expected: get %s", err.Error())
				}
			}
		})
	}

}
