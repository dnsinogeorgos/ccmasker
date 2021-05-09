# CC masker

This is a lightweight tool intended to be used as a filter with Rsyslog.  
It spawns a singlethreaded process and accepts log messages through stdin.  
Rsyslog will spawn more processes as needed, and expects to receive messages in the same
order.  
If the message contains a PAN number, it returns the message with a masked PAN in a JSON
key named "msg".  
Otherwise, an empty JSON is returned.  

For more information regarding the plugin nature of this tool:  
https://github.com/rsyslog/rsyslog/blob/master/plugins/external/INTERFACE.md#external-message-modification-modules  
https://github.com/rsyslog/rsyslog/blob/master/plugins/external/messagemod/anon_cc_nbrs/anon_cc_nbrs.py

### How to use
Add the following to your rsyslog config and restart.
More information here https://www.rsyslog.com/doc/master/configuration/modules/mmexternal.html
```
module(load="mmexternal")
action(type="mmexternal" binary="/path/to/ccmasker")
```

### Feedback
This has been a learning excercise, constructive feedback is always appreciated.

### False positives and rewrite
Filtering for PAN data without context is a process prone to false positives.  
Further steps to reduce false positives were required and it was a tricky process due to
variable length of matches.  

### ccmasker.py
Wrote an equivalent python script for comparison and it is actually faster ¯\_(ツ)_/¯  
A test with 748mb of real logs
```
timing ccmasker written in go (1.16)

real    0m34,526s
user    0m40,160s
sys     0m1,600s

timing ccmasker written in python (3.9)

real    0m28,051s
user    0m31,487s
sys     0m1,303s
```

### TODO
* experiment with benchmarks  
* write some tests
