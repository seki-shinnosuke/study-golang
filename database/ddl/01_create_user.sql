CREATE USER 'golang-user'@'%' IDENTIFIED BY 'golang-pass';
GRANT SELECT,INSERT,UPDATE,DELETE,EXECUTE,SHOW VIEW ON study_golang.* TO 'golang-user'@'%';