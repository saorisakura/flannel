前提
1. 会使用IP命令设置路由、IP和子网络接口
    在大多数基于 Linux 的系统上，`ip` 命令是由 `iproute2` 包提供的。你可以使用包管理器来安装它。以下是一些常见的 Linux 发行版的安装命令：

    ### Ubuntu/Debian
    ```sh
    sudo apt update
    sudo apt install iproute2
    ```

    ### CentOS/RHEL
    ```sh
    sudo yum install iproute
    ```

    ### Fedora
    ```sh
    sudo dnf install iproute
    ```

    ### Arch Linux
    ```sh
    sudo pacman -S iproute2
    ```

    安装完成后，你可以使用 `ip` 命令来管理网络接口、路由和其他网络配置。

2. 会使用brctl相关命令设置网桥相关参数
    在大多数基于 Linux 的系统上，`brctl` 命令是由 `bridge-utils` 包提供的。你可以使用包管理器来安装它。以下是一些常见的 Linux 发行版的安装命令：

    ### Ubuntu/Debian
    ```sh
    sudo apt update
    sudo apt install bridge-utils
    ```

    ### CentOS/RHEL
    ```sh
    sudo yum install bridge-utils
    ```

    ### Fedora
    ```sh
    sudo dnf install bridge-utils
    ```

    ### Arch Linux
    ```sh
    sudo pacman -S bridge-utils
    ```

    安装完成后，你可以使用 `brctl` 命令来管理网桥。
3. 要使用 `ip` 命令设置网卡子接口，可以按照以下步骤操作：

   1. 创建子接口。
   2. 为子接口分配 IP 地址。
   3. 启用子接口。

   以下是具体的命令示例：

   ### 创建子接口
   假设你要在 `eth0` 接口上创建一个名为 `eth0.10` 的子接口：
   ```sh
   sudo ip link add link eth0 name eth0.10 type vlan id 10
   ```

   ### 为子接口分配 IP 地址
   为 `eth0.10` 分配一个 IP 地址，例如 `192.168.1.1/24`：
   ```sh
   sudo ip addr add 192.168.1.1/24 dev eth0.10
   ```

   ### 启用子接口
   启用 `eth0.10` 子接口：
   ```sh
   sudo ip link set dev eth0.10 up
   ```

   通过以上步骤，你就可以使用 `ip` 命令设置网卡子接口。
4. 要使用 `bridge` 命令设置网卡的 FDB（Forwarding Database），可以按照以下步骤操作：

   1. 添加一个 FDB 条目。
   2. 删除一个 FDB 条目。

   以下是具体的命令示例：

   ### 添加一个 FDB 条目
   假设你要在 `br0` 网桥上添加一个 MAC 地址为 `00:11:22:33:44:55`，并且该地址通过 `eth0` 接口转发：
   ```sh
   sudo bridge fdb add 00:11:22:33:44:55 dev eth0 master br0
   ```

   ### 删除一个 FDB 条目
   假设你要在 `br0` 网桥上删除一个 MAC 地址为 `00:11:22:33:44:55` 的 FDB 条目：
   ```sh
   sudo bridge fdb del 00:11:22:33:44:55 dev eth0 master br0
   ```

   通过以上步骤，你就可以使用 `bridge` 命令设置网卡的 FDB。

   交换机之间的 MAC 地址信息交换通常通过以下几种方式实现：

   1. **生成和学习**：交换机通过监听网络中的数据帧来学习 MAC 地址。当交换机接收到一个数据帧时，它会记录下源 MAC 地址和接收端口的对应关系，并将其存储在 MAC 地址表中。

   2. **生成和泛洪**：当交换机接收到一个目标 MAC 地址未知的数据帧时，它会将该数据帧泛洪到所有端口（除了接收端口），以确保目标设备能够接收到数据帧并进行响应。目标设备响应后，交换机就能学习到目标 MAC 地址和对应的端口。

   3. **生成和老化**：交换机会定期检查 MAC 地址表中的条目，并删除那些在一定时间内没有活动的条目。这一过程称为老化，确保 MAC 地址表保持最新。

   4. **生成和协议**：一些高级交换机使用协议（如 Spanning Tree Protocol, STP）来交换和更新 MAC 地址信息，以避免网络环路和优化网络路径。

   以下是一个简单的示例，展示了交换机如何通过监听数据帧来学习 MAC 地址：

   ```plaintext
   1. 设备 A 发送数据帧到设备 B。
   2. 交换机 S1 接收到数据帧，并记录下源 MAC 地址（设备 A 的 MAC 地址）和接收端口。
   3. 交换机 S1 检查目标 MAC 地址（设备 B 的 MAC 地址），如果在 MAC 地址表中没有找到对应条目，则将数据帧泛洪到所有端口。
   4. 设备 B 接收到数据帧并进行响应。
   5. 交换机 S1 接收到设备 B 的响应数据帧，并记录下设备 B 的 MAC 地址和接收端口。
   ```

   通过以上步骤，交换机能够动态学习和更新 MAC 地址表，从而实现高效的数据帧转发。

5. ### Plan
   1. Use the `subprocess` module to execute the `ip link show` command.
   2. Capture the output and decode it from bytes to a string.
   3. Parse the output to extract relevant information about network interfaces.

   ### Code
   ```python
   import subprocess
   import re

   def get_ip_link_show():
       result = subprocess.run(['ip', 'link', 'show'], capture_output=True, text=True)
       return result.stdout

   def parse_ip_link_show(output):
       interfaces = []
       interface = {}
       for line in output.splitlines():
           if re.match(r'^\d+:', line):
               if interface:
                   interfaces.append(interface)
               interface = {}
               parts = line.split(': ', 2)
               interface['index'] = parts[0]
               interface['name'] = parts[1].split('@')[0]
               interface['flags'] = parts[2].strip()
           elif 'link/' in line:
               parts = line.split()
               interface['type'] = parts[0]
               interface['mac'] = parts[1]
               interface['state'] = parts[2] if len(parts) > 2 else None
       if interface:
           interfaces.append(interface)
       return interfaces

   if __name__ == "__main__":
       output = get_ip_link_show()
       interfaces = parse_ip_link_show(output)
       for iface in interfaces:
           print(iface)
   ```

   This code will parse the output of `ip link show` and print a list of dictionaries, each containing information about a network interface.