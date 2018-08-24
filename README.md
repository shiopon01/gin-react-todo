# gin-react-todo

This repositry is sapmle project (use gin)

- gin
- goose
- xorm

# How to run this sample

- 1. Install node and go packages

```
yarn install
dep ensure
```

- 2. Create `bindata.go`

```
make
```

- 3. Start servers

```
yarn run dev
```

and

```
fresh
```

access to `localhost:8080`

# Supplement

- Release

```
make build
```

## memo

- create migrate file

`goose create create_users sql`
