package global

import (
	"file_rotatelogs/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	// GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
)
