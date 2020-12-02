## Consensus using Hashicorp's Raft

This is an example of using Hashicorp's Raft library that implements the RAFT protocol.


<br/>

### Issues

This is still WIP.

```shell
$ ./run_main.sh                                                                                   1 ✘ 
2020-12-02T17:51:42.346+0200 [INFO]  raft: initial configuration: index=0 servers=[]
2020-12-02T17:51:42.346+0200 [INFO]  raft: initial configuration: index=0 servers=[]
2020-12-02T17:51:42.346+0200 [INFO]  raft: initial configuration: index=0 servers=[]
2020-12-02T17:51:42.346+0200 [INFO]  raft: entering follower state: follower="Node at fd5ae6e4-bcfc-21e1-b0f5-f61897dbc396 [Follower]" leader=
2020-12-02T17:51:42.346+0200 [INFO]  raft: entering follower state: follower="Node at 8c4b4502-6a4f-6443-1d12-e3563a109f41 [Follower]" leader=
2020-12-02T17:51:42.346+0200 [INFO]  raft: entering follower state: follower="Node at 56a2570f-c53d-932f-eb58-0f5faa2d60f3 [Follower]" leader=
2020-12-02T17:51:43.640+0200 [WARN]  raft: no known peers, aborting election
2020-12-02T17:51:44.114+0200 [WARN]  raft: no known peers, aborting election
2020-12-02T17:51:44.152+0200 [WARN]  raft: no known peers, aborting election


```
