package goutils_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoutils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goutils Suite")
}
