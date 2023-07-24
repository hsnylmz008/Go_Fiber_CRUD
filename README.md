# Go_Fiber_CRUD
Go Fiber Framework'ü ve PostgreSQL veritabanı üzerinde temel CRUD (Create, Read, Update, Delete) İşlemleri

Öncelikle, Docker içinde PostgreSQL veritabanınızı çalıştırın:

## Adım 1: PostgreSQL Veritabanını Docker İçinde Çalıştırma

```bash
  docker run -d \
  --name postgres-db \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_PASSWORD=mypassword \
  -e POSTGRES_DB=mydatabase \
  -p 5432:5432 \
  postgres
```

