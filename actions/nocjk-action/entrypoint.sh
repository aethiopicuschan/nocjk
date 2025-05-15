#!/bin/bash
set -e

ARGS=""

if [[ "${INPUT_IGNORE_CHINESE}" == "true" ]]; then
  ARGS+=" --ignore-chinese"
fi

if [[ "${INPUT_IGNORE_JAPANESE}" == "true" ]]; then
  ARGS+=" --ignore-japanese"
fi

if [[ "${INPUT_IGNORE_KOREAN}" == "true" ]]; then
  ARGS+=" --ignore-korean"
fi

# Add additional user-specified arguments
ARGS+=" ${INPUT_ARGS}"

# Run the nocjk command with the constructed args
exec nocjk ${ARGS}
