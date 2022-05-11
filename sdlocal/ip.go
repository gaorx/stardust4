package sdlocal

import (
	"net"

	"github.com/gaorx/stardust4/sderr"
)

type IPPredicate func(net.Interface, net.IP) bool

func (p IPPredicate) Not() IPPredicate {
	return func(iface net.Interface, ip net.IP) bool {
		return !p(iface, ip)
	}
}

func IPs(predicates ...IPPredicate) ([]net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, sderr.Wrap(err, "get net interfaces error")
	}
	var ips []net.IP
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ip := extractIP(iface, addr)
			if len(ip) > 0 && predicateIP(iface, ip, predicates) {
				ips = append(ips, ip)
			}
		}
	}
	return ips, nil
}

func IP(predicates ...IPPredicate) (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, sderr.Wrap(err, "get net interfaces error")
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ip := extractIP(iface, addr)
			if len(ip) > 0 && predicateIP(iface, ip, predicates) {
				return ip, nil
			}
		}
	}
	return nil, sderr.New("not found ip")
}

func IPString(predicates ...IPPredicate) string {
	ip, err := IP(predicates...)
	if err != nil || len(ip) <= 0 {
		return ""
	}
	return ip.String()
}

func PrivateIP4String(ifaceNames ...string) string {
	if len(ifaceNames) > 0 {
		return IPString(Is4(), IsPrivate(), IfaceNameIn(ifaceNames...))
	} else {
		return IPString(Is4(), IsPrivate())
	}
}

// IP predicates

func Is4() IPPredicate {
	return func(_ net.Interface, ip net.IP) bool {
		ip4 := ip.To4()
		return len(ip4) > 0
	}
}

func IfaceNameIs(ifaceName string) IPPredicate {
	return func(iface net.Interface, ip net.IP) bool {
		return iface.Name == ifaceName
	}
}

func IfaceNameIn(ifaceNames ...string) IPPredicate {
	return func(iface net.Interface, ip net.IP) bool {
		for _, ifaceName := range ifaceNames {
			if iface.Name == ifaceName {
				return true
			}
		}
		return false
	}
}

func IsLoopback() IPPredicate {
	return func(_ net.Interface, ip net.IP) bool {
		return ip.IsLoopback()
	}
}

func IsPrivate() IPPredicate {
	return func(_ net.Interface, ip net.IP) bool {
		return ip.IsPrivate()
	}
}

func IsMulticast() IPPredicate {
	return func(_ net.Interface, ip net.IP) bool {
		return ip.IsMulticast()
	}
}

// helpers

func predicateIP(iface net.Interface, ip net.IP, predicates []IPPredicate) bool {
	ok := true
	for _, pred := range predicates {
		if pred != nil && !pred(iface, ip) {
			ok = false
			break
		}
	}
	return ok
}

func extractIP(_ net.Interface, addr net.Addr) net.IP {
	switch v := addr.(type) {
	case *net.IPNet:
		return v.IP
	case *net.IPAddr:
		return v.IP
	default:
		return nil
	}
}