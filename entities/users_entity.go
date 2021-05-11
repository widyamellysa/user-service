package entities

type User struct {
	ID           string      `json:"_id"`
	Email        string      `json:"email"`
	Password     string      `json:"password"`
	Nama         string      `json:"nama"`
	Phone        string      `json:"phone"`
	JenisKelamin int         `json:"jenis_kelamin"`
	Foto         string      `json:"foto"`
	Alamat       string      `json:"alamat"`
	RegisteredAt int64       `json:"registered_at"`
	IsRegister   int         `json:"is_register"`
	IsAktif      int         `json:"is_aktif"`
	ForgetToken  string      `json:"forget_token"`
	TokenFcm     interface{} `json:"token_fcm"`
	Roles        []int       `json:"roles"`
	Bookmark     []struct {
		IDBerita  string `json:"id_berita"`
		IsAktif   int    `json:"is_aktif"`
		CreatedAt int64  `json:"created_at"`
	} `json:"bookmark"`
}
