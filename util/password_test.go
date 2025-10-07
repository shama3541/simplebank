package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	password := "hello123"
	hashed_password, err := HashedPassword(password)
	fmt.Println("The password is:", password)
	fmt.Println("the hashed password is:", hashed_password)

	require.NoError(t, err)
	require.NotEmpty(t, hashed_password)

	err = CheckHashesPassword(hashed_password, password)
	require.NoError(t, err)

}
