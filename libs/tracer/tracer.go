package tracer

import "github.com/uber/jaeger-client-go/config"

func InitGlobal(service string) error {
	cfg, err := config.FromEnv()
	if err != nil {
		return err
	}

	if _, err := cfg.InitGlobalTracer(service); err != nil {
		return err
	}
	return nil
}
