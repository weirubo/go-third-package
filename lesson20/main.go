package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func main() {
	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	//provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	//client, err := oss.New("oss-rg-china-mainland.aliyuncs.com", "", "", oss.SetCredentialsProvider(&provider))

	endpoint := "your endpoint"
	ak := "your ak"
	sk := "your sk"
	client, err := GetOssClient(endpoint, ak, sk)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	//bucket, err := client.Bucket("thenextoss")
	bucketName := "your bucket name"
	bucket, err := GetBucket(client, bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	//err = bucket.PutObjectFromFile("a2.jpeg", "/Users/frank/Downloads/15801198.jpeg")
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}

	// 上传字符串
	//blogConetnt := "This is my first blog"
	//err = bucket.PutObject("my-first-blog.txt", strings.NewReader(blogConetnt))
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}

	// 上传字节切片
	//user := struct {
	//	UserId   int64  `json:"user_id"`
	//	UserName string `json:"user_name"`
	//	Email    string `json:"email"`
	//}{
	//	UserId:   10001,
	//	UserName: "frank",
	//	Email:    "gopher@88.com",
	//}
	//userData, _ := json.Marshal(user)
	//err = bucket.PutObject("user.txt", bytes.NewReader(userData))
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}

	// 下载到本地文件
	err = bucket.GetObjectToFile("a.jpeg", "/Users/frank/Downloads/oss.jpeg")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func GetOssClient(endpoint, ak, sk string) (client *oss.Client, err error) {
	client, err = oss.New(endpoint, ak, sk)
	return
}

func GetBucket(client *oss.Client, bucketName string) (bucket *oss.Bucket, err error) {
	bucket, err = client.Bucket(bucketName)
	return
}
