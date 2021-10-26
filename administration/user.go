package administration

type User struct {
	Id       int
	Nickname string
	password string
	AboutMe  string
}

var TestUsers = [...]User{
	User{0, "Dolbaeb", "1", "Ya dolbaeb"},
	User{1, "Svarshik", "1", "Ya tozhe dolbaeb"},
	User{2, "Yurin Roman", "1", "Lublu parenia i massage"},
	User{3, "Billy Harington", "1", "Fisting is 300 bucks"},
	User{4, "Master Van", "1", "Boss of this gym"},
}
