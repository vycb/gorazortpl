package main

import (
	//"fmt"
	//"time"
	"net/http"
	//."github.com/vycb/gorazortpl/tpl/bench"
	"github.com/vycb/gorazortpl/tpl/helper"
	."github.com/vycb/gorazortpl/tpl"
)


func main() {

	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)

	//fmt.Println(Home(2, &helper.User{"UserName", "User@email", "User Intro"}))
	//fmt.Println(Bench(&Page{"UserName", []string{"User Intro","red","black"}, time.Now().Year()}))
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(Home(2, &helper.User{"UserName", "User@email", "User Intro"})))
	//fmt.Println(Bench(&Page{"UserName", []string{"User Intro","red","black"}, time.Now().Year()}))

	//page := &Page{
	//	Title: "Bob",
	//	FavoriteColors: []string{"blue", "green", "mauve"},
	//	Year: time.Now().Year(),
	//}

}

