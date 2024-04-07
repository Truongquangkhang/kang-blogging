package logs

import (
	"github.com/sirupsen/logrus"
)

func LogCommandExecution(commandName string, cmd interface{}, err error) {
	log := logrus.WithField("cmd", cmd)

	if err == nil {
		log.Info(commandName + " command succeeded")
	} else {
		log.WithError(err).Error(commandName + " command failed")
	}
}

func LogQueryExecution(queryName string, query interface{}, err error) {
	log := logrus.WithField("query", query)

	if err == nil {
		log.Info(queryName + " query succeeded")
	} else {
		log.WithError(err).Error(queryName + " query failed")
	}
}

func LogUsecaseExecution(usecaseName string, usecase interface{}, err error) {
	log := logrus.WithField("usecase", usecase)

	if err == nil {
		log.Info(usecaseName + " usecase succeeded")
	} else {
		log.WithError(err).Error(usecaseName + " usecase failed")
	}
}
