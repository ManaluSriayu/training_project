package talk_training

import "fmt"

func init() {
	connection()
}

func GetTalks() []Talks {
	rows, err := MainDB.Query("SELECT * FROM ws_talk limit 10")
	checkErr(err)
	var talk_id int64
	var product_id int64
	var message string

	var list = make([]Talks, 10)
	var i = 0
	for rows.Next() {
		err = rows.Scan(&talk_id, &product_id, &message)

		talks := Talks{
			ID:        talk_id,
			ProductID: product_id,
			Message:   message,
			//CreateTime: time.Now(),
		}
		list[i] = talks
		i++
	}
	return list
}

func GetTalksParameter(productID int64) []Talks {
	rows, err := MainDB.Query("SELECT * FROM ws_talk limit 10 where product_id= %d", productID)
	checkErr(err)
	var talk_id int64
	var product_id int64
	var message string

	var list = make([]Talks, 10)
	var i = 0
	for rows.Next() {
		err = rows.Scan(&talk_id, &product_id, &message)

		talks := Talks{
			ID:        talk_id,
			ProductID: product_id,
			Message:   message,
			//CreateTime: time.Now(),
		}
		list[i] = talks
		i++
	}
	return list
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InsertTalks(productID int64, userID int64, message string, createTime string) {

	check, err := MainDB.Query("INSERT INTO public.ws_talk(product_id, user_id, message, create_time) VALUES (%d, %d, %d, %d", productID, userID, message, createTime)
	checkErr(err)
	if check != nil {
		fmt.Println("Sukses Insert")
	} else {
		fmt.Println("Gagal Insert")
	}
}
