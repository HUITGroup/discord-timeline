# discord-timeline

# .envの例
Discord のトークンが必須です。

https://discord.com/developers/applications/ からBotのトークンを取得してください。
## 使用方法
1. .envを設定する
   
    ```.env
    DISCORD_TOKEN=<Discord Bot Token>
    MYSQL_ROOT_PASSWORD=rootpassword
    MYSQL_USER=dctimeline
    MYSQL_PASSWORD=dctimeline
    MYSQL_DATABASE=dctimeline
    MYSQL_ADDRESS=localhost
    MYSQL_PORT=3306
    ```

1. コンテナを build する
    
    ```bash
    $ docker build -t discord-timeline:1.0.0 .
    ```

1. docker-compose で MYSQL と一緒に起動

    ```bash
    $ docker-compose up -d
    ```
