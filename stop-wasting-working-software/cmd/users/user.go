package users

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joshbrgs/mongorm/cmd/mongorm"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	mongorm.Model

	ID         primitive.ObjectID `bson:"_id,omitempty"`
	ScreenName string             `bson:"screen_name"`
	Username   string             `bson:"username"`
	Password   string             `bson:"password"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

type UserService struct {
	db *mongo.Database
}

var (
	secretKey         = []byte("your_secret_key") // Replace with a secure secret key
	jwtExpiryDuration = 24 * time.Hour
)

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	userModel := &user                                // Cast User to a pointer for mongorm.Model
	return user.Create(ctx, s.db, "users", userModel) // Use mongorm.Create
}

func (s *UserService) FetchUserByID(ctx context.Context, userID primitive.ObjectID) (User, error) {
	var user User
	filter := bson.M{"_id": userID}                     // Create filter for fetching by ID
	err := user.Read(ctx, s.db, "users", filter, &user) // Use mongorm.Read

	return user, err
}

func (s *UserService) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	var user User
	filter := bson.M{"_id": userID}                // Create filter for deletion
	return user.Delete(ctx, s.db, "users", filter) // Use mongorm.Delete
}

func (s *UserService) UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser User) error {
	var user User

	updatedUser.UpdatedAt = time.Now()    // Update timestamp
	filter := bson.M{"_id": userID}       // Create filter for update
	update := bson.M{"$set": updatedUser} // Update data

	return user.Update(ctx, s.db, "users", filter, update) // Use mongorm.Update
}

func (s *UserService) authenticateUser(username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	if err := user.Read(ctx, s.db, "users", bson.M{"username": username}, &user); err != nil {
		return "", fmt.Errorf("user not found")
	}
	// err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.Hex(),
		"exp": time.Now().Add(jwtExpiryDuration).Unix(),
	})

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
