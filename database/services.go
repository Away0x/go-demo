package main

import (
	"database/sql"
	"log"
)

func getOne(id int) (a app, err error) {
	err = db.QueryRow(
		"SELECT Id, Name, Status, Level, [Order] FROM dbo.App WHERE Id=@Id",
		sql.Named("Id", id),
	).Scan(&a.ID, &a.name, &a.status, &a.level, &a.order)
	return
}

func getMany(id int) (apps []app, err error) {
	rows, err := db.Query(
		"SELECT Id, Name, Status, Level, [Order] FROM dbo.App WHERE Id>@Id",
		sql.Named("Id", id),
	)

	for rows.Next() {
		a := app{}
		err = rows.Scan(&a.ID, &a.name, &a.status, &a.level, &a.order)
		if err != nil {
			log.Fatalln(err.Error())
		}
		apps = append(apps, a)
	}

	return
}

func (a *app) Update() (err error) {
	_, err = db.Exec(
		"UPDATE dbo.App SET Name=@Name, [Order]=@Order WHERE Id=@Id",
		sql.Named("Name", a.name),
		sql.Named("Order", a.order),
		sql.Named("Id", a.ID),
	)

	return
}

func (a *app) Delete() (err error) {
	_, err = db.Exec(
		"DELETE FROM dbo.App WHERE Id=@Id",
		sql.Named("Id", a.ID),
	)
	return
}

func (a *app) Insert() (err error) {
	sqlStr := `INSERT INTO dbo.App
(NAME,
STATUS,
Level,
[Order])
VALUES
(@Name,
@Status,
@Level,
@Order);
SELECT isNull(SCOPE_IDENTITY(), -1);` // 插入数据并获取 id

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		sql.Named("Name", a.name),
		sql.Named("Status", a.status),
		sql.Named("Level", a.level),
		sql.Named("Order", a.order),
	).Scan(&a.ID)

	if err != nil {
		return
	}

	return
}
