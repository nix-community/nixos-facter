package ephem_test

import (
	"strings"
	"testing"

	"github.com/numtide/nixos-facter/pkg/ephem"
	"github.com/stretchr/testify/require"
)

var (
	empty   = "Filename                                Type            Size    Used    Priority\n"
	corrupt = `Filename                                Type            Size    Used    Priority
foo bar baz
`
	badPartition = `Filename				Type		Size	Used	Priority
/var/lib/swap-1							foo        1048576	123		-3
`

	sample = `Filename				Type		Size	Used	Priority
/dev/sda6                               partition   4194300	0		-1
/var/lib/swap-1							file        1048576	123		-3
/var/lib/swap-2 						file        2097152	4567	-2`
)

func TestReadSwapFile(t *testing.T) {
	t.Parallel()
	as := require.New(t)

	_, err := ephem.ReadSwapFile(strings.NewReader(""))
	as.Error(err, "swaps file is empty")

	_, err = ephem.ReadSwapFile(strings.NewReader("foo bar baz hello world\n"))
	as.Errorf(err, "header in swaps file is malformed: '%s'", "foo bar baz hello world\n")

	swaps, err := ephem.ReadSwapFile(strings.NewReader(empty))
	as.NoError(err)
	as.Empty(swaps)

	_, err = ephem.ReadSwapFile(strings.NewReader(corrupt))
	as.Errorf(err, "malformed entry in swaps file: '%s'", "foo bar baz")

	_, err = ephem.ReadSwapFile(strings.NewReader(badPartition))
	as.Errorf(err, "malformed entry in swaps file: '%s'", `/var/lib/swap-1							foo        1048576	123		-3`)

	swaps, err = ephem.ReadSwapFile(strings.NewReader(sample))
	as.NoError(err)
	as.Len(swaps, 3)

	as.Equal("/dev/sda6", swaps[0].Filename)
	as.Equal(ephem.SwapTypePartition, swaps[0].Type)
	as.Equal(uint64(4194300), swaps[0].Size)
	as.Equal(uint64(0), swaps[0].Used)
	as.Equal(int32(-1), swaps[0].Priority)

	as.Equal("/var/lib/swap-1", swaps[1].Filename)
	as.Equal(ephem.SwapTypeFile, swaps[1].Type)
	as.Equal(uint64(1048576), swaps[1].Size)
	as.Equal(uint64(123), swaps[1].Used)
	as.Equal(int32(-3), swaps[1].Priority)

	as.Equal("/var/lib/swap-2", swaps[2].Filename)
	as.Equal(ephem.SwapTypeFile, swaps[2].Type)
	as.Equal(uint64(2097152), swaps[2].Size)
	as.Equal(uint64(4567), swaps[2].Used)
	as.Equal(int32(-2), swaps[2].Priority)
}
