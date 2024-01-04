# Dither Filter

This is a simple image filter made for fun. It uses the GIMP Arithmetic ADD dither filter to create a dithering effect, maps all the colors to a palette in a hex file and does a 0,7 gaussian blur. This can give medium res images an 80s effect if you use pastel colors.

## Using the program

To use the program it must be compiled as go run can't handle arguments. Use -i for the input -o for the output and -p for the palette.

## Install

To install the program is very simple. You can copy the binary to a PATH folder like /usr/local/bin and done. To be safer you can put the program in /opt and create a symbolic link in the PATH folder instead.

## Examples ( Hover the mouse on the images for info )

### Credits to ArseniXC for the images

![Default image](imgs/wallpaper.jpg "This is the default image for control")
![Maze image](imgs/wallpaper_maze.png "This is same image with the deep maze palette")
![Retro image](imgs/wallpaper_retro.png "This is same image with the retro 115 palette")
![Rgr image](imgs/wallpaper_rgr.png "This is same image with the rgr papercut palette")
![cc image](imgs/wallpaper_cc.png "This is same image with the cc 29 palette")

