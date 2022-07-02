# Crawler_TRT04
Project for crawlling information on [TRT04](https://pesquisatextual.trt4.jus.br/pesquisas/acordaos?0).
CSV file is generate with collected data.
 
## Dependencies
- [Selenium](https://github.com/tebeka/selenium#readme)
- [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/)
- [Selenium-server-standalone](https://selenium-release.storage.googleapis.com/index.html?path=3.5/)
- [html package](https://pkg.go.dev/golang.org/x/net/html)
- [htmlquery](https://github.com/antchfx/htmlquery)

## Run using brew

```brew install chrome driver ``` (not needed if you alredy have chrome driver)

```java -jar selenium-server-standalone.jar```

```go run main.go```

- Sometimes chromedriver need a clearnce in security options on MacOS.
- Don't forget to previus install Java.
