@echo off
title Minecraft Setup Environment
cls

mcsetupenv.exe -forgeurl=https://files.minecraftforge.net/maven/net/minecraftforge/forge/1.12.2-14.23.3.2666/forge-1.12.2-14.23.3.2666-mdk.zip -filename=myfile.zip -fp-extract=mdk -delfile=false

PAUSE