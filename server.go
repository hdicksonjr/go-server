package main

import (
  "flag"
  "html/template"
  "log"
  "net/http"
)

// what is the flag library in Go?
// Var addr = flag.String('addr', ':1718', 'http service address') //sets a default port for our 

var addr = flag.String("addr", ":1718", "http service address") // create command libe flag called 'addr'

var templ = template.Must(template.New("qr").Parse(templateStr))
// template.Must() is a wrapper. It is meant to take any function that returns a template and an error. It 'panics' if the error is non-nil.
// Parse() is in the text/template package (which html/template wraps, which is why we can use it here) 
// Parse() parses a template definition string so that it can be executed
// New() allocates a new HTML file with the given name. A name must be passed which seems kind of odd since I never reference the template by tha name
// and it seems like the convention to always wrap new inside template.Must(), so that name is useless.

func main() {
    flag.Parse() // parse command line arguments
    http.Handle("/", http.HandlerFunc(ShowPage)) //HandleFunc registers the handler function for the given pattern in the DefaultServeMux
    // DefaultServerMux is Go's fancy name for http multiplexor created by default when you instantiate the 'net/http' package.
    // http.HandlerFunc(f) straight from docs: The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. 
    // If f is a function with the appropriate signature, HandlerFunc(f) is a Handler object that calls f.
    err := http.ListenAndServe(*addr, nil)
    // where we have passed this nil, you can pass a handler. Nil is common though, it means it will immediately try to match the route against the DefaultServeMux
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func ShowPage(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>hey dudes</title>
</head>
<body>
<h1>hello</h1>
</body>
</html>
`
