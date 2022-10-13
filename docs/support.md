## Currently Windows and WSL2 do not support `yast`.

## Issue with Windows support

As the config utilities were developed keeping Unix FileSystem in mind, the config files of yast had no support in the Windows File System.

## Issue with WSL2 support

Since yast works with targets that require VPN access to fetch results we need to set up a VPN connection in WSL2. However, WSL2 currently does not support VPN.

## Approach To Solve This
The current approach is to create a Docker image for `yast` to extend its support for multiple platforms. As discussed in #58.