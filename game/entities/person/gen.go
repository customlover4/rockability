package person

import "math/rand/v2"

var MaleFirstNames = []string{
	"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph",
	"Thomas", "Charles", "Christopher", "Daniel", "Matthew", "Anthony", "Mark", "Paul",
	"Steven", "Andrew", "Kenneth", "Joshua", "Kevin", "Brian", "George", "Edward",
	"Ronald", "Timothy", "Jason", "Jeffrey", "Ryan", "Jacob", "Gary", "Nicholas",
	"Eric", "Jonathan", "Stephen", "Larry", "Justin", "Scott", "Brandon", "Benjamin",
	"Samuel", "Gregory", "Frank", "Alexander", "Raymond", "Patrick", "Jack", "Dennis",
	"Jerry", "Tyler", "Aaron", "Adam", "Nathan", "Henry", "Zachary", "Douglas",
	"Peter", "Kyle", "Ethan", "Arthur", "Carl", "Austin", "Joe", "Jesse", "Albert",
	"Bryan", "Billy", "Bruce", "Christian", "Randy", "Willie", "Lawrence", "Gabriel",
	"Wayne", "Louis", "Jeremy", "Craig", "Keith", "Sean", "Philip", "Shane", "Philip",
}

var FemaleFirstNames = []string{
	"Mary", "Patricia", "Jennifer", "Linda", "Barbara", "Elizabeth", "Susan", "Jessica",
	"Sarah", "Karen", "Lisa", "Nancy", "Betty", "Margaret", "Sandra", "Ashley", "Dorothy",
	"Kimberly", "Emily", "Donna", "Michelle", "Carol", "Amanda", "Melissa", "Deborah",
	"Stephanie", "Rebecca", "Laura", "Sharon", "Cynthia", "Kathleen", "Amy", "Angela",
	"Shirley", "Anna", "Brenda", "Pamela", "Nicole", "Emma", "Samantha", "Katherine",
	"Christine", "Debra", "Rachel", "Carolyn", "Janet", "Catherine", "Maria", "Heather",
	"Diane", "Ruth", "Olivia", "Julia", "Grace", "Sophia", "Hannah", "Megan", "Victoria",
	"Evelyn", "Lauren", "Chloe", "Natalie", "Lily", "Eleanor", "Stella", "Aria", "Zoe",
	"Mila", "Aubrey", "Clara", "Violet", "Lucy", "Audrey", "Maya", "Alice", "Isla", "Eliza",
}

var Surnames = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis",
	"Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson",
	"Thomas", "Taylor", "Moore", "Jackson", "Martin", "Lee", "Perez", "Thompson",
	"White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson", "Walker",
	"Young", "Allen", "King", "Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores",
	"Green", "Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell",
	"Carter", "Roberts", "Evans", "Turner", "Phillips", "Parker", "Edwards", "Collins",
	"Stewart", "Morris", "Murphy", "Cook", "Rogers", "Morgan", "Cooper", "Peterson",
	"Bailey", "Reed", "Kelly", "Howard", "Cox", "Ward", "Richardson", "Wood", "Watson",
	"Brooks", "Bennett", "Gray", "James", "Hughes", "Price", "Myers", "Long", "Ross",
	"Foster", "Sanders", "Powell", "Jenkins", "Perry", "Butler", "Barnes", "Fisher",
}

func MaleName(rnd *rand.Rand) string {
	name := MaleFirstNames[rnd.IntN(len(MaleFirstNames))]
	surname := Surnames[rnd.IntN(len(Surnames))]
	return name + " " + surname
}

func FemaleName(rnd *rand.Rand) string {
	name := FemaleFirstNames[rnd.IntN(len(FemaleFirstNames))]
	surname := Surnames[rnd.IntN(len(Surnames))]
	return name + " " + surname
}
