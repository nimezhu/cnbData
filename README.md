# Nucleome Browser Data Server 

## Introduction

[*Nucleome Browser Data Server*](http://v.nucleome.org/data/server) is a component for [*Nucleome Browser Platform*](http://v.nucleome.org/home). It is a software for user to host their own [bigWig](https://genome.ucsc.edu/goldenpath/help/bigWig.html), [bigBed](https://genome.ucsc.edu/goldenpath/help/bigBed.html) and [.hic](https://github.com/aidenlab/Juicebox/blob/master/HiC_format_v8.docx) format data. When user start a data server in their local machine, their private data can be integrated with public available data which are hosted by other public data servers and rendered in [*Nucleome Browser Progressive Web Application*](https://vis.nucleome.org). 

User's private data are not accessible by other users or web application administrator if their data server is in localhost.

The input for this software is a Google Sheet or a Excel file which has the information such as file location(URI), short label(shortLabel), long label(longLabel) and weblink(metaLink) of future description of tracks. These data files can be either located in local drive or just a weblink.


## Install
This software is implemented in [GoLang](https://golang.org/) and works in Linux, Windows and Mac OS without installation. 

Please download executable binary file

[![Linux64](https://img.shields.io/badge/binary-linux-green.svg?style=flat)](https://vis.nucleome.org/static/dist/current/linux/cnbData)
[![Windows](https://img.shields.io/badge/binary-win-blue.svg?style=flat)](https://vis.nucleome.org/static/dist/current/win64/cnbData.exe)
[![MacOS](https://img.shields.io/badge/binary-macos-yellow.svg?style=flat)](https://vis.nucleome.org/static/dist/current/mac/cnbData)

And change the mode of this file into excutable.

In Linux or Mac OS, this can be done in a terminal, using command *chmod*.

```shell
chmod 755 cnbData
```

Then,you can run cnbData as a command line tool in terminal.

## Quick Start 

Start a Data Server

```shell
./cnbData start -i [google sheet id or excel file] -p [port default:8611]
```
[Example Input: Google sheet](https://docs.google.com/spreadsheets/d/1nJwOozr4EL4gnx37hzF2Jmv-HPsgFMA9jN-lbUj1GvM/edit#gid=1744383077)


> The google sheet id is part of the url in your google sheet webpage.
> ![Google Sheet ID Demo](https://nbrowser.github.io/image/google_sheet_id_demo.png)

The track configuration input for cnbData could be a google sheet or a excel file.


Two reserved sheets is needed for start cnb data browser.  One is “Config”,  which store the configuration variable values, currently, “root” variable is the only variable needed for cnbData. It stores the root path for you store all data. It is designed for user conveniently migrating data between servers. The other sheet is “Index”, which stores the configuration information of all other sheets which are needed to use in cnbData server. The sheet titles which are not in Index sheet will be ignored by cnbData.


For track format data sheet, if using four columns, the columns name should be “shortLabel” , “uri,metaLink,longLabel”, and the corresponding column header such as A,B et al. should put into the 4th and 5th column.

 

If using two columns, the column name could be any string user defined. Just filled in the column index into the fourth and the fifth column accordingly.


For those entries which Id starts with “#” will be ignored when loading.


When first time use cnbData with google sheet, it will prompt a link to ask for authorize permission, copy this link to browser and get back a token, then copy and paste this token to command terminal, a credential token will be stored in ~/.cnbData/credentials/gsheet.json.


Data file uris can be web url or local file relative path to “root” in Config sheet. Currently,  in type “track” cnbData support three format self indexed file, bigwig format, bigbed format and .hic format.


The localhost http://127.0.0.1:8611 is one of default server in CNB. If user starts a data server in localhost and the port is default 8611, user doesn’t need to configure the server list. Just reload server content or add new genome browser panel after the local server start, the custom data will show in this genome browser config panel.


If Data server is in other port or other web address instead of localhost, user need to add the server into server lists. Open the CNB main website in your chrome browser. Add a genome panel.  Then Click Config tracks → Click Config Servers → Input Server URI and any Id into table → Click Refresh Button to reload.

The last time loading server configure will be saved into user’s application variable, then next time open a new genome panel user doesn’t need configure the server list.
