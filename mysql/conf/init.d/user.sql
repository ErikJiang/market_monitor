-- Set mysql account permissions
grant all privileges on *.* to root@"%" identified by '123456' with grant option;
grant all privileges on *.* to monitor@"%" identified by '123456' with grant option;