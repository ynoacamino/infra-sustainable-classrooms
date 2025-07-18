// Code generated by goa v3.21.1, DO NOT EDIT.
//
// knowledge-http gRPC client CLI support package
//
// Command:
// $ goa gen
// github.com/ynoacamino/infra-sustainable-classrooms/services/knowledge/design/api
// -o ./services/knowledge/

package cli

import (
	"flag"
	"fmt"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return ``
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return ""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	cc *grpc.ClientConn,
	opts ...grpc.CallOption,
) (goa.Endpoint, any, error) {
	var ()

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}
