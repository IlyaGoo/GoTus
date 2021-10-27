package administration

type User struct {
	Id       int
	Nickname string
	Slug     string
	password string
	AboutMe  string
}

var TestUsers = [...]User{
	User{0, "Dolbaeb", "dolbaeb", "1", "Ya dolbaeb"},
	User{1, "Svarshik", "svarshik", "1", "Ya tozhe dolbaeb"},
	User{2, "Yurin Roman", "yurin_roman", "1", "Lublu parenia i massage"},
	User{3, "Billy Harington", "billy_harington", "1", "Fisting is 300 bucks"},
	User{4, "Master Van", "master_van", "1", "Boss of this gym"},
}
