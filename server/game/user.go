package game

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID string `json:"id"`
}

type UserFact struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Value       string    `json:"value"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

var ErrCantFindUser = errors.New("can't find user")

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", beego.AppConfig.String("DBConnection"))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	checkDatabase()
}

func Enroll(name string) (map[string]string, error) {
	id := uuid.NewV4().String()
	_, err := getUser(id)
	if err != nil {
		return nil, fmt.Errorf("enroll user error: %v", err)
	}
	factID := uuid.NewV4().String()
	_, err = addUserFact(factID, id, "name", name, time.Now())
	if err != nil {
		return nil, fmt.Errorf("enroll user error: %v", err)
	}
	result := make(map[string]string)
	result["id"] = id
	result["name"] = name
	return result, nil
}

func getUser(id string) (User, error) {
	_, err := db.Query("select id from users where id = ? limit 1", id)
	if err != nil && err == sql.ErrNoRows {
		return addUser(id)
	}
	return User{ID: id}, nil
}

func addUser(id string) (User, error) {
	stmt, err := db.Prepare("insert into users(id) values(?)")
	if err != nil {
		return User{}, fmt.Errorf("add user error: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return User{}, fmt.Errorf("add user error: %v", err)
	}
	return User{
		ID: id,
	}, nil
}

func getUserFacts(userId string) ([]UserFact, error) {
	facts := []UserFact{}
	rows, err := db.Query("select id, name, value, created_time, updated_time from user_facts where user_id = ?", userId)
	if err != nil {
		return nil, fmt.Errorf("get user facts error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		var value string
		var createdTime time.Time
		var updatedTime time.Time

		err = rows.Scan(&id, &name, &value, &createdTime, &updatedTime)
		if err != nil {
			return nil, fmt.Errorf("get user facts error: %v", err)
		}

		facts = append(facts, UserFact{
			ID:          id,
			UserID:      userId,
			Name:        name,
			Value:       value,
			CreatedTime: createdTime,
			UpdatedTime: updatedTime,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("get user facts error: %v", err)
	}
	return facts, nil
}

func addUserFact(id, userID, name, value string, createdTime time.Time) (UserFact, error) {
	stmt, err := db.Prepare("insert into user_facts(id, user_id, name, value, created_time) values(?, ?, ?, ?, ?)")
	if err != nil {
		return UserFact{}, fmt.Errorf("add user facts error: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, userID, name, value, createdTime)
	if err != nil {
		return UserFact{}, fmt.Errorf("add user facts error: %v", err)
	}
	return UserFact{
		ID:          id,
		UserID:      userID,
		Name:        name,
		Value:       value,
		CreatedTime: createdTime,
	}, nil
}

func updateUserFact(id, value string) error {
	stmt, err := db.Prepare("update user_facts set value = ?, updated_time = ? where id = ?")
	if err != nil {
		return fmt.Errorf("update user[%v] fact error: %v", id, err)
	}
	_, err = stmt.Exec(value, time.Now(), id)
	if err != nil {
		return fmt.Errorf("update user[%v] fact error: %v", id, err)
	}
	return nil
}

func checkDatabase() {
	_, err := db.Query("select 1 from users limit 1;")
	if err != nil {
		beego.Warn("The database has no users and user_facts tables, now is going to create them.")
		initDatabase()
	}
}
func initDatabase() {
	script, err := ioutil.ReadFile(beego.AppConfig.String("InitScriptPath"))
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		panic(err)
	}
}
