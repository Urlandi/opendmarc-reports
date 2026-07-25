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

const table_domains = "domains"

type Domain struct {
	Generic
}

func NewDomain(Name string) *Domain {
	obj := &Domain{
		Generic{
			table: table_domains,
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

func GetDomain(Id int) (*Domain, error) {
	obj := NewDomain("")

	var err error

	if Id != 0 {
		obj.Id = Id
		err = obj.Load()
	}

	return obj, err
}
