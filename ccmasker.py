#!/usr/bin/env python3

import sys
import re
import json


def compile_filters():
    patterns = {
        "XXXX-VISA-XXXX":       "4[0-9]{3}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?([0-9]{4}|[0-9]{1})",  # noqa: E501
        "XXXX-Master5xxx-XXXX": "5[1-5]{1}[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}",  # noqa: E501
        "XXXX-Maestro-XXXX":    "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{0,4}[ +-_]?[0-9]{0,3}",  # noqa: E501
        "XXXX-MaestroUK-XXXX":  "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]?[0-9]{0,4}[ +-_]?[0-9]{0,3}",  # noqa: E501
        "XXXX-Master2xxx-XXXX": "2[2-7]{1}[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}",  # noqa: E501
        "XXXX-AmEx-XXXX":       "(34|37)[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{3}",  # noqa: E501
        "XXXX-DinersUSC-XXXX":  "54[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}",  # noqa: E501
        "XXXX-DinersInt-XXXX":  "36[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{2,4}[ +-_]?[0-9]{0,3}",  # noqa: E501
    }
    filters = {}
    for mask, pattern in patterns.items():
        filters[mask] = re.compile(pattern)

    return filters


def process_message(msg, filters):
    matched = False
    for mask, filter in filters.items():
        if filter.search(msg):
            matched = True
            msg = filter.sub(mask, msg)

    if matched:
        return json.dumps({'msg': msg})
    return json.dumps({})


def __main__():
    filters = compile_filters()
    stop = False
    while not stop:
        msg = sys.stdin.readline()
        if msg:
            msg = msg.rstrip("\n")
            print(process_message(msg, filters))
            sys.stdout.flush()
        else:
            stop = True
    sys.stdout.flush()


if __name__ == "__main__":
    __main__()
