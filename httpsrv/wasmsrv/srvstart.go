package wasmsrv

import "net/http"

func Srvstart() {
	http.ListenAndServe(`:8083`, http.FileServer(http.Dir(`./www`)))
}
