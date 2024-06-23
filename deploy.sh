#!/bin/bash


# 检查是否提供了参数
if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <HOST> <PORT> <TOKEN>"
    exit 1
fi


GITHUB_API_URL="https://api.github.com/repos/imblowsnow/server-monitor/releases/latest"

SERVICE_NAME="monitor-client"
SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"
WORK_DIR=$(pwd)
SERVER_HOST=$1
SERVER_PORT=$2
TOKEN=$3
SERVICE_ARGS="-host $SERVER_HOST -port $SERVER_PORT -token $TOKEN"

# 获取系统类型
OS=$(uname -s)
ARCH=$(uname -m)
# 将系统和架构转换为 GitHub 上的文件命名规则
case $OS in
    Linux)
        OS="linux"
        ;;
    Darwin)
        OS="darwin"
        ;;
    MINGW* | CYGWIN* | MSYS*)
        OS="windows"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    arm64 | aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# 获取最新版本信息
echo "Fetching latest release information..."

response=$(curl -s $GITHUB_API_URL)
# response 为空时，说明请求失败
if [ -z "$response" ]; then
    echo "Failed to fetch latest release information"
    exit 1
fi
# 查找匹配的下载 URL
# https://github.com/imblowsnow/server-monitor/releases/download/0.0.7/monitor-client-0.0.7-darwin-amd64.tar.gz
echo "$response" | grep -o "\"browser_download_url.*\".*\"" | grep "${OS}-${ARCH}" | grep "${SERVICE_NAME}"
download_url=$(echo "$response" | grep -o "\"browser_download_url.*\".*\"" | grep "${SERVICE_NAME}" | grep "${OS}-${ARCH}" | cut -d '"' -f 4)
if [ -z "$download_url" ]; then
    echo "No suitable download found for OS: $OS and architecture: $ARCH"
    exit 1
fi

file_name=$(basename $download_url)

echo "Latest release URL: $download_url"
echo "File name: $file_name"

# 下载最新版本文件
echo "Downloading latest release..."
curl -L $download_url -o $file_name

# 解压下载的 tar.gz 文件 (假设文件为 tar.gz 格式)
#echo "Extracting the downloaded file..."
#tar -xzf $DOWNLOAD_DIR/$file_name -C $DOWNLOAD_DIR

# 找到解压后的可执行文件 (假设解压后的目录结构固定)
#extracted_file=$(find . -type f -name 'monitor-client*' | head -n 1)
#
#if [ -z "$extracted_file" ]; then
#    echo "Extracted file not found"
#    exit 1
#fi

# 获取绝对路径
extracted_file=$(realpath $file_name)

# 赋予执行权限
echo "Setting execute permissions..."
chmod +x $extracted_file

# 创建 systemd 服务文件
echo "Creating systemd service file..."
cat <<EOT > $SERVICE_FILE
[Unit]
Description=Client Monitor Service
After=network.target

[Service]
ExecStart=$extracted_file $SERVICE_ARGS
Restart=always
User=root
Group=root

[Install]
WantedBy=multi-user.target
EOT

# 重新加载 systemd 管理配置
echo "Reloading systemd daemon..."
systemctl daemon-reload

# 启动并启用服务
echo "Starting and enabling the service..."
systemctl start $SERVICE_NAME
systemctl enable $SERVICE_NAME

echo "Deployment completed. ${extracted_file}"
