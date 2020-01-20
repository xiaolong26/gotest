package api

import (
    //"fmt"
    "github.com/gin-gonic/gin"
    "k8s.io/client-go/kubernetes"
   // "gotest/db"
)

var clientset *kubernetes.Clientset

func Test(c *gin.Context){
   /* k8sconfig := flag.String("k8sconfig","./config","kubernetes config file path")
    flag.Parse()
    config , err := clientcmd.BuildConfigFromFlags("",*k8sconfig)
    if err != nil {
        log.Println(err)
    }
    clientset , err = kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatalln(err)
    } else {
        fmt.Println("connect k8s success")
    }

    //获取POD
    pods,err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
    if err != nil {
        log.Println(err.Error())
    }

    fmt.Println(pods.Items[1].Name)
    fmt.Println(pods.Items[1].CreationTimestamp)
    fmt.Println(pods.Items[1].Labels)
    fmt.Println(pods.Items[1].Namespace)
    fmt.Println(pods.Items[1].Status.HostIP)
    fmt.Println(pods.Items[1].Status.PodIP)
    fmt.Println(pods.Items[1].Status.StartTime)
    fmt.Println(pods.Items[1].Status.Phase)
    fmt.Println(pods.Items[1].Status.ContainerStatuses[0].RestartCount)   //重启次数
    fmt.Println(pods.Items[1].Status.ContainerStatuses[0].Image) //获取重启时间

    //获取NODE
    fmt.Println("##################")
    nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
    fmt.Println(nodes.Items[0].Name)
    fmt.Println(nodes.Items[0].CreationTimestamp)    //加入集群时间
    fmt.Println(nodes.Items[0].Status.NodeInfo)
    fmt.Println(nodes.Items[0].Status.Conditions[len(nodes.Items[0].Status.Conditions)-1].Type)
    fmt.Println(nodes.Items[0].Status.Allocatable.Memory().String())
    */
    c.String(200, "pong")
}
