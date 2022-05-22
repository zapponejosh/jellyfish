package models

import "github.com/jackc/pgtype"

type Test struct {
	TestStr pgtype.Varchar
}

type Project struct {
	ID          int
	Admin       int
	Start_date  pgtype.Date
	End_date    pgtype.Date
	Institution int
	Created_at  pgtype.Timestamptz
	Updated_at  pgtype.Timestamptz
	Category    int
	Name        pgtype.Varchar
	Description pgtype.Varchar
}
type GetProject struct {
	ID            int
	AdminID       int
	Admin         pgtype.Varchar
	Start_date    pgtype.Date
	End_date      pgtype.Date
	InstitutionID int
	Institution   pgtype.Varchar
	Created_at    pgtype.Timestamptz
	Updated_at    pgtype.Timestamptz
	CategoryID    int
	Category      pgtype.Varchar
	Name          pgtype.Varchar
	Description   pgtype.Varchar
}

type ProjectPreview struct {
	Name          pgtype.Varchar
	ID            int
	Admin         pgtype.Varchar
	AdminID       int
	Institution   pgtype.Varchar
	InstitutionID int
	Category      pgtype.Varchar
	CategoryID    int
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
	Name pgtype.Varchar
}

type Institution struct {
	ID         int
	Name       pgtype.Varchar
	Email      pgtype.Varchar
	Website    pgtype.Varchar
	Owner      int
	Created_at pgtype.Timestamptz
	Updated_at pgtype.Timestamptz
}
