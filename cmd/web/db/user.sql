CREATE USER 'developers'@'%';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'developers';
ALTER USER 'developers' IDENTIFIED BY 'd3v3l0p3rs';

CREATE USER 'developers'@'%' IDENTIFIED BY 'd3v3l0p3rs';
GRANT ALL ON snippetbox.* TO 'developers'@'%';
FLUSH PRIVILEGES;

REVOKE ALL PRIVILEGES, GRANT OPTION FROM 'developers'@'%';

GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'developers'@'%';
FLUSH PRIVILEGES;

SHOW GRANTS FOR 'developers'@'%';