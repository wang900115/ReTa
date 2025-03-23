package irepositoryweb2

import entitiesweb2 "backend/internal/domain/entities/web2"

type ILogoutRepository interface {
	Verify(user entitiesweb2.User, uuid string) (entitiesweb2.User, error)
	Success(entitiesweb2.User) error
	Fail(entitiesweb2.User, error) error
}
