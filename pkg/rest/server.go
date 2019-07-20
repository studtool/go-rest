package rest

import (
	"github.com/valyala/fasthttp"

	"github.com/studtool/go-errs/pkg/errs"
	"github.com/studtool/go-logs/pkg/logs"
)

type Server struct {
	address string
	server  *fasthttp.Server
	logs    LogPack
}

type ServerParams struct {
	Address      string
	RouterConfig RouterConfig
	LogPack      LogPack
}

type LogPack struct {
	StructLogger     *logs.StructLogger
	StackTraceLogger *logs.StackTraceLogger
	RequestLogger    *logs.RequestLogger
}

func NewServer(params ServerParams) *Server {
	router := NewRouter(
		RouterParams{
			logs:   &params.LogPack,
			config: params.RouterConfig,
		},
	)

	backend := &fasthttp.Server{
		Handler: router.handler,
	}

	return &Server{
		server:  backend,
		address: params.Address,
		logs:    params.LogPack,
	}
}

func (srv *Server) Run() *errs.Error {
	srv.logs.StructLogger.
		Infof("started [address = '%s']", srv.address)

	go func() {
		e := srv.server.
			ListenAndServe(srv.address)
		if e != nil {
			panic(e)
		}
	}()

	return nil
}

func (srv *Server) Shutdown() *errs.Error {
	srv.logs.StructLogger.
		Info("stopped")

	err := srv.server.
		Shutdown()
	if err != nil {
		return errs.Wrap(err)
	}

	return nil
}
