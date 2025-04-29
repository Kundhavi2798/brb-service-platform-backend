package service_test

import (
	"brb-service-platform-backend/internal/service"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

// TestUpdateServiceAvailability tests the UpdateServiceAvailability method
func TestUpdateServiceAvailability(t *testing.T) {
	// Step 1: Create a mock DB connection using sqlmock
	db, mock, err := sqlmock.New()
	require.NoError(t, err, "failed to open mock database connection")
	defer db.Close()

	// Step 2: Wrap the sqlmock DB into a gorm DB object
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err, "failed to open gorm DB")

	// Step 3: Initialize your service with the mocked DB
	service := &service.Service{DB: gdb}

	// Step 4: Set up expectations on the mock DB (define the query you expect)
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT .+ FROM services WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status"}).
			AddRow(1, "active"))
	mock.ExpectExec("UPDATE services SET availability = ? WHERE id = ?").
		WithArgs("available", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Step 5: Call the method you're testing
	err = service.UpdateServiceAvailability(1, "available")
	require.NoError(t, err, "expected no error during service update")

	// Step 6: Assert expectations
	err = mock.ExpectationsWereMet()
	require.NoError(t, err, "there were unmet expectations")
}
