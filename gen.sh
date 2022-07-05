#!/bin/bash
p=`pwd`
echo $p
rm -rf pkg/client
bash ../../../k8s.io/code-generator/generate-groups.sh all \
github.com/maomaoguodragonplus/k8s-gmm-controller/pkg/client \
github.com/maomaoguodragonplus/k8s-gmm-controller/pkg/apis \
gmmcrd:v1
tree
