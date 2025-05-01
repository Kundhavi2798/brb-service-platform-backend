package auth

import (
	"brb-service-platform-backend/pkg/db"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

// TestMain to setup the database before running tests
func TestMain(m *testing.M) {
	// Setup the database once before all tests
	setupTestDB()

	// Run the tests
	m.Run()
}

// Setup the in-memory test database
func setupTestDB() {
	var err error
	// Create an in-memory SQLite DB
	testDB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the test database")
	}

	// Assign it to db.DB for global use in the application
	db.DB = testDB

	// Migrate the schema
	err = db.DB.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate the database")
	}
}

// Test function to create a user
func TestCreateUser(t *testing.T) {
	// No need to call setupTestDB here; it's already done in TestMain
	user := &User{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "hashed-password",
	}

	// Assume CreateUser uses db.DB and saves the user
	err := CreateUser(user)
	assert.NoError(t, err)

	var found User
	err = db.DB.First(&found, user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "alice@example.com", found.Email)
}

// Test function to get a user by email
func TestGetUserByEmail(t *testing.T) {
	// No need to call setupTestDB here
	user := &User{
		Name:     "Bob",
		Email:    "bob@example.com",
		Password: "secure-pass",
	}
	db.DB.Create(user)

	result, err := GetUserByEmail("bob@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "Bob", result.Name)
	assert.Equal(t, "bob@example.com", result.Email)

	// Test non-existent user
	_, err = GetUserByEmail("notfound@example.com")
	assert.Error(t, err)
}

// Test password hashing and checking
func TestHashPasswordAndCheck(t *testing.T) {
	password := "mysecurepassword"

	hashed, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	match := CheckPasswordHash(password, hashed)
	assert.True(t, match)

	// Negative test
	wrong := CheckPasswordHash("wrongpassword", hashed)
	assert.False(t, wrong)
}

// Test JWT generation
func TestGenerateJWT(t *testing.T) {
	tokenStr, err := GenerateJWT(1, "admin")
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, float64(1), claims["user_id"])
	assert.Equal(t, "admin", claims["role"])

	// Check expiration claim is within 24 hours
	expUnix := int64(claims["exp"].(float64))
	assert.Greater(t, expUnix, time.Now().Unix())
	assert.Less(t, expUnix, time.Now().Add(25*time.Hour).Unix())
}
