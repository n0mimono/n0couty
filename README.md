# n0couty

AIベースのスカウト支援アプリ（仮、WIP）

# Screenshots

https://github.com/n0mimono/n0couty/wiki

---

# Setup / Start

## 環境変数

```
# http: application server
export PORT=8080

# MySQL
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_USER=username
export DB_NAME=dbname
export DB_PASSWORD=dbpassword
export DB_ROOT_PASSWORD=dbrootpassword

# http: ml server
export ML_HOST=127.0.0.1
export ML_PORT=5000
```

## go

```
go run main.go
```

## MySQL

```
mysql.server start
```

## python

```
python ./ml/main.py
```
