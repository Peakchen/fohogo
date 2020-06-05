@echo off
call make_ExternalWebGateway.bat
cd ..
cd bin
ExternalWebGateway.exe
cd ..
pause