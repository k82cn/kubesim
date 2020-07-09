1. 提供一个新的配置文件，能够动态配置虚拟机器上的资源，如下所示：
```yaml
nodeClasses:
- name: large
  labels:      // 虚拟节点上附加的label
    host.type: virtual
  resources:   // 虚拟节点的资源容量
    capacity:
      cpu: "8"
      nvidia.com/gpu: "4"
      memory: "128Gi"
- name: small
  labels:
    host.type: virtual
  resources:
    capacity:
      cpu: "1"
      nvidia.com/gpu: "2"
      memory: "4Gi"
```
相比kubemark，kubesim启动时添加两个新参数--nodefile和--nodetemplatename：
```
kubemark \
--morph=kubelet \
--name=$(NODE_NAME) \
--kubeconfig=/kubeconfig/kubelet.kubeconfig \
$(CONTENT_TYPE) \
--log-file=/var/log/kubelet-$(NODE_NAME).log \
--logtostderr=false \
--alsologtostderr \
--v=2 \
--nodefile=hollow_node.config \  <========
--nodetemplatename=large         <========
```
2. 提供模拟pod生命周期的功能，kubesim会根据以下两个pod label来修改Pod状态：
* simulation.runDuration -- Pod在虚拟节点上的运行时间
* simulation.terminalPhase -- Pod的结束状态

下述例子表示Pod会运行10秒钟，运行结束后Pod状态为Succeeded：
```yaml
labels:
  simulation.runDuration: 10s
  simulation.terminalPhase: Succeeded
```