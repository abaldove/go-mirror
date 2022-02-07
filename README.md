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
[ ] - Better asset download control, right now I do have a way to control already visited pages, but the same is not true for assets download. <br/>
[ ] - Automated tests
[ ] - Maybe, I'd like to try a version with no colly framework, and for that I probably would use a BTree to store site's hierarchy. With that, I could traverse the tree in pre and post order to accelerate the crawler process, also I could store these struct to use as some kind of progress control, in addition to that, I also would have an Map to control assets download and prevent possible duplications. In the boostrap process I would load those extra sctructs to memory and persist them by the end or while the Gracefull Shutdown is happening.
