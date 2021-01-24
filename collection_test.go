package roseServer

import (
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = GinkgoDescribe("Collection tests", func() {
	GinkgoIt("Should create a collection", func() {
		conn := testUnixConnect()

		req := testCreateSocketRequest("createCollection", []uint8("myColl"))

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)
		b := testReadUnixResponse(conn)

		var res socketResponse
		err := json.Unmarshal(b, &res)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal response: %s", err.Error()))
		}

		gomega.Expect(res.Method).To(gomega.Equal(createCollection))
		gomega.Expect(res.Status).To(gomega.Equal(1))
		gomega.Expect(res.Data).To(gomega.BeNil())

		testCloseUnixConn(conn)
	})
})
