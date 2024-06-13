# Restaurant

**Project Description:**
Restoran bron qilish uchun qulaylashtirilgan, keng imkoniyatlarga ega kichkina API service.
Ushbu tizim orqali foydalanuvchilar restoran stollarini bron qilishlari, ovqatlarni tanlashlari, 
to'lov qilishlari va restoran ma'muriyati esa bronlarni boshqarishi mumkin bo'ladi.

## Installation

1. Initialize a git repository and clone the project:
    ```sh
    git init
    git clone git@github.com:Mubinabd/reservation_service.git
    ```
2. Create a database named `restaurant` on port `5432`.
3. Update the `.env` file with the appropriate configuration.
   ```.env
   DB_HOST=localhost
   DB_USER=postgres
   DB_NAME=restaurant
   DB_PASSWORD=pass
   DB_PORT=5432
   LOGPATH=logs/info.log
   ```

4. Use the following Makefile commands to manage the database migrations and set up the project:
    ```makefile
    # Set the database URL
    exp:
        export DBURL="postgres://postgres:1234@localhost:5432/restaurant?sslmode=disable'"

    # Run migrations
    mig-up:
        migrate -path migrations -database ${DBURL} -verbose up

    # Rollback migrations
    mig-down:
        migrate -path migrations -database ${DBURL} -verbose down

    # Create a new migration
    mig-create:
        migrate create -ext sql -dir migrations -seq create_table

    # Create an insert migration
    mig-insert:
        migrate create -ext sql -dir migrations -seq insert_table

    # Generate Swagger documentation
    swag:
        swag init -g api/api.go -o api/docs

    # Clean up migrations (commented out by default)
    # mig-delete:
    #   rm -r db/migrations
    ```
5. Set the environment variable and run the project:
    ```sh
    make exp
    make mig-up
    go run main.go
    ```
6. Open the following URL to access the Swagger documentation:
    ```
    http://localhost:8080/api/swagger/index.html#/
    ```

## Features and Usages
1. Auth serviceda User uchun Register va Login bo'limi bor va Get bo'limida user faqat o'zining profile ni ko'ra oladi.
2. Restaurant,Reservation,Order Create bo'limidan faqat ma'muriyat restaurant,reservation,order qo'sha oladi.
3. Restaurant,Reservation,Order serviceda Get,Update,Delete bo'limida faqat id lari orqali restaurant,reservation,orderni
ma'lumotlarini olish,yangilash yokida o'chirish mumkin.
4. Reservation serviceda  POST metodi orqali bron mavjud yoki yo'qligini tekshirsa bo'ladi.
5. Joylarni bron qilishda restaurant ma'muriyati nazorat qilib boradi,bron qilingan bo'lsa qabul qilinmaydi,bron qilinmagan bo'lsa
bron qilinib menyu tanlash uchun o'tiladi
6. Bron qilishda ovqatlarni tanlash va bron qilish orqali payments servicega ulangan holda to'lovni amalga oshirish mumkin.
7. Bron qilingan joy va menyudagi taomlarga qarab jami total summa hisoblanadi va payent servicega ulanadi.
8. Payments serviceda reservation tablega ulangan holda reservationID berish orqali toatl summani avtomatik ravishda amount qismidan ayira boshlaydi.
9. Payments serviceda to'lov holatini tekshiradi kartada mablag' yetarli yoki yetarli emasligini,yetarli bo'lmasa ogohlantiradi,
agar yetarli bo'lsa kerakli summani yechib oladi.
10. Bron qilingan joy va menyuga qarab narxlar o'zgarib turadi. 

## Dependencies

- **Scheduling**: [github.com/go-co-op/gocron](https://github.com/go-co-op/gocron)
- **Swagger**: [github.com/swaggo/swag](https://github.com/swaggo/swag)
- **Database**:
    - [database/sql](https://golang.org/pkg/database/sql/)
    - [github.com/lib/pq](https://github.com/lib/pq)
- **Environment Variables**: [github.com/joho/godotenv](https://github.com/joho/godotenv)
- **API Framework**: [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- **Migrations**: [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
****
## Acknowledgments

- Mubina Bahodirova
- Shamsiddin Okilov
- Azizbek Sobirov
- Abdumajid Abdukarimov
- Qodir

## Known Issues
1. Ba'zi testlarda kamchilik va xatoliklar bo'lishi mumkin;