package main
import ("fmt"
"database/sql"
 "github.com/mattn/go-sqlite3"
)

func addIngredient(){
	var name string
	var weight float64

	fmt.Print("Introduzca nombre: \n")
	fmt.Scan(&name)
	fmt.Print("Introduzca peso: \n")
	fmt.Scan(&weight)
	//add to database
}


func addRecipe(){
	var name string
	var weight float64

	fmt.Print("Introduzca nombre: \n")
	fmt.Scan(&name)
	fmt.Print("Introduzca peso: \n")
	fmt.Scan(&weight)
	//add to database
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
// func calculatePortion(ingredient string, amount float64)

func main(){
	db, err := sql.Open("sqlite3", "./recipes.db")

	checkErr(err)
	fmt.Print(db)

	statement, err := db.Prepare("INSERT INTO ingredients(name, amount) values(?,?)")
	res,err:=statement.Exec("test",1)
	id,err := res.LastInsertId()
	fmt.Print(id)

	var option int
	fmt.Print("1) Agregar ingrediente: \n")
	fmt.Print("2) Agregar receta: \n")
	fmt.Print("3) Ver: \n")

	fmt.Print("Escoja una opcion: \n")
	fmt.Scan(&option)
	fmt.Print(option)

	// db.Close()
}
