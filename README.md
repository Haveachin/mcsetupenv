# Minecraft setup environment
This program is setting up the standard minecraft modding environment.

## Supported IDEs
* Eclipse

## Necessary flags
`ff-url` - The URL to the Zip-file of the forge version (mdk)
```
Type:    string
Default: ""   

Example:
mcsetupenv.exe -ff-url=https://files.minecraftforge.net/maven/net/minecraftforge/forge/[...]-mdk.zip
```

## Optional flags
`fp-dl` - The path + name of the forge file that is downlaoded
```
Type:    string
Default: "temp.zip"   

Example:
mcsetupenv.exe -fp-dl=myfile.zip
```

`fp-extract` - The path where the forge file should extract
```
Type:    string
Default: "."   

Example:
mcsetupenv.exe -fp-extract=mdk
```

`ff-del` - If the forge file should be deleted or not
```
Type:    bool
Default: true

Example:
mcsetupenv.exe -ff-del=false
```
---
## Example Usage
For Windows the best solution is to use a batch file:

**setup.bat**
```
@echo off
title Minecraft Setup Environment
cls

mcsetupenv.exe -forgeurl=https://files.minecraftforge.net/maven/net/minecraftforge/forge/1.12.2-14.23.3.2666/forge-1.12.2-14.23.3.2666-mdk.zip -filename=myfile.zip -fp-extract=mdk -delfile=false

PAUSE
```