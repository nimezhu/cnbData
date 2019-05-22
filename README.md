# NBrowser Data Server 

## Add Custom Tracks to CNB

Download executable binary file
- [Linux Version](https://vis.nucleome.org/static/dist/current/linux/cnbData)
- [Windows Version](https://vis.nucleome.org/static/dist/current/win64/cnbData.exe)
- [Mac OS version](https://vis.nucleome.org/static/dist/current/mac/cnbData)


Start a Data Server

```shell
./cnbData start -i [google sheet id or excel file] -p [port default:8611]
```

[Example Input: Google sheet](https://docs.google.com/spreadsheets/d/1nJwOozr4EL4gnx37hzF2Jmv-HPsgFMA9jN-lbUj1GvM/edit#gid=1744383077)

The google sheet id is part of the url in your google sheet webpage.


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
