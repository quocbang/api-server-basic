package logging

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// customWriter definition.
type CustomResponseWriter struct {
	echo.Response
	Data []byte
}

func (c *CustomResponseWriter) Write(data []byte) (int, error) {
	c.Data = append(c.Data, data...)
	return c.Writer.Write(data)
}

func CustomLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			reqLoggerFields := []zap.Field{}
			logger.Info("start request", reqLoggerFields...)
			startTime := time.Now()

			// custom data.
			customWriter := &CustomResponseWriter{
				Response: *ctx.Response(),
				Data:     []byte{},
			}

			// replace the writer default to writer custom.
			ctx.Response().Writer = customWriter

			// next to next into handle method.
			if err = next(ctx); err != nil {
				ctx.Error(err)
			}

			// Calculate the response time
			responseTime := time.Since(startTime)

			// Log the response information
			statusCode := ctx.Response().Status
			requestMethod := ctx.Request().Method
			requestPath := ctx.Request().URL.Path

			responseFields := []zap.Field{
				zap.Int("status", statusCode),
				zap.String("method", requestMethod),
				zap.String("path", requestPath),
				zap.Duration("responseTime", responseTime),
				zap.String("response", string(customWriter.Data)),
			}

			loggingWithLevel(statusCode, responseFields, logger)
			return
		}
	}
}

func loggingWithLevel(statusCode int, responseFields []zap.Field, logger *zap.Logger) {
	switch {
	case statusCode >= http.StatusInternalServerError:
		logger.Error("Response", responseFields...)
	case statusCode >= http.StatusBadRequest:
		logger.Warn("Response", responseFields...)
	default:
		logger.Info("Response", responseFields...)
	}
}

func InitializeLogger(devMode bool) (logger *zap.Logger, err error) {
	if devMode {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	}
	return logger, nil
}
