package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	Name            string
	Age             int
	FavouriteNumber int
	Chance          int
}

type Game struct {
	Number int
	Player Player
}

func main() {

	f := NewPlayer()
	newuse := Player{}
	newuse = newuse.Newplayers(f)
	game := Game{}
	game = game.NewGame(newuse)

	game.StartGame()

}

func (g Game) NewGame(p Player) Game {
	perdict := p.Chance

	if perdict == 0 {
		perdict = 1
	}
	pre := perdict * perdict

	return Game{
		Number: rand.Intn(pre),
		Player: p,
	}

}

func (c Player) Newplayers(p Player) Player {

	return Player{
		Name:            p.Name,
		Chance:          p.Chance,
		Age:             p.Age,
		FavouriteNumber: p.FavouriteNumber,
	}
}
func NewPlayer() Player {
	var playerName string
	var age, favouriteNum, chance int

	fmt.Println("Game is starting!")
	fmt.Print("Please enter your name: ")
	_, err := fmt.Scan(&playerName)
	if err != nil {
		fmt.Println("Error reading name input.")
		return Player{}
	}

	fmt.Print("Please enter your age: ")
	_, err = fmt.Scan(&age)

	if err != nil {
		fmt.Println("Invalid input! Please enter a number.")
		//return Player{}
	}

	if age <= 18 {
		fmt.Println("You must be at least 18 years old to play.")
		return Player{}
	}

	fmt.Print("Please enter your favourite number: ")
	_, err = fmt.Scan(&favouriteNum)
	if err != nil {
		fmt.Println("Error reading favourite number.")
		return Player{}
	}

	fmt.Print("Enter the number of chances you want but .")
	fmt.Println("Note: Your predicted number range will be squared (e.g., 4 → 16, 5 → 25).")
	fmt.Print("Please enter a number: ")

	_, err = fmt.Scan(&chance)

	if err != nil || chance <= 0 {
		fmt.Println("Invalid number of chances.")
		return Player{}
	}

	return Player{
		Name:            playerName,
		Age:             age,
		FavouriteNumber: favouriteNum,
		Chance:          chance,
	}
}

func (g Game) StartGame() {
	var numb int
	chance := g.Player.Chance
	if g.Number%2 != 0 {
		fmt.Println("predict numbe is  odd number ")
	}
	for i := 1; i <= int(chance); i++ {
		fmt.Print("Please enter your guees number: ")
		fmt.Scan(&numb)
		if g.Number == numb {
			fmt.Println("you win, Congratulation ", g.Player.Name)
			if g.Player.FavouriteNumber == numb {
				fmt.Println("you win, Congratulation and also your favourite number is eqaual to predict number bravo", g.Player.Name)
			}
			return
		} else if i == int(chance) {
			fmt.Printf("you lost this game :( predict number was %v \n", g.Number)
		} else {
			g.Player.Chance--
			fmt.Printf("you did not find, you have %v chances \n", g.Player.Chance)
		}

	}
}
