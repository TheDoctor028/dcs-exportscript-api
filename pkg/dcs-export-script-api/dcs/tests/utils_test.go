package utils_tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/thedoctor028/dcsexportscriptapi/dcs"
	"testing"
)

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utils Suite")
}

var _ = Describe("Utils Tests", func() {
	var sampleText string

	BeforeEach(func() {
		sampleText = "62b74d61*146=0.0017:141=0.0074 \n:128=-0.0682:146=0.0008:180=0.4000:141=0.0018\n=0:184=0:200=0"
	})

	Context("When the decoded string", func() {
		It("should map it to an int string map", func() {
			expected := getExtractUIDAndValueExpectedValues()

			res := DCS.ExtractUIDAndValue(sampleText, ":")

			Expect(res.Data).To(Equal(expected))
		})
	})

})

func getExtractUIDAndValueExpectedValues() map[int]string {
	keyValueMap := make(map[int]string)
	keyValueMap[200] = "0"
	keyValueMap[184] = "0"
	keyValueMap[141] = "0.0074"
	keyValueMap[180] = "0.4000"
	keyValueMap[146] = "0.0008"
	keyValueMap[128] = "-0.0682"

	return keyValueMap
}
