package method

import (
	"User/userpb"
	"database/sql"
	"log"
)

func StoreNewUser(db *sql.DB, req *userpb.Request) (*userpb.Responce, error) {
	var user userpb.Responce
	query := "INSERT INTO userservice (name, email age) VALUES($1,$2,$3) RETURNING id;"
	row:= db.QueryRow(query, req.Name, req.Email, req.Age)
	if err := row.Scan(&user.Id); err!=nil{
		log.Println("unable to insert user:", err)
		return nil, err
	}

	return &user, nil
}

func GetbyIdUser(db *sql.DB, req *userpb.Responce) (*userpb.User, error){
	var user userpb.User
	query := "SELECT *FROM userservice WHERE id = $1;"

	row := db.QueryRow(query, req.Id)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Age); err!=nil{
		log.Println("unable to get user:", err)
		return nil, err
	}

	return  &user, nil

}

func UpdateUser(db *sql.DB, req *userpb.User) error{
	query := "UPDATE userservice SET name = $1, email = $2, age = $3 WHERE id = %4;"
	_, err := db.Exec(query, req.Name, req.Email, req.Age, req.Id)
	if err!=nil{
		log.Println("unable to update user:", err)
		return err
	}

	return  nil
}


func DeleteUser(db *sql.DB, req *userpb.Responce) error{
	query := "DELETE FROM userservice WHERE id = $1;"
	_, err := db.Exec(query, req.Id)
	if err!=nil{
		log.Println("unable to delete user:", err)
		return err
	}

	return nil
}