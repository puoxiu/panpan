package init_db

import (
	"context"
	"strings"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaWriter struct {
	Pusher *kq.Pusher
}

func NewKafkaWriter(pusher *kq.Pusher) *KafkaWriter {
	return &KafkaWriter{
		Pusher: pusher,
	}
}

func (w *KafkaWriter) Write(p []byte) (n int, err error) {
	// writing log with newlines, trim them.
	if err := w.Pusher.Push(context.Background(), strings.TrimSpace(string(p))); err != nil {
		return 0, err
	}

	return len(p), nil
}

func LogxKafka() *logx.Writer {
	pusher := kq.NewPusher([]string{"localhost:9092"}, "panpan-log")
	defer pusher.Close()

	writer := logx.NewWriter(NewKafkaWriter(pusher))
	return &writer
}
