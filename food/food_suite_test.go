package food_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFood(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Food Suite")
}
