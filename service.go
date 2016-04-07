package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"gitlab.com/buzz/user/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	// SecretKey is the key to hash the JWT token
	SecretKey = "hello,ladies,lolz"
)

// UserService is an interface for controlling users
type UserService interface {
	Create(newUser *model.CreateUser) (*model.User, error)
	GetAll() ([]model.User, error)
	GetByID(id string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	Login(username, password string) (model.JWTToken, error)
	Remove(id string) error
	Update(updatedUser *model.UpdateUser) (*model.User, error)
}

type userService struct{}

func (userService) Create(newUser *model.CreateUser) (*model.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:        bson.NewObjectId().Hex(),
		Email:     newUser.Email,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Password:  string(hashedPassword),
		Role:      newUser.Role,
		Username:  newUser.Username,
	}

	//Grab a copy of our session
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	//Get our collection of applications
	db := session.DB("buzz-test-user")
	collection := db.C("users")

	// Create a unique index for email and username
	index := mgo.Index{
		Key:    []string{"username", "email"},
		Unique: true,
	}
	err = collection.EnsureIndex(index)
	if err != nil {
		return nil, err
	}

	//Insert our application
	err = collection.Insert(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userService) GetAll() ([]model.User, error) {
	//Grab a copy of our session
	session, err := getSession()
	if err != nil {
		return []model.User{}, err
	}
	defer session.Close()

	//Get our collection of applications
	db := session.DB("buzz-test-user")
	collection := db.C("users")

	//Get our applications from the collection
	var retrievedUsers []model.User
	err = collection.Find(bson.M{}).All(&retrievedUsers)
	if err != nil {
		return []model.User{}, err
	}

	return retrievedUsers, nil
}

func (userService) GetByID(id string) (*model.User, error) {
	//Grab a copy of our session
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	//Get our collection of applications
	db := session.DB("buzz-test-user")
	collection := db.C("users")

	//Get our applications from the collection
	var retrievedUser *model.User
	err = collection.Find(bson.M{"_id": id}).One(&retrievedUser)
	if err != nil {
		return nil, err
	}

	return retrievedUser, nil
}

func (userService) GetByUsername(username string) (*model.User, error) {
	//Grab a copy of our session
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	//Get our collection of applications
	db := session.DB("buzz-test-user")
	collection := db.C("users")

	//Get our applications from the collection
	var retrievedUser *model.User
	err = collection.Find(bson.M{"username": username}).One(&retrievedUser)
	if err != nil {
		return nil, err
	}

	return retrievedUser, nil
}

func (u userService) Login(username, password string) (model.JWTToken, error) {
	// try to retrive the user by the username
	user, err := u.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// compare the passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate the JWT token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["userid"] = user.ID
	// Expire in 5 mins
	token.Claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return model.JWTToken(tokenString), nil
}

func (userService) Remove(id string) error {
	//Grab a copy of our session
	session, err := getSession()
	if err != nil {
		return err
	}
	defer session.Close()

	//Get our collection of applications
	db := session.DB("buzz-test-user")
	collection := db.C("users")

	//remove our applications from the collection
	err = collection.Remove(bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (userService) Update(updatedUser *model.UpdateUser) (*model.User, error) {
	user := &model.User{
		ID:        updatedUser.ID,
		Email:     updatedUser.Email,
		FirstName: updatedUser.FirstName,
		LastName:  updatedUser.LastName,
		Password:  updatedUser.Password,
		Role:      updatedUser.Role,
		Username:  updatedUser.Username,
	}

	//Grab a copy of our session
	session, err := getSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	//Get our collection of applications
	db := session.DB("buzz-test-user")
	collection := db.C("users")

	// Create a unique index for email and username
	index := mgo.Index{
		Key:    []string{"username", "email"},
		Unique: true,
	}
	err = collection.EnsureIndex(index)
	if err != nil {
		return nil, err
	}

	//Insert our application
	err = collection.Update(bson.M{"_id": updatedUser.ID}, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

var globalSession *mgo.Session

func getSession() (*mgo.Session, error) {
	//Establish our database connection
	if globalSession == nil {
		var err error
		globalSession, err = mgo.Dial(":27017")
		if err != nil {
			return nil, err
		}

		//Optional. Switch the session to a monotonic behavior.
		globalSession.SetMode(mgo.Monotonic, true)
	}

	return globalSession.Copy(), nil
}
