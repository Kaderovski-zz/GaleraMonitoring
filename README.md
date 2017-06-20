# Galera Monitoring
This program is a Golang cluster mariaDB/Mysql Monitoring.\
Alerts are sending on Slack Channel.

##### Notes
Currently still working on it.

  * [Installation](#installation)
  * [Exemple](#exemple)
  * [Docker](#docker)
    * [Docker-run](#docker-run)
  

### Installation
* Run the following :
```
> go get https://github.com/F00b4rch/GaleraMonitoring
```
* Create you SlackApp : https://api.slack.com/apps/
* Add feature Incoming Webhooks
* Put your WebhookURL/Username/Channel inside slackApp/payload.go
```
	// Add here your webhookurl
	webhookUrl := "https://hooks.slack.com/services/your_key"

	payload := slack.Payload{
		Text:    text,
		Username: "Galera_alerte",
		Channel:  "#infra",

```


### Exemple

This exemple will show you normal output only if you **uncomment** fmt in import() and good parts in the code.
```
> go run main.go

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

#### Docker-run :
```
$ sudo docker run --detach=true --name node1 -h node1 erkules/galera:latest --wsrep-cluster-name=local-test --wsrep-cluster-address=gcomm://
$ sudo docker run --detach=true --name node2 -h node2 --link node1:node1 erkules/galera:latest --wsrep-cluster-name=local-test --wsrep-cluster-address=gcomm://node1
$ sudo docker run --detach=true --name node3 -h node3 --link node1:node1 erkules/galera:latest --wsrep-cluster-name=local-test --wsrep-cluster-address=gcomm://node1
```
