package zone

import (
	"fmt"
	"testing"
)

func TestMakeZone(t *testing.T) {
	tests := []struct {
		Board			[19][19]byte
		Zone			[19][19]byte
		radius			int
		threshold		int
	}{
		{
			Board:		[19][19]byte{},
			Zone:		[19][19]byte{},
			radius: 	2,
			threshold:	2,

		},{
			Board:		[19][19]byte{	[19]byte{5,0,0,0,0,0,0,0,0,5,0,0,0,0,0,0,0,0,6},},
			Zone:		[19][19]byte{	[19]byte{5,2,1,0,0,0,0,1,2,5,2,1,0,0,0,0,1,2,6},
										[19]byte{2,2,1,0,0,0,0,1,2,2,2,1,0,0,0,0,1,2,2},
										[19]byte{1,1,1,0,0,0,0,1,1,1,1,1,0,0,0,0,1,1,1},
							},
			radius:		2,
			threshold:	2,
		},{
			Board:		[19][19]byte{	[19]byte{5,0,0,5,0,0,0,0,0,5,0,0,0,0,0,0,0,0,6},},
			Zone:		[19][19]byte{	[19]byte{5,2,2,5,2,1,0,1,2,5,2,1,0,0,0,0,1,2,6},
										[19]byte{2,2,2,2,2,1,0,1,2,2,2,1,0,0,0,0,1,2,2},
										[19]byte{1,2,2,1,1,1,0,1,1,1,1,1,0,0,0,0,1,1,1},
							},
			radius:		2,
			threshold:	2,
		},{
			Board:		[19][19]byte{	[19]byte{5,6,0,5,0,0,0,0,0,5,0,0,0,0,0,0,0,0,6},},
			Zone:		[19][19]byte{	[19]byte{5,6,2,5,2,1,0,1,2,5,2,1,0,0,0,0,1,2,6},
										[19]byte{2,2,2,2,2,1,0,1,2,2,2,1,0,0,0,0,1,2,2},
										[19]byte{2,2,2,2,1,1,0,1,1,1,1,1,0,0,0,0,1,1,1},
							},
			radius:		2,
			threshold:	2,
		},{
			Board:		[19][19]byte{	[19]byte{5,0,0,5,0,0,0,0,0,6,0,0,0,0,0,0,0,0,5},},

			Zone:		[19][19]byte{	[19]byte{5,3,3,5,3,2,2,2,3,6,3,2,1,0,0,1,2,3,5},
										[19]byte{3,3,3,3,3,2,2,2,3,3,3,2,1,0,0,1,2,3,3},
										[19]byte{3,3,3,3,2,2,2,2,2,2,2,2,1,0,0,1,2,2,2},
										[19]byte{2,2,2,2,1,1,2,1,1,1,1,1,1,0,0,1,1,1,1},
							},
			radius:		3,
			threshold:	3,
		},{
			Board:		[19][19]byte{	[19]byte{5,0,0,5,0,0,0,0,0,6,0,0,0,0,0,0,0,0,5},},

			Zone:		[19][19]byte{	[19]byte{5,2,2,5,2,2,2,2,2,6,2,2,1,0,0,1,2,2,5},
										[19]byte{2,2,2,2,2,2,2,2,2,2,2,2,1,0,0,1,2,2,2},
										[19]byte{2,2,2,2,2,2,2,2,2,2,2,2,1,0,0,1,2,2,2},
										[19]byte{2,2,2,2,1,1,2,1,1,1,1,1,1,0,0,1,1,1,1},
							},
			radius:		3,
			threshold:	2,
		},
	}
	for _,test := range tests {
		zone := MakeZone(test.Board, test.radius, test.threshold)
		if zone != test.Zone {
			fmt.Println("Fail")
			for _, val := range zone {
				fmt.Println(val)
			}
			fmt.Println()
			for _, val := range test.Zone {
				fmt.Println(val)
			}
			t.Fatalf("zone: makeZone: wrong result")
		}
	}
}

func TestUpdateZone(t *testing.T) {
	tests := []struct {
		zoneBefore	[19][19]byte
		zoneAfter	[19][19]byte
		x			int
		y			int
		player		bool
		radius		int
		threshold	int
		err			error
		}{
		{
			zoneBefore: [19][19]byte{	[19]byte{5,2,1,0,0,0,0,1,2,5,2,1,0,0,0,0,1,2,6},
										[19]byte{2,2,1,0,0,0,0,1,2,2,2,1,0,0,0,0,1,2,2},
										[19]byte{1,1,1,0,0,0,0,1,1,1,1,1,0,0,0,0,1,1,1},},

			zoneAfter: [19][19]byte{	[19]byte{5,2,2,2,6,2,1,1,2,5,2,1,0,0,0,0,1,2,6},
										[19]byte{2,2,2,2,2,2,1,1,2,2,2,1,0,0,0,0,1,2,2},
										[19]byte{1,1,2,1,1,1,1,1,1,1,1,1,0,0,0,0,1,1,1},
							},
			x: 0,
			y: 4,
			player: true,
			err: nil,
			radius:		2,
			threshold:	2,
		},{
			zoneBefore: [19][19]byte{	[19]byte{5,2,2,1,0,0,1,2,2,5,2,2,1,0,0,1,2,2,6},
										[19]byte{2,2,2,1,0,0,1,2,2,2,2,2,1,0,0,1,2,2,2},
										[19]byte{2,2,2,1,0,0,1,2,2,2,2,2,1,0,0,1,2,2,2},
										[19]byte{1,1,1,1,0,0,1,1,1,1,1,1,1,0,0,1,1,1,1},},

			zoneAfter: [19][19]byte{	[19]byte{5,2,2,2,6,2,2,2,2,5,2,2,1,0,0,1,2,2,6},
										[19]byte{2,2,2,2,2,2,2,2,2,2,2,2,1,0,0,1,2,2,2},
										[19]byte{2,2,2,2,2,2,2,2,2,2,2,2,1,0,0,1,2,2,2},
										[19]byte{1,2,2,2,1,1,2,2,1,1,1,1,1,0,0,1,1,1,1},
							},
			x: 0,
			y: 4,
			player: true,
			err: nil,
			radius:		3,
			threshold:	2,
		},
	}
	for _,test := range tests {
		err := UpdateZone(&test.zoneBefore, test.x, test.y, test.player, test.radius, test.threshold)
		if err != test.err {
			t.Fatalf("zone: updateZone: wrong error")
		}
		if test.zoneBefore != test.zoneAfter {
			fmt.Println("Fail")
			for _, val := range test.zoneBefore {
				fmt.Println(val)
			}
			fmt.Println()
			for _, val := range test.zoneAfter {
				fmt.Println(val)
			}
			t.Fatalf("zone: updateZone: wrong zone")
		}
	}
}
