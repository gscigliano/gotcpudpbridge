# gotcpudpbridge
very rudimentary traffic bridge with zero error checking

can bridge tcp into udp and viceversa
endpoints can be a mix of ipv4/ipv6

EXAMPLES:

**Usage: gotcpudpbridge TYPE SRC_SOCKET DST_SOCKET UDP_BUF_SIZE TCP_BUF_SIZE

# bridge udp6 into tcp4 
_(v6 addresses have to be enclosed in [])_

./gotcpudpbridge udp2tcp [::1]:161 127.0.0.1:9000 1024 1024

# bridge udp into tcp

./gotcpudpbridge udp2tcp 127.0.0.1:161 127.0.0.1:9000 1024 1024

# bridge tcp into udp

./gotcpudpbridge tcp2udp 127.0.0.1:65003 switch:161 1024 1024
