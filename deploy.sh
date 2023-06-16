#!/bin/bash

# 发布代码时,减少停机时间,快速替换新版本
containerName="qcxx-admin-server"
#containerName="test"

versionNum=$(docker images | grep ^$containerName | awk '{print $2}')
if [ -z "$versionNum" ]; then
   versionNum=1
else
  # 自动增加版本号
  versionNum=$(($versionNum + 1))
fi

echo "=== container current versionNum=$versionNum, previous versionNum=$(($versionNum - 1)) ==="
echo "=== building docker image of $containerName:$versionNum ==="


cur_sec=$(date '+%s')

docker build -t $containerName:$versionNum .

echo "============== built for $(($(date '+%s') - cur_sec))s =================="
echo "=== built docker image of $containerName:$versionNum ==="

if [ $versionNum != 1 ]; then
  # 不是第一次发布,删除旧版本容器
  docker stop $containerName
  docker rm $containerName

  echo "=== stop and remove $containerName container ==="
fi

docker run --name $containerName \
--network my-bridge \
-v /data/$containerName/logs:/app/logs \
-v /data/$containerName/static:/app/static \
-p 8081:8081 \
-d $containerName:$versionNum

echo "=== run $containerName container ==="


# delete useless images due to multi-build

if [ $versionNum != 1 ]; then
  # 不是第一次发布,清理旧版本镜像
  # 最后清理,如果前面失败,可以快速回到上一个版本
  docker rmi -f $containerName:$(($versionNum - 1))

  echo "=== cleared $containerName container and image ==="
fi

noneList=$(docker images | grep \<none\> | awk '{print $3}')
for params in $noneList; do
  if [ -z "$params" ]; then
    continue
  fi
  docker rmi -f "$params"
  echo "delete $(docker images | grep "$params" | awk '{print $1}' ) image"
done
