package dcs_tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	DCS "github.com/thedoctor028/dcsexportscriptapi/dcs"
)

var _ = Describe("Utils Tests", func() {
	var exportData DCS.ExportData

	BeforeEach(func() {
		exportData = DCS.NewDcsData()
		exportData.Data[100] = "100"
		exportData.Data[101] = "101"
		exportData.Data[102] = "102"
	})

	Describe("NewDcsData", func() {
		Context("when creating a new ExportData", func() {
			It("should return an export data with empty map", func() {
				data := DCS.NewDcsData()
				Expect(data.Data).To(Equal(map[int]string{}))
			})
		})
	})

	Describe("GetDataByUid", func() {
		Context("when getting a data by uid", func() {
			It("should return nil if it is not exists", func() {
				res := exportData.GetDataByUid(99)
				Expect(res).To(BeNil())
			})

			It("should return the string under the id if it exists", func() {
				res := exportData.GetDataByUid(100)
				Expect(*res).To(Equal("100"))
			})
		})
	})

	Describe("ToString", func() {
		Context("when converting to string", func() {
			It("should return the data in a formatted string", func() {
				Expect(exportData.ToString()).To(Equal("100:100\n101:101\n102:102\n"))
			})
		})
	})
})
