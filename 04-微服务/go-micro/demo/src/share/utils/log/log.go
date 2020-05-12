package log

import (
	"bytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path"
	"time"
)

var instance *zap.Logger

// Instance 唯一实例
func Instance() *zap.Logger {
	return instance
}

// Init 初始化,srvName 生成的日志文件夹名字
func Init(srvName string) *zap.Logger {
	instance = NewLogger(srvName)
	return instance
}

// NewLogger 新建日志
func NewLogger(srvName string) *zap.Logger {
	// log和path文件只需要改这个路径即可
	directory := "D:\\go_work\\src\\go\\go-micro\\"
	if len(directory) == 0 {
		directory = path.Join("..", "log", srvName)
	} else {
		directory = path.Join(directory, "log", srvName)
	}
	writers := []zapcore.WriteSyncer{newRollingFile(directory)}
	writers = append(writers, os.Stdout)
	logger, _ := newZapLogger(true, zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(logger)
	return logger
}

func newRollingFile(directory string) zapcore.WriteSyncer {
	if err := os.MkdirAll(directory, 0766); err != nil {
		log.Println("failed create log directory:", directory, ":", err)
		return nil
	}

	return newLumberjackWriteSyncer(&lumberjack.Logger{
		Filename:  path.Join(directory, "output.log"),
		MaxSize:   100, //megabytes
		MaxAge:    7,   //days
		LocalTime: true,
		Compress:  false,
	})
}

func newZapLogger(isProduction bool, output zapcore.WriteSyncer) (*zap.Logger, *zap.AtomicLevel) {
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
	}

	var encoder zapcore.Encoder
	dyn := zap.NewAtomicLevel()
	if isProduction {
		dyn.SetLevel(zap.InfoLevel)
		encCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encCfg)
	} else {
		dyn.SetLevel(zap.DebugLevel)
		encCfg.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encCfg)
	}

	return zap.New(zapcore.NewCore(encoder, output, dyn), zap.AddCaller()), &dyn
}

type lumberjackWriteSyncer struct {
	*lumberjack.Logger
	buf       *bytes.Buffer
	logChan   chan []byte
	closeChan chan interface{}
	maxSize   int
}

func newLumberjackWriteSyncer(l *lumberjack.Logger) *lumberjackWriteSyncer {
	ws := &lumberjackWriteSyncer{
		Logger:    l,
		buf:       bytes.NewBuffer([]byte{}),
		logChan:   make(chan []byte, 5000),
		closeChan: make(chan interface{}),
		maxSize:   1024,
	}
	go ws.run()
	return ws
}

func (l *lumberjackWriteSyncer) run() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			if l.buf.Len() > 0 {
				l.sync()
			}
		case bs := <-l.logChan:
			_, err := l.buf.Write(bs)
			if err != nil {
				continue
			}
			if l.buf.Len() > l.maxSize {
				l.sync()
			}
		case <-l.closeChan:
			l.sync()
			return
		}
	}
}

func (l *lumberjackWriteSyncer) Stop() {
	close(l.closeChan)
}

func (l *lumberjackWriteSyncer) Write(bs []byte) (int, error) {
	b := make([]byte, len(bs))
	for i, c := range bs {
		b[i] = c
	}
	l.logChan <- b
	return 0, nil
}

func (l *lumberjackWriteSyncer) Sync() error {
	return nil
}

func (l *lumberjackWriteSyncer) sync() error {
	defer l.buf.Reset()
	_, err := l.Logger.Write(l.buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
