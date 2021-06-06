package main

const (
	defaultJeagerAddr = "gitea.inyu.in:6831"
)

func main() {
	initJaeger()
	runHttp(":8087")
}
