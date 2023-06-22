package vo

import (
	libuuid "github.com/google/uuid"
)

func UUID() string {
	return libuuid.NewString()
}
