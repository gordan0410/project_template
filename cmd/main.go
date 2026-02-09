package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"project_template/internal/config"
	"sync"
	"syscall"
	"time"

	common_config "github.com/gordan0410/common/config"
	common_dto "github.com/gordan0410/common/dto"
	common_enum "github.com/gordan0410/common/enum"
	common_helper "github.com/gordan0410/common/helper"
	common_log "github.com/gordan0410/common/log"
)

func init() {
	// 讀取版本資訊
	file, err := os.Open("version.json")
	if err != nil {
		log.Panicln("Error opening version.json:", err)
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Panicln("Error closing version.json:", err)
		}
	}()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&version)
	if err != nil {
		log.Panicln("Error  decoding version.json:", err)
		return
	}
}

var (
	version common_dto.VersionInfo
)

func main() {
	log.Printf("current project : %s \n", version.ProjectName)
	log.Printf("current version : %s \n", version.Version)
	log.Printf("current environment : %s \n", version.Environment)

	baseConfig := InitConfig()
	commonLogger := common_log.NewLogger(common_log.LogField{ProjectName: version.ProjectName, Version: version.Version, Level: baseConfig.Config.SystemConfig.LogLevel})

	app := InitializeApplication(
		version,
		commonLogger,
		baseConfig.Config,
		&baseConfig.Config.MasterDBConfig,
		&baseConfig.Config.RedisConfig,
	)

	runCtx := context.Background()
	runCtx = context.WithValue(runCtx, common_enum.TraceId.ToString(), common_helper.GetUUID())
	runCtx = context.WithValue(runCtx, common_enum.SubTraceId.ToString(), common_helper.GetUUID())

	app.httpServer.Run(runCtx, fmt.Sprintf(":%d", baseConfig.Config.SystemConfig.HTTPPort))

	// 開始執行 graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx = context.WithValue(ctx, common_enum.TraceId.ToString(), common_helper.GetUUID())
	ctx = context.WithValue(ctx, common_enum.SubTraceId.ToString(), common_helper.GetUUID())
	commonLogger.WithCtx(ctx).Info().Msg("start graceful shutdown")

	defer cancel()

	sdc := make(chan bool, 1)
	go func() {
		var wg sync.WaitGroup
		wg.Add(0)
		wg.Wait()
		sdc <- true
	}()

	select {
	case <-ctx.Done():
		commonLogger.Error().Msg("shutdown timeout")
		fmt.Println("shutdown timeout")
	case <-sdc:
		commonLogger.Info().Msg("shutdown gracefully")
	}
}

// InitConfig 初始化環境變數
func InitConfig() *common_config.BaseConfig[config.EnvConfig] {
	common_config.InitDotEnv(version.Environment, ".env")
	baseConfig := common_config.NewConfig[config.EnvConfig]()
	if baseConfig.Config == nil {
		log.Panic("Error initializing config")
	}

	return baseConfig
}
