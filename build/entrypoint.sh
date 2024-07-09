#!/bin/bash

if [ "$1" = "" ]; then
	/bin/bash
else
	exec "$@"
fi
