BOOK STORE

Pre-requisites
Install golang(1.9+). Follow the instructions for your environment from https://golang.org/dl/
Install Go GIN: "github.com/gin-gonic/gin"
Install gorm: "go get gorm.io/gorm"
Install gorm driver: "gorm.io/driver/mysql"
Add $GOPATH/bin to your PATH
Mysql

POSTMAN
Configured the file with postman to see the results of the requests that is being called
Postman download: "https://www.postman.com/downloads/"

MYSQL 
All the tables and database are created here.
Workbench download: "https://dev.mysql.com/downloads/workbench/"

Work Flow:
User inputs the title, author, and genre of the required book - User Friendly
storage of new books in library - librarian frindly

Run tests:
All the unit tests are stored here: "https://rb.gy/cfe4x"
Run tests for author table, move to "FINAL\API\internal\lms\author\service_test.go" and run: go test
Run tests for book table, move to "FINAL\API\internal\lms\book\service_test.go" and run: go test
Run tests for genre table, move to "FINAL\API\internal\lms\genre\service_test.go" and run: go test
Run tests for library table, move to "FINAL\API\internal\lms\library\service_test.go" and run: go test