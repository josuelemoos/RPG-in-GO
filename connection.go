package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

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

func main() {
	
	db, err := sql.Open("sqlite3", "Data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	
	createTables(db)

	
	insertUsuarios(db)
	insertBosses(db)

	
	monster, err := spawnMonster(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Monstro spawnou: %+v\n", monster)
}

func createTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS usuarios (
			id INTEGER PRIMARY KEY,
			name TEXT,
			class TEXT,
			Strong INT,
			Intelligence INT,
			Health INT,
			Furtive INT,
			Charisma INT
		);`,
		`CREATE TABLE IF NOT EXISTS boss (
			id INTEGER PRIMARY KEY,
			name TEXT,
			class TEXT,
			Strong INT,
			Intelligence INT,
			Health INT,
			Furtive INT,
			Charisma INT
		);`,
		`CREATE TABLE IF NOT EXISTS player (
			id INTEGER PRIMARY KEY,
			name TEXT,
			class TEXT,
			Strong INT,
			Intelligence INT,
			Health INT,
			Furtive INT,
			Charisma INT
		);`,
		`CREATE TABLE IF NOT EXISTS attack (
			id INTEGER PRIMARY KEY,
			name TEXT,
			class TEXT,
			Strong INT,
			Intelligence INT,
			Furtive INT,
			Charisma INT,
			Damage INT
		);`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			log.Fatal(err)
		}
	}
}

func insertUsuarios(db *sql.DB) {
	usuarios := []Character{
		{Name: "Goblin Scout", Class: "Rogue", Strong: 5, Intelligence: 4, Health: 14, Furtive: 10, Charisma: 3},
		{Name: "Orc Warrior", Class: "Warrior", Strong: 11, Intelligence: 3, Health: 22, Furtive: 4, Charisma: 2},
		{Name: "Skeleton Soldier", Class: "Undead", Strong: 7, Intelligence: 2, Health: 18, Furtive: 3, Charisma: 1},
		{Name: "Shadow Wolf", Class: "Beast", Strong: 8, Intelligence: 3, Health: 16, Furtive: 11, Charisma: 2},
		{Name: "Dark Apprentice", Class: "Mage", Strong: 3, Intelligence: 10, Health: 12, Furtive: 5, Charisma: 6},
		{Name: "Cave Troll", Class: "Giant", Strong: 12, Intelligence: 2, Health: 35, Furtive: 2, Charisma: 1},
		{Name: "Giant Spider", Class: "Beast", Strong: 7, Intelligence: 2, Health: 15, Furtive: 12, Charisma: 1},
		{Name: "Ghost Knight", Class: "Undead", Strong: 9, Intelligence: 7, Health: 24, Furtive: 4, Charisma: 8},
		{Name: "Green Slime", Class: "Elemental", Strong: 2, Intelligence: 1, Health: 28, Furtive: 1, Charisma: 1},
		{Name: "Young Wyvern", Class: "Dragon", Strong: 10, Intelligence: 6, Health: 30, Furtive: 7, Charisma: 5},
		{Name: "Bandit Raider", Class: "Rogue", Strong: 6, Intelligence: 5, Health: 17, Furtive: 8, Charisma: 4},
		{Name: "Forest Spirit", Class: "Elemental", Strong: 4, Intelligence: 9, Health: 20, Furtive: 9, Charisma: 10},
	}

	insertCharacters(db, "usuarios", usuarios)
}

func insertBosses(db *sql.DB) {
	bosses := []Character{
		{Name: "Souless: the Lost Warrior", Class: "Human", Strong: 5, Intelligence: 9, Health: 100, Furtive: 9, Charisma: 10},
		{Name: "Atronach", Class: "Elemental", Strong: 12, Intelligence: 10, Health: 80, Furtive: 3, Charisma: 3},
		{Name: "BigMouth", Class: "Human", Strong: 12, Intelligence: 11, Health: 100, Furtive: 9, Charisma: 10},
	}

	insertCharacters(db, "boss", bosses)
}

func insertCharacters(db *sql.DB, table string, chars []Character) {
	query := fmt.Sprintf(`INSERT INTO %s (name, class, Strong, Intelligence, Health, Furtive, Charisma)
		VALUES (?, ?, ?, ?, ?, ?, ?)`, table)

	for _, c := range chars {
		_, err := db.Exec(query, c.Name, c.Class, c.Strong, c.Intelligence, c.Health, c.Furtive, c.Charisma)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func insertUsuarios(db *sql.DB) {
	usuarios := []Character{
		{Name: "Goblin Scout", Class: "Rogue", Strong: 5, Intelligence: 4, Furtive: 10, Charisma: 3, Damage: 1},
		{Name: "Orc Warrior", Class: "Warrior", Strong: 11, Intelligence: 3, Furtive: 4, Charisma: 2, Damage: 1},
		{Name: "Skeleton Soldier", Class: "Undead", Strong: 7, Intelligence: 2, Furtive: 3, Charisma: 1, Damage: 1},
	}

	insertttacks(db, "usuarios", usuarios)
}

func insertAttacks(db *sql.DB, table string, chars []Character) {
	query := fmt.Sprintf(`INSERT INTO %s (name, class, Strong, Intelligence, Furtive, Charisma, Damage)
		VALUES (?, ?, ?, ?, ?, ?, ?)`, table)

	for _, c := range chars {
		_, err := db.Exec(query, c.Name, c.Class, c.Strong, c.Intelligence, c.Furtive, c.Charisma, c.Damage)
		if err != nil {
			log.Fatal(err)
		}
	}
}


func spawnMonster(db *sql.DB) (*Character, error) {
	row := db.QueryRow("SELECT * FROM usuarios ORDER BY RANDOM() LIMIT 1")

	var c Character
	err := row.Scan(&c.ID, &c.Name, &c.Class, &c.Strong, &c.Intelligence, &c.Health, &c.Furtive, &c.Charisma)
	if err != nil {
		return nil, err
	}

	return &c, nil
}