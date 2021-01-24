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

		gomega.Expect(res.Data.Method).To(gomega.Equal(rose.WriteMethodType))
		gomega.Expect(res.Data.Status).To(gomega.Equal(rose.OkResultStatus))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should read something with uds server", func() {
		ginkgo.Skip("")
		collName := "writeColl_2"

		testCreateCollection(collName)

		conn := testUnixConnect()

		testWrite(conn, rose.WriteMetadata{
			CollectionName: collName,
			Data:           nil,
		})

		testCloseUnixConn(conn)

		conn = testUnixConnect()

		wm := rose.ReadMetadata{
			CollectionName: collName,
			ID: 1,
		}

		b, err := json.Marshal(wm)

		if err != nil {
			gomega.Expect(err).To(gomega.BeNil())
		}

		req := testCreateSocketRequest("read", b)

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)

		b = testReadUnixResponse(conn)

		var res socketResponse
		err = json.Unmarshal(b, &res)

		fmt.Println(string(b))

		gomega.Expect(err).To(gomega.BeNil())

		fmt.Println(string(b))
	})
})
