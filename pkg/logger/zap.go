package logger

import (
	"fmt"
	"go/types"
	"time"

	"github.com/gramilul123/test-makves/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	Log *zap.Logger
}

func NewZapLogger(cfg *config.Config) (*ZapLogger, error) {
	var (
		logger *zap.Logger
		err    error
	)

	if cfg.IsProduction() && !cfg.Debug { // production without debug -- silent mode
		logger = zap.NewNop().WithOptions()
	} else {
		switch cfg.Environment {
		case config.DefEnvProd:
			logger, err = zap.NewProduction()
			if err != nil {

				return nil, err
			}
		case config.DefEnvDev:
			logger, err = zap.NewDevelopment()
			if err != nil {

				return nil, err
			}
		default:
			logger = zap.NewExample()
			logger.Info(fmt.Sprintf("new app environment: %s", cfg.Environment))
		}
	}

	return &ZapLogger{Log: logger}, nil
}

func (t *ZapLogger) Sync() error {
	return t.Log.Sync()
}

func (t *ZapLogger) Debug(msg string, args ...interface{}) {
	fields := ZapConvertArgs(args)
	t.Log.Debug(msg, fields...)
}

func (t *ZapLogger) Info(msg string, args ...interface{}) {
	fields := ZapConvertArgs(args)
	t.Log.Info(msg, fields...)
}

func (t *ZapLogger) Warn(msg string, args ...interface{}) {
	fields := ZapConvertArgs(args)
	t.Log.Warn(msg, fields...)
}

func (t *ZapLogger) Error(msg string, args ...interface{}) {
	fields := ZapConvertArgs(args)
	t.Log.Error(msg, fields...)
}

func (t *ZapLogger) Fatal(msg string, args ...interface{}) {
	fields := ZapConvertArgs(args)
	t.Log.Fatal(msg, fields...)
}

func ZapConvertArgs(args []interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(args))
	for _, f := range args {
		field, ok := f.(zap.Field)
		if !ok {
			var zapField zapcore.Field

			switch v := f.(type) {
			case error:
				zapField.Interface = v
				zapField.Key = "error"
				zapField.Type = zapcore.ErrorType
			case types.Slice:
				zapField.Interface = v
				zapField.Key = "array"
				zapField.Type = zapcore.ArrayMarshalerType
			case types.Object:
				zapField.Interface = v
				zapField.Key = "object"
				zapField.Type = zapcore.ObjectMarshalerType
			case []byte:
				zapField.Interface = v
				zapField.Key = "binary"
				zapField.Type = zapcore.BinaryType
			case string:
				zapField.String = v
				zapField.Key = "string"
				zapField.Type = zapcore.StringType
			case bool:
				var fb int64
				if v {
					fb = 1
				} else {
					fb = 0
				}
				zapField.Integer = fb
				zapField.Key = "bool"
				zapField.Type = zapcore.BoolType
			case complex128:
				zapField.Interface = v
				zapField.Key = "complex128"
				zapField.Type = zapcore.Complex128Type
			case complex64:
				zapField.Interface = v
				zapField.Key = "complex64"
				zapField.Type = zapcore.Complex64Type
			case float64:
				zapField.Integer = int64(v) // zap хранит float в этом поле!!!
				zapField.Key = "float64"
				zapField.Type = zapcore.Float64Type
			case float32:
				zapField.Integer = int64(v)
				zapField.Key = "float32"
				zapField.Type = zapcore.Float32Type
			case int64:
				zapField.Integer = v
				zapField.Key = "int64"
				zapField.Type = zapcore.Int64Type
			case int32:
				zapField.Integer = int64(v)
				zapField.Key = "int32"
				zapField.Type = zapcore.Int32Type
			case int16:
				zapField.Integer = int64(v)
				zapField.Key = "int16"
				zapField.Type = zapcore.Int16Type
			case int8:
				zapField.Integer = int64(v)
				zapField.Key = "int8"
				zapField.Type = zapcore.Int8Type
			case uint64:
				zapField.Integer = int64(v)
				zapField.Key = "uint64"
				zapField.Type = zapcore.Uint64Type
			case uint32:
				zapField.Integer = int64(v)
				zapField.Key = "uint32"
				zapField.Type = zapcore.Uint32Type
			case uint16:
				zapField.Integer = int64(v)
				zapField.Key = "uint16"
				zapField.Type = zapcore.Uint16Type
			case uint8:
				zapField.Integer = int64(v)
				zapField.Key = "uint8"
				zapField.Type = zapcore.Uint8Type
			case uintptr:
				zapField.Integer = int64(v)
				zapField.Key = "unsafe ptr"
				zapField.Type = zapcore.UintptrType
			case time.Time:
				zapField.Interface = v
				zapField.Key = "time"
				zapField.Type = zapcore.TimeType
			case time.Duration:
				zapField.Integer = int64(v)
				zapField.Key = "duration"
				zapField.Type = zapcore.DurationType
			}
			fields = append(fields, zapField)
		} else {
			fields = append(fields, field)
		}
	}
	return fields
}
