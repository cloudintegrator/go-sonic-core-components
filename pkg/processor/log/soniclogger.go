package main

import (
	"go.uber.org/zap"
	"log"
)

type SonicLogger struct {
	SonicChan chan interface{}
}

func (l *SonicLogger) Init(ch chan interface{}) {
	l.SonicChan = ch
}
func (l *SonicLogger) Process() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	defer logger.Sync()
	sugar.Info("Hello")
}

var INSTANCE SonicLogger
