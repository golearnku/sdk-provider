# provider
mysql 、redis  provider

# `Mysql` Provider 使用方式:

## 初始化 mysql 实例对象:

```go
config := mysql.Config{
    Debug:    false,
    DBName:   "db",
    User:     "root",
    Password: "root",
    Adds:     "127.0.01",
}
client, err := mysql.NewClient(config)
if err != nil {
    panic(err)
}
```

## 使用 mysql 实例对象进行数据库操作:

```go
db := mysql.GetConnDB()
count := 0
if err := db.Table("order").Where("id = ?", 201906050027).Count(&count).Error; err != nil {
    log.Fatal(err)
}
fmt.Printf("count: %d \n", count)
```

# `Redis`  Provider 使用方式:

## 初始化 Redis 实例对象:

```go
config := redis.Config{
    Addr:     "127.0.0.1:6379",
    Password: "",
    PoolSize: 100,
    DB:       0,
}
client, err = redis.NewClient(config)
if err != nil {
    panic(err)
}
fmt.Printf("client: %+v \n", client)
```

## 使用 Redis 实例对象进行 Redis 操作:

```go
db := redis.GetConnDB()
if err := db.Set("lock", "1", time.Second*2).Err(); err != nil {
    log.Fatal(err)
}

val := db.Get("lock").Val()
fmt.Printf("val: %s \n", val)
```