#!/bin/bash

for x in $(cat /etc/environment); do export $x; done
/go/bin/ewp-rest