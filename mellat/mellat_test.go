package mellat

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		terminalId   int64
		userName     string
		userPassword string
	}
	tests := []struct {
		name string
		args args
		want Mellat
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.terminalId, tt.args.userName, tt.args.userPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMellat_PayRequest(t *testing.T) {
	type fields struct {
		terminalId   int64
		userName     string
		userPassword string
	}
	type args struct {
		amount      int64
		orderId     int64
		callBackUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr string
	}{
		{
			name: "21",
			fields: fields{
				terminalId:   123,
				userName:     "test",
				userPassword: "test",
			},
			args: args{
				amount:      1000,
				orderId:     1,
				callBackUrl: "",
			},
			wantErr: "21",
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := Mellat{
				terminalId:   tt.fields.terminalId,
				userName:     tt.fields.userName,
				userPassword: tt.fields.userPassword,
			}
			got, err := service.PayRequest(tt.args.amount, tt.args.orderId, tt.args.callBackUrl)
			if err != nil {
				if err, ok := err.(Error); ok {
					if err.ErrCode() != tt.wantErr {
						t.Errorf("Mellat.PayRequest() error = %v, wantErr %v", err.ErrCode(), tt.wantErr)
						return
					}
				} else {
					t.Errorf("Mellat.PayRequest() error = %v,", err)
					return
				}
			}
			if got != tt.want {
				t.Errorf("Mellat.PayRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMellat_VerifyRequest(t *testing.T) {
	type fields struct {
		terminalId   int64
		userName     string
		userPassword string
	}
	type args struct {
		orderId         int64
		saleOrderId     int64
		saleReferenceId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := Mellat{
				terminalId:   tt.fields.terminalId,
				userName:     tt.fields.userName,
				userPassword: tt.fields.userPassword,
			}
			got, err := service.VerifyRequest(tt.args.orderId, tt.args.saleOrderId, tt.args.saleReferenceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mellat.VerifyRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Mellat.VerifyRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMellat_SettleRequest(t *testing.T) {
	type fields struct {
		terminalId   int64
		userName     string
		userPassword string
	}
	type args struct {
		orderId         int64
		saleOrderId     int64
		saleReferenceId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := Mellat{
				terminalId:   tt.fields.terminalId,
				userName:     tt.fields.userName,
				userPassword: tt.fields.userPassword,
			}
			got, err := service.SettleRequest(tt.args.orderId, tt.args.saleOrderId, tt.args.saleReferenceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mellat.SettleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Mellat.SettleRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
