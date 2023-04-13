@echo off
start /B "" room_service.exe
start /B "" user_service.exe
ping -n 1 127.0.0.1 >nul
start /B "" main_service.exe