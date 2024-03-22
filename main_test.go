package main

import (
	"context"
	"fmt"
	"permission/ent"
	"permission/ent/enttest"
	"testing"

	_ "github.com/lib/pq"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	groups, err := client.Group.CreateBulk(
		client.Group.Create().SetName("admin"),
		client.Group.Create().SetName("viewer"),
	).Save(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	roles, err := client.Role.CreateBulk(
		client.Role.Create().SetName("direct"),
		client.Role.Create().SetName("admin").AddGroups(groups[0]),
		client.Role.Create().SetName("viewer").AddGroups(groups[1]),
	).Save(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.User.CreateBulk(
		client.User.Create().SetName("a").AddWithin(groups[0]),
		client.User.Create().SetName("b").AddWithin(groups[1]),
		client.User.Create().SetName("c").AddDirectRoles(roles[2]),
	).Save(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return func(t *testing.T) {
		client.Group.Delete().ExecX(context.Background())
		client.Role.Delete().ExecX(context.Background())
		client.User.Delete().ExecX(context.Background())
		client.Close()
	}
}

func TestAssociatedRoles_no_duplicates(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	users, err := client.User.Query().All(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for _, user := range users {
		t.Run(user.Name, func(t *testing.T) {
			roles, err := AssociatedRoles(user)
			if err != nil {
				t.Fatal(err)
			}
			if len(roles) != 2 {
				t.Errorf("got %d roles, want 2", len(roles))
			}
		})

		// AssociatedRoles(client, user)
	}
}

func S(user *ent.User) {
	user_roles := user.QueryDirectRoles().AllX(context.Background())
	fmt.Println(user_roles)
}

func TestSomehting(t *testing.T) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}

	group := client.Group.Create().SetName("admin").SaveX(context.Background())
	_ = client.Role.Create().SetName("indirect").AddGroups(group).SaveX(context.Background())
	role := client.Role.Create().SetName("direct").SaveX(context.Background())
	user := client.User.Create().SetName("a").AddDirectRoles(role).AddWithin(group).SaveX(context.Background())

	role_user := role.QueryUsers().AllX(context.Background())
	fmt.Println(role_user)

	user_role := user.QueryDirectRoles().AllX(context.Background())
	fmt.Println(user_role)

	func() {
		user_role2 := user.QueryDirectRoles().AllX(context.Background())
		user_role3 := user.QueryDirectRoles().AllX(context.Background())
		fmt.Println(user_role2, user_role3)
	}()

	S(user)

	roles, err := AssociatedRoles(user)
	if err != nil {
		t.Fatal(err)
	}
	if len(roles) != 2 {
		t.Errorf("got %d roles, want 2", len(roles))
	}
}
