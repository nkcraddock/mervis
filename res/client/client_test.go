package client_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nkcraddock/mervis/res/client"
	"github.com/nkcraddock/mervis/testhelp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Resource Tests")
}

var _ = Describe("client.handler", func() {
	Context("", func() {
		var srv *testhelp.TestServer
		var c *client.Handler

		BeforeEach(func() {
			router := mux.NewRouter()
			srv = &testhelp.TestServer{router}
			c = client.NewHandler(&client.BinDataLocator{})
			c.Handle(router)
		})

		It("Serves up data from bindata", func() {
			res := srv.GET("/")
			Ω(res.Code).Should(Equal(http.StatusOK))
		})
	})
})
