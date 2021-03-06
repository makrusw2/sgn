#! /usr/bin/env python
import os
import glob

dir_path = os.path.dirname(os.path.realpath(__file__))
filenames = glob.glob(dir_path+"/../../../docker-volumes/node*/sgnd/sgnd.log")
files = []
for filename in filenames:
    files.append(open(filename, "r"))

def isTendermintLog(line):
    if (len(line) > 27 and line[1] == '[' and line[12] == "|" and line[25] == ']'):
        return True
    return False

def readnext(f):
    line = f.readline()
    while isTendermintLog(line):
        line = f.readline()
    return line.rstrip('\n')


def select(lines):
    time, line, index = "", "", -1
    for i in range(len(lines)):
        t = lines[i][:23]
        if t != "" and (time == "" or t < time):
            line, time, index = lines[i], t, i
    return line, index


def merge(files):
    lines = []
    for f in files:
        lines.append(readnext(f))
    while 1:
        line, index = select(lines)
        if index == -1:
            break
        print("n%d: %s" % (index, line))
        lines[index] = readnext(files[index])


if __name__ == '__main__':
    merge(files)