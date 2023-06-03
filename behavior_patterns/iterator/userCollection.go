package iterator

//具体集合

type UserCollection struct {
	Users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: u.Users,
	}
}
