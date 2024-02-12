@echo off 
mkdir "C:\Program Files\Exporter" 
icacls "C:\Program Files\Exporter" /grant Everyone:(OI)(CI)F 
copy main.exe "C:\Program Files\Exporter" 
sc create "Exporter" binPath= "C:\Program Files\Exporter\main.exe Exporter" start= auto 
sc start "Exporter" 
