package log

import "go.uber.org/zap"

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelFatal Level = "fatal"
)

type Config struct {
	Env   string `json:"env" mapstructure:"env"`
	Level Level  `json:"level" mapstructure:"level"`
}

type Logger struct {
	*zap.Logger
}

var mapLevel = map[Level]zap.AtomicLevel{
	LevelDebug: zap.NewAtomicLevelAt(zap.DebugLevel),
	LevelInfo:  zap.NewAtomicLevelAt(zap.InfoLevel),
	LevelWarn:  zap.NewAtomicLevelAt(zap.WarnLevel),
	LevelError: zap.NewAtomicLevelAt(zap.ErrorLevel),
	LevelFatal: zap.NewAtomicLevelAt(zap.FatalLevel),
}

func DefaulConfig() Config {
	return Config{
		Level: LevelDebug,
	}
}

func (c Config) Build(opts ...zap.Option) (*Logger, error) {
	zapConfig := zap.NewProductionConfig()
	if c.Env == "development" {
		zapConfig = zap.NewDevelopmentConfig()
	}

	zapConfig.Level = mapLevel[c.Level]
	zapLog, err := zapConfig.Build(opts...)
	if err != nil {
		return nil, err
	}

	return &Logger{zapLog}, nil
}
