package dictApp

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "new_it/app/dictApp/DictcentApp",
		Iface: reflect.TypeOf((*DictcentApp)(nil)).Elem(),
		New:   func() any { return &dictcent_App{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return dictcentApp_local_stub{impl: impl.(DictcentApp), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return dictcentApp_client_stub{stub: stub, registerTablesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "new_it/app/dictApp/DictcentApp", Method: "RegisterTables"}), registerRouterMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "new_it/app/dictApp/DictcentApp", Method: "RegisterRouter"}), registerDBdataMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "new_it/app/dictApp/DictcentApp", Method: "RegisterDBdata"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return dictcentApp_server_stub{impl: impl.(DictcentApp), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type dictcentApp_local_stub struct {
	impl   DictcentApp
	tracer trace.Tracer
}

func (s dictcentApp_local_stub) RegisterTables(ctx context.Context) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "dictApp.DictcentApp.RegisterTables", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.RegisterTables(ctx)
}

func (s dictcentApp_local_stub) RegisterRouter(ctx context.Context) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "dictApp.DictcentApp.RegisterRouter", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.RegisterRouter(ctx)
}

func (s dictcentApp_local_stub) RegisterDBdata(ctx context.Context) (err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "dictApp.DictcentApp.RegisterDBdata", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.RegisterDBdata(ctx)
}

// Client stub implementations.

type dictcentApp_client_stub struct {
	stub                  codegen.Stub
	registerTablesMetrics *codegen.MethodMetrics
	registerRouterMetrics *codegen.MethodMetrics
	registerDBdataMetrics *codegen.MethodMetrics
}

func (s dictcentApp_client_stub) RegisterTables(ctx context.Context) (err error) {
	// Update metrics.
	start := time.Now()
	s.registerTablesMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "dictApp.DictcentApp.RegisterTables", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.registerTablesMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.registerTablesMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.registerTablesMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 2, nil, shardKey)
	if err != nil {
		return
	}
	s.registerTablesMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s dictcentApp_client_stub) RegisterRouter(ctx context.Context) (err error) {
	// Update metrics.
	start := time.Now()
	s.registerRouterMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "dictApp.DictcentApp.RegisterRouter", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.registerRouterMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.registerRouterMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.registerRouterMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 1, nil, shardKey)
	if err != nil {
		return
	}
	s.registerRouterMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s dictcentApp_client_stub) RegisterDBdata(ctx context.Context) (err error) {
	// Update metrics.
	start := time.Now()
	s.registerDBdataMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "dictApp.DictcentApp.RegisterDBdata", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.registerDBdataMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.registerDBdataMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.registerDBdataMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	if err != nil {
		return
	}
	s.registerDBdataMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

// Server stub implementations.

type dictcentApp_server_stub struct {
	impl    DictcentApp
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s dictcentApp_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "RegisterTables":
		return s.registerTables
	case "RegisterRouter":
		return s.registerRouter
	case "RegisterDBdata":
		return s.registerDBdata
	default:
		return nil
	}
}

func (s dictcentApp_server_stub) registerTables(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.RegisterTables(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s dictcentApp_server_stub) registerRouter(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.RegisterRouter(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s dictcentApp_server_stub) registerDBdata(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.RegisterDBdata(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}
