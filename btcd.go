// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	//"github.com/btcsuite/btclog"
	"github.com/btcsuite/btcd/blockchain/indexers"
	"github.com/btcsuite/btcd/database"
	"github.com/btcsuite/btcd/limits"
)

const (
	// blockDbNamePrefix is the prefix for the block database name.  The
	// database type is appended to this value to form the full block
	// database name.
	blockDbNamePrefix = "blocks"
)

var (
	cfg *config
)

// winServiceMain is only invoked on Windows.  It detects when btcd is running
// as a service and reacts accordingly.
var winServiceMain func() (bool, error)


//对于os.Exit()被调用,而被延迟执行的函数并没按预期执行的事实,是很有必要去解决的.
// btcdMain is the real main function for btcd.  It is necessary to work around
// the fact that deferred functions do not run when os.Exit() is called.
//serverChan通道主要被服务代码使用,一旦设置,服务器就会被通知,当服务控制器传来请求时,服务会被优雅地停止
// The optional serverChan parameter is mainly used by the service code to be
// notified with the server once it is setup, so it can gracefully stop it when
// requested from the service control manager.
func btcdMain(serverChan chan<- *server) error {
	// Load configuration and parse command line.  This function also
	// initializes logging and configures it accordingly.
	//加载配置信息，使用了github.com/jessevdk/go-flags这个第三方库，
	// 可以实现从配置文件或命令行获取参数值，功能挺强大的库，近期也有一直在维护。
	tcfg, _, err := loadConfig()
	if err != nil {
		return err
	}
	cfg = tcfg
	defer func() {
		if logRotator != nil {
			logRotator.Close()
		}
	}()

	// Get a channel that will be closed when a shutdown signal has been
	// triggered either from an OS signal such as SIGINT (Ctrl+C) or from
	// another subsystem such as the RPC server.
	//这个方法主要是使用signal.Notify(interruptChannel, interruptSignals...)来监听系统interrupt的signal，如有中断则close channel。
	interrupt := interruptListener()
	defer btcdLog.Info("Shutdown complete")

	// Show version at startup.
	btcdLog.Infof("Version %s", version())

	// Enable http profiling server if requested.
	// 开启http服务器性能监测
	if cfg.Profile != "" {
		go func() {
			listenAddr := net.JoinHostPort("", cfg.Profile)
			btcdLog.Infof("Profile server listening on %s", listenAddr)
			profileRedirect := http.RedirectHandler("/debug/pprof",
				http.StatusSeeOther)
			http.Handle("/", profileRedirect)
			btcdLog.Errorf("%v", http.ListenAndServe(listenAddr, nil))
		}()
	}
	// 开启cpu性能监测
	// Write cpu profile if requested.
	if cfg.CPUProfile != "" {
		f, err := os.Create(cfg.CPUProfile)
		if err != nil {
			btcdLog.Errorf("Unable to create cpu profile: %v", err)
			return err
		}
		pprof.StartCPUProfile(f)
		defer f.Close()
		defer pprof.StopCPUProfile()
	}

	// Perform upgrades to btcd as new versions require it.
	// 升级老版本数据库相关文件路径
	if err := doUpgrades(); err != nil {
		btcdLog.Errorf("%v", err)
		return err
	}

	// Return now if an interrupt signal was triggered.
	// 是否有中断信号
	if interruptRequested(interrupt) {
		return nil
	}

	// Load the block database.
	//根据cfg.DbType配置打开数据库连接。通过代码得知支持数据库有dbTypes := []string{"ffldb", "leveldb", "sqlite"}
	// 加载区块数据库
	db, err := loadBlockDB()
	if err != nil {
		btcdLog.Errorf("%v", err)
		return err
	}
	defer func() {
		// Ensure the database is sync'd and closed on shutdown.
		btcdLog.Infof("Gracefully shutting down the database...")
		db.Close()
	}()

	// Return now if an interrupt signal was triggered.
	// 是否有中断信号
	if interruptRequested(interrupt) {
		return nil
	}

	// Drop indexes and exit if requested.
	//
	// NOTE: The order is important here because dropping the tx index also
	// drops the address index since it relies on it.
	// 删除地址索引，并退出程序
	//
	// 注意： 在这里删除顺序很重要，因为删除交易索引同时也会删除地址索引。
	// 因为交易索引会依赖地址索引
	if cfg.DropAddrIndex {
		if err := indexers.DropAddrIndex(db, interrupt); err != nil {
			btcdLog.Errorf("%v", err)
			return err
		}

		return nil
	}
	// 删除交易索引，并退出程序
	if cfg.DropTxIndex {
		if err := indexers.DropTxIndex(db, interrupt); err != nil {
			btcdLog.Errorf("%v", err)
			return err
		}

		return nil
	}
	// 删除交易索引，并退出程序
	if cfg.DropCfIndex {
		if err := indexers.DropCfIndex(db, interrupt); err != nil {
			btcdLog.Errorf("%v", err)
			return err
		}

		return nil
	}

	// Create server and start it.
	//使用配置的端口、刚连接的db及网络参数初始一个server对象。
	//activeNetParams.Params是硬编码参数，如下：
	//var MainNetParams = Params{
	//	Name:        "mainnet",
	//	Net:         wire.MainNet,
	//	DefaultPort: "8333",
	//	DNSSeeds: []DNSSeed{
	//		{"seed.bitcoin.sipa.be", true},
	//		{"dnsseed.bluematt.me", true},
	//		{"dnsseed.bitcoin.dashjr.org", false},
	//		{"seed.bitcoinstats.com", true},
	//		{"seed.bitnodes.io", false},
	//		{"seed.bitcoin.jonasschnelli.ch", true},
	//	},
	//
	//	// Chain parameters
	//	GenesisBlock:             &genesisBlock,
	//}
	btcdLog.Infof("new  %s", "新建server")
	// 创建服务器，并启动
	server, err := newServer(cfg.Listeners, cfg.AgentBlacklist,
		cfg.AgentWhitelist, db, activeNetParams.Params, interrupt)
	if err != nil {
		// TODO: this logging could do with some beautifying.
		btcdLog.Errorf("Unable to start server on %v: %v",
			cfg.Listeners, err)
		return err
	}
	defer func() {
		btcdLog.Infof("Gracefully shutting down the server...")
		server.Stop()
		server.WaitForShutdown()
		srvrLog.Infof("Server shutdown complete")
	}()
	server.Start()
	if serverChan != nil {
		serverChan <- server
	}

	// Wait until the interrupt signal is received from an OS signal or
	// shutdown is requested through one of the subsystems such as the RPC
	// server.
	//挂起主线程，等待中断信号
	<-interrupt
	return nil
}

// removeRegressionDB removes the existing regression test database if running
// in regression test mode and it already exists.
func removeRegressionDB(dbPath string) error {
	// Don't do anything if not in regression test mode.
	if !cfg.RegressionTest {
		return nil
	}

	// Remove the old regression test database if it already exists.
	fi, err := os.Stat(dbPath)
	if err == nil {
		btcdLog.Infof("Removing regression test database from '%s'", dbPath)
		if fi.IsDir() {
			err := os.RemoveAll(dbPath)
			if err != nil {
				return err
			}
		} else {
			err := os.Remove(dbPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// dbPath returns the path to the block database given a database type.
func blockDbPath(dbType string) string {
	// The database name is based on the database type.
	dbName := blockDbNamePrefix + "_" + dbType
	if dbType == "sqlite" {
		dbName = dbName + ".db"
	}
	dbPath := filepath.Join(cfg.DataDir, dbName)
	return dbPath
}

// warnMultipleDBs shows a warning if multiple block database types are detected.
// This is not a situation most users want.  It is handy for development however
// to support multiple side-by-side databases.
func warnMultipleDBs() {
	// This is intentionally not using the known db types which depend
	// on the database types compiled into the binary since we want to
	// detect legacy db types as well.
	dbTypes := []string{"ffldb", "leveldb", "sqlite"}
	duplicateDbPaths := make([]string, 0, len(dbTypes)-1)
	for _, dbType := range dbTypes {
		if dbType == cfg.DbType {
			continue
		}

		// Store db path as a duplicate db if it exists.
		dbPath := blockDbPath(dbType)
		if fileExists(dbPath) {
			duplicateDbPaths = append(duplicateDbPaths, dbPath)
		}
	}

	// Warn if there are extra databases.
	if len(duplicateDbPaths) > 0 {
		selectedDbPath := blockDbPath(cfg.DbType)
		btcdLog.Warnf("WARNING: There are multiple block chain databases "+
			"using different database types.\nYou probably don't "+
			"want to waste disk space by having more than one.\n"+
			"Your current database is located at [%v].\nThe "+
			"additional database is located at %v", selectedDbPath,
			duplicateDbPaths)
	}
}

// loadBlockDB loads (or creates when needed) the block database taking into
// account the selected database backend and returns a handle to it.  It also
// contains additional logic such warning the user if there are multiple
// databases which consume space on the file system and ensuring the regression
// test database is clean when in regression test mode.
func loadBlockDB() (database.DB, error) {
	// The memdb backend does not have a file path associated with it, so
	// handle it uniquely.  We also don't want to worry about the multiple
	// database type warnings when running with the memory database.
	if cfg.DbType == "memdb" {
		btcdLog.Infof("Creating block database in memory.")
		db, err := database.Create(cfg.DbType)
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	warnMultipleDBs()

	// The database name is based on the database type.
	dbPath := blockDbPath(cfg.DbType)

	// The regression test is special in that it needs a clean database for
	// each run, so remove it now if it already exists.
	removeRegressionDB(dbPath)

	btcdLog.Infof("Loading block database from '%s'", dbPath)
	db, err := database.Open(cfg.DbType, dbPath, activeNetParams.Net)
	if err != nil {
		// Return the error if it's not because the database doesn't
		// exist.
		if dbErr, ok := err.(database.Error); !ok || dbErr.ErrorCode !=
			database.ErrDbDoesNotExist {

			return nil, err
		}

		// Create the db if it does not exist.
		err = os.MkdirAll(cfg.DataDir, 0700)
		if err != nil {
			return nil, err
		}
		db, err = database.Create(cfg.DbType, dbPath, activeNetParams.Net)
		if err != nil {
			return nil, err
		}
	}

	btcdLog.Info("Block database loaded")
	return db, nil
}

func main() {
	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	//设置垃圾回收比率
	// Block and transaction processing can cause bursty allocations.  This
	// limits the garbage collector from excessively overallocating during
	// bursts.  This value was arrived at with the help of profiling live
	// usage.
	//设置触发GC垃圾收集器阀值： 10% ， 即 最新分布的内存大小/上次GC之后活动内存大小*100
	debug.SetGCPercent(10)

	// Up some limits.
	//limits包中有针对不同平台的三个文件：limits_plan9.go limits_unix.go limits_windows.go
	//设置系统资源限制
	//设置进程能打开的最大文件资源,即file descriptor数量。
	// 做linux系统中，所有资源(pipe, socket, file等等)都以文件形式出现。
	if err := limits.SetLimits(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to set limits: %v\n", err)
		os.Exit(1)
	}

	// Call serviceMain on Windows to handle running as a service.  When
	// the return isService flag is true, exit now since we ran as a
	// service.  Otherwise, just fall through to normal operation.
	if runtime.GOOS == "windows" {
		isService, err := winServiceMain()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if isService {
			os.Exit(0)
		}
	}

	// Work around defer not working after os.Exit()
	if err := btcdMain(nil); err != nil {
		os.Exit(1)
	}
}
