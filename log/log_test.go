package log

import (
	"errors"
	"testing"
)

/*func TestLogger_printf(t *testing.T) {

	apexLog.SetHandler(cli.Default)
	apexLog.SetLevel(apexLog.InfoLevel)
	apexLog.SetHandler(json.New(os.Stderr))

	ctx := apexLog.WithFields(apexLog.Fields{
		"file": "something.png",
		"type": "image/png",
	})

	for range time.Tick(time.Millisecond * 200) {
		ctx.Info("upload")
		ctx.Info("upload complete")
		ctx.Debug("msg dbg")
		ctx.Warn("upload retry")
		ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	}

}*/

/*func _TestLoggger(t *testing.T) {

	ctx := Logger("log_test.go")

	testVar := "VarString"

	for range time.Tick(time.Millisecond * 200) {
		ctx.Info("upload :" + testVar)
		ctx.Info("upload complete")
		ctx.Debug("msg dbg")
		ctx.Warn("upload retry")
		ctx.WithError(errors.New("unauthorized")).Error("upload failed")
	}

}*/

func TestLogggerJson(t *testing.T) {

	ctx := LoggerJSON().WithField("user", "Tobi")

	testVar := "VarString"

	ctx.WithField("hola", 2).Info("upload :" + testVar)
	ctx.Info("upload complete")
	ctx.Debug("msg dbg")
	ctx.Warn("upload retry")
	ctx.WithError(errors.New("unauthorized")).Error("upload failed")

}
