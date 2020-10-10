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

# デバッグ

```bash
$ kubectl get pods
NAME                       READY   STATUS    RESTARTS   AGE
graphql-5fb6b4d595-rx9kv   1/1     Running   0          43s
mysql-85cbfcfc7c-b4bd6     1/1     Running   1          7h59m
$ kubectl port-forward graphql-57dcbbdcf-nl8zc 56269:56269
```

localhost:56269に向けてremote debug