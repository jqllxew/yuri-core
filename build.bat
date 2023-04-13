@echo off
cd .\main_service\
go build -ldflags="-w -s" -buildmode=pie -o ..\output\main_service.exe
cd ..\room_service\
go build -ldflags="-w -s" -buildmode=pie -o ..\output\room_service.exe
cd ..\user_service\
go build -ldflags="-w -s" -buildmode=pie -o ..\output\user_service.exe
cd ..
xcopy /Y /Q script\* output\
copy /Y conf\server.conf output\server.conf