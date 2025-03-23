package irepositoryweb2

import entitiesweb2 "backend/internal/domain/entities/web2"

type ILoginRepository interface {
	Verify(username string, password string) (entitiesweb2.User, error)
	Success(entitiesweb2.User) error
	Fail(entitiesweb2.User, error) error
}
