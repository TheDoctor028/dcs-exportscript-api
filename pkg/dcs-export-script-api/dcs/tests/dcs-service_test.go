package dcs_tests

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	DCS "github.com/thedoctor028/dcsexportscriptapi/dcs"
	udpConnection "github.com/thedoctor028/dcsexportscriptapi/udp-connection"
	mock_udp_connection "github.com/thedoctor028/dcsexportscriptapi/udp-connection/mock"
)

var _ = Describe("DcsService Tests", func() {
	var service *DCS.Service

	var (
		mockCtrl      *gomock.Controller
		mockUDPClient udpConnection.IUDPClient
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockUDPClient = mock_udp_connection.NewMockIUDPClient(mockCtrl)
	})

	BeforeEach(func() {
		service = DCS.NewService()

	})

	Describe("NewService", func() {

	})

})
