## drill-odbc
Some simple scripts with a tiny parquet dataset to test ODBC connectivity to Apache Drill from Python and Go.

One thing I found frustrating was that I needed to make the /opt/mapr/drillodbc/Setup/mapr.drillodbc.ini file setting for ODBCInstLib match that of my local home directory .mapr.drillodbc.ini setting.  The local version is supposed to take priority but for whatever reason on CentOS 7 ODBC connections would not work until both of those files matched.  

*Note:* Python attempt works, still trying to get Go working, it connects but the query is failing at the moment.
