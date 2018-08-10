package main

// import (
// 	"fmt"
//     "net/http"
// )

// type ReqUser struct{
// 	UserId string `json:"userId"`
// };

// func user(w http.ResponseWriter, req *http.Request) {
// 	userId := req.URL.Query().Get("userId")
// 	s := ReqUser{userId}
// 	fmt.Printf("==== %v", s)
// 	w.Write([]byte("Hello"))
// }
// func say(w http.ResponseWriter, req *http.Request) {
// 	w.Write([]byte("Hello"))
// }

// func main() {
//     http.HandleFunc("/user/", user)
//     http.HandleFunc("/login/", say)

//     http.ListenAndServe(":8001", nil)
// }


// struct to json
import (
    "fmt"
    "encoding/json"
)

type User struct {
    Name string `json:"name"`
}

func main() {
    user := &User{Name: "Frank"}
    b, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
}
