package main

import "fmt"

type Character struct {
	ID           int
	Name         string
	Class        string
	Strong       int
	Intelligence int
	Health       int
	Furtive      int
	Charisma     int
}

var player *Character

func more() {
	fmt.Println("i can tell you a tiny secret, and u can tell me one of yours little player...")
	fmt.Println("what is it? - says the player")
	fmt.Println("something from a beyond life of elsewhere, but tell me, are you strong enough to hear?")
	if player.Strong < 5 {
		fmt.Println("you are not strong enough!!!... get back were when u becamed capable...")
	} else {
		fmt.Println("in did, you are a great warrior, so, go ahead, come closer...")
	}
}

func anwsers(input int) {
	if input == 1 {
		fmt.Println("SO U WANT a boss... WELL i dont think i've seen one")
	} else if input == 2 {
		fmt.Println("SO U WANT more, good choice, what else do u wanna know?")
		more()
	} else if input == 3 {
		fmt.Println("no.")
	}
}

func creatingCharacter() *Character {
	p := &Character{}
	point := 25
	p.Health = 80

	var name string
	fmt.Println("WHO ARE YOU?")
	fmt.Scan(&name)
	p.Name = name
	fmt.Println("well ", name, " what are your atributes?")

	for {
		fmt.Println("strong? (MAX: 10. you have", point, "points left)")
		fmt.Scan(&p.Strong)
		if p.Strong >= 1 && p.Strong <= 10 && p.Strong <= point {
			break
		}
		fmt.Println("i dont believe that! try again.")
	}
	point -= p.Strong

	for {
		fmt.Println("intelligence? (MAX: 10. you have", point, "points left)")
		fmt.Scan(&p.Intelligence)
		if p.Intelligence >= 1 && p.Intelligence <= 10 && p.Intelligence <= point {
			break
		}
		fmt.Println("i dont believe that! try again.")
	}
	point -= p.Intelligence

	for {
		fmt.Println("charisma? (MAX: 10. you have", point, "points left)")
		fmt.Scan(&p.Charisma)
		if p.Charisma >= 1 && p.Charisma <= 10 && p.Charisma <= point {
			break
		}
		fmt.Println("i dont believe that! try again.")
	}
	point -= p.Charisma

	for {
		fmt.Println("furtive? (MAX: 10. you have", point, "points left)")
		fmt.Scan(&p.Furtive)
		if p.Furtive >= 1 && p.Furtive <= 10 && p.Furtive <= point {
			break
		}
		fmt.Println("i dont believe that! try again.")
	}

	return p
}

func beginning() {
	fmt.Println("the sun is setting. the world? well... is closing his doors, preparing to sleep")
	fmt.Println("the day is getting darker, the cold is getting bigger than the orange\nof and sleepy sun that is goig to bed")
	fmt.Println("who are you?... u just arrived here")

	player = creatingCharacter()

	fmt.Println("well well, u are", player.Name, " i can tell")
	fmt.Println("what are u hopping to find here? the day is getting darker...")
	fmt.Println("but i see. u are a traveler, i can see in your eyes, u want to go to\nNIRVYNYA dont you?")
	fmt.Println("u want to see the beaches, the dogs running, the sound of the water waves hitting the rocks")
	fmt.Println("but is impossible right now... The Big Bull of time is on the path")
	fmt.Println("U NEED TO DEFEAT HIM!!!!!!!!!!!!!!!!")
}

func theMiddle(p *Character) {
	// to be implemented
}

func main() {

	beginning()
	
	fmt.Println("type a choice")
	fmt.Println("1 - say boss")
	fmt.Println("2 - say more")
	fmt.Println("3 - say less")

	var input int
	fmt.Scan(&input)

	anwsers(input)
}