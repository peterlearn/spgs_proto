#to update proto to global-game-proto : copy .ssh from user folder and run docker build -t update .    
FROM firerocksg-registry.ap-southeast-1.cr.aliyuncs.com/firerocksg/protoc:3.19.4

WORKDIR /src

COPY game/*.proto game/
COPY client/enums.proto client/
COPY client/common.proto client/
COPY client/item.proto client/

RUN sed -i "s|gitlab.com\/firerocksg\/global-proto|gitlab.com\/firerocksg\/global-game-proto|g" client/*.proto
RUN sed -i "s|gitlab.com\/firerocksg\/global-proto|gitlab.com\/firerocksg\/global-game-proto|g" game/*.proto

RUN protoc --go_out=. client/*.proto -I./client -I.
RUN protoc --go_out=. game/*.proto -I./client -I.

COPY ssh /root/.ssh
RUN chmod 0700 -R /root/.ssh

RUN git clone git@gitlab.com:firerocksg/global-game-proto.git
RUN rm -rf global-game-proto/client
RUN rm -rf global-game-proto/game
RUN mv gitlab.com/firerocksg/global-game-proto/client global-game-proto/
RUN mv gitlab.com/firerocksg/global-game-proto/game global-game-proto/

WORKDIR global-game-proto

RUN pwd

RUN git add .
RUN git commit -m "更新协议"
RUN git push

# RUN date "+%y.%m%d.%H%M%S" > tag; sed 's/^/v&/' -i tag; cat tag | xargs git tag
# RUN cat tag | xargs git push origin

RUN pwd

