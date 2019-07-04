#! /bin/bash

echo "$(find -name '*.sh' -exec basename \ {} .sh \;  )" | cut -d "." -f2 | cut -d "/" -f2-