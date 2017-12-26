# Planet THC

[![Build Status](https://travis-ci.org/teresinahc/planet.svg?branch=master)](https://travis-ci.org/teresinahc/planet)

## Mas como assim não é apenas um planet?

Isso é uma aposta!!!

Você tem um mês para escrever pelo menos 1 blogpost. Se você não fizer, você terá que pagar cerveja pra todo mundo :).

## Quem são os blogueiros?

* Gustavo Carvalho: [Blog](http://blog.gtsalles.com.br) ~ [Feed](http://blog.gtsalles.com.br/tags/beerblogging/index.xml)
* Filipe Saraiva: [Blog](http://blog.filipesaraiva.info/) ~ [Feed](http://blog.filipesaraiva.info/?tag=planet-thc&feed=rss2)
* Jonhnny Weslley: [Blog](http://raciocinio-lateral.jonhnnyweslley.net) ~ [Feed](http://feeds.feedburner.com/RaciocinioLateralBlogSpot)
* Ruan Aragão: [Blog](http://ruanaragao.github.io) ~ [Feed](http://ruanaragao.github.io/feed)

## Se juntando a nós

1. Você precisa forkar esse repositório
2. Adicione seus dados ao arquivo `members.json` seguindo o modelo:

  ```json
    "seu username": {
      "name": "seu nome",
      "email": "seuemail@domain.com",
      "blog": "http://linkdoseublog.com",
      "feed": "http://linkdofeeddoseublog.com",
      "twitter": "seuusuario",
      "date_joined": "2015-11-30"
    }
  ```

3. Edite o arquivo README.md, e adicione a url do seu blog e feed.
4. Envie um Pull request.
5. ESCREVA!

## Falando de código...

Estamos usando:

* [Golang](http://golang.org/)

## Instalação

    git clone git@github.com:teresinahc/planet.git
    cd planet
    go build
    ./planet

Visite o endereço [localhost:9000](http://localhost:9000).

Para mais opções execute:

    ./planet -h


OBS: Esse planet é um fork do [Beerblogging](https://github.com/avelino/beerblogging) e foi adaptado para as necessidades do Teresina Hacker Clube.
