@echo off
set Version=1.2.2
echo v%Version%

md builds

set goos=windows
set goarch=amd64
echo goos=windows, goarch=amd64
go build
md builds\ascf_v%Version%_Windows_x64
move ascf.exe builds\ascf_v%Version%_Windows_x64

set goos=windows
set goarch=386
echo goos=windows, goarch=386
go build
md builds\ascf_v%Version%_Windows_x86
move ascf.exe builds\ascf_v%Version%_Windows_x86

set goos=darwin
set goarch=amd64
echo goos=darwin, goarch=amd64
go build
md builds\ascf_v%Version%_macOS_x64
move ascf builds\ascf_v%Version%_macOS_x64

set goos=darwin
set goarch=386
echo goos=darwin, goarch=386
go build
md builds\ascf_v%Version%_macOS_x86
move ascf builds\ascf_v%Version%_macOS_x86

set goos=linux
set goarch=amd64
echo goos=linux, goarch=amd64
go build
md builds\ascf_v%Version%_Linux_x64
move ascf builds\ascf_v%Version%_Linux_x64

set goos=linux
set goarch=386
echo goos=linux, goarch=386
go build
md builds\ascf_v%Version%_Linux_x86
move ascf builds\ascf_v%Version%_Linux_x86

set goos=linux
set goarch=arm
set goarm=6
echo goos=linux, goarch=arm, goarm=6
go build
md builds\ascf_v%Version%_Linux_ARMv6
move ascf builds\ascf_v%Version%_Linux_ARMv6

set goos=linux
set goarch=arm
set goarm=7
echo goos=linux, goarch=arm, goarm=7
go build
md builds\ascf_v%Version%_Linux_ARMv7
move ascf builds\ascf_v%Version%_Linux_ARMv7

set goos=linux
set goarch=arm64
echo goos=linux, goarch=arm64, goarm=8
go build
md builds\ascf_v%Version%_Linux_ARMv8
move ascf builds\ascf_v%Version%_Linux_ARMv8

pause