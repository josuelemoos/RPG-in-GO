package main

import "fmt"

func more() {
	fmt.Println("i can tell you a tiny secret, and u can tell me one of yours little player...")
	fmt.Println("what is it? - says the player")
	fmt.Println("something from a beyond life of elsewhere, but tell me, are you strong enough to hear?")
	if player.strong < 5 {
		fmt.Println("you are not strong enough!!!... get back were when u becamed capable...")
	}else if player.strong >= 5 {
		fmt.Println("in did, you are a great warrior, so, go ahead, come closer...")
	}
}


func anwsers(input int) {
if input == 1 {
		fmt.Println("SO U WANT a boss... WELL i dont think i've seen one")
	} else if input == 2{
		fmt.Println("SO U WANT more, good choice, what else do u wanna know?")

		more();
	}else if input == 3{
		fmt.Println("no.")

	}
}
func creatingCharacter() *Character{
	
	p := &Character{}
	var point int = 25
	var name string
	var value int = 5
	p.health = 80

	fmt.Println("WHO ARE YOU?")
	fmt.Scan(&name)
	p.name = name
	fmt.Println("well ", name, " what are your atributes?")

	for value < 11 && value > 0{
			for {
				fmt.Println("strong? (MAX: 10. you have", point, "points left)")
				fmt.Scan(&value)
				if value >= 1 && value <= 10 { break }
				fmt.Println("i dont believe that! try again.")
			}
		p.strong = value
		point = point - value
			
			for{
				fmt.Println("intelligence? (MAX: 10. you have", point, "points left)")
				fmt.Scan(&value)
				if value >= 1 &&  value <= 10 { break }
				fmt.Println("i dont believe that! try again.")	
			}
		p.intelligence = value
		point = point - value	
			

			for{
				fmt.Println("charisma? (MAX: 10. you have", point, "points left)")
				fmt.Scan(&value)
				if value >= 1 && value <= 10 { break }
				fmt.Println("i dont believe that! try again.")
			}
		p.charisma = value
		point = point - value
			
			for{
				fmt.Println("furtive? (MAX: 10. you have", point, "points left)")
				fmt.Scan(&value)
				if value >= 1 && value <= 10 { break }
				fmt.Println("i dont believe that! try again.")
			}
		p.furtive = value
		point = point - value
		break	
	}
	return p
}
func beginning(){
	fmt.Println("the sun is setting. the world? well... is closing his doors, preparing to sleep")
	fmt.Println("the day is getting darker, the cold is getting bigger than the orange\nof and sleepy sun that is goig to bed")
	fmt.Println("who are you?... u just arrived here")

	player = creatingCharacter()

	fmt.Println("well well, u are", player.name, " i can tell")
	fmt.Println("what are u hopping to find here? the day is getting darker...")
	fmt.Println("but i see. u are a traveler, i can see in your eyes, u want to go to\nNIRVYNYA dont you?")
	fmt.Println("u want to see the beaches, the dogs running, the sound of the water waves hitting the rocks")
	fmt.Println("but is impossible right now... The Big Bull of time is on the path")
	fmt.Println("U NEED TO DEFEAT HIM!!!!!!!!!!!!!!!!")
}

func theMiddle(p* Character){
	fmt
}

func main(){
	fmt.Println("type a choice")
	fmt.Println("1 - say boss")
	fmt.Println("2 - say more")
	fmt.Println("3 - say less")

	var input int 
	fmt.Scan(&input)

	anwsers(input)
	
}