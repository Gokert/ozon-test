package profile

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"
	"ozon-test/configs"
	utils "ozon-test/pkg"
	"ozon-test/pkg/models"
	"strconv"
	"time"
)

//go:generate mockgen -source=profile_repo.go -destination=../../mocks/repo_mock.go -package=mocks
type IRepository interface {
	GetUser(login string, password []byte) (*models.UserItem, bool, error)
	GetUserName(id uint64) (string, error)
	FindUser(login string) (bool, error)
	CreateUser(login string, password []byte) error
	GetUserId(login string) (uint64, error)
	GetRole(userId uint64) (string, error)
}

type Repository struct {
	db *sql.DB
}

func GetPsxRepo(config *configs.DbPsxConfig, logger *logrus.Logger) (*Repository, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password= %s host=%s port=%d sslmode=%s",
		config.User, config.Dbname, config.Password, config.Host, config.Port, config.Sslmode)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql open error: %s", err.Error())
	}

	repo := &Repository{db: db}

	errs := make(chan error)
	go func() {
		errs <- repo.pingDb(3, logger)
	}()

	if err := <-errs; err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	logger.Info("Successfully connected to database")

	return repo, nil
}

func (r *Repository) pingDb(timer uint32, log *logrus.Logger) error {
	var err error
	var retries int

	for retries < utils.MaxRetries {
		err = r.db.Ping()
		if err == nil {
			return nil
		}

		retries++
		log.Errorf("sql ping error: %s", err.Error())
		time.Sleep(time.Duration(timer) * time.Second)
	}

	return fmt.Errorf("sql max pinging error: %s", err.Error())
}

func (repo *Repository) GetUser(login string, password []byte) (*models.UserItem, bool, error) {
	post := &models.UserItem{}

	err := repo.db.QueryRow("SELECT profile.id, profile.login FROM profile "+
		"WHERE profile.login = $1 AND profile.password = $2 ", login, password).Scan(&post.Id, &post.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("get query user error: %s", err.Error())
	}

	return post, true, nil
}

func (repo *Repository) GetUserName(id uint64) (string, error) {
	var name string

	str := strconv.FormatUint(id, 10)

	err := repo.db.QueryRow("SELECT profile.login FROM profile WHERE profile.id = $1 ", str).Scan(&name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("get query user name error: %s", err.Error())
	}

	return name, nil
}

func (repo *Repository) FindUser(login string) (bool, error) {
	post := &models.UserItem{}

	err := repo.db.QueryRow(
		"SELECT login FROM profile "+
			"WHERE login = $1", login).Scan(&post.Login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("find user query error: %s", err.Error())
	}

	return true, nil
}

func (repo *Repository) CreateUser(login string, password []byte) error {
	var userID uint64
	err := repo.db.QueryRow("INSERT INTO profile(login, password) VALUES($1, $2) RETURNING id", login, password).Scan(&userID)
	if err != nil {
		return fmt.Errorf("create user error: %s", err.Error())
	}

	return nil
}

func (repo *Repository) GetUserId(login string) (uint64, error) {
	var userID uint64

	err := repo.db.QueryRow(
		"SELECT profile.id FROM profile WHERE profile.login = $1", login).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("user not found for login: %s", login)
		}
		return 0, fmt.Errorf("select user profile id error: %s", err.Error())
	}

	return userID, nil
}

func (repo *Repository) GetRole(userId uint64) (string, error) {
	var roleName string

	err := repo.db.QueryRow("SELECT profile.role FROM profile  WHERE profile.id = $1", userId).Scan(&roleName)
	if err != nil {
		return "", fmt.Errorf("get user role err: %s", err.Error())
	}

	return roleName, nil
}
