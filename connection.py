import sqlite3

connection = sqlite3.connect('Data.db')

cursor = connection.cursor()

cursor.execute("""CREATE TABLE IF NOT EXISTS usuarios (
    id INTEGER PRIMARY KEY,
    name TEXT,
    class TEXT,
    Strong INT,
    Intelligence INT,
    Health INT,
    Furtive INT,
    Charisma INT
);""")

cursor.execute("""CREATE TABLE IF NOT EXISTS boss (
    id INTEGER PRIMARY KEY,
    name TEXT,
    class TEXT,
    Strong INT,
    Intelligence INT,
    Health INT,
    Furtive INT,
    Charisma INT
);""")

cursor.executemany("""
INSERT INTO usuarios (name, class, Strong, Intelligence, Health, Furtive, Charisma)
VALUES (?, ?, ?, ?, ?, ?, ?)
""", [

    ("Goblin Scout", "Rogue", 5, 4, 14, 10, 3),
    ("Orc Warrior", "Warrior", 11, 3, 22, 4, 2),
    ("Skeleton Soldier", "Undead", 7, 2, 18, 3, 1),
    ("Shadow Wolf", "Beast", 8, 3, 16, 11, 2),
    ("Dark Apprentice", "Mage", 3, 10, 12, 5, 6),
    ("Cave Troll", "Giant", 12, 2, 35, 2, 1),
    ("Giant Spider", "Beast", 7, 2, 15, 12, 1),
    ("Ghost Knight", "Undead", 9, 7, 24, 4, 8),
    ("Green Slime", "Elemental", 2, 1, 28, 1, 1),
    ("Young Wyvern", "Dragon", 10, 6, 30, 7, 5),
    ("Bandit Raider", "Rogue", 6, 5, 17, 8, 4),
    ("Forest Spirit", "Elemental", 4, 9, 20, 9, 10)

])

cursor.executemany("""
INSERT INTO boss (name, class, Strong, Intelligence, Health, Furtive, Charisma)
VALUES (?, ?, ?, ?, ?, ?, ?)
""", [

    ("Souless: the Lost Warrior", "Human", 5, 9, 100, 9, 10),
    ("Atronach", "Elemental", 12, 10, 80, 3, 3),
    ("BigMouth", "Human", 12, 11, 100, 9, 10)

])

connection.commit()

def spawn_monster(cursor):
    cursor.execute("SELECT * FROM usuarios ORDER BY RANDOM() LIMIT 1")
    return cursor.fetchone()