package main

import (
	"fmt"

	"github.com/iwindy/face-engine/v4/util"
)

var imageInfo = util.GetResizedImageInfo("./mask.jpg")

func main() {
	// 激活SDK
	if err := OnlineActivation("8EKZpEAeQCnD2ysKdxSrY6TNzog5kjQLHtNYzAGxkRtF", "8YnJR7PmW1yd12wxCA4sXzxop1ipRd9FMGmdmAKdjqct", "86J1-11LD-V13U-7T6Q"); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 初始化引擎
	engine, err := NewFaceEngine(DetectModeVideo,
		OrientPriority0,
		50, // 4.0最大支持10个人脸
		EnableFaceDetect|EnableFaceRecognition|EnableLiveness|EnableIRLiveness|EnableAge|EnableGender|EnableMaskDetect)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	deviceInfo, err := GetActiveDeviceInfo()
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("设备信息：%s\n", deviceInfo)
	// 检测人脸
	info, err := engine.DetectFaces(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8)
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	if info.FaceNum > 0 {
		fmt.Printf("3D FaceId:%d pitch:%f yaw:%f rall:%f \n", info.FaceID[0], info.Face3DAngle.Pitch[0], info.Face3DAngle.Yaw[0], info.Face3DAngle.Roll[0])
		fmt.Printf("RECT FaceId:%d Top:%d Right:%d  \n", info.FaceID[0], info.ForeheadRect[0].Top, info.ForeheadRect[0].Right)
	}
	// 处理人脸数据
	if err = engine.Process(imageInfo.Width, imageInfo.Height, ColorFormatBGR24, imageInfo.DataUInt8, info, EnableAge|EnableGender|EnableLiveness|EnableMaskDetect); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	// 获取年龄
	ageInfo, err := engine.GetAge()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("ageInfo: %v\n", ageInfo)
	// 获取口罩信息
	maskInfo, err := engine.GetMask()
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	fmt.Printf("口罩信息：%#v\n", maskInfo)
	// 销毁引擎
	if err = engine.Destroy(); err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
}
