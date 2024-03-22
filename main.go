package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"permission/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	client := Open("postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	defer client.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users, err := client.User.Query().All(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, user := range users {
			w.Write([]byte(fmt.Sprintf("%v\n", user)))
		}

	})
	slog.Info("server started at :8080")
	http.ListenAndServe(":8080", nil)

	// client.User.Create().SetName("Tim").SetBirthday(time.Now()).SaveX(context.Background())

	// users := client.User.Query().AllX(context.Background())
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	// users[3].Update().SetName("Tom").SaveX(context.Background())

	// groups, err := client.Group.CreateBulk(
	// 	client.Group.Create().SetName("admin"),
	// 	client.Group.Create().SetName("viewer"),
	// ).Save(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// roles, err := client.Role.CreateBulk(
	// 	client.Role.Create().SetName("direct"),
	// 	client.Role.Create().SetName("admin").AddGroups(groups[0]),
	// 	client.Role.Create().SetName("viewer").AddGroups(groups[1]),
	// ).Save(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// users, err := client.User.CreateBulk(
	// 	client.User.Create().SetName("a").AddWithin(groups[0]).AddDirectRoles(roles[2]),
	// 	client.User.Create().SetName("b").AddWithin(groups[1]).AddDirectRoles(roles[2]),
	// 	client.User.Create().SetName("c").AddDirectRoles(roles[2]).AddWithin(groups[1]),
	// ).Save(context.Background())
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(users)
}

func AssociatedRoles(user *ent.User) ([]*ent.Role, error) {

	direct := user.QueryDirectRoles().AllX(context.Background())
	fmt.Println("direct", direct)
	indirect, err := user.QueryWithin().QueryRoles().All(context.Background())
	if err != nil {
		return nil, err
	}
	return removeDuplicate(append(direct, indirect...)), nil
	// return append(direct, indirect...), nil

}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
