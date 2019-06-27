# Nucleome Data Server 

## Introduction

[*Nucleome Data Server*](http://v.nucleome.org/data/server) is a command line tool for [*Nucleome Platform*](http://v.nucleome.org/home). It is designed for users to host their data in servers. 

If you want to host your data in you local machine. Please refer to our Graphic User Interface tool [Nucleome Data](https://github.com/nimezhu/ndata).

It is a software for user to host their own [bigWig](https://genome.ucsc.edu/goldenpath/help/bigWig.html), [bigBed](https://genome.ucsc.edu/goldenpath/help/bigBed.html) and [.hic](https://github.com/aidenlab/Juicebox/blob/master/HiC_format_v8.docx) format data. 

When user start a data server in their local machine, their private data can be integrated with public available data which are hosted by other public data servers and rendered in [*Nucleome Browser*](https://vis.nucleome.org). 

User's private data are not accessible by other users or web application administrator if their data server is in localhost or local network.

The input for this software is a Google Sheet or a Excel file which has the information such as file location(URI), short label(shortLabel), long label(longLabel) and weblink(metaLink) of future description of tracks. These data files can be either located in local drive or just a weblink.


## Install
This software is implemented in [GoLang](https://golang.org/) and works in Linux, Windows and Mac OS without installation. 

Please download executable binary file

[![Linux64](https://img.shields.io/badge/binary-linux-green.svg?style=flat)](https://vis.nucleome.org/static/dist/current/linux/cnbData)
[![Windows](https://img.shields.io/badge/binary-win-blue.svg?style=flat)](https://vis.nucleome.org/static/dist/current/win64/cnbData.exe)
[![MacOS](https://img.shields.io/badge/binary-macos-yellow.svg?style=flat)](https://vis.nucleome.org/static/dist/current/mac/cnbData)

And change the mode of this file into executable.

In Linux or Mac OS, this can be done in a terminal, using command `chmod`.

```shell
chmod 755 cnbData
```

If you are using Windows and not familiar with runnning command line tool in Windows, please read [this article](https://www.computerhope.com/issues/chusedos.htm) first. Then,you can run `cnbData` as a command line tool in terminal.




## Quick Start 

Start a Data Server

```shell
./cnbData start -i [google sheet id or excel file] -p [port default:8611]
```

The track configuration input for cnbData could be a google sheet or a excel file.

### Start with Google Sheet
[Example Input: Google sheet](https://docs.google.com/spreadsheets/d/1nJwOozr4EL4gnx37hzF2Jmv-HPsgFMA9jN-lbUj1GvM/edit#gid=1744383077)

For easy start, you could download it as an excel file and named it as `cnb.xlsx`.
Then you can run the command below.

`cnbData start -i cnb.xlsx`

You can also skip downloading and use **Google Sheet ID** directly like this.
`cnbData start -i 1nJwOozr4EL4gnx37hzF2Jmv-HPsgFMA9jN-lbUj1GvM`

> The **Google Sheet ID** is part of the url in your google sheet webpage. It is in blue background in the following demostration image.
> ![Google Sheet ID Demo](https://nbrowser.github.io/image/google_sheet_id_demo.png)

When **first time** use `cnbData` with google sheet, it will prompt a link to ask for authorize permission, copy this link to browser and get back a token, then copy and paste this token to command terminal, a credential token will be stored in `[Your Home Dir]/.cnbData/credentials/gsheet.json`.








### Input Google Sheet or Xlsx Format

Two reserved sheets are required for start this data server.  

One is “Config”,  which store the configuration variable values. Currently, `root` variable is the only variable needed for cnbData. It stores the root path for you store all track data files. It is designed for user conveniently migrating data between servers. All the URI in other sheets will be the relative path to `root` if their URI are not start with `http` or `https`.
![Sheet Config Example](https://nbrowser.github.io/image/sheetConfig.png)

The other sheet is “Index”, which stores the configuration information of all other sheets which are needed to use in cnbData server. The sheet titles which are not in Index sheet will be ignored by cnbData.

![Sheet Index Example](https://nbrowser.github.io/image/sheetIndex.png)

For track format data sheet, if using four columns, the columns name should be “shortLabel” , “uri,metaLink,longLabel”, and the corresponding column header such as A,B et al. should put into the 4th and 5th column.


 
If using two columns, the column name could be any string user defined. Just filled in the column index into the fourth and the fifth column accordingly. In sheet "Index", those entries which Id starts with “#” will be ignored when loading.
Column "Type" is a reserve entry for future data server. Currently, just use "track" in this column. It support bigWig, bigBed and .hic.
#### Simple Name and URI
![Sheet Data Example](https://nbrowser.github.io/image/sheetSimpleData.png)

#### With Long Label and Meta Link
![Sheet Data Example](https://nbrowser.github.io/image/sheetData4.png)


The localhost http://127.0.0.1:8611 is one of default servers in CNB. If user starts a data server in localhost and the port is default 8611, user doesn’t need to configure the server list. Just reload server content or add new genome browser panel after the local server start, the custom data will show in this genome browser config panel.


If Data server is in other port or other web servers instead of localhost, user need to add the server into server lists. Open the CNB main website in your chrome browser. If user don't have a genome browser panel, please add a genome browser panel, the add button is in submenu of panels in the menu bar. Then, in this genome browser, then Click Config tracks → Click Config Servers → Input Server URI and any Id into table → Click Refresh Button to reload.


![Config Servers](https://nbrowser.github.io/image/configServers.png)

If user open a new genome browser panel , it will loading servers as last configuration. Servers configuration is stored as settings for this panel, if user duplicate this panel, the servers setting will be automatically copied too.
