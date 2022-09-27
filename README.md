
## 1.머리말
> 하나의 봇에 여러개의 어플리케이션을 연결이 안되는것으로 확인이 됩니다(Interactivity가 하나). 그래서 중간에 kafka를 두고 이를 통해서 micro service를 만들어 보겠습니다. 


실행법 
1. 각서비스에 config.json파일을 알맞게 세팅해주세요
2. docker-compose를 설치
3. root directory에서 docker-compose up 명령어 실행


---

## 2.구상도

![concept](https://velog.velcdn.com/images/divan/post/8b5c8b0c-7d12-4cdb-b9de-7e78dbb15418/image.png)

----

## 3.설명

> 1. slack bot에 대한 응답은 intractivty 서비스가 받아서 payload를 그데로 kafka에 메세지를 발행한다. 
2. 야근 서비스와 탄련근무 서비스는 kafka를 구동하고 있기에 각각의 서비스가 처리해야할 메시지이면 처리한다.
3. 추후에 확장을 원한다면 야근서비스, 탄력근무처럼 kafka만 구독하면 서비스를 늘리는것이 가능하다.

---

## 4.docker-compose.yml
> reversProxy(nginx)를 사용합니다. 그렇기에 다른 서비스가 켜지고 나면 후에 서비스가 동작하도록 하였음. nginx에서 다른 컨테이너와 연결이 용이하도록 같은 network이여야 합니다.

~~~ yml
version: "3"

services:
  reverse-proxy:
    image: nginx
    ports:
      - "80:80"
    restart: always
    depends_on:
      - bot-listener
      - extra-work
      - time-loan
    networks:
      - slackEvent
    volumes:
      - ./proxy/nginx.conf:/etc/nginx/nginx.conf

#  DB는 별도로 운영
#  mongodb:
#    image: mongo
#    container_name: mongodb
#    restart: always
#    volumes:
#      - ./mongodb:/data/db
#    networks:
#      - slackEvent


  bot-listener:
    container_name: slackbot_kafka_bot_listener
    build:
      context: .
      dockerfile: Dockerfile.botlistener
#    ports:
#      - "8282:8181"
    networks:
      - slackEvent

  extra-work:
    container_name: slackbot_kafka_extra_work
    build:
      context: .
      dockerfile: Dockerfile.extraworkservice
#    ports:
#      - "8383:8181"
    depends_on:
      - bot-listener
    networks:
      - slackEvent

  time-loan:
    container_name: slackbot_kafka_time_loan
    build:
      context: .
      dockerfile: Dockerfile.timeloanservice
#    ports:
#      - "8484:8181"
    depends_on:
      - bot-listener
    networks:
      - slackEvent

networks:
  slackEvent:

~~~


---
## Config file


~~~ json
{
  "databasetype": "mongodb", // 디비종류선택
  "dbconnection": <connectionstring>,
  "restfulapi_endpoint": "localhost:3000"// restapi 오픈번호,
  "slack_token": <slack_token>,
  "kafka_message_brokers": [<kafka_sever1>, .....]
}

~~~

---- 

블로그 주소 
https://velog.io/write?id=86d4dd88-3b0c-4e2d-86f9-b9827ff09e1b

