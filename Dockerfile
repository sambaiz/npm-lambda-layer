FROM lambci/lambda-base:build

WORKDIR /opt

RUN curl -sL https://rpm.nodesource.com/setup_12.x | bash - && \
    yum install -y nodejs && \
    mkdir bin nodejs && \
    cp /usr/bin/node bin/node && \
    cp -r /usr/lib/node_modules ./nodejs/node_modules && \
    ln -s ../nodejs/node_modules/npm/bin/npm-cli.js ./bin/npm && \
    zip -yr /tmp/npm-layer.zip ./*