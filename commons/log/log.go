package log

import (
	"context"
	"os"
	"runtime/debug"

	"trendtracker/constants"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger is used to initialize logger
func InitLogger() {
	zerolog.ErrorStackMarshaler = getErrorStackMarshaller()
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func Trace(ctx context.Context) *zerolog.Event {
	return withValue(ctx, log.Trace())
}

// Debug is the for debug log
func Debug(ctx context.Context) *zerolog.Event {
	return withValue(ctx, log.Debug())
}

// Info is the for info log
func Info(ctx context.Context) *zerolog.Event {
	return withValue(ctx, log.Info())
}

// Warn is the for warn log
func Warn(ctx context.Context) *zerolog.Event {
	return withValue(ctx, log.Warn())
}

// Error is the for error log
func Error(ctx context.Context) *zerolog.Event {
	return withValue(ctx, log.Error().Stack())
}

// Panic is the for panic log
func Panic(ctx context.Context) *zerolog.Event {
	return withValue(ctx, log.Panic().Stack())
}

// Fatal is the for fatal log
func Fatal(ctx context.Context) *zerolog.Event {
	return withValue(ctx, log.Fatal().Stack())
}

func getErrorStackMarshaller() func(err error) interface{} {
	return func(err error) interface{} {
		return string(debug.Stack())
	}
}

func withValue(ctx context.Context, event *zerolog.Event) *zerolog.Event {
	if ctx == nil {
		return event
	}
	// path := ctx.Value(constants.PathLogParam)
	// if path != nil {
	// 	event.Interface(constants.PathLogParam, path)
	// }

	requestId := ctx.Value(constants.HeaderRequestID)
	if requestId != nil {
		event.Interface(constants.HeaderRequestID, requestId)
	}

	// ipAddress := ctx.Value(constants.ClientIPLogParam)
	// if ipAddress != nil {
	// 	event.Interface(constants.ClientIPLogParam, ipAddress)
	// }
	return event
}
