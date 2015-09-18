all:

~/Desktop/declips/1078.mp3:
	ffmpeg -y -i /Users/erin/Downloads/de/baader-meinhof/baader-meinhof.avi  -ss 00:17:58.208 -to 00:18:17.026 ~/Desktop/declips/1078.mp3 &> /dev/null
all: ~/Desktop/declips/1078.mp3

