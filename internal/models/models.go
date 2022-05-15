package models

import "github.com/jackc/pgtype"

type Test struct {
	TestStr string
}

type Project struct {
	ID          int
	Admin       User
	Start_date  string
	End_date    string
	Institution Institution
	Created_at  string
	Updated_at  string
	Category    Category
	Name        string
	Description string
}

type User struct {
	ID            int
	Name          pgtype.Varchar
	Email         pgtype.Varchar
	Password      pgtype.Varchar
	Created_at    pgtype.Timestamptz
	Updated_at    pgtype.Timestamptz
	Profile_image pgtype.Varchar
}

type Category struct {
	ID   int
	Name string
}

type Institution struct {
	ID         int
	Name       string
	Email      string
	Website    string
	Owner      User
	Created_at string
	Updated_at string
}
