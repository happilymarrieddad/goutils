package goutils_test

import (
	"github.com/happilymarrieddad/goutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("ref", func() {
	It("should successfully reference a value", func() {
		Expect(goutils.Ref(1)).To(PointTo(Equal(1)))
		Expect(goutils.Ref("some text")).To(PointTo(Equal("some text")))
	})
})
