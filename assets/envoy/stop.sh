#!/bin/bash
a=$(ps -aux | grep envoy | grep -v grep | awk '{print $2}')
if [[ -n "${a}" ]];then
echo $a
    kill -9 $a
fi