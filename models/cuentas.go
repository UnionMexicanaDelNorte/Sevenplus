package models

var (
	CuentasList map[string]*Cuenta
)

func init() {
	CuentasList = make(map[string]*Cuenta)
}

type Cuenta struct {
	ACNT_CODE string
	DESCR string
}


func AddCuenta(u Cuenta, index string)  {
	CuentasList[index] = &u	
}
/*
func GetUser(uid string) (u *User, err error) {
	if u, ok := CuentasList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}
*/
func GetAllCuentas() map[string]*Cuenta {
	return CuentasList
}
/*
func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := CuentasList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range CuentasList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(CuentasList, uid)
}
*/