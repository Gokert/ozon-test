package httpResponse

import (
	"github.com/sirupsen/logrus"
	"time"
)

func SendLog(status int, query string, firstTime time.Time, log *logrus.Logger) {
	timeNow := time.Now()
	responseTime := timeNow.Sub(firstTime)

	log.Infof("Time: %v, Response time: %v, Status: %d, Query: %s", timeNow, responseTime, status, query)
}
