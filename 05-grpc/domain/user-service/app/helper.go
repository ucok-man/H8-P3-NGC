package app

import (
	"fmt"
)

func (app *Application) background(fn func()) {
	app.wg.Add(1)

	go func() {

		defer app.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				app.logger.Error(fmt.Errorf("%s", err), "failed processing background task", nil)
			}
		}()

		fn()
	}()
}
