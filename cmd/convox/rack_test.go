package main

import (
	"fmt"
	"testing"

	"github.com/convox/rack/client"
	"github.com/convox/rack/test"
	"github.com/convox/version"
	"github.com/stretchr/testify/require"
)

func TestRackUpdateStable(t *testing.T) {
	versions, err := version.All()
	require.Nil(t, err)

	stable, err := versions.Resolve("stable")
	require.Nil(t, err)

	ts := testServer(t,
		test.Http{Method: "PUT", Body: fmt.Sprintf("version=%s", stable.Version), Path: "/system", Code: 200, Response: client.System{
			Name:    "mysystem",
			Version: "ver",
			Count:   1,
			Type:    "type",
		}},
	)

	defer ts.Close()

	test.Runs(t,
		test.ExecRun{
			Command: "convox rack update",
			Exit:    0,
			Stdout:  fmt.Sprintf("Updating to %s... UPDATING\n", stable.Version),
		},
	)
}

func TestRackUpdateSpecified(t *testing.T) {
	ts := testServer(t,
		test.Http{Method: "PUT", Body: "version=20150909014908", Path: "/system", Code: 200, Response: client.System{
			Name:    "mysystem",
			Version: "ver",
			Count:   1,
			Type:    "type",
		}},
	)

	defer ts.Close()

	test.Runs(t,
		test.ExecRun{
			Command: "convox rack update 20150909014908",
			Exit:    0,
			Stdout:  "Updating to 20150909014908... UPDATING\n",
		},
	)
}

func TestRackUpdateWait(t *testing.T) {
	ts := testServer(t,
		test.Http{Method: "PUT", Body: "version=20150909014908", Path: "/system", Code: 200, Response: client.System{
			Name:    "mysystem",
			Version: "ver",
			Count:   1,
			Type:    "type",
		}},
		test.Http{Method: "GET", Path: "/system", Code: 200, Response: client.System{
			Name:    "mysystem",
			Status:  "running",
			Version: "ver",
			Count:   1,
			Type:    "type",
		}},
	)

	defer ts.Close()

	test.Runs(t,
		test.ExecRun{
			Command: "convox rack update 20150909014908 --wait",
			Exit:    0,
			Stdout:  "Updating to 20150909014908... UPDATING\nWaiting for completion... OK\n",
		},
	)
}
