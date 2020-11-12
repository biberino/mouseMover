### mouseMover
Moves the cursor randomly on the screen for ever. Once you move the cursor fast by hand, the program will terminate.

### Installation
Grab the latest binary from the releases and execute it.

### Build
The compilation is tested on Ubuntu 18.04. For crosscompilation you need mingw-w64 compiler and minGW-zlib:
```
sudo apt install mingw-w64  
sudo apt install libz-mingw-w64-dev  
```
Use make to build
```
make window64
make windows32
make linux64
make linux32
```