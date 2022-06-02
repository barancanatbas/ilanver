package repositorytest

import (
	"database/sql"
	"database/sql/driver"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func setup(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *gorm.DB) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//dsn := "root:mysql123@tcp(127.0.0.1:3306)/dbilanver2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       mockDb,
	}), &gorm.Config{})

	return mockDb, mock, db
}

func setupQM(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *gorm.DB) {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//dsn := "root:mysql123@tcp(127.0.0.1:3306)/dbilanver2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       mockDb,
	}), &gorm.Config{})

	return mockDb, mock, db
}

func TestGetAllAddress(t *testing.T) {

	mockDb, mock, db := setup(t)
	defer mockDb.Close()

	address := []model.Adress{
		{
			Districtfk: 1,
			Detail:     "test address",
		},
		{
			Districtfk: 2,
			Detail:     "test address2",
		},
	}

	repo := repository.NewAddressRepository(db)

	mock.ExpectQuery("SELECT (.+) FROM `adresses`").
		WillReturnRows(sqlmock.NewRows([]string{"id", "districtfk", "detail"}).
			AddRow(address[0].ID, address[0].Districtfk, address[0].Detail).AddRow(address[1].ID, address[1].Districtfk, address[1].Detail))

	getAddress, err := repo.GetAll()

	assert.Equal(t, err, nil)
	assert.Equal(t, getAddress[0].Detail, address[0].Detail)

	assert.Equal(t, getAddress[1].Detail, address[1].Detail)

}

func TestInsertAddress(t *testing.T) {
	address := model.Adress{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Districtfk: 1,
		DeletedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Detail: "test address",
	}

	mockDb, mock, db := setupQM(t)
	defer mockDb.Close()

	repo := repository.NewAddressRepository(db)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `adresses` (`created_at`,`updated_at`,`deleted_at`,`detail`,`districtfk`) VALUES (?,?,?,?,?)").
		WithArgs(address.CreatedAt, address.UpdatedAt, address.DeletedAt.Time, address.Detail, address.Districtfk).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Save(&address)

	assert.Equal(t, err, nil)
}

func TestGetById(t *testing.T) {
	mockDb, mock, db := setupQM(t)
	defer mockDb.Close()

	repo := repository.NewAddressRepository(db)

	address := model.Adress{
		ID:         1,
		Districtfk: 1,
		Detail:     "test address",
	}

	mock.ExpectQuery("SELECT * FROM `adresses` WHERE id = ? ORDER BY `adresses`.`id` LIMIT 1").WithArgs(address.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "districtfk", "detail"}).
			AddRow(address.ID, address.Districtfk, address.Detail))

	getAddress, err := repo.GetByID(address.ID)

	assert.Equal(t, err, nil)
	assert.Equal(t, getAddress.Detail, address.Detail)
}

func TestUpdateAddress(t *testing.T) {
	mockDb, mock, db := setupQM(t)
	defer mockDb.Close()

	date := time.Now().UTC()

	repo := repository.NewAddressRepository(db)

	address := model.Adress{
		UpdatedAt:  date,
		Districtfk: 1,
		Detail:     "test address",
		ID:         1,
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `adresses` SET `detail`=?,`districtfk`=?,`updated_at`=? WHERE `id` = ?").
		WithArgs(address.Detail, address.Districtfk, address.UpdatedAt, address.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Update(&address)

	assert.Equal(t, err, nil)
}
