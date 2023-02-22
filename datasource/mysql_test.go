package datasource

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_mysql_MySQLConnectionInfo(t *testing.T) {
	type args struct {
		ip     string
		port   int
		id     string
		pw     string
		dbName string
	}
	tests := []struct {
		name string
		m    mysql
		args args
		want mysql
	}{
		{
			name: "test1",
			m:    mysqlInstance,
			args: args{
				ip:     "127.0.0.1",
				port:   3306,
				id:     "root",
				pw:     "pwd",
				dbName: "testdb",
			},
			want: mysql{
				driverName:     "mysql",
				dataSourceName: "root:pwd@tcp(127.0.0.1:3306)/testdb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.MySQLConnectionInfo(tt.args.ip, tt.args.port, tt.args.id, tt.args.pw, tt.args.dbName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysql.MySQLConnectionInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
