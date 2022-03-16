package main
// wooo i'm a programmer 
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Handler struct{}
// some weird magic that reads JSON i guess idk i got it from SO
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path
	
// Sometimes things like this fall through the review process and end up in
// production. Pretend that's why this servername is hardcoded like it is 
// and try to work around it in your containers :)
	url := "http://server:8000/item"
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
// best variable names are the best
	xkcd := "https://xkcd.com/" + string(responseData)

	if uri == "/" {
		http.Redirect(w, r, xkcd, http.StatusSeeOther)
	}
	fmt.Fprintf(w, uri)
	return
}

func main() {
	handler := new(Handler)
// I'll just keep this port for teting purpouses and change later
	http.ListenAndServe(":8090", handler)
}
