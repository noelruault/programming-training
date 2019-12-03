package stripe

import (
	"fmt"
	"testing"
)

type StripeClientMock struct {
	err error
}

func (scm *StripeClientMock) Charge(amount int, source, desc string) (*Charge, error) {
	return nil, scm.err
}

func TestApp_Run(t *testing.T) {
	type fields struct {
		sc StripeClient
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "test1",
			fields:  fields{sc: &StripeClientMock{}},
			wantErr: false,
		},
		{
			name:    "test2",
			fields:  fields{sc: &StripeClientMock{err: fmt.Errorf("mock error")}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				sc: tt.fields.sc,
			}
			if err := a.Run(); (err != nil) != tt.wantErr {
				t.Errorf("App.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}