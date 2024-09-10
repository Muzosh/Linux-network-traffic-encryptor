package utils

import (
	"log"

	"github.com/songgao/water"
	"github.com/vishvananda/netlink"
)

func EnsureTun(gw string) (*water.Interface, error) {
	if tun, err := netlink.LinkByName("tun99"); err == nil {
		err = netlink.LinkDel(tun)
		if err != nil {
			return nil, err
		}
	}

	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
		PlatformSpecificParams: water.PlatformSpecificParams{
			Name: "tun99",
		},
	})
	if err != nil {
		return nil, err
	}

	tun, err := netlink.LinkByName("tun99")
	if err != nil {
		return nil, err
	}

	addr1, err := netlink.ParseAddr("192.168.1.1/31")
	if err != nil {
		log.Fatal(err)
	}

	if err = netlink.AddrAdd(tun, addr1); err != nil {
		return nil, err
	}

	if err = netlink.LinkSetUp(tun); err != nil {
		return nil, err
	}

	addr2, err := netlink.ParseAddr("192.168.1.2/31")
	if err != nil {
		return nil, err
	}

	gwAddr, err := netlink.ParseIPNet(gw)
	if err != nil {
		return nil, err
	}

	if err = netlink.RouteAdd(&netlink.Route{
		LinkIndex: tun.Attrs().Index,
		Dst: gwAddr,
		Via: &netlink.Via{
			AddrFamily: netlink.FAMILY_V4,
			Addr:       addr2.IP,
		},
	}); err != nil {
		return nil, err
	}

	return ifce, nil
}
