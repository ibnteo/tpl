# tpl

Template engine similar to sprintf

%e = html.EscapeString(s)
%q = url.QueryEscape(s)
%s = string
%n = text<br>
%% = %

Usage:

	package main

	import (
		"net/http"

		"github.com/ibnteo/tpl"
	)

	func main() {
		http.HandleFunc("/", handlerMain)
		http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
		http.ListenAndServe(":80", nil)
	}

	func handlerMain(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		var info string
		if r.URL.Query().Get("error") != "" {
			info = tpl.Format(`<div class="error">%e</div>`, r.URL.Query().Get("error"))
		}

		tpl.Write(w, `<b style="font-weight:%ept">%e</b>%s`, 10, "<abc>", info)

		tpl.Print("%s\n", info)
	}
