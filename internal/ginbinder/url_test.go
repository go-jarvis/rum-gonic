package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestQuery(t *testing.T) {
	// https://www.v2ex.com/t/596134

	data := []string{
		`http://api.example.com/users?name=zhangsan,wangwu,zhaoliu`,
		`http://api.example.com/users?name=zhangsan&name=wangwu&name=zhaoliu`,
		`http://api.example.com/users?name[]=zhangsan&name[]=wangwu&name[]=zhaoliu`,
		`http://api.example.com/users?name[0]=zhangsan&name[1]=wangwu&name[2]=zhaoliu`,
	}

	for _, u := range data {

		ur, err := url.Parse(u)
		if err != nil {
			panic(err)
		}

		fmt.Println(ur.Query())
	}

	// map[name:[zhangsan,wangwu,zhaoliu]]
	// map[name:[zhangsan wangwu zhaoliu]]
	// map[name[]:[zhangsan wangwu zhaoliu]]
	// map[name[0]:[zhangsan] name[1]:[wangwu] name[2]:[zhaoliu]]
}
