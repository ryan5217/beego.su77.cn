package goAdmin

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "users" => http://localhost:8080/go_admin/info/users
//
var Generators = map[string]table.Generator{
	"users": GetUsersTable,
}
