package repository

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func Test_Repository(t *testing.T) {
	suite.Run(t, &RepositoryTest{})
}
