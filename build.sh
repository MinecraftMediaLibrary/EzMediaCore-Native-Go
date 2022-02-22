# Windows i386
env GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=/Users/bli24/Downloads/toolchains/mingw-w64/9.0.0_2/bin/i686-w64-mingw32-gcc go build -buildmode=c-shared -o binary/windows/i386/dither.dll

# Windows amd64
# GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=/Users/bli24/Downloads/toolchains/mingw-w64/9.0.0_2/bin/i686-w64-mingw32-gcc
# go build -buildmode=c-shared -o dither.dll

# Darwin amd64
env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -buildmode=c-shared -o binary/darwin/amd64/libdither.dylib

# Darwin arm64
env GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -buildmode=c-shared -o binary/darwin/arm64/libdither.dylib

# Linux amd64
env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=/Users/bli24/Downloads/toolchains/musl-cross/0.9.9_1/libexec/bin/x86_64-linux-musl-gcc go build -buildmode=c-shared -o binary/linux/amd64/libdither.so

export GOGCCFLAGS=""

# Linux arm64
env GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=/Users/bli24/Downloads/toolchains/gcc-arm-none-eabi-10.3-2021.10/bin/arm-none-eabi-gcc go build -buildmode=c-shared -o binary/linux/arm64/libdither.so

# Linux arm
env GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=/Users/bli24/Downloads/toolchains/gcc-arm-linux-gnueabihf/6.5.0/bin/arm-linux-gnueabihf-gcc go build -buildmode=c-shared -o binary/linux/arm/libdither.so

# Linux i386
# GOOS=linux GOARCH=386 CGO_ENABLED=1 CC=/Users/bli24/Downloads/toolchains/i386-elf-gcc/bin/i386-elf-gcc GOGCCFLAGS="-fPIC -arch x86_64 -m64 -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/jh/f2mbr4dn0296rxkscq7xg1gr0000gp/T/go-build4140354949=/tmp/go-build -gno-record-gcc-switches -fno-common"
# go build -buildmode=c-shared -o libdither.so

