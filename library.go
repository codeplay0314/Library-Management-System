package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
	"github.com/asaskevich/govalidator"

	// mysql connector
	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

const (
	User     = "root"
	Password = "123456"
	DBName   = "SLMS"
)

var lib Library
type Library struct {
	db *sqlx.DB
}

type Admin struct {
	id, name, password, status, info string
}
func (admin *Admin) scan(rows *sqlx.Rows) error {
	return rows.Scan(&admin.id, &admin.name, &admin.password, &admin.status, &admin.info)
}
func (admin *Admin) insert() error {
	str := `INSERT INTO admin(id, aname, password, status, info) VALUES(?, ?, ?, ?, ?)`
	_, err := lib.db.Exec(str, admin.id, admin.name, admin.password, admin.status, admin.info)
	return err
}
func (admin *Admin) show() {
	fmt.Println("\n-------------")
	fmt.Println("[ID] ", admin.id)
	fmt.Println("[Name] ", admin.name)
	fmt.Println("[Status] ", admin.status)
	fmt.Println("[info] ", admin.info)
	fmt.Println("-------------")
}

type Student struct {
	id, name, password, status, info string
}
func (stu *Student) scan(rows *sqlx.Rows) error {
	return rows.Scan(&stu.id, &stu.name, &stu.password, &stu.status, &stu.info)
}
func (stu *Student) insert() error {
	str := `INSERT INTO students(id, sname, password, status, info) VALUES(?, ?, ?, ?, ?)`
	_, err := lib.db.Exec(str, stu.id, stu.name, stu.password, stu.status, stu.info)
	return err
}
func (stu *Student) show() {
	fmt.Println("\n-------------")
	fmt.Println("[ID] ", stu.id)
	fmt.Println("[Name] ", stu.name)
	fmt.Println("[Status] ", stu.status)
	fmt.Println("[info] ", stu.info)
	fmt.Println("-------------")
}

type Book struct {
	id int
	ISBN, title, author, status, info string
}
func (book *Book) scan(rows *sqlx.Rows) error {
	return rows.Scan(&book.id, &book.ISBN, &book.title, &book.author, &book.status, &book.info)
}
func (book *Book) insert() error {
	str := `INSERT INTO books(id, ISBN, title, author, status, info) VALUES(?, ?, ?, ?, ?, ?)`
	_, err := lib.db.Exec(str, book.id, book.ISBN, book.title, book.author, book.status, book.info)
	return err
}
func (book *Book) show() {
	fmt.Println("\n-------------")
	fmt.Println("[ID] ", book.id)
	fmt.Println("[ISBN] ", book.ISBN)
	fmt.Println("[Title] ", book.title)
	fmt.Println("[Author] ", book.author)
	fmt.Println("[Status] ", book.status)
	fmt.Println("[Info] ", book.info)
	fmt.Println("-------------")
}

type BorrowRecord struct {
	id, bid, delay int
	stu, borDate, retDate, status, info string
}
func (book *BorrowRecord) scan(rows *sqlx.Rows) error {
	return rows.Scan(&book.id, &book.bid, &book.stu, &book.borDate, &book.delay, &book.retDate,
		&book.status, &book.info)
}
func (book *BorrowRecord) insert() error {
	str := `INSERT INTO borrow_record(record, bid, stu, borrowedDate, delay, returnDate, status, info)
VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := lib.db.Exec(str, book.id, book.bid, book.stu, book.borDate, book.delay, book.retDate,
		book.status, book.info)
	return err
}
func (book *BorrowRecord) show() {
	fmt.Println("\n-------------")
	fmt.Println("[Record ID] ", book.id)
	fmt.Println("[Book ID] ", book.bid)
	fmt.Println("[Borrowed By]", book.stu, "(Student ID)")
	fmt.Println("[Borrowed Time] ", book.borDate)
	if book.retDate != "2999-12-31" {
		fmt.Println("[Returned Time] ", book.retDate)
	}
	t, _ := time.Parse("2006-01-02", book.borDate)
	t = t.Add(time.Duration((3 + book.delay) * 24 * 7) * time.Hour)
	fmt.Println("[Expected Return Time] ", t.Format("2006-01-02"))
	fmt.Println("[Delayed Times] ", book.delay)
	fmt.Println("[Status] ", book.status)
	fmt.Println("[Info] ", book.info)
	fmt.Println("-------------")
}
func (rec *BorrowRecord) overdue() bool {
	t, _ :=  time.Parse("2006-01-02", rec.borDate)
	sub := time.Now().Sub(t)
	return sub.Hours() > float64((3 + rec.delay) * 24 * 7)
}

var config map[string]interface{}

func ClearScreen(str string) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Print(str)
}

func Use(t ...interface{}) {
}
func readLine() string {
	var s []byte
	for {
		var c byte
		fmt.Scanf("%c", &c);
		if c == '\n' {
			break
		}
		s = append(s, c)
	}
	return string(s)
}
func readInt(l int, r int) (int, bool) {
	var res int
	var str string
	var s []byte
	for {
		var c byte
		fmt.Scanf("%c", &c);
		if c == '\n' {
			break
		}
		s = append(s, c)
	}
	str = string(s)
	_, err := fmt.Sscanf(str, "%d", &res)
	if err != nil || fmt.Sprintf("%d", res) != str || (l >= 0 && res < l) || (r >= 0 && res > r) {
		return 0, false
	}
	return res, true
}

func (lib *Library) ConnectDB() {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", User, Password))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + DBName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = db.Exec("USE " + DBName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	lib.db = db
}

// CreateTables created the tables in MySQL
func (lib *Library) CreateTables() error {

	result, err := lib.db.Exec(`
		CREATE TABLE IF NOT EXISTS admin (
			id VARCHAR(20) PRIMARY KEY,
			aname VARCHAR(32) NOT NULL,
			password VARCHAR(32) NOT NULL,
			status VARCHAR(20) NOT NULL,
			info VARCHAR(100)
		)
	`) // admin
	if err != nil {
		panic(err.Error())
	}
	result, err = lib.db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id VARCHAR(20) PRIMARY KEY,
			sname VARCHAR(32) NOT NULL,
			password VARCHAR(32) NOT NULL,
			status VARCHAR(20) NOT NULL,
			info VARCHAR(100)
		)
	`) // students
	if err != nil {
		panic(err.Error())
	}
	result, err = lib.db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INT PRIMARY KEY,
			ISBN VARCHAR(20) NOT NULL,
			title VARCHAR(100) NOT NULL,
			author VARCHAR(30) NOT NULL,
			status VARCHAR(20) NOT NULL,
			info VARCHAR(100)
		)
	`) // books
	if err != nil {
		panic(err.Error())
	}
	result, err = lib.db.Exec(`
		CREATE TABLE IF NOT EXISTS borrow_record (
			record INT AUTO_INCREMENT PRIMARY KEY,
			bid INT,
			stu VARCHAR(20),
			borrowedDate DATE,
			delay SMALLINT,
			returnDate DATE,
			status VARCHAR (20),
			info VARCHAR(100),
			FOREIGN KEY (bid) REFERENCES books(id),
			FOREIGN KEY (stu) REFERENCES students(id)
		)
	`) // borrow_record
	if err != nil {
		panic(err.Error())
	}
	use := result
	result = use

	return nil
}

func (lib *Library) init() error {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	json.Unmarshal([]byte(byteValue), &config)

	lib.db.Exec(`DELETE FROM admin WHERE id = 0`)
	_, err = lib.db.Exec("INSERT IGNORE INTO admin(id, aname, password, status, info) VALUES (?, ?, ?, ?, ?)",
		0, config["rootName"], config["password"], config["status"], "protected")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var stu Student
	var stulist []Student
	rows, err := lib.db.Queryx(`SELECT * FROM students WHERE status = "active"`)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for rows.Next() {
		err = stu.scan(rows)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		stulist = append(stulist, stu)
	}

	for i, v := range(stulist) {
		rows, err = lib.db.Queryx(`SELECT * FROM borrow_record WHERE stu = ? AND status = "being used"`, v.id)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		var cnt int
		for rows.Next() {
			var rec BorrowRecord
			err = rec.scan(rows)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			if rec.overdue() {
				cnt++
			}
		}
		if cnt > 3 {
			_, err = lib.db.Exec(`UPDATE students SET status = "suspended", info = ? WHERE id = ?`,
				"suspended by root because of having more than 3 books overdue", v.id)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		}
		Use(i)
	}

	return nil
}

var commands = []string {
	"Add an admin account",
	"Add a student account",
	"Change the status of an admin account",
	"Change the status of a student account",

	"Borrow a book",
	"Return a book",
	"Add a book to library",
	"Remove a book from library",

	"Query info of a book in library",
	"Query borrow histroy info of a student",
	"Query all overdue books a student has borrowed",
	"Query all books a student has borrowed and not returned yet",
	"Check the deadline of a borrowed book",
	"Extend the deadline of returning a book",

	"Log out",
}
var access = [][]bool {
	//root
	{true, true, true, true, false, false, true, true, true, true, true, true, true, true, true},
	//student
	{false, false, false, false, true, true, false, false, true, true, true, true, true, true, true},
	//admin
	{false, true, false, true, false, false, true, true, true, true, true, true, true, true, true},
}

// Add an admin account
func Command_0(typ int, id string) {
	for {
		ClearScreen("[" + commands[0] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[0] + "]\n")
		var admin Admin

		ok := true

		mp := make(map[string]bool)
		rows, err := lib.db.Queryx("SELECT * FROM admin")
		if err != nil {
			if err.Error() != "Error 1046: No database selected" {
				fmt.Println(err.Error())
				time.Sleep(2 * time.Second)
				return
			}
		} else {
			for rows.Next() {
				err = admin.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					ok = false
					break
				}
				mp[admin.id] = true
			}
		}
		rows.Close()
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}
		for i := 1; ; i++ {
			if mp[fmt.Sprintf("%d", i)] == false {
				admin.id = fmt.Sprintf("%d", i)
				break
			}
		}

		fmt.Printf("Name: ")
		admin.name = readLine()
		if len(admin.name) > 32 {
			fmt.Println("Admin name should be less than 32 digits.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("Password: ")
		admin.password = readLine()
		if len(admin.password) > 32 {
			fmt.Println("Password should be less than 32 digits.")
			time.Sleep(2 * time.Second)
			continue
		}

		admin.status = "active"
		admin.info = ""

		for {
			ClearScreen("[" + commands[0] + "]\n")
			fmt.Printf("\n\nPlease check:\n[ID]: %s\n[Name]: %s\n[Password]: %s\n" +
				"Are you sure to add this admin account? (y/N)", admin.id, admin.name, admin.password)
			str := readLine()
			if str == "y" {
				ok = true
				break
			} else if str == "N" {
				ok = false
				break
			}
			fmt.Println("\nPlease enter y/N")
			time.Sleep(2 * time.Second)
		}

		if !ok {
			continue
		}

		err = admin.insert()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\nSuccessfully added!")
		}
		time.Sleep(3 * time.Second)
	}
}

// Add a student account
func Command_1(typ int, id string) {
	for {
		ClearScreen("[" + commands[1] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[1] + "]\n")
		var stu Student

		ok := true

		fmt.Printf("Student ID: ")
		stu.id = readLine()
		for i, v := range(stu.id) {
			if v == ' ' || v < '0' || v > '9' {
				fmt.Println("Student ID should only contain numbers.")
				ok = false
				break
			}
			Use(i)
		}
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}
		if len(stu.id) > 20 {
			fmt.Println("Student ID should be less than 20 digits.")
			time.Sleep(2 * time.Second)
			continue
		}
		rows, err := lib.db.Queryx("SELECT * FROM students WHERE id = \"" + stu.id + "\"")
		if err != nil {
			if err.Error() != "Error 1046: No database selected" {
				fmt.Println(err.Error())
				time.Sleep(2 * time.Second)
				return
			}
		} else {
			for rows.Next() {
				fmt.Println("The ID", stu.id, "already exists.")
				ok = false
				break
			}
		}
		rows.Close()
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("Name: ")
		stu.name = readLine()
		if len(stu.name) > 32 {
			fmt.Println("Student name should be less than 32 digits.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("Password: ")
		stu.password = readLine()
		if len(stu.password) > 32 {
			fmt.Println("Password should be less than 32 digit.")
			time.Sleep(2 * time.Second)
			continue
		}

		stu.status = "active"
		stu.info = "add by admin " + id


		for {
			ClearScreen("[" + commands[1] + "]\n")
			fmt.Printf("\n\nPlease check:\n[ID]: %s\n[Name]: %s\n[Password]: %s\n" +
				"Are you sure to add this student account? (y/N)", stu.id, stu.name, stu.password)
			str := readLine()
			if str == "y" {
				ok = true
				break
			} else if str == "N" {
				ok = false
				break
			}
			fmt.Println("\nPlease enter y/N")
			time.Sleep(2 * time.Second)
		}

		if !ok {
			continue
		}

		err = stu.insert()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\nSuccessfully added!")
		}
		time.Sleep(3 * time.Second)
	}
}

// Change the status of an admin account
func Command_2(typ int, id string) {
	for {
		ClearScreen("[" + commands[2] + "]\n\nAccount Lists:")

		rows, err := lib.db.Queryx("SELECT * FROM admin")
		if err != nil {
			if err.Error() != "Error 1046: No database selected" {
				fmt.Println(err.Error())
				time.Sleep(2 * time.Second)
				return
			}
		} else {
			for rows.Next() {
				var user Admin
				err = user.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					time.Sleep(2 * time.Second)
					return
				}
				user.show()
			}
		}
		rows.Close()
		fmt.Println("1. Begin\n0. Exit")

		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[2] + "]\n")

		ok := true
		var user, nuser Admin

		fmt.Printf("The ID of the admin account that you want to change: ")
		uid := readLine()

		rows, err = lib.db.Queryx("SELECT * FROM admin WHERE id = \"" + uid + "\"")
		if err != nil {
			if err.Error() != "Error 1046: No database selected" {
				fmt.Println(err.Error())
				time.Sleep(2 * time.Second)
				return
			}
		} else {
			found := false
			for rows.Next() {
				found = true
				err = user.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					time.Sleep(2 * time.Second)
					return
				}
			}
			if !found {
				fmt.Println("\nID not found! Please check again.")
				ok = false
			}
		}
		rows.Close()
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("You want to change the status of this account to: (suspended/active/...)")
		aim := readLine()

		fmt.Printf("The reason for changing account status: ")
		reason := readLine()

		for {
			ClearScreen("[" + commands[2] + "]\n")
			fmt.Printf("\n\nPlease check:\nYou want to change the admin account from")
			user.show()
			fmt.Printf("\nto")
			nuser = user
			nuser.status = aim
			nuser.info = "set" + aim + " by admin " + id + " because " + reason
			nuser.show()
			fmt.Printf("Do you confirm? (y/N)")
			str := readLine()
			if str == "y" {
				ok = true
				break
			} else if str == "N" {
				ok = false
				break
			}
			fmt.Println("\nPlease enter y/N")
			time.Sleep(2 * time.Second)
		}

		str := `UPDATE admin SET status = ?, info = ? WHERE id = ?`
		_, err = lib.db.Exec(str, aim, nuser.info, uid)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			return
		}
		fmt.Println("\nSuccessfully changed!")
		time.Sleep(3 * time.Second)
	}
}

// Change the status of a student account
func Command_3(typ int, id string) {
	for {
		ClearScreen("[" + commands[3] + "]\n\nAccount Lists:")

		rows, err := lib.db.Queryx("SELECT * FROM students")
		if err != nil {
			if err.Error() != "Error 1046: No database selected" {
				fmt.Println(err.Error())
				time.Sleep(2 * time.Second)
				return
			}
		} else {
			for rows.Next() {
				var user Student
				err = user.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					time.Sleep(2 * time.Second)
					return
				}
				user.show()
			}
		}
		rows.Close()
		fmt.Println("1. Begin\n0. Exit")

		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[3] + "]\n")

		ok := true
		var user, nuser Student

		fmt.Printf("The ID of the student account that you want to change: ")
		uid := readLine()

		rows, err = lib.db.Queryx("SELECT * FROM students WHERE id = \"" + uid + "\"")
		if err != nil {
			if err.Error() != "Error 1046: No database selected" {
				fmt.Println(err.Error())
				time.Sleep(2 * time.Second)
				return
			}
		} else {
			found := false
			for rows.Next() {
				found = true
				err = user.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					time.Sleep(2 * time.Second)
					return
				}
			}
			if !found {
				fmt.Println("\nID not found! Please check again.")
				ok = false
			}
		}
		rows.Close()
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("You want to change the status of this account to: (suspended/active/...)")
		aim := readLine()

		fmt.Printf("The reason for changing account status: ")
		reason := readLine()

		for {
			ClearScreen("[" + commands[3] + "]\n")
			fmt.Printf("\n\nPlease check:\nYou want to change the student account from")
			user.show()
			fmt.Printf("\nto")
			nuser = user
			nuser.status = aim
			nuser.info = aim + " by admin " + id + " because " + reason
			nuser.show()
			fmt.Printf("Do you confirm? (y/N)")
			str := readLine()
			if str == "y" {
				ok = true
				break
			} else if str == "N" {
				ok = false
				break
			}
			fmt.Println("\nPlease enter y/N")
			time.Sleep(2 * time.Second)
		}

		str := `UPDATE students SET status = ?, info = ? WHERE id = ?`
		_, err = lib.db.Exec(str, aim, nuser.info, uid)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			return
		}
		fmt.Println("\nSuccessfully changed!")
		time.Sleep(3 * time.Second)
	}
}

// Borrow a book
func Command_4(typ int, id string) {
	for {
		ClearScreen("[" + commands[4] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[4] + "]\n")
		var book Book

		ok := true

		fmt.Printf("Please enter ID number of the book: ")
		str := readLine()

		rows, err := lib.db.Queryx("SELECT * FROM books WHERE id = \"" + str + "\" AND status = \"available\"")
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			return
		} else {
			found := false
			for rows.Next() {
				found = true
				err = book.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					ok = false
					break
				}
			}
			if !found {
				fmt.Println("The book is not available currently.")
				time.Sleep(2 * time.Second)
				return
			}
		}
		rows.Close()

		for {
			ClearScreen("[" + commands[4] + "]\n")
			fmt.Printf("\n\nPlease check")
			book.show()
			fmt.Printf("Are you sure to borrow this book? (y/N)")
			str := readLine()
			if str == "y" {
				ok = true
				break
			} else if str == "N" {
				ok = false
				break
			}
			fmt.Println("\nPlease enter y/N")
			time.Sleep(2 * time.Second)
		}

		if !ok {
			continue
		}

		var n int
		rows, err = lib.db.Queryx("SELECT COUNT(*) FROM borrow_record")
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			continue
		}
		for rows.Next() {
			err = rows.Scan(&n)
			if err != nil {
				fmt.Println(err.Error())
				time.Sleep(2 * time.Second)
				return
			}
		}
		rows.Close()

		_, err = lib.db.Exec(`
			UPDATE books
			SET status = "borrowed", info = ?
			WHERE id = ?`, "borrowed by student " + id, book.id)
		if err != nil {
			fmt.Println(err.Error())
			ok = false
		} else {
			var rec BorrowRecord
			rec = BorrowRecord{n + 1, book.id, 0, id,
				time.Now().Format("2006-01-02"), "2999-12-31", "being used", ""}
			err = rec.insert()
			if err != nil {
				fmt.Println(err.Error())
				ok = false
			}
		}
		if !ok {
			_, err = lib.db.Exec(`DELETE FROM books WHERE ID = ?`, book.id)
			book.insert()
		} else {
			fmt.Println("Successfully borrowed!")
		}
		time.Sleep(3 * time.Second)

	}

}

// Return a book
func Command_5(typ int, id string) {
	for {
		ClearScreen("[" + commands[5] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[5] + "]\n")
		var rec BorrowRecord
		var recs []BorrowRecord

		ok := true

		rows, err := lib.db.Queryx(`SELECT * FROM borrow_record WHERE stu = ? AND status = "being used"`, id)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			return
		} else {
			found := false
			for rows.Next() {
				found = true
				err = rec.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					ok = false
					break
				}
				recs = append(recs, rec)
			}
			if !found {
				fmt.Println("You have no book to be returned!")
				time.Sleep(2 * time.Second)
				return
			}
		}
		rows.Close()

		var no int
		for {
			ClearScreen("[" + commands[5] + "]\n")
			fmt.Printf("\n\nPlease check the book you have borrowed: ")
			for i, v := range(recs) {
				fmt.Printf("\n\nNO.%d", i + 1)
				v.show()
			}
			fmt.Printf("You want to return book NO.")
			no, ok = readInt(1, len(recs))
			if !ok {
				if len(recs) == 1 {
					fmt.Printf("\nPlease enter a valid number(1)!\n")
				} else {
					fmt.Printf("\nPlease enter a valid number(1-%d)!\n", len(recs))
				}
				time.Sleep(2 * time.Second)
				continue
			}
			fmt.Printf("Are you sure to return the book NO.%d? (y/N)", no)
			str := readLine()
			if str == "y" {
				ok = true
				break
			} else if str == "N" {
				ok = false
				break
			}
			fmt.Println("\nPlease enter y/N")
			time.Sleep(2 * time.Second)
		}
		if !ok {
			continue
		}

		rec = recs[no - 1]

		_, err = lib.db.Exec(`UPDATE borrow_record SET returnDate = ?, status = "returned" WHERE record = ?`,
			time.Now().Format("2006-01-02"), rec.id)
		if err != nil {
			fmt.Println(err.Error())
			ok = false
		} else {
			_, err = lib.db.Exec(`UPDATE books SET status = "available" WHERE id = ?`, rec.bid)
			if err != nil {
				fmt.Println(err.Error())
				ok = false
			} else {
				fmt.Println("Successfully returned!")
			}
		}
		time.Sleep(3 * time.Second)

	}

}

// Add a book to library
func Command_6(typ int, id string) {
	for {
		ClearScreen("[" + commands[6] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[6] + "]\n")
		var book Book

		ok := true

		mp := make(map[int]bool)
		rows, err := lib.db.Queryx("SELECT * FROM books")
		if err != nil {
			if err.Error() != "Error 1046: No database selected" {
				fmt.Println(err.Error())
				time.Sleep(2 * time.Second)
				return
			}
		} else {
			for rows.Next() {
				err = book.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					ok = false
					break
				}
				mp[book.id] = true
			}
		}
		rows.Close()
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}
		for i := 1; ; i++ {
			if mp[i] == false {
				book.id = i
				break
			}
		}

		fmt.Printf("ISBN: ")
		book.ISBN = readLine()
		if len(book.ISBN) > 20 {
			fmt.Println("Book ISBN should be less than 20 digits.")
			time.Sleep(2 * time.Second)
			continue
		}
		if !govalidator.IsISBN10(book.ISBN) && !govalidator.IsISBN13(book.ISBN) {
			fmt.Println("Please enter valid ISBN number.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("Title: ")
		book.title = readLine()
		if len(book.title) > 100 {
			fmt.Println("Book Title should be less than 100 digits.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Printf("Authour: ")
		book.author = readLine()
		if len(book.author) > 30 {
			fmt.Println("Book Author should be less than 30 digits.")
			time.Sleep(2 * time.Second)
			continue
		}

		book.status = "available"
		book.info = ""

		for {
			ClearScreen("[" + commands[6] + "]\n")
			fmt.Printf("\n\nPlease check")
			book.show()
			fmt.Printf("Are you sure to add this book? (y/N)")
			str := readLine()
			if str == "y" {
				ok = true
				break
			} else if str == "N" {
				ok = false
				break
			}
			fmt.Println("\nPlease enter y/N")
			time.Sleep(2 * time.Second)
		}

		if !ok {
			continue
		}

		err = book.insert()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\nSuccessfully added!")
		}
		time.Sleep(3 * time.Second)
	}

}

// Remove a book from library
func Command_7(typ int, id string) {
	for {
		ClearScreen("[" + commands[7] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[7] + "]\n")
		var book Book

		ok := true

		fmt.Printf("Please enter ID number of the book: ")
		str := readLine()

		rows, err := lib.db.Queryx("SELECT * FROM books WHERE id = \"" + str + "\" AND status != \"removed\"")
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			return
		} else {
			found := false
			for rows.Next() {
				found = true
				err = book.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					ok = false
					break
				}
			}
			if !found {
				fmt.Println("The BOOK ID doesn't exist! Please check.")
				time.Sleep(2 * time.Second)
				return
			}
		}
		rows.Close()

		for {
			ClearScreen("[" + commands[7] + "]\n")
			fmt.Printf("\n\nPlease check")
			book.show()
			fmt.Printf("Are you sure to remove this book? (y/N)")
			str := readLine()
			if str == "y" {
				ok = true
				break
			} else if str == "N" {
				ok = false
				break
			}
			fmt.Println("\nPlease enter y/N")
			time.Sleep(2 * time.Second)
		}

		if !ok {
			continue
		}

		fmt.Print("The reason for removal: ")
		str = "removed by admin " + id + " because " + readLine()

		_, err = lib.db.Exec(`UPDATE books SET status = "removed", info = ? WHERE id = ?`, str, book.id)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_, err = lib.db.Exec(`UPDATE borrow_record SET status = "removed", info = ?
				WHERE bid = ? AND status = "being used";`, str, book.id)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("\nSuccessfully removed!")
			}
		}
		time.Sleep(3 * time.Second)
	}

}

// Query info of a book in library
func Command_8(typ int, id string) {
	for {
		ClearScreen("[" + commands[8] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[8] + "]\n")

		ok := true

		for {
			fmt.Println("1. Search by book ID")
			fmt.Println("2. Search by book ISBN")
			fmt.Println("3. Search by book title")
			fmt.Println("4. Search by author")
			fmt.Println("0. Exit")
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 4)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			continue
		}

		var book Book
		var booklist []Book

		var rows *sqlx.Rows
		var err error

		switch com {
		case 1:
			fmt.Println("Please enter book ID: ")
			rows, err = lib.db.Queryx(`SELECT * FROM books WHERE id = ?`, readLine())
		case 2:
			fmt.Println("Please enter book ISBN: ")
			rows, err = lib.db.Queryx(`SELECT * FROM books WHERE ISBN = ?`, readLine())
		case 3:
			fmt.Println("Please enter book title: ")
			rows, err = lib.db.Queryx(`SELECT * FROM books WHERE title = ?`, readLine())
		case 4:
			fmt.Println("Please enter author name: ")
			rows, err = lib.db.Queryx(`SELECT * FROM books WHERE author = ?`, readLine())
		}
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			return
		} else {
			for rows.Next() {
				err = book.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					ok = false
					break
				}
				booklist = append(booklist, book)
			}
			if len(booklist) == 0 {
				fmt.Println("No books found.")
				ok = false
			}
		}
		rows.Close()
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}

		ClearScreen("[" + commands[8] + "]\n")
		fmt.Printf("\n\nAll books found:")
		for i, v := range(booklist) {
			v.show()
			Use(i)
		}
		fmt.Println("\nType enter to conitnue...")
		readLine()
	}
}

// Query borrow histroy info of a student
func Command_9(typ int, id string) {
	for {
		ClearScreen("[" + commands[9] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[9] + "]\n")

		ok := true

		var stu Student

		var rows *sqlx.Rows
		var err error
		if typ == 0 {
			fmt.Println("Please enter student ID: ")
			rows, err = lib.db.Queryx(`SELECT * FROM students WHERE id = ?`, readLine())
			if err != nil {
				fmt.Println(err.Error())
				ok = false
			} else {
				found := false
				for rows.Next() {
					found = true
					stu.scan(rows)
					if err != nil {
						fmt.Println(err.Error())
						ok = false
					}
				}
				if ! found {
					fmt.Println("Invalid Student ID.")
					ok = false
				}
			}
		} else {
			rows, err = lib.db.Queryx(`SELECT * FROM students WHERE id = ?`, id)
			if err != nil {
				fmt.Println(err.Error())
				ok = false
			} else {
				for rows.Next() {
					stu.scan(rows)
					if err != nil {
						fmt.Println(err.Error())
						ok = false
					}
				}
			}
		}
		rows.Close()
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}

		rows, err = lib.db.Queryx(`SELECT * FROM borrow_record WHERE stu = ?`, stu.id)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			continue
		}

		ClearScreen("[" + commands[9] + "]\n")
		if typ == 0 {
			fmt.Printf("\n\nBorrow history of student %s:", stu.id)
		} else {
			fmt.Printf("\n\nYour borrow history:")
		}
		for rows.Next() {
			var rec BorrowRecord
			rec.scan(rows)
			rec.show()
		}
		rows.Close()
		fmt.Println("\nType enter to conitnue...")
		readLine()
	}
}

// Query all overdue books a student has borrowed
func Command_10(typ int, id string) {
	for {
		ClearScreen("[" + commands[10] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[10] + "]\n")

		ok := true

		var stu Student

		var rows *sqlx.Rows
		var err error
		if typ == 0 {
			fmt.Println("Please enter student ID: ")
			rows, err = lib.db.Queryx(`SELECT * FROM students WHERE id = ?`, readLine())
			if err != nil {
				fmt.Println(err.Error())
				ok = false
			} else {
				found := false
				for rows.Next() {
					found = true
					stu.scan(rows)
					if err != nil {
						fmt.Println(err.Error())
						ok = false
					}
				}
				if ! found {
					fmt.Println("Invalid Student ID.")
					ok = false
				}
			}
		} else {
			rows, err = lib.db.Queryx(`SELECT * FROM students WHERE id = ?`, id)
			if err != nil {
				fmt.Println(err.Error())
				ok = false
			} else {
				for rows.Next() {
					stu.scan(rows)
					if err != nil {
						fmt.Println(err.Error())
						ok = false
					}
				}
			}
		}
		rows.Close()
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}

		rows, err = lib.db.Queryx(`SELECT * FROM borrow_record WHERE stu = ?`, stu.id)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			continue
		}

		ClearScreen("[" + commands[10] + "]\n")
		if typ == 0 {
			fmt.Printf("\n\nOverdue borrowing of student %s:", stu.id)
		} else {
			fmt.Printf("\n\nYour overdue borrowing:")
		}
		for rows.Next() {
			var rec BorrowRecord
			rec.scan(rows)
			if rec.overdue() {
				rec.show()
			}
		}
		rows.Close()
		fmt.Println("\nType enter to conitnue...")
		readLine()
	}
}

// Query all books a student has borrowed and not returned yet
func Command_11(typ int, id string) {
	for {
		ClearScreen("[" + commands[11] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ClearScreen("[" + commands[11] + "]\n")

		ok := true

		var stu Student

		var rows *sqlx.Rows
		var err error
		if typ == 0 {
			fmt.Println("Please enter student ID: ")
			rows, err = lib.db.Queryx(`SELECT * FROM students WHERE id = ?`, readLine())
			if err != nil {
				fmt.Println(err.Error())
				ok = false
			} else {
				found := false
				for rows.Next() {
					found = true
					stu.scan(rows)
					if err != nil {
						fmt.Println(err.Error())
						ok = false
					}
				}
				if ! found {
					fmt.Println("Invalid Student ID.")
					ok = false
				}
			}
		} else {
			rows, err = lib.db.Queryx(`SELECT * FROM students WHERE id = ?`, id)
			if err != nil {
				fmt.Println(err.Error())
				ok = false
			} else {
				for rows.Next() {
					stu.scan(rows)
					if err != nil {
						fmt.Println(err.Error())
						ok = false
					}
				}
			}
		}
		rows.Close()
		if !ok {
			time.Sleep(2 * time.Second)
			continue
		}

		rows, err = lib.db.Queryx(`SELECT * FROM borrow_record WHERE stu = ? AND status = "being used"`, stu.id)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
			continue
		}

		ClearScreen("[" + commands[11] + "]\n")
		if typ == 0 {
			fmt.Printf("\n\nBorrorwed books of student %s:", stu.id)
		} else {
			fmt.Printf("\n\nAll your borrowed books:")
		}
		for rows.Next() {
			var rec BorrowRecord
			rec.scan(rows)
			rec.show()
		}
		rows.Close()
		fmt.Println("\nType enter to conitnue...")
		readLine()
	}
}

// Check the deadline of a borrowed book
func Command_12(typ int, id string) {
	for {
		ClearScreen("[" + commands[12] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ok := true
		var rec BorrowRecord

		for {
			ClearScreen("[" + commands[12] + "]\n")
			fmt.Print("Book ID: ")
			bid, valid := readInt(1, -1)
			rows, err := lib.db.Queryx(`SELECT * FROM borrow_record WHERE bid = ?`, bid)
			found := false
			for rows.Next() {
				err = rec.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					ok = false
				}
				found = true
			}
			if !ok {
				break
			}
			if !valid || !found {
				fmt.Println("Please enter valid Book ID!")
				time.Sleep(2 * time.Second)
				continue
			}
			break
		}
		if !ok {
			continue
		}

		if rec.status != "being used" {
			fmt.Println("The book is not being borrowed. Please go to check its status.")
		} else {
			t, _ := time.Parse("2006-01-02", rec.borDate)
			t = t.Add(time.Duration((3 + rec.delay) * 24 * 7) * time.Hour)
			fmt.Println("The book is expected to be return on", t.Format("2006-01-02"))
		}
		fmt.Println("\nType enter to conitnue...")
		readLine()

	}

}

// Extend the deadline of returning a book
func Command_13(typ int, id string) {
	for {
		ClearScreen("[" + commands[13] + "]\n1. Begin\n0. Exit\n")
		var com int
		for {
			fmt.Printf("You want to(1/0): ")
			var ok bool
			com, ok = readInt(0, 2)
			if !ok {
				fmt.Println("\nPlease enter a valid number!")
				continue
			}
			break
		}
		if com == 0 {
			break
		}

		ok := true
		var rec BorrowRecord

		for {
			ClearScreen("[" + commands[13] + "]\n")
			fmt.Print("Book ID: ")
			bid, valid := readInt(1, -1)
			rows, err := lib.db.Queryx(`SELECT * FROM borrow_record WHERE bid = ?`, bid)
			found := false
			for rows.Next() {
				err = rec.scan(rows)
				if err != nil {
					fmt.Println(err.Error())
					ok = false
				}
				found = true
			}
			if !ok {
				break
			}
			if !valid || !found {
				fmt.Println("Please enter valid Book ID!")
				time.Sleep(2 * time.Second)
				continue
			}
			break
		}
		if !ok {
			continue
		}

		if rec.status != "being used" {
			fmt.Println("The book is not being borrowed. Please go to check its status.")
		} else if typ == 1 && rec.stu != id {
			fmt.Println("You can only extend the deadline of a book that you have borrowed.")
		} else if rec.delay >= 3 {
			fmt.Println("Sorry, the return date has been already extended for 3 times.")
		} else {
			_, err := lib.db.Exec(`UPDATE borrow_record SET delay = ? WHERE record = ?`, rec.delay + 1, rec.id)
			if err != nil {
				fmt.Println(err.Error())
				ok = false
			} else {
				t, _ := time.Parse("2006-01-02", rec.borDate)
				t = t.Add(time.Duration((3 + rec.delay + 1) * 24 * 7) * time.Hour)
				fmt.Println("Successfully extended! Now the book should be returned by", t.Format("2006-01-02"))
			}
		}
		fmt.Println("\nType enter to conitnue...")
		readLine()

	}
}

func run_command(no int, typ int, id string) {
	switch no {
	case 0: Command_0(typ, id)
	case 1: Command_1(typ, id)
	case 2: Command_2(typ, id)
	case 3: Command_3(typ, id)
	case 4: Command_4(typ, id)
	case 5: Command_5(typ, id)
	case 6: Command_6(typ, id)
	case 7: Command_7(typ, id)
	case 8: Command_8(typ, id)
	case 9: Command_9(typ, id)
	case 10: Command_10(typ, id)
	case 11: Command_11(typ, id)
	case 12: Command_12(typ, id)
	case 13: Command_13(typ, id)
	}
}

func work_student(exit *bool) {
	ClearScreen("Please log in to a student account\nStudent ID: ")
	var user Student
	uid := readLine()
	fmt.Print("Password: ")
	upwd := readLine()
	rows, err := lib.db.Queryx(`SELECT * FROM students WHERE id = "` + uid + `" AND password = "` + upwd + `"`)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(3 * time.Second)
		return
	}
	found := false
	for rows.Next() {
		err = user.scan(rows)
		found = true
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(3 * time.Second)
			return
		}
	}
	rows.Close()
	if !found {
		fmt.Println("\nLogin Failed!\nPlease enter correct ID and password...\n")
		time.Sleep(3 * time.Second)
		return
	}
	if user.status != "active" {
		fmt.Println("\nYour account is not activated. Please contact the librarian!")
		if user.info != "" {
			fmt.Println("More info:", user.info)
		}
		time.Sleep(5 * time.Second)
		return
	}

	var acc []int
	for i, v := range(access[1]) {
		if v {
			acc = append(acc, i)
		}
	}
	n := len(acc)

	first := true
	for {
		ClearScreen("")
		if first {
			first = false
			fmt.Println("Welcome", user.name + "!")
		}
		fmt.Println("What do you want to do?")
		for i, v := range(acc) {
			fmt.Printf("%2d: %s\n", (i + 1) % n, commands[v])
		}

		var com int
		for {
			fmt.Printf("You need to do(0-%d): ", n - 1)
			var ok bool
			com, ok = readInt(0, n - 1)
			if !ok {
				fmt.Printf("\nPlease enter a valid number(0-%d)!\n", n - 1)
				continue
			}
			break
		}
		if com == 0 {
			break
		}
		run_command(acc[(com - 1 + n) % n], 1, user.id)
	}

}
func work_admin(exit *bool) {
	var user Admin
	ClearScreen("Please log in to an admin account\nAdmin ID: ")
	uid := readLine()
	fmt.Print("Password: ")
	upwd := readLine()
	rows, err := lib.db.Queryx(`SELECT * FROM admin WHERE id = "` + uid + `" AND password = "` + upwd + `"`)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(3 * time.Second)
		return
	}
	found := false
	for rows.Next() {
		err = user.scan(rows)
		found = true
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(3 * time.Second)
			return
		}
	}
	if !found {
		fmt.Println("\nLogin Failed!\nPlease enter correct ID and password...\n")
		time.Sleep(3 * time.Second)
		return
	}
	if user.status != "active" {
		fmt.Println("\nYour account is not activated.")
		if user.info != "" {
			fmt.Println("More info:", user.info)
		}
		time.Sleep(5 * time.Second)
		return
	}

	var acc []int
	for i, v := range(access[2]) {
		if v {
			acc = append(acc, i)
		}
	}
	n := len(acc)

	first := true
	for {
		ClearScreen("")
		if first {
			first = false
			fmt.Println("Welcome", user.name + "!")
		}
		fmt.Println("What do you want to do?")
		for i, v := range(acc) {
			fmt.Printf("%2d: %s\n", (i + 1) % n, commands[v])
		}

		var com int
		for {
			fmt.Printf("You need to do(0-%d): ", n - 1)
			var ok bool
			com, ok = readInt(0, n - 1)
			if !ok {
				fmt.Printf("\nPlease enter a valid number(0-%d)!\n", n - 1)
				continue
			}
			break
		}
		if com == 0 {
			break
		}
		run_command(acc[(com - 1 + n) % n], 0, user.id)
	}

}
func work_root(exit *bool) {
	ClearScreen("Please log in to the root account\npassword: ")
	str := readLine()
	if str != config["password"] {
		fmt.Println("\nLogin Failed: Wrong Password...\n")
		time.Sleep(3 * time.Second)
		return
	}

	var acc []int
	for i, v := range(access[0]) {
		if v {
			acc = append(acc, i)
		}
	}
	n := len(acc)

	first := true
	for {
		ClearScreen("")
		if first {
			first = false
			fmt.Println("Successfully log in to root account!\n")
		}
		fmt.Println("Please select service below:")
		for i, v := range(acc) {
			fmt.Printf("%2d: %s\n", (i + 1) % n, commands[v])
		}

		var com int
		for {
			fmt.Printf("You need to do(0-%d): ", n - 1)
			var ok bool
			com, ok = readInt(0, n - 1)
			if !ok {
				fmt.Printf("\nPlease enter a valid number(0-%d)!\n", n - 1)
				continue
			}
			break
		}
		if com == 0 {
			break
		}
		run_command(acc[(com - 1 + n) % n], 0, "0")
	}

}

func main() {

	lib.ConnectDB()
	err := lib.CreateTables()
	if err != nil {
		return
	}
	err = lib.init()
	if err != nil {
		return
	}
	if config["status"] != "active" {
		fmt.Println("System Not Activated.\nStatus: ", config["status"])
		return
	}
	defer lib.db.Close()

	ClearScreen("Welcome to the Library Management System!\n")
	var mode int
	for {
		fmt.Println("Please choose login mode:\n1. student\n2. admin\n3. root\n0. exit")
		var str string
		for {
			fmt.Print("\nPlease select a number(0-3): ")
			var ok bool
			mode, ok = readInt(0, 3)
			if !ok {
				continue
			}
			break
		}
		var exit bool
		switch mode {
		case 0:
			for {
				fmt.Print("Are you sure to exit the system? (y/N)")
				str = readLine()
				switch str {
				case "y": exit = true
				case "N": exit = false
				default: continue
				}
				break
			}
			if exit {
				break
			}
		case 1:
			work_student(&exit)
		case 2:
			work_admin(&exit)
		case 3:
			work_root(&exit)
		}
		ClearScreen("")
		if exit {
			break
		}
	}

	fmt.Println("Goodbye~ Hoping see you next time!")

	return
}
