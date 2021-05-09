#!/usr/bin/env python3
"""
This is a PAN number masking plugin for Rsyslog.
"""

import json
import re
import sys
import fast_luhn as fl


def compile_patterns():
    """
    Compiles regex patterns.
    Patterns base on https://en.wikipedia.org/wiki/Payment_card_number.
    Only VISA, Electron, Mastercard, Maestro, American Express and Diners are being
    implemented.

    :return: dict of patterns with values of type Pattern.
    """
    # fmt: off
    # pylint: disable=C0301
    regex_patterns = {
        "XXXX-VISA-XXXX": {
            "variable": "4[0-9]{3}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?([0-9]{4}|[0-9]{1})",  # noqa: E501
            "fixed": [
                "4[0-9]{3}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{1}",
                "4[0-9]{3}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}"
            ]
        },
        "XXXX-Master5xxx-XXXX": {
            "variable": "5[1-5]{1}[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}",  # noqa: E501
            "fixed": [
                "5[1-5]{1}[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}"
            ]
        },
        "XXXX-Maestro-XXXX": {
            "variable": "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{0,4}[ +-_]?[0-9]{0,3}",  # noqa: E501
            "fixed": [
                "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]",  # noqa: E501
                "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{1}",  # noqa: E501
                "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{2}",  # noqa: E501
                "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{3}",  # noqa: E501
                "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}",  # noqa: E501
                "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{1}",  # noqa: E501
                "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{2}",  # noqa: E501
                "(5018|5020|5038|5893|6304|6759|6761|6762|6763)[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{3}"  # noqa: E501
            ]
        },
        "XXXX-MaestroUK-XXXX": {
            "variable": "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]?[0-9]{0,4}[ +-_]?[0-9]{0,3}",  # noqa: E501
            "fixed": [
                "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]",
                "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]?[0-9]{1}",  # noqa: E501
                "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]?[0-9]{2}",  # noqa: E501
                "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]?[0-9]{3}",  # noqa: E501
                "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}",  # noqa: E501
                "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{1}",  # noqa: E501
                "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{2}",  # noqa: E501
                "(6767[ +-_]?70[0-9]{2}|6767[ +-_]?74[0-9]{2})[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{3}"  # noqa: E501
            ]
        },
        "XXXX-Master2xxx-XXXX": {
            "variable": "2[2-7]{1}[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}",  # noqa: E501
            "fixed": [
                "2[2-7]{1}[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}"
            ]
        },
        "XXXX-AmEx-XXXX": {
            "variable": "(34|37)[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{3}",
            "fixed": [
                "(34|37)[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{3}"
            ]
        },
        "XXXX-DinersUSC-XXXX": {
            "variable": "54[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}",
            "fixed": [
                "54[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}"
            ]
        },
        "XXXX-DinersInt-XXXX": {
            "variable": "36[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{2,4}[ +-_]?[0-9]{0,3}",  # noqa: E501
            "fixed": [
                "36[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{2}",
                "36[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{3}",
                "36[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}",
                "36[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{2,4}[ +-_]?[0-9]{1}",  # noqa: E501
                "36[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{2,4}[ +-_]?[0-9]{2}",  # noqa: E501
                "36[0-9]{2}[ +-_]?[0-9]{4}[ +-_]?[0-9]{4}[ +-_]?[0-9]{2,4}[ +-_]?[0-9]{3}"  # noqa: E501
            ]
        }
    }
    # fmt: on

    compiled_patterns = {}
    for group, patterns_dict in regex_patterns.items():
        compiled_patterns[group] = {}
        compiled_patterns[group]["variable"] = re.compile(patterns_dict["variable"])

        compiled_patterns[group]["fixed"] = []
        for pattern in patterns_dict["fixed"]:
            compiled_patterns[group]["fixed"].append(re.compile(pattern))

    return compiled_patterns


def process_message(msg, patterns, re_filter):
    """
    Processes a message from Rsyslog, returns appropriate json value.

    :param msg: original log message
    :param patterns: dict of compiled patterns
    :param re_filter: compiled pattern that matched non-digits
    :return: appropriate json value for Rsyslog
    """
    validated = False
    for group, patterns_dict in patterns.items():
        matched_outer = patterns_dict["variable"].search(msg)
        if matched_outer:
            for pattern in patterns_dict["fixed"]:
                matched_inner = pattern.search(msg)
                if matched_inner:
                    match_dirty = matched_inner.group()
                    match_clean = re_filter.sub("", match_dirty)
                    if fl.validate(match_clean):
                        validated = True
                        msg = pattern.sub(group, msg)

    if validated:
        return json.dumps({"msg": msg})
    return json.dumps({})


def __main__():
    patterns = compile_patterns()
    re_filter = re.compile("[^0-9]")
    stop = False
    while not stop:
        msg = sys.stdin.readline()
        if msg:
            msg = msg.rstrip("\n")
            print(process_message(msg, patterns, re_filter))
            sys.stdout.flush()
        else:
            stop = True
    sys.stdout.flush()


if __name__ == "__main__":
    __main__()
