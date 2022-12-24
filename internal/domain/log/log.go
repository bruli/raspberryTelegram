package log

import "github.com/bruli/raspberryRainSensor/pkg/common/vo"

type Log struct {
	executedAt vo.Time
	seconds    int
	zoneName   string
}

func (l Log) ExecutedAt() vo.Time {
	return l.executedAt
}

func (l Log) Seconds() int {
	return l.seconds
}

func (l Log) ZoneName() string {
	return l.zoneName
}

func (l *Log) Hydrate(executedAt vo.Time, seconds int, zoneName string) {
	l.executedAt = executedAt
	l.seconds = seconds
	l.zoneName = zoneName
}
