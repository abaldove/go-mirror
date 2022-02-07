# Go-Mirror

A simple attempt to implement a web crawler using Golang that behaves like `wget --mirror`. Built on top of [go-colly](https://github.com/gocolly/colly), a `Lightning Fast and Elegant Scraping Framework for Gophers`, according with their own github repo comments.  

## How to Run

This is a [Golang](https://golang.org) application, it's assumed you have Go installed. If you don't, please read [official installation instructions](https://golang.org/doc/install), be aware that `brew` and you linux distro package manage probably provides a go package. Also, be sure you have [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) set in your environment.

Then, just execute the command below

```sh
go run main.go mirror --url={provide the url to be crawled}
go run main.go mirror --url=http://go-colly.org/articles/
```


## TO-DO

[ ] - Gracefull Shutdown <br/>
[ ] - Better asset download control,including a resume control option, since it's only working for already visited urls
