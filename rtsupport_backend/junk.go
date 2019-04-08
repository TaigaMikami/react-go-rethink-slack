package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
)

type User struct {
	Id string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}
func main() {
	session, err := r.Connect(r.ConnectOpts {
		Address: "localhost:28015",
		Database: "rtsupport",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//user := User {
	//	Name: "Taiga Mikami",
	//}

	//response, err := r.Table("user").Insert(user).RunWrite(session)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//response, _ := r.Table("user").Get("dbd4ac87-fb1a-4feb-bca0-cd365ca2bc76").Update(user).RunWrite(session)
	//response, _ := r.Table("user").Get("dbd4ac87-fb1a-4feb-bca0-cd365ca2bc76").Delete().RunWrite(session)
	//fmt.Printf("%#v\n", response)


}