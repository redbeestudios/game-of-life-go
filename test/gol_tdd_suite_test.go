package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGolTdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GolTdd Suite")
}