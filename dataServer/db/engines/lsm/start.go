package lsm

import (
	"github.com/huiming23344/mindfs/dataServer/db/engines/lsm/config"
	"github.com/huiming23344/mindfs/dataServer/db/engines/lsm/ssTable"
	"log"
	"os"
)

// Start 启动数据库
func Start(con config.Config) {
	if database != nil {
		return
	}
	// 将配置保存到内存中
	log.Println("Loading a Configuration File")
	config.Init(con)
	// 初始化数据库
	log.Println("Initializing the database")
	initDatabase(con.DataDir)

	// 检查内存
	checkMemory()
	// 检查压缩数据库文件
	database.TableTree.Check()
	// 启动后台线程
	go Check()
	go CompressMemory()
}

// 初始化 Database，从磁盘文件中还原 SSTable、WalF、内存表等
func initDatabase(dir string) {
	database = &Database{
		MemTable:  &MemTable{},
		iMemTable: &ReadOnlyMemTables{},
		TableTree: &ssTable.TableTree{},
	}
	// 从磁盘文件中恢复数据
	// 如果目录不存在，则为空数据库
	if _, err := os.Stat(dir); err != nil {
		log.Printf("The %s directory does not exist. The directory is being created\r\n", dir)
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			log.Println("Failed to create the database directory")
			panic(err)
		}
	}
	database.iMemTable.Init()
	database.MemTable.InitMemTree()
	log.Println("Loading all wal.log...")
	database.loadAllWalFiles(dir)
	database.MemTable.InitWal(dir)
	log.Println("Loading database...")
	database.TableTree.Init(dir)
}
