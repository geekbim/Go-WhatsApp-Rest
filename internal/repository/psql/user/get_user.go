package postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"gokomodo/domain/entity"
	"gokomodo/internal/repository/psql/mapper"
	"gokomodo/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq"
)

func (repository *userRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	stmt := fmt.Sprintf(`select id, email, name, role, password from %s where email = $1 limit 1`, models.User{}.TableName())

	opts := &dbq.Options{SingleResult: true, ConcreteStruct: models.User{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	result, err := dbq.Q(ctx, repository.db, stmt, opts, email)
	if err != nil {
		return nil, err
	}

	if result != nil {
		user := mapper.ToDomainUser(result.(*models.User))
		return user, nil
	} else {
		return nil, errors.New("account not found")
	}
}
