#!/usr/bin/env python3

import sys
import re
import json

def compileFilters():
	s = " +\-_"
	patterns = {
		"XXXX-VISA-XXXX":       "4[0-9]{3}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?([0-9]{4}|[0-9]{1})",
		"XXXX-Master5xxx-XXXX": "5[1-5]{1}[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}",
		"XXXX-Maestro-XXXX":    "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{0,4}[" + s + "]?[0-9]{0,3}",
		"XXXX-MaestroUK-XXXX":  "(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{0,4}[" + s + "]?[0-9]{0,3}",
		"XXXX-Master2xxx-XXXX": "2[2-7]{1}[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}",
		"XXXX-AmEx-XXXX":       "(34|37)[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}",
		"XXXX-DinersUSC-XXXX":  "54[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}",
		"XXXX-DinersInt-XXXX":  "36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2,4}[" + s + "]?[0-9]{0,3}",
	}
	filters = {}
	for mask, pattern in patterns.items():
		filters[mask] = re.compile(pattern)

	return filters

def filterMessage(msg, filters):
	for mask, filter in filters.items():
		if filter.search(msg):
			return json.dumps({'msg': filter.sub(mask, msg)})
	return json.dumps({})

def __main__():
	filters = compileFilters()
	stop = False
	while not stop:
		msg = sys.stdin.readline()
		if msg:
			msg = msg.rstrip("\n")
			print(filterMessage(msg, filters))
			sys.stdout.flush()
		else:
			stop = True
	sys.stdout.flush()

if __name__ == "__main__":
	__main__()
