# CC masker

This is a lightweight tool that is intended to be used as a filter with
Rsyslog. It accepts log messages through stdin.
If the message contains a PAN number, it returns the message with a
masked PAN in a JSON key named "msg".
Otherwise an empty JSON is returned.

Possible improvements:  
matching with groups to eliminate false positives  
map of expressions through json config

For more regarding the plugin nature of this tool  
https://github.com/rsyslog/rsyslog/blob/master/plugins/external/INTERFACE.md#external-message-modification-modules  
https://github.com/rsyslog/rsyslog/blob/master/plugins/external/messagemod/anon_cc_nbrs/anon_cc_nbrs.py

### Feedback
Constructive feedback is always very much appreciated.

### How to use

Add the following to your rsyslog config and restart
```
module(load="mmexternal")
action(type="mmexternal" binary="/path/to/ccmasker" interface.input="msg")
```

### Docker compose
Build with
```
go build .
```

Let it run in the foreground in order to see the output
```
docker-compose up
```

Rebuild the image with
```
docker-compose build
```

Send logs to the listening port  
I am assuming you work on a linux environment, if not, look for an alternative
```
logger -d -n localhost "this log message 5311111111111111 contains a PAN"
```
