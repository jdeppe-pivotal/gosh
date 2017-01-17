#!/usr/bin/env python3
import os
import subprocess


pid = os.fork()
if pid == 0:  # child
    pid2 = os.fork()
    if pid2 != 0:  # parent
        print('The zombie pid will be: {}'.format(pid2))
else:  # parent
    os.waitpid(pid, 0)
    subprocess.check_call(('ps', 'xawuf'))