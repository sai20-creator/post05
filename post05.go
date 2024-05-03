// package post05

// import (
// 	"database/sql"
// 	"errors"
// 	"fmt"
// 	"strings"

// 	_ "github.com/lib/pq"
// )

// // Connection details
// var (
// 	Hostname = ""
// 	Port     = 2345
// 	Username = ""
// 	Password = ""
// 	Database = ""
// )

// // Userdata is for holding full user data
// // Userdata table + Username
// type Userdata struct {
// 	ID          int
// 	Username    string
// 	Name        string
// 	Surname     string
// 	Description string
// }

// func openConnection() (*sql.DB, error) {
// 	// connection string
// 	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		Hostname, Port, Username, Password, Database)

// 	// open database
// 	db, err := sql.Open("postgres", conn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

// // The function returns the User ID of the username
// // -1 if the user does not exist
// func exists(username string) int {
// 	username = strings.ToLower(username)

// 	db, err := openConnection()
// 	if err != nil {
// 		fmt.Println(err)
// 		return -1
// 	}
// 	defer db.Close()

// 	userID := -1
// 	statement := fmt.Sprintf(`SELECT "id" FROM "users" where username = '%s'`, username)
// 	rows, err := db.Query(statement)

// 	for rows.Next() {
// 		var id int
// 		err = rows.Scan(&id)
// 		if err != nil {
// 			fmt.Println("Scan", err)
// 			return -1
// 		}
// 		userID = id
// 	}
// 	defer rows.Close()
// 	return userID
// }

// // AddUser adds a new user to the database
// // Returns new User ID
// // -1 if there was an error
// func AddUser(d Userdata) int {
// 	d.Username = strings.ToLower(d.Username)

// 	db, err := openConnection()
// 	if err != nil {
// 		fmt.Println(err)
// 		return -1
// 	}
// 	defer db.Close()

// 	userID := exists(d.Username)
// 	if userID != -1 {
// 		fmt.Println("User already exists:", Username)
// 		return -1
// 	}

// 	insertStatement := `insert into "users" ("username") values ($1)`
// 	_, err = db.Exec(insertStatement, d.Username)
// 	if err != nil {
// 		fmt.Println(err)
// 		return -1
// 	}

// 	userID = exists(d.Username)
// 	if userID == -1 {
// 		return userID
// 	}

// 	insertStatement = `insert into "userdata" ("userid", "name", "surname", "description")
// 	values ($1, $2, $3, $4)`
// 	_, err = db.Exec(insertStatement, userID, d.Name, d.Surname, d.Description)
// 	if err != nil {
// 		fmt.Println("db.Exec()", err)
// 		return -1
// 	}

// 	return userID
// }

// // DeleteUser deletes an existing user
// func DeleteUser(id int) error {
// 	db, err := openConnection()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	// Does the ID exist?
// 	statement := fmt.Sprintf(`SELECT "username" FROM "users" where id = %d`, id)
// 	rows, err := db.Query(statement)

// 	var username string
// 	for rows.Next() {
// 		err = rows.Scan(&username)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	defer rows.Close()

// 	if exists(username) != id {
// 		return fmt.Errorf("User with ID %d does not exist", id)
// 	}

// 	// Delete from Userdata
// 	deleteStatement := `delete from "userdata" where userid=$1`
// 	_, err = db.Exec(deleteStatement, id)
// 	if err != nil {
// 		return err
// 	}

// 	// Delete from Users
// 	deleteStatement = `delete from "users" where id=$1`
// 	_, err = db.Exec(deleteStatement, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// // ListUsers lists all users in the database
// func ListUsers() ([]Userdata, error) {
// 	Data := []Userdata{}
// 	db, err := openConnection()
// 	if err != nil {
// 		return Data, err
// 	}
// 	defer db.Close()

// 	rows, err := db.Query(`SELECT "id","username","name","surname","description"
// 		FROM "users","userdata"
// 		WHERE users.id = userdata.userid`)
// 	if err != nil {
// 		return Data, err
// 	}

// 	for rows.Next() {
// 		var id int
// 		var username string
// 		var name string
// 		var surname string
// 		var description string
// 		err = rows.Scan(&id, &username, &name, &surname, &description)
// 		temp := Userdata{ID: id, Username: username, Name: name, Surname: surname, Description: description}
// 		Data = append(Data, temp)
// 		if err != nil {
// 			return Data, err
// 		}
// 	}
// 	defer rows.Close()
// 	return Data, nil
// }

// // UpdateUser is for updating an existing user
// func UpdateUser(d Userdata) error {
// 	db, err := openConnection()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	userID := exists(d.Username)
// 	if userID == -1 {
// 		return errors.New("User does not exist")
// 	}
// 	d.ID = userID
// 	updateStatement := `update "userdata" set "name"=$1, "surname"=$2, "description"=$3 where "userid"=$4`
// 	_, err = db.Exec(updateStatement, d.Name, d.Surname, d.Description, d.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

package post05

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connection details
var (
	Hostname = ""
	Port     = 2345
	Username = ""
	Password = ""
	Database = ""
)

// MSDSCourse struct defines the structure of an MSDS course
type MSDSCourse struct {
	CID     string
	CNAME   string
	CPREREQ string
}

// openConnection opens a connection to the PostgreSQL database
func openConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Hostname, Port, Username, Password, Database)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// AddCourse adds a new course to the database and returns the new course ID
func AddCourse(course MSDSCourse) (int, error) {
	db, err := openConnection()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	// Insert into "MSDSCourseCatalog" table
	insertStmt := `INSERT INTO "MSDSCourseCatalog" ("CID", "CNAME", "CPREREQ") VALUES ($1, $2, $3) RETURNING "ID"`
	var courseID int
	err = tx.QueryRow(insertStmt, course.CID, course.CNAME, course.CPREREQ).Scan(&courseID)
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return courseID, nil
}

// DeleteCourse deletes an existing course from the database
func DeleteCourse(courseID int) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Delete from "MSDSCourseCatalog" table
	deleteStmt := `DELETE FROM "MSDSCourseCatalog" WHERE "ID" = $1`
	_, err = db.Exec(deleteStmt, courseID)
	if err != nil {
		return err
	}

	return nil
}

// ListCourses lists all courses in the database
func ListCourses() ([]MSDSCourse, error) {
	var courses []MSDSCourse

	db, err := openConnection()
	if err != nil {
		return courses, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT "CID", "CNAME", "CPREREQ" FROM "MSDSCourseCatalog"`)
	if err != nil {
		return courses, err
	}
	defer rows.Close()

	for rows.Next() {
		var course MSDSCourse
		err := rows.Scan(&course.CID, &course.CNAME, &course.CPREREQ)
		if err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

// UpdateCourse updates an existing course in the database
func UpdateCourse(courseID int, newCourse MSDSCourse) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	// Update "MSDSCourseCatalog" table
	updateStmt := `UPDATE "MSDSCourseCatalog" SET "CID" = $1, "CNAME" = $2, "CPREREQ" = $3 WHERE "ID" = $4`
	_, err = db.Exec(updateStmt, newCourse.CID, newCourse.CNAME, newCourse.CPREREQ, courseID)
	if err != nil {
		return err
	}

	return nil
}
