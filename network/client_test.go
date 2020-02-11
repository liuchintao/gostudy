package network

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		cfg config
	}
	tests := []struct {
		name string
		args args
		want client
	}{
		{"TestCase1", args{config{}}, newClient(config{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newClient(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
