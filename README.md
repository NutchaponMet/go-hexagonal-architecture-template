# Go Go Hexagonal Architecture Template
### Go Hexagonal Architecture Template สำหรับผู้ที่สนใจอยากจะนำไปใช้เป็น template สำหรับขึ้นโปรเจคในการทำงาน API หลังบ้านต่างกับ Web application

#### คำสั่งสำหรับการ clone repository นี้
```
$ git clone https://github.com/NutchaponMet/go-hexagonal-architecture-template.git
```

ภายในโปรเจคนี้จะประกอบไปด้วย library ที่สำคัญอยู่หลายตัวที่ใช้สำหรับสร้าง Backend Web application
ซึ่งผุ้่ที่นำ template นี้ไปใช้งานจะต้องมีความเข้าใจใน library พื้นฐานดังกล่าวก่อน ทั้งนี้ สามารถเข้าไปอ่านทำความเข้าใจเพิ่มเติมได้จาก

1. [Go fiber](https://docs.gofiber.io/) | web framework
```
go get github.com/gofiber/fiber/v2
```
2. [Gorm](https://gorm.io/index.html) | ตัวจัดการงานเกี่ยวกับฐานข้อมูล ติดกับฐานข้อมูล
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite // optional
go get -u gorm.io/driver/mysql // optional
go get -u gorm.io/driver/postgres // optional
```
3. [zap](https://pkg.go.dev/go.uber.org/zap) | ตัวจัดการ log
```
go get go.uber.org/zap
```
4. [viper](https://pkg.go.dev/github.com/spf13/viper) | ตัวจัดการ configuration
```
go get github.com/spf13/viper
```
5. [mongodb driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo) | ใช้สำหรับเป็นที่เก็บ logs
```
go get go.mongodb.org/mongo-driver/mongo
```
