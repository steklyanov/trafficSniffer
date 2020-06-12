
You should create a go application (go version >= 1.13) that will do the following:

Run on a 64-bit Linux distribution (Centos, Ubuntu, Debian).  
Sniff tcp/ip packets.  
Detect among the sniffed packets detect SSL (https) handshake packets.  
Print to stdout each detection in the following format: IP_SRC,TCP_SRC,IP_DST,TCP_DST,COUNT(TCP_OPTIONS).  