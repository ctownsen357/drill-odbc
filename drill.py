import pyodbc
import os

#a quick test to make sure apache drill and the ODBC driver work correctly
#tested on Ubuntu 14, 16, Fedora 23, Centos 7.x, RHEL 7.x

conn = pyodbc.connect("""Driver=/opt/mapr/drillodbc/lib/64/libmaprdrillodbc64.so;
			Catalog=DRILL;
			Schema=hivestg; 
			ConnectionType=Direct;
			Host=localhost;
			Port=31010;
			AuthenticationType=No Authentication""", autocommit=True)

cur = conn.cursor()

#Figuring out the path dynamically so the query works wherever it may be downloaded to
data_directory = os.path.join(os.path.dirname(os.path.realpath(__file__)), 'test_data')

sql = 'select * from dfs.`{data}` '.format(data=data_directory)

cur.execute(sql)

data = cur.fetchall()

print(data)
