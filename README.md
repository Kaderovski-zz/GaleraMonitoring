# Galera_Monitoring
Golang cluster mariaDB/Mysql monitoring

### Notes
Currently still working on it.

### Exemple 

```
go run main.go

### Version ####
2017/06/16 14:19:48 Serveur n1 - version 5.6.35-1xenial
2017/06/16 14:19:48 Serveur n2 - version 5.6.35-1xenial
2017/06/16 14:19:48 Serveur n3 - version 5.6.35-1xenial
### UUID ###
2017/06/16 14:19:48 n1 a1c404a9-51b4-11e7-b057-237cc5970d38
2017/06/16 14:19:48 n2 a1c404a9-51b4-11e7-b057-237cc5970d38
2017/06/16 14:19:48 n3 a1c404a9-51b4-11e7-b057-237cc5970d38
### Nodes ###
2017/06/16 14:19:48 Total Nodes : 3
2017/06/16 14:19:48 Number of Nodes counts : 3
2017/06/16 14:19:48 Number of Nodes counts : 3
2017/06/16 14:19:48 Number of Nodes counts : 3
### STATUS ###
2017/06/16 14:19:48 n1 status : Primary
2017/06/16 14:19:48 n2 status : Primary
2017/06/16 14:19:48 n3 status : Primary
2017/06/16 14:19:48 n1 is ready : [ON]
2017/06/16 14:19:48 n2 is ready : [ON]
2017/06/16 14:19:48 n3 is ready : [ON]
2017/06/16 14:19:48 n1 is connected : [ON]
2017/06/16 14:19:48 n2 is connected : [ON]
2017/06/16 14:19:48 n3 is connected : [ON]
### Average Replication ###
2017/06/16 14:19:48 Average on n2 : 0.000000
2017/06/16 14:19:48 Average on n3 : 0.000000
2017/06/16 14:19:48 Average on n1 : 0.100000
```

## Docker

You can try it with 3 Galera Docker containers.

Please refer to [Galera Documentation](http://galeracluster.com/2015/05/getting-started-galera-with-docker-par-1/)

I will make a docker-compose.yml asap.