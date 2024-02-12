@echo off 
sc stop "Exporter" 
sc delete "Exporter" 
rmdir /s /q "C:\Program Files\Exporter" 
