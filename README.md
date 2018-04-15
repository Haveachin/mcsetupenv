# Minecraft setup environment
This program is setting up the standard minecraft modding environment.

## Supported IDEs
* Eclipse

##Necessary flags
`forgeurl` - The URL to the Zip-fiel of the forge version (mdk)
```
Type:    string
Default: ""   

Example:
mcsetupenv.exe -forgeurl=https://files.minecraftforge.net/maven/net/minecraftforge/forge/[...]-mdk.zip
```
---
##Optional flags
`filename` - The name of the forge file that is downlaoded
```
Type:    string
Default: "temp.zip"   

Example:
mcsetupenv.exe -filename=myfile.zip
```

`delfile` - If the forge file should be deleted or not
```
Type:    bool
Default: true   

Example:
mcsetupenv.exe -delfile=true
```

## Example Usage
For Windows the best solution is to use a batch file:

**setup.bat**
```
@echo off
title Minecraft Setup Environment
cls

mcsetupenv.exe -forgeurl=https://files.minecraftforge.net/maven/net/minecraftforge/forge/1.12.2-14.23.3.2666/forge-1.12.2-14.23.3.2666-mdk.zip -filename=myfile.zip -delfile=false

PAUSE
```