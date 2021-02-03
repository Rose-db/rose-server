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

		var requestBody string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam eu metus cursus, consectetur leo suscipit, fringilla lectus. Donec rutrum finibus mauris condimentum efficitur. Aenean efficitur in urna ut vehicula. Sed ornare ligula tortor. Pellentesque sagittis elit dolor, non consequat sem elementum ut. In vitae nulla eu nisl vehicula gravida at quis ex. Sed arcu lacus, feugiat in velit at, lacinia vulputate sem. Praesent eu rhoncus nunc. Praesent eleifend suscipit imperdiet. Sed hendrerit viverra feugiat. Aenean pretium augue vitae risus maximus, vel scelerisque urna venenatis. Praesent id hendrerit turpis. Proin nulla dui, sollicitudin nec maximus et, aliquam sed lectus. Nulla lacinia ante eu risus dapibus tincidunt. Pellentesque tortor sem, venenatis eget magna sed, laoreet pellentesque velit. Phasellus malesuada eu nisi eu accumsan. Curabitur fermentum justo at vehicula porttitor. Aliquam sit amet dictum diam. Fusce eu tristique leo, sed volutpat est. Phasellus eu auctor turpis. Donec gravida sodales quis. ";

		wm := rose.WriteMetadata{
			CollectionName: collName,
			Data:           testAsJson(requestBody),
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

		testCloseUnixConn(conn)

		conn = testUnixConnect()

		dm := rose.DeleteMetadata{
			CollectionName: collName,
			ID: 1,
		}

		b, err := json.Marshal(dm)

		if err != nil {
			gomega.Expect(err).To(gomega.BeNil())
		}

		req := testCreateSocketRequest("delete", b)

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)

		b = testReadUnixResponse(conn)

		var res socketResponse
		err = json.Unmarshal(b, &res)

		gomega.Expect(err).To(gomega.BeNil())

		gomega.Expect(res.Error).To(gomega.BeNil())
		gomega.Expect(res.Method).To(gomega.Equal(deleteMethod))
		gomega.Expect(res.Data).To(gomega.Not(gomega.BeNil()))
		gomega.Expect(res.Status).To(gomega.Equal(OperationSuccessCode))
		gomega.Expect(res.ReadData).To(gomega.BeNil())

		gomega.Expect(res.Data.Status).To(gomega.Equal(rose.DeletedResultStatus))
	})

	GinkgoIt("Should replace a document with uds server", func() {
		collName := "writeColl_4"

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

		testCloseUnixConn(conn)

		conn = testUnixConnect()

		replacedData := "replaced data"

		dm := rose.ReplaceMetadata{
			CollectionName: collName,
			ID: 1,
			Data: testAsJson(replacedData),
		}

		b, err := json.Marshal(dm)

		if err != nil {
			gomega.Expect(err).To(gomega.BeNil())
		}

		req := testCreateSocketRequest("replace", b)

		testWriteUnixServer(conn, req)

		testCloseUnixWriteConn(conn)

		b = testReadUnixResponse(conn)

		var res socketResponse
		err = json.Unmarshal(b, &res)

		gomega.Expect(err).To(gomega.BeNil())

		gomega.Expect(res.Error).To(gomega.BeNil())
		gomega.Expect(res.Method).To(gomega.Equal(replaceMethod))
		gomega.Expect(res.Data).To(gomega.Not(gomega.BeNil()))
		gomega.Expect(res.Status).To(gomega.Equal(OperationSuccessCode))
		gomega.Expect(res.ReadData).To(gomega.BeNil())

		gomega.Expect(res.Data.Status).To(gomega.Equal(rose.ReplacedResultStatus))

		testCloseUnixConn(conn)

		conn = testUnixConnect()

		rm = rose.ReadMetadata{
			CollectionName: collName,
			ID: 1,
		}

		testRead(conn, rm, replacedData)

		testCloseUnixConn(conn)
	})
})
