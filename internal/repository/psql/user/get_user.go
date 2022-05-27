package postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"go-rest-ddd/domain/entity"
	"go-rest-ddd/internal/repository/psql/mapper"
	"go-rest-ddd/internal/repository/psql/models"

	"github.com/rocketlaunchr/dbq"
)

func (repository *userRepository) GetUserByUserNameAndPassword(ctx context.Context, userName, password string) (*entity.User, error) {
	stmt := fmt.Sprintf(`select * from %s where user_name = $1 and password = MD5($2) limit 1`, models.User{}.TableName())

	opts := &dbq.Options{SingleResult: true, ConcreteStruct: models.User{}, DecoderConfig: dbq.StdTimeConversionConfig()}

	result, err := dbq.Q(ctx, repository.db, stmt, opts, userName, password)
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
