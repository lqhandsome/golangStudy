package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int
	Skill string
}

func main() {
  monster := Monster{
  	"牛魔王",
  	5000,
  	"牛角",
  }
	err := monster.Store()
	if err != nil {
		fmt.Println("格式化保存失败",err)
	} else {
		fmt.Println("保存成功")
	}

}

func (monster *Monster) Store() error {

	jsonMonster, err := json.Marshal(monster)
	if err != nil {
		return err
	}
	fileName := monster.Name + ".json"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 777)

	if err != nil {
		return err
	}
	//关闭文件
	defer file.Close()

	_, writerRes := file.WriteString(string(jsonMonster))
	if writerRes != nil {
		return writerRes
	}
	return nil
}

func (monster *Monster) ReStore() error{
	monsterJson ,err := ioutil.ReadFile(monster.Name + ".json")
	if err != nil {
		return  err
	}
 	marshal := json.Unmarshal(monsterJson,&monster)
	if marshal != nil {
		return marshal
	}
	return  nil
}
