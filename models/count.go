package models

type Count struct {
	Lines        int
	Words        int
	Vowels       int
	Punctuations int
	Spaces       int
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
