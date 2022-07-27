package usps

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_response_isError(t *testing.T) {
	type fields struct {
		number  string
		details []string
	}
	type args struct {
		xmlBody []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &response{
				number:  tt.fields.number,
				details: tt.fields.details,
			}
			tt.wantErr(t, r.isError(tt.args.xmlBody), fmt.Sprintf("isError(%v)", tt.args.xmlBody))
		})
	}
}
