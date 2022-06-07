package constant

import (
	"io/ioutil"
	"text/template"

	"github.com/sirupsen/logrus"

	"backend/utils/logger"
)

var IosProfileTemplate *template.Template

func init() {
	if bytes, err := ioutil.ReadFile("./resources/ios_profile.plist"); err != nil {
		logger.Log(logrus.Fatal, "UNABLE TO READ IOS PROFILE TEMPLATE FILE")
	} else {
		IosProfileTemplate = template.Must(template.New("ios_profile").Parse(string(bytes)))
	}
}
