# Planet THC

[![Build Status](https://travis-ci.org/teresinahc/thc-blog.svg?branch=master)](https://travis-ci.org/teresinahc/thc-blog)

## Mas como assim não é apenas um planet?

Isso é uma aposta!!!

Você tem um mês para escrever pelo menos 1 blogpost. Se você não fizer, você terá que pagar cerveja pra todo mundo :).

## Quem são os blogueiros?

* Gustavo Carvalho: [Blog](http://blog.gtsalles.com.br) ~ [Feed](http://blog.gtsalles.com/index.xml)
* Filipe Saraiva: [Blog](http://blog.filipesaraiva.info/) ~ [Feed](http://blog.filipesaraiva.info/?tag=THC-blog&feed=rss2)
* Jonhnny Weslley: [Blog](http://raciocinio-lateral.jonhnnyweslley.net) ~ [Feed](http://feeds.feedburner.com/RaciocinioLateralBlogSpot)

## Se juntando a nós

1. Você precisa forkar esse repositório
2. Adicione seus dados ao arquivo `members.yaml` seguindo o modelo:

  ```YAML
    ---
    name: seu nome
    email: seuemail@domain.com
    blog: http://linkdoseublog.com
    feed: http://linkdofeeddoseublog.com
    twitter: seuusuario
    date_joined: !!timestamp 'Y-m-d H:M:s'
    tags: tags, do, seu, blog
    id: membroAnterior.id++
  ```

3. Edite o arquivo README.md, e adicione a url do seu blog e feed.
4. Envie um Pull request.
5. ESCREVA!

## Falando de código...

Estamos usando:

* [Python](http://python.org/) linguagem de programação
* [Flask](http://flask.pocoo.org/) microframework web
* [Feed Parser](http://www.feedparser.org/) biblioteca

## Instalação

1. Primeiro você precisa clonar esse repositório: `git clone git@github.com:teresinahc/planet.git`
* Então instale todos os requisitos rodando `pip install -r requirements.txt`
* Execute `./manager.py create_db` para criar o banco de dados
* Execute `./manager.py fetch_posts` para popular o banco de dados
* Execute `./manager.py run` para rodar o servidor em [localhost:5000](http://localhost:5000)


## Testando

Apenas execute:

```
$ make test
```

E isso deve executar os testes para você.

## Realizando Deploy

Para fazer o deploy no Heroku:

* `heroku create`
* `heroku addons:add heroku-postgresql:dev`
* `heroku pg:promote HEROKU_POSTGRESQL_COLOR_URL`
* `git push heroku master`
* `heroku run python manager.py create_db`
* `heroku run ./update_posts.sh`

**OBS:** O script `update_posts.sh` deve ser executado a todo tempo, para coletar os novos posts.

OBS 2: Esse planet é um fork do [Beerblogging](https://github.com/avelino/beerblogging) e foi adaptado para as necessidades do Teresina Hacker Clube.
