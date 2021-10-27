package administration

type User struct {
	Id       int
	Nickname string
	Slug     string
	password string
	AboutMe  string
}

var TestUsers = [...]User{
	{0, "Dolbaeb", "dolbaeb", "1", "Ya dolbaeb"},
	{1, "Svarshik", "svarshik", "1", "Ya tozhe dolbaeb"},
	{2, "Yurin Roman", "yurin_roman", "1", "Lublu parenia i massage"},
	{3, "Billy Harington", "billy_harington", "1", "Fisting is 300 bucks"},
	{4, "Master Van", "master_van", "1", "Boss of this gym"},
}
