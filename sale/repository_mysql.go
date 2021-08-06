package sale

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/keriwisnu/dblocking/config"
	"github.com/keriwisnu/dblocking/models"
	"log"
)

const (
	table          = "flashsale"
	id			   = "1"
	newstock	   = 1
)

//Db Row Locking
func Update(ctx context.Context, fsale models.Sales) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	//rowlocking
	rows, err := db.Query("SELECT stock FROM flashsale FOR UDPATE")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var resultStock []models.Sales

	for rows.Next() {
		var each = models.Sales{}
		var err = rows.Scan(&each.Stock)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		resultStock = append(resultStock, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(resultStock)
	fmt.Sprintf("%v", resultStock)

	var x = resultStock - newstock

	//update stock
	queryUpdate:= fmt.Sprintf("UPDATE %v set stock = %d where id = '%d'",
		table,
		x,
		id,
	)
	fmt.Println(queryUpdate)

	_, err = db.ExecContext(ctx, queryUpdate)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

