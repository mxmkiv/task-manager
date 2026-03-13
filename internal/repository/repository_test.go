package repository

import (
	"errors"
	"slices"
	"testing"
)

type TestsObject struct {
	testName     string
	data         map[string]string
	reqId        int
	correctQuery string
	correctArgs  []any
	wantError    error
}

func TestUpdateUserData(t *testing.T) {

	tests := []TestsObject{
		{
			testName: "all fields",
			data: map[string]string{
				"login":         "NewLogin",
				"password_hash": "NewPassword",
				"role":          "newRole",
			},
			reqId:        1,
			correctQuery: "UPDATE users SET login=$1, password_hash=$2, role=$3 WHERE id=$4",
			correctArgs:  []any{"NewLogin", "NewPassword", "newRole", 1},
			wantError:    nil,
		},
		{
			testName:     "no fields",
			data:         map[string]string{},
			reqId:        2,
			correctQuery: "",
			correctArgs:  nil,
			wantError:    errors.New("no fields to update"),
		},
		{
			testName: "only role",
			data: map[string]string{
				"role": "newRole",
			},
			reqId:        3,
			correctQuery: "UPDATE users SET role=$1 WHERE id=$2",
			correctArgs:  []any{"newRole", 3},
			wantError:    nil,
		},
		{
			testName: "only password",
			data: map[string]string{
				"password_hash": "newPassword",
			},
			reqId:        3,
			correctQuery: "UPDATE users SET password_hash=$1 WHERE id=$2",
			correctArgs:  []any{"newPassword", 3},
			wantError:    nil,
		},
		{
			testName: "only login",
			data: map[string]string{
				"login": "newLogin",
			},
			reqId:        3,
			correctQuery: "UPDATE users SET login=$1 WHERE id=$2",
			correctArgs:  []any{"newLogin", 3},
			wantError:    nil,
		},
	}

	for _, test := range tests {

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			query, args, err := UpdateQueryForm(test.data, test.reqId)

			if query != test.correctQuery {
				t.Errorf("incorrect query want %s, get %s", test.correctQuery, query)
			}
			if !slices.Equal(args, test.correctArgs) {
				t.Errorf("incorrect args want %s, get %s", test.correctArgs, args)
			}
			if test.wantError != nil {
				if err.Error() != test.wantError.Error() {
					t.Errorf("want error %s, get %s", err.Error(), test.wantError.Error())
				}
			}
		})
	}

}
