# codex-collation-workbench-117

项目用途：这是一个服务数字人文研究者的 Go 全栈应用，用来整理不同传本的卷册坐标、段落摘录、字词异文和底本候选。系统采用分片文档存储：每个摘录按文本摘要写入内容寻址目录，修订会产生新的文本块，不覆盖原始记录。校勘页面使用 Go 模板渲染，浏览器再通过 JSON API 增量读取数据。项目源代码、依赖描述和评测专用 Docker 文件共同构成自包含任务；不依赖本机预编译二进制。

## 标准构建、运行和测试命令

```bash
go build ./...
go run ./cmd/collation
go test ./...
```
## 评测容器

评测专用 Dockerfile 为 `benzhi.Dockerfile`，构建脚本为 `build_benzhi_docker.sh`。

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh my-go-task linux/arm64
./build_benzhi_docker.sh my-go-task linux/amd64
docker run -it my-go-task:latest
```
