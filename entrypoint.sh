#!/bin/sh
set -eu

args="${INPUT_ARGS:-.}"

flags=""
[ "${INPUT_IGNORE_CHINESE:-false}" = "true" ] && flags="$flags --ignore-chinese"
[ "${INPUT_IGNORE_JAPANESE:-false}" = "true" ] && flags="$flags --ignore-japanese"
[ "${INPUT_IGNORE_KOREAN:-false}" = "true" ] && flags="$flags --ignore-korean"

# shellcheck disable=SC2086
exec nocjk $flags $args
