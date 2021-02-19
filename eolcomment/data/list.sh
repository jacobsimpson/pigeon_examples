#! /bin/bash
# A simple script to do nothing in particular.

echo "Show all running instances of eolcomment."
ps -ef | grep "[e]olcomment" # List eolprocesses.

if [[ ! $? -eq 0 ]]; then
    echo 'None found.'
fi
