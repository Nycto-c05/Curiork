## routing
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method { //method is get/post etc
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("index page"))
			return
		case "/users":
			w.Write([]byte("Users page"))

		}
	default:
		http.NotFound(w, r)
	}
}
