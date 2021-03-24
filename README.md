# CC masker

This is an experimental lightweight tool, intended to be used as a filter
with Rsyslog. It accepts log messages through stdin.
If the message contains a PAN number, it returns the message with a
masked PAN in a JSON key named "msg".
Otherwise an empty JSON is returned.

For more information regarding the plugin nature of this tool  
https://github.com/rsyslog/rsyslog/blob/master/plugins/external/INTERFACE.md#external-message-modification-modules  
https://github.com/rsyslog/rsyslog/blob/master/plugins/external/messagemod/anon_cc_nbrs/anon_cc_nbrs.py

### Feedback
This was a learning experiment, constructive feedback is always appreciated.

### How to use

Add the following to your rsyslog config and restart  
More information here: https://www.rsyslog.com/doc/master/configuration/modules/mmexternal.html
```
module(load="mmexternal")
action(type="mmexternal" binary="/path/to/ccmasker")
```

### Docker compose
Build and let it run in the foreground in order to see the output
```
cd docker
go build .
docker-compose up
```

Rebuild the image with
```
docker-compose build
```

Send logs to the listening port  
```
logger -d -n localhost "this log message 5311111111111111 contains a PAN"
```

### ccmasker.py
Wrote an equivalent python script for comparison  
Took me 10 minutes to write and is actually faster ¯\_(ツ)_/¯  
A test with 92mb of real logs finished in 35 seconds

### TODO
##### write tests  
- experiment with benchmarks, write some tests as well  
##### send messages to goroutines?  
- **no**, rsyslog will spawn processes as needed
