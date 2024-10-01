package repository

import (
	"fmt"
	"strconv"

	"github.com/PhamNgocHoangLong/greendeco-be.git/app/models"
	"github.com/PhamNgocHoangLong/greendeco-be.git/platform/database"
)

type AdminRepository interface {
	GetCustomer(limit, offset int) ([]*models.User, error)
}

type AdminRepo struct {
	db *database.DB
}

var _ AdminRepository = (*AdminRepo)(nil)

const (
	RoleTable     = "roles"
	UserRoleTable = "user_role"
)

func NewAdminRepo(db *database.DB) AdminRepository {
	return &AdminRepo{db}
}

func (repo *AdminRepo) GetCustomer(limit, offset int) ([]*models.User, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE admin = false`, UserTable)
	if limit > 0 {
		query += " LIMIT " + strconv.Itoa(limit)
	}

	if offset > 0 {
		query += " OFFSET" + strconv.Itoa(offset)
	}

	users := []*models.User{}
	if err := repo.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}
