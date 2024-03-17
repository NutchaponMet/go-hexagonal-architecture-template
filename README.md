# Go Hexagonal Architecture Template
### Go Hexagonal Architecture Template สำหรับผู้ที่สนใจอยากจะนำไปใช้เป็น template สำหรับขึ้นโปรเจคในการทำงาน API หลังบ้านต่าง ๆ ของ Web application
ความรู้เกี่ยวกับภาษา go อื่นๆ สามารถศึกษาเพิ่มเติมได้ที่ [codebangkok](https://github.com/codebangkok/golang) ซึ่งถือเป็นแหล่งความรู้ที่ทำให้ repository นี้เกิดขึ้น
#### คำสั่งสำหรับการ clone repository นี้
```
git clone https://github.com/NutchaponMet/go-hexagonal-architecture-template.git
```
> [!NOTE]
> ภายในโปรเจคนี้จะประกอบไปด้วย library ที่สำคัญอยู่หลายตัวที่ใช้สำหรับสร้าง Backend Web application
> ซึ่งผู้ที่นำ template นี้ไปใช้งานจะต้องมีความเข้าใจใน library พื้นฐานดังกล่าวก่อน ทั้งนี้ สามารถเข้าไปอ่านทำความเข้าใจเพิ่มเติมได้ตามหัวข้อด้านล่างนี้ได้เลย

1. [Go fiber](https://docs.gofiber.io/) | web framework
```
go get github.com/gofiber/fiber/v2
```
-----------
2. [Gorm](https://gorm.io/index.html) | ตัวจัดการงานเกี่ยวกับฐานข้อมูล ติดกับฐานข้อมูล
```
go get -u gorm.io/gorm
```
> ### Driver
> sqlite driver
> ```
> go get -u gorm.io/driver/sqlite
> ```
> MySql driver
> ```
> go get -u gorm.io/driver/mysql
> ```
> Postgres driver
> ```
> go get -u gorm.io/driver/postgres
> ```
------------
3. [zap](https://pkg.go.dev/go.uber.org/zap) | ตัวจัดการ log
```
go get go.uber.org/zap
```
------------
4. [viper](https://pkg.go.dev/github.com/spf13/viper) | ตัวจัดการ configuration
```
go get github.com/spf13/viper
```
------------
5. [mongodb driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo) | ใช้สำหรับเป็นที่เก็บ logs
```
go get go.mongodb.org/mongo-driver/mongo
```
------------
## Configuration File
##### config.yaml ---> dev configuration
```yaml
# dev config file
app:
  port: 5555
db: 
  username: "admin0001"
  password: "admin1234"
  host: "localhost"
  port: 4444
  dbname: "pgdb"
mongodb:
  driver: "mongodb"
  host: "10.1.1.96"
  port: 27017
  username: "api_user"
  password: "api1234"
  dbname: "api_dev_db"
```
