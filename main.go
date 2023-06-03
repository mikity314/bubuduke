package main

import (
	"log"

	"github.com/electricbubble/gadb"
	"github.com/manifoldco/promptui"
)

func main(){
	// クライアントを生成
	adbClient, err := gadb.NewClient()
	if err != nil {
		log.Fatalln(err)
	}

	// デバイス一覧を取得
	devices, err := adbClient.DeviceList()
	if err != nil{
		log.Fatalln(err)
	}

	// デバイスを選択
	if len(devices) == 0{
		log.Fatalln("デバイスが接続されていません。")
	}
	if len(devices) >= 2{
		log.Fatalln("複数のデバイスが接続されています。現在は複数デバイスはサポートしていません。")
	}
	device := devices[0]

	// 方法を選択
	allAppWhichHasCareerName := "キャリア名を含むアプリを一括で削除"
	fromBlackList := "ブラックリストに指定されているアプリを削除"
	prompt := promptui.Select{
		Label: "方法を選択してください",
		Items: []string{
			allAppWhichHasCareerName,
			fromBlackList,
		},
	}
	_, result, err := prompt.Run()
	if err !=nil{
		log.Fatalln(err)
	}

	switch result{
	case allAppWhichHasCareerName:
		remove_all_app(&adbClient, &device)
	case fromBlackList:
		remove_from_blacklist(&adbClient, &device)
	default:
		log.Fatalln("何コレ（ゴロリ）")
	}

}

func remove_all_app(client *gadb.Client, dev *gadb.Device){

}

func remove_from_blacklist(client *gadb.Client, dev *gadb.Device){
	log.Fatalln("この方法はまだサポートされていません、ごめんね。")
}
