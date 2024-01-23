package app

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ucok-man/h8-p3-ngc/05-grpc/domain/user-service/internal/logging"
	"github.com/ucok-man/h8-p3-ngc/05-grpc/pb"
	"google.golang.org/grpc"
)

func (app *Application) Serve() error {
	srv := grpc.NewServer(grpc.ChainUnaryInterceptor())

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", app.config.Port))
	if err != nil {
		return err
	}
	done := make(chan struct{})

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		app.logger.Info("caught signal", logging.Meta{
			"signal": s.String(),
		})

		srv.GracefulStop()

		app.logger.Info("completing background tasks", logging.Meta{
			"addr": listener.Addr(),
		})

		app.wg.Wait()
		close(done)
	}()

	app.logger.Info("starting server", logging.Meta{
		"port": app.config.Port,
		"env":  app.config.Environment,
	})

	pb.RegisterUserServiceServer(srv, app)

	err = srv.Serve(listener)
	if err != nil {
		return err
	}

	<-done
	app.logger.Info("stopped server", logging.Meta{
		"addr": listener.Addr(),
	})

	return nil
}
