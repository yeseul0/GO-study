package main

import (
	"net/http"
)

func main() {
	s := "Hello, world!"

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// HTML로 웹 페이지 작성
		html := `
		<html>
		<head>
			<title>Hello</title>
			<script type="text/javascript" src="/assets/hello.js"></script>
			<link href="/assets/hello.css" rel="stylesheet" />
		</head>
		<body>
			<span class="hello">` + s + `</span>
		</body>
		</html>
		`

		res.Header().Set("Content-Type", "text/html") // HTML 헤더 설정
		res.Write([]byte(html))                       // 웹 브라우저에 응답
	})

	http.Handle( // /assets/ 경로에 접근했을 때 파일 서버를 동작시킴
		"/assets/",
		http.StripPrefix( // 파일 서버를 실행할 때 assets // 디렉터리를 지정했으므로
			"/assets/",                          //URL 경로에서 /assets/ 삭제
			http.FileServer(http.Dir("assets")), // 웹 서버에서 assets 를 루트로 보겟다.
			// -> 그럼 assets/hello.js가 아니라 hello.js로 접근해야해서 앞 prefix를 제거하는것!
		),
	)
	//http.HandleFunc, http.Handle로 요청 처리할 핸들러 설정했으므로.. 2nd 매개변수 nil
	http.ListenAndServe(":80", nil) // 80번 포트에서 웹 서버 실행
}
