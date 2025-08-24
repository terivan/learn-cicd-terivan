package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input   http.Header
		want    string
		wantErr error
	}

	headers1 := make(http.Header)
	headers1.Add("Content-Type", "application/json")
	headers1.Add("Authorization", "NotKey ABCD")

	headers2 := make(http.Header)
	headers2.Add("Content-Type", "application/json")
	headers2.Add("Authorization", "ApiKey ABCD")

	headers3 := make(http.Header)
	headers3.Add("Content-Type", "application/json")

	var malformedHeader = errors.New("malformed authorization header")
	tests := []test{
		{input: headers1, want: "", wantErr: malformedHeader},
		{input: headers2, want: "ABCD", wantErr: nil},
		{input: headers3, want: "", wantErr: ErrNoAuthHeaderIncluded},
	}

	for _, tc := range tests {
		gotString, gotError := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want, gotString) || !reflect.DeepEqual(tc.wantErr, gotError) {
			t.Fatalf("expected: %v and %v, got: %v and %v", tc.want, gotString, tc.wantErr, gotError)
		}
	}
}
