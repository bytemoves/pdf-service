package app

import (
	"log/slog"

	"github.com/bytesmoves/internal/handler"
	"github.com/bytesmoves/internal/logger"
	"github.com/bytesmoves/internal/service/core"
	"github.com/bytesmoves/internal/service/invoice"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/bytesmoves/internal/server"
	"github.com/gin-gonic/gin"
)

func Run() {
	fx.New(
	 fx.Provide(
	  logger.New,
	  invoice.NewService,
	  core.NewTemplateService,
	  core.NewPdfService,
	  handler.NewPdf,
	  handler.NewHealth,
	  server.New,
	 ),
	 fx.Invoke(func(e *gin.Engine) {
	  err := e.Run(":8080")
	  if err != nil {
	   return
	  }
	 }),
	 fx.WithLogger(func(log *slog.Logger) fxevent.Logger {
	  return &fxevent.SlogLogger{Logger: log}
	 }),
	).Run()
   }