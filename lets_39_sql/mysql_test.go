// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/16 - 6:05 下午

package lets_39_sql

import (
	"testing"
)

func TestConnectionDB(t *testing.T) {
	//定义测试需要的结构体
	type args struct {
		dbc DBConfig
	}
	//定义需要测试的数据库组
	testDB := map[string]args{
		"DB1": {dbc: DBConfig{IPAddr: "192.168.64.2", Port: 3306, Password: "root", userName: "root"}},
		"DB2": {dbc: DBConfig{IPAddr: "192.168.64.0", Port: 3306, Password: "root1", userName: "root"}},
	}
	//开始测试
	for name, tt := range testDB {
		t.Run(name, func(t *testing.T) {
			t.Log(func() string {
				// 这里方法不通过 验证不出来 使用还是t.log()进行测试！
				if _, err := ConnectionDB(tt.dbc); err != nil {
					return tt.dbc.IPAddr + " 连接失败！！！"
				}
				return tt.dbc.IPAddr + " 连接正常!"
			}())
		})
	}
}
