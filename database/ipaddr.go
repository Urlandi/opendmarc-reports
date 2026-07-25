package database

/*
Copyright 2018 Nicolas JUHEL

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

const table_ipaddr = "ipaddr"

type IpAddr struct {
	Generic
}

func NewIpAddr(Name string) *IpAddr {
	obj := &IpAddr{
		Generic{
			table: table_ipaddr,
			fctField: func() FieldList {
				return FieldList{
					"id":   "INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT",
					"name": "TEXT NOT NULL DEFAULT '' UNIQUE",
					"date": "datetime NOT NULL DEFAULT CURRENT_TIMESTAMP",
				}
			},
		},
	}

	if Name != "" {
		obj.Name = Name
	}

	return obj
}

func GetIpAddr(Id int) (*IpAddr, error) {
	obj := NewIpAddr("")

	var err error

	if Id != 0 {
		obj.Id = Id
		err = obj.Load()
	}

	return obj, err
}
