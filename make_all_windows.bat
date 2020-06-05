@echo off

echo "welcome use fohogo!"

go fmt ./src/...

call .\\windows\\make_GameServer.bat

call .\\windows\\make_Externalgws.bat

call .\\windows\\make_Innergws.bat

call .\\windows\\make_loginserver.bat

::call .\\windows\\make_simulate.bat

call .\\windows\\make_dbserver.bat

call .\\windows\\make_MMGServer.bat

pause