package config

import (
	"reflect"
	"testing"
)

func Test_loadConfigJSON(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name  string
		args  args
		wantC *Config
	}{
		{
			name: "test1",
			args: args{path: "config.json"},
			wantC: &Config{
				Service: service{
					Port: "8080",
					Root: "/api",
				},
				Mysql: mysql{
					Host:         "127.0.0.1",
					Port:         3306,
					DatabaseName: "dbName",
					User:         "root",
					Password:     "1234",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := loadConfigJSON(tt.args.path); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("loadConfigJSON() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
