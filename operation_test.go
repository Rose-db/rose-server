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

		gomega.Expect(res.Method).To(gomega.Equal(writeMethod))
		gomega.Expect(res.Status).To(gomega.Equal(OperationSuccessCode))

		gomega.Expect(res.Data.Method).To(gomega.Equal(rose.WriteMethodType))
		gomega.Expect(res.Data.Status).To(gomega.Equal(rose.OkResultStatus))

		testCloseUnixConn(conn)
	})

	GinkgoIt("Should read something with uds server", func() {
		collName := "writeColl_2"

		testCreateCollection(collName)

		conn := testUnixConnect()

		writeData := "something written"

		testWrite(conn, rose.WriteMetadata{
			CollectionName: collName,
			Data:           testAsJson(writeData),
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

		gomega.Expect(err).To(gomega.BeNil())

		gomega.Expect(res.Error).To(gomega.BeNil())
		gomega.Expect(res.Method).To(gomega.Equal(readMethod))
		gomega.Expect(res.Data).To(gomega.Not(gomega.BeNil()))
		gomega.Expect(res.Status).To(gomega.Equal(OperationSuccessCode))
		gomega.Expect(res.ReadData).To(gomega.Equal(writeData))

		gomega.Expect(res.Data.Status).To(gomega.Equal(rose.FoundResultStatus))
	})

	GinkgoIt("Should delete a document with uds server", func() {
		collName := "writeColl_3"

		testCreateCollection(collName)

		conn := testUnixConnect()

		writeData := "something written"

		testWrite(conn, rose.WriteMetadata{
			CollectionName: collName,
			Data:           testAsJson(writeData),
		})

		testCloseUnixConn(conn)

		conn = testUnixConnect()

		rm := rose.ReadMetadata{
			CollectionName: collName,
			ID: 1,
		}

		testRead(conn, rm, writeData)
	})
})
