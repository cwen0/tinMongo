package utils

import "github.com/Sirupsen/logrus"

var (
	TinMongoBuildTS = "None"
	TinMongoGitHash = "None"
)

func PrintTinMongoInfo() {
	logrus.Infof("Welcome to TinMongo.")
	logrus.Infof("Version:")
	logrus.Infof("Git Commit Hash: %s", TinMongoGitHash)
	logrus.Infof("UTC Build Time:  %s", TinMongoBuildTS)
}
