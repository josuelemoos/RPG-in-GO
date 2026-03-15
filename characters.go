package GO

type Character struct{
	Name string
	strong int
	intelligence int
	furtive int 
	charisma int 
	health int
}

func state(p *Personagen) Class() string{
	if p.strong > 5 {
		return " powerfull! "
	}else if p.inteligence > 6 {
		return " wise!"
	}else if p.charisma > 6{
		return " charming"
	}else if p.health>30{
		return " weak..."
	}else {
		return " normal..."
	}
}