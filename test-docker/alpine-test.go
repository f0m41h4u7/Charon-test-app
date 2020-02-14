package main

import (
    "io"
    "os"
    "github.com/docker/go-connections/nat"
    "github.com/docker/docker/api/types/container"
    "context"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
) 

func main() {
    ctx := context.Background()
    cli, err := client.NewEnvClient()
    if err != nil {
        panic(err)
    }

    reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, reader)

    hostBinding := nat.PortBinding{
        HostIP:   "0.0.0.0",
        HostPort: "8000",
    }
    containerPort, err := nat.NewPort("tcp", "80")
    if err != nil {
        panic("Unable to get the port")
    }
    portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: "alpine",
        Cmd:   []string{"echo", "hello world"},
        Tty:   true,
    }, &container.HostConfig{
        PortBindings: portBinding,
    }, nil, "")
    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }

    out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, out)
}
