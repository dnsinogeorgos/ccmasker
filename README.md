# CC masker

This is an experimental lightweight tool, intended to be used as a filter
with Rsyslog. It accepts log messages through stdin.
If the message contains a PAN number, it returns the message with a
masked PAN in a JSON key named "msg".
Otherwise an empty JSON is returned.

##### Possible improvements:  
matching with groups to eliminate false positives  
map of expressions through json config

For more information regarding the plugin nature of this tool  
https://github.com/rsyslog/rsyslog/blob/master/plugins/external/INTERFACE.md#external-message-modification-modules  
https://github.com/rsyslog/rsyslog/blob/master/plugins/external/messagemod/anon_cc_nbrs/anon_cc_nbrs.py

### Feedback
Constructive feedback is always appreciated.

### How to use

Add the following to your rsyslog config and restart  
More information here: https://www.rsyslog.com/doc/master/configuration/modules/mmexternal.html
```
module(load="mmexternal")
action(type="mmexternal" binary="/path/to/ccmasker")
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

### Python equivalent
Wrote an equivalent python script for comparison  
Took me 10 minutes to write and is almost as fast ¯\_(ツ)_/¯  
Update: upon testing with real logs, the python script was faster (34s to 27s for same dataset)  
Update 2: upon testing with rubex and jsoniter, go got down to 25s time. changes on faster-with-non-stdlib branch  

### Todo
##### write tests  
experiment with benchmarks, write some tests as well  
##### send messages to goroutines?  
**no**, rsyslog will spawn processes as needed  
