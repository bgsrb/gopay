package mellat

import "testing"

func TestMellatError_ErrCode(t *testing.T) {
	type fields struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &MellatError{
				code: tt.fields.code,
			}
			if got := e.ErrCode(); got != tt.want {
				t.Errorf("MellatError.ErrCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMellatError_Error(t *testing.T) {
	type fields struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := MellatError{
				code: tt.fields.code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("MellatError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
