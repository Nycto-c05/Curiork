package main

import (
	"errors"
	"fmt"
)

//
// ========== MODELS ============
//

type User struct {
	ID   int
	Name string
}

//
// ========== REPOSITORY LAYER ============
//

// Interface
type UserRepo interface {
	Save(u User) error
	Get(id int) (User, error)
}

// In-memory implementation (for example)
type MemoryUserRepo struct {
	data map[int]User
}

func NewMemoryUserRepo() *MemoryUserRepo {
	return &MemoryUserRepo{data: make(map[int]User)}
}

func (r *MemoryUserRepo) Save(u User) error {
	r.data[u.ID] = u
	return nil
}

func (r *MemoryUserRepo) Get(id int) (User, error) {
	u, ok := r.data[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	return u, nil
}

//
// ========== SERVICE LAYER ============
//

// Service depends on interface — NOT actual DB!
type UserService struct {
	repo UserRepo
}

func NewUserService(r UserRepo) *UserService {
	return &UserService{repo: r}
}

// Business logic here
func (s *UserService) RegisterUser(id int, name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	return s.repo.Save(User{ID: id, Name: name})
}

func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.repo.Get(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

//
// ========== MAIN APP ============
//

func main() {
	var repo UserRepo = NewMemoryUserRepo() // could be postgres repo
	service := NewUserService(repo)         // inject dependency

	service.RegisterUser(1, "Bob")

	name, _ := service.GetUserName(1)
	fmt.Println(name) // -> Bob
}
