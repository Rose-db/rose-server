package roseServer

import (
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"rose/rose"
)

var _ = GinkgoDescribe("Write tests", func() {
	GinkgoIt("Should write something with uds server", func() {
		collName := "writeColl_1"

		testCreateCollection(collName)

		conn := testUnixConnect()

		wm := rose.WriteMetadata{
			CollectionName: collName,
			Data:           testAsJson("something to write"),
		}

		b, err := json.Marshal(wm)

		if err != nil {
			gomega.Expect(err).To(gomega.BeNil())
		}

		req := testCreateSocketRequest("write", b)

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)

		b = testReadUnixResponse(conn)

		var res socketResponse
		err = json.Unmarshal(b, &res)

		if err != nil {
			ginkgo.Fail(fmt.Sprintf("Unable to unmarshal response: %s", err.Error()))
		}

		gomega.Expect(res.Method).To(gomega.Equal(write))
		gomega.Expect(res.Status).To(gomega.Equal(OperationSuccessCode))

		gomega.Expect(res.Result.Method).To(gomega.Equal(rose.WriteMethodType))
		gomega.Expect(res.Result.Status).To(gomega.Equal(rose.OkResultStatus))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should write something with uds server", func() {
		collName := "writeColl_2"

		testCreateCollection(collName)

		conn := testUnixConnect()

		testWrite(conn, rose.WriteMetadata{
			CollectionName: collName,
			Data:           nil,
		})

		testCloseUnixConn(conn)
	})
})
