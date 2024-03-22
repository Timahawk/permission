// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "birthday", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// GroupMembersColumns holds the columns for the "group_members" table.
	GroupMembersColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// GroupMembersTable holds the schema information for the "group_members" table.
	GroupMembersTable = &schema.Table{
		Name:       "group_members",
		Columns:    GroupMembersColumns,
		PrimaryKey: []*schema.Column{GroupMembersColumns[0], GroupMembersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_members_group_id",
				Columns:    []*schema.Column{GroupMembersColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_members_user_id",
				Columns:    []*schema.Column{GroupMembersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// GroupRolesColumns holds the columns for the "group_roles" table.
	GroupRolesColumns = []*schema.Column{
		{Name: "group_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// GroupRolesTable holds the schema information for the "group_roles" table.
	GroupRolesTable = &schema.Table{
		Name:       "group_roles",
		Columns:    GroupRolesColumns,
		PrimaryKey: []*schema.Column{GroupRolesColumns[0], GroupRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_roles_group_id",
				Columns:    []*schema.Column{GroupRolesColumns[0]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "group_roles_role_id",
				Columns:    []*schema.Column{GroupRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoleUsersColumns holds the columns for the "role_users" table.
	RoleUsersColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// RoleUsersTable holds the schema information for the "role_users" table.
	RoleUsersTable = &schema.Table{
		Name:       "role_users",
		Columns:    RoleUsersColumns,
		PrimaryKey: []*schema.Column{RoleUsersColumns[0], RoleUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_users_role_id",
				Columns:    []*schema.Column{RoleUsersColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_users_user_id",
				Columns:    []*schema.Column{RoleUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GroupsTable,
		RolesTable,
		UsersTable,
		GroupMembersTable,
		GroupRolesTable,
		RoleUsersTable,
	}
)

func init() {
	GroupMembersTable.ForeignKeys[0].RefTable = GroupsTable
	GroupMembersTable.ForeignKeys[1].RefTable = UsersTable
	GroupRolesTable.ForeignKeys[0].RefTable = GroupsTable
	GroupRolesTable.ForeignKeys[1].RefTable = RolesTable
	RoleUsersTable.ForeignKeys[0].RefTable = RolesTable
	RoleUsersTable.ForeignKeys[1].RefTable = UsersTable
}