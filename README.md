# 概要
golangを用いたアプリケーション開発用リポジトリ

# 実行手順

```bash
$ # build and run container
$ make build
$
$ # exec application
$ make create-user
$
$ # delete container
$ make delete
$
$ # add package
$ make get dep=github.com/sirupsen/logrus
```