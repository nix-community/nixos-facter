package hwinfo

// original list taken from systemd.network
var virtualNetworkDevices = map[string]bool {
	"bonding": true,
	"bridge": true,
	"dummy": true,
	"ip_gre": true,
	"ip6_gre": true,
	"ipip": true,
	"ipvlan": true,
	"macvlan": true,
	"sit": true,
	"tun": true,
	"veth": true,
	"8021q": true,
	"ip_vti": true,
	"ip6_vti": true,
	"vxlan": true,
	"geneve": true,
	"macsec": true,
	"vrf": true,
	"vcan": true,
	"vxcan": true,
	"wireguard": true,
	"nlmon": true,
	"fou": true,
	"ifb": true,
	"bareudp": true,
	"batman-adv": true,
	"ib_ipoib": true,
	// FIXME: could not confirm the following netdev drivers:
	// "l2tp": true,
	// "xfrm": true,
	// "erspan": true,
	// "ip6tnl": true,
}

