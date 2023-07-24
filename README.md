# Go_Fiber_CRUD
Go Fiber Framework'ü ve PostgreSQL veritabanı üzerinde temel CRUD (Create, Read, Update, Delete) İşlemleri

Öncelikle, Docker içinde PostgreSQL veritabanınızı çalıştırın:
Adım 1: PostgreSQL Veritabanını Docker İçinde Çalıştırma
docker run -d \
  --name postgres-db \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_PASSWORD=mypassword \
  -e POSTGRES_DB=mydatabase \
  -p 5432:5432 \
  postgres
Adım 2: Go Fiber ve Gerekli Paketlerin Kurulumu
go get -u github.com/gofiber/fiber/v2
go get -u github.com/gofiber/fiber/v2/middleware/cors
go get -u github.com/gofiber/fiber/v2/middleware/logger
go get -u github.com/jackc/pgx/v4/pgxpool
Adım 3: MVC (Model-View-Controller) mantığına uygun olarak kodumuzu bölümlere ayıralım
main.go: Uygulamayı başlatan ve temel yapılandırmayı yapan ana dosya.
models/task.go: Veritabanı işlemleri için Model dosyası.
controllers/task.go: HTTP istekleri için Controller dosyası.
routes/routes.go: Tüm route tanımlarının yer aldığı dosya.

Proje Dizin Yapısı
|- main.go
|- models
|   |- task.go
|- controllers
|   |- task.go
|- routes
|   |- routes.go
