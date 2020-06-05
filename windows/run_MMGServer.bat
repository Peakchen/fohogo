@echo off
call make_MMGServer.bat
cd ..
cd bin
MMGServer.exe
cd ..
pause