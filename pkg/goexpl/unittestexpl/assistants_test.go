package unittestexpl

import (
	"fmt"
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"

	"gorm.io/gorm"
)

func Test_UpsertTestModelToDb(t *testing.T) {
	mockTx := &gorm.DB{}

	// test case
	PatchConvey("Test UpsertTestModelToDb", t, func() {
		fmt.Println("start to mock db")
		// 模拟 tx 的 Model 方法
		Mock(mockTx.Model).Return(mockTx).Build()
		// 模拟 tx 的 Where 方法
		Mock(mockTx.Where).Return(mockTx).Build()
		// 模拟 tx 的 Assign 方法
		Mock(mockTx.Assign).Return(mockTx).Build()
		// 模拟 tx 的 FirstOrCreate 方法，成功返回
		//Mock(mockTx.FirstOrCreate).Return(&gorm.DB{}).Build()

		// initialize a testModel
		tm := &testModel{
			name: "nexus",
			val:  73,
		}
		err := UpsertTestModelToDb(tm, mockTx)
		So(err, ShouldBeNil)
		fmt.Println("done unit test")
	})
}
