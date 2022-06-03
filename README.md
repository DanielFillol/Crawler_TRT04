# Crawler_TRT04

Projeto que visa capturar informações do site [TRT04](https://pesquisatextual.trt4.jus.br/pesquisas/acordaos?0). O objetivo é simplesmente obter as informações de forma a possibilitar o consumo das informações de maneira ágil.

Ao final são gerados um arquivo CSV com todos os tados pesquisados.
 
## Dependências
Para o esse crawler somos dependentes do projeto [selenium](https://github.com/tebeka/selenium#readme), sendo necessário a pré-instalação do [ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/) e do [selenium-server-standalone](https://selenium-release.storage.googleapis.com/index.html?path=3.5/)
Também somos dependentes do [html package](https://pkg.go.dev/golang.org/x/net/html) e do [htmlquery] (https://github.com/antchfx/htmlquery).


## Run
Basta dar o start no selenium server e iniciar o app em GO: 

```java -jar selenium-server-standalone.jar```

```go run main.go```

## Instalar - MacOS
```
Instalar o Google Chrome
Instalar o JAVA
Instalar o Brew

Baixar o selenium server stand alone\
Permitir o acesso ao arquivo (configurações de segurança)

brew install chrome river
Permitir o chromedriver no painel de segurança
```
