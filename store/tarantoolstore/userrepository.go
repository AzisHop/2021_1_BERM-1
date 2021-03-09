package tarantoolstore

import (
	"errors"
	"fl_ru/model"
	"github.com/tarantool/go-tarantool"
)

type UserRepository struct {
	store *Store
}
func (u *UserRepository)Create(user *model.User) error{
	resp, err := u.store.conn.Insert("user", userToTarantoolData(user))
	*user = *tarantoolDataToUser(resp.Tuples()[0])
	return err
}

func (r *UserRepository)FindByEmail(user *model.User) error{
	resp, err := r.store.conn.Select("user", "email_key",
		0, 1,  tarantool.IterEq, []interface{}{
		user.Email,
		})
	if err != nil{
		return err
	}
	if len(resp.Tuples()) == 0{
		return errors.New("Bad password")
	}
	*user = *tarantoolDataToUser(resp.Tuples()[0])
	return nil
}

//{name = 'id', type = 'unsigned'},
//{name = 'email', type = 'string'},
//{name = 'password', type = 'string'},
//{name = 'user_name', type = 'string'},
//{name = 'first_name', type = 'string'},
//{name = 'second_name', type = 'string'},
//{name = 'executor', type = 'boolean'},
//{name = 'description', type = 'string', is_nullable=true},
//{name = 'specializes', type = 'array', is_nullable=true},
//{name = 'img_url', type = 'string', is_nullable = true},

func userToTarantoolData(user *model.User) []interface{}{
	data := []interface{}{nil}
	if user.Email == ""{
		data = append(data, nil)
	} else{
		data = append(data, user.Email)
	}
	if len(user.Password) == 0{
		data = append(data, nil)
	} else{
		data = append(data, user.Password )
	}
	if user.UserName == ""{
		data = append(data, nil)
	} else{
		data = append(data, user.UserName)
	}
	if len(user.FirstName) == 0{
		data = append(data, nil)
	} else{
		data = append(data, user.FirstName )
	}
	if len(user.SecondName) == 0{
		data = append(data, nil)
	} else{
		data = append(data, user.SecondName )
	}
	data = append(data, user.Executor)
	if user.Description == ""{
		data = append(data, nil)
	} else{
		data = append(data, user.Description)
	}
	if len(user.Specializes) == 0{
		data = append(data, nil)
	} else{
		data = append(data, user.Specializes)
	}
	if user.ImgUrl == ""{
		data = append(data, nil)
	} else{
		data = append(data, user.ImgUrl)
	}
	return data
}



func tarantoolDataToUser(data []interface{}) *model.User{
	u := &model.User{}
	u.Id, _ =          data[0].(uint64)
	u.Email, _ =       data[1].(string)
	u.Password, _ =    data[2].(string)
	u.UserName, _ =    data[3].(string)
	u.FirstName, _ =   data[4].(string)
	u.SecondName, _ =  data[5].(string)
	u.Executor, _ =    data[6].(bool)
	u.Description, _ = data[7].(string)
	u.Specializes, _ = data[8].([]string)
	u.ImgUrl, _ =      data[9].(string)
	return u
}
