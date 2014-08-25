#!/bin/bash

ARG=$1

vlc -vvv http://admin:password@192.168.3.100:8080/videostream.asf --sout "#transcode{acodec=mp4a,ab=128,channels=2,samplerate=44100}:std{access=file,mux=mp4,dst=OUTPUT${ARG}}"
