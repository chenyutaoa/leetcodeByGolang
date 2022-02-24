package leetcode

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/golang/protobuf/proto"
)

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/container-with-most-water
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func maxArea(height []int) int {
	//javaPath := "C:\\Program Files\\ELEF\\ANT\\tools\\android\\java\\bin"
	pp := "D:\\Desktop\\Phone_Log\\20211126164702\\temp\\10\\data\\data\\com.tencent.mm\\MicroMsg"
	cfg := path.Join(pp, "autoauth.cfg")
	//jar := path.Join(pp, "javacode.jar")
	file, err := os.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return 0
	}

	test := &Student{
		Name:   "geektutu",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}
	data, err := proto.Marshal(test)
	var item proto.Message
	err = proto.Unmarshal(content, item)
	if err != nil {
		fmt.Println("反序列化失败", data)
	}
	return 0
}

type Student struct {
	Name   string
	Male   bool
	Scores []int32
}

func (s Student) Reset() {

}

func (s Student) String() string {
	return s.Name
}

func (s Student) ProtoMessage() {
}

type o struct {
}
