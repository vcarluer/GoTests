package main
import(
	"fmt"
	"net/http"
)

type Struct struct {
	Greeting, Punct, Who string
}

type String string

func (st Struct) ServeHTTP(
	w http.ResponseWriter, 
	r *http.Request) {
	fmt.Fprint(w, st.Greeting + st.Punct + st.Who);	
}

func (s String) ServeHTTP(
	w http.ResponseWriter, 
	r *http.Request) {
	fmt.Fprint(w, s);
}

func main() {
	
	http.Handle("/string", String("I'm a frayed knot."));
	http.Handle("/struct", &Struct{ "Greetings", ":", "Gophers!"})
	fmt.Println("Listening on port 4000")
	fmt.Println("Available routes: /string and /struct")
	http.ListenAndServe("localhost:4000", nil)

}

