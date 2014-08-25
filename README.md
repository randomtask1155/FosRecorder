<h1>FosRecorder</h1>

  A simple wrapper to control and capture a VLC stream.  My primary use case is to record wired/wireless foscams planted around the house. 
  A 30 minitue recording results in about 230MB of data give depending on video stream
  
Assuming 100GB of free space and average 230mb for 30 minutes
1xcam = 222 hours per cam ( 9 days )
2xcam = 111 hours per cam ( 4 days )
3xcam = 74  hours per cam ( 3 days )

Usage:
```
Usage for FosRecorder:

  -n DEFAULT="48":
     #~ Number of recording to keep before rolling over the first

  -r DEFAULT="1800":
     #~ how many seconds to record for each recording

  -s DEFAULT="":
     #~ shell script used to execute capture of stream.  Must take exactly one int argument

  -v
     #~ enable verbose output for debugging
```

Example:
```
FosRecorder -n 3 -r 10 -s ./capture.sh
```
