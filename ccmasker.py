#!/usr/bin/env python3
"""
This is a PAN number masking plugin for Rsyslog.
"""

import sys
import re
import json
import fast_luhn as fl


def compile_patterns(s):  # pylint: disable=C0103
    """
    Compiles regex patterns.

    :param separators: string of characters used as separators
    :return: dict of patterns with values of type Pattern.
    """
    regex_patterns = {
        "XXXX-VISA-XXXX": "4[0-9]{3}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?([0-9]{4}|[0-9]{1})",  # pylint: disable=C0301  # noqa: E501
        "XXXX-Master5xxx-XXXX": "5[1-5]{1}[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}",  # pylint: disable=C0301  # noqa: E501
        "XXXX-Maestro-XXXX": "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{0,4}[" + s + "]?[0-9]{0,3}",  # pylint: disable=C0301  # noqa: E501
        "XXXX-MaestroUK-XXXX": "(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{0,4}[" + s + "]?[0-9]{0,3}",  # pylint: disable=C0301  # noqa: E501
        "XXXX-Master2xxx-XXXX": "2[2-7]{1}[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}",  # pylint: disable=C0301  # noqa: E501
        "XXXX-AmEx-XXXX": "(34|37)[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}",  # pylint: disable=C0301  # noqa: E501
        "XXXX-DinersUSC-XXXX": "54[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}",  # pylint: disable=C0301  # noqa: E501
        "XXXX-DinersInt-XXXX": "36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2,4}[" + s + "]?[0-9]{0,3}",  # pylint: disable=C0301  # noqa: E501
    }
    compiled_patterns = {}
    for mask, pattern in regex_patterns.items():
        compiled_patterns[mask] = re.compile(pattern)

    return compiled_patterns


def process_message(msg, patterns, separators):
    """
    Processes a message from Rsyslog, returns appropriate json value.

    :param msg: original log message
    :param patterns: dict of compiled patterns
    :param separators: string of characters used as separators
    :return: appropriate json value for Rsyslog
    """
    matched = False
    for mask, pattern in patterns.items():
        match = pattern.search(msg)
        if match:
            match = match.group().translate({ord(i): None for i in separators})
            if fl.validate(match):
                matched = True
                msg = pattern.sub(mask, msg)

    if matched:
        return json.dumps({"msg": msg})
    return json.dumps({})


def __main__():
    separators = " +-_"
    filters = compile_patterns(separators)
    stop = False
    while not stop:
        msg = sys.stdin.readline()
        if msg:
            msg = msg.rstrip("\n")
            print(process_message(msg, filters, separators))
            sys.stdout.flush()
        else:
            stop = True
    sys.stdout.flush()


if __name__ == "__main__":
    __main__()
