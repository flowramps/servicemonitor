
Iniciar um novo projeto e após construir a aplicação:
```
go mod init bifrost
go build -o bifrost
```

Após compilar sua aplicação e executá-la localmente, você poderá acessá-la por meio de um navegador da web ou usando ferramentas de linha de comando como curl ou wget. Aqui está como você pode fazer isso:

Executar sua aplicação: Depois de compilar sua aplicação com sucesso, execute-a no terminal usando o seguinte comando (supondo que você esteja no diretório onde o executável foi criado):

```
./bifrost
```

Este comando iniciará sua aplicação e ela estará ouvindo na porta especificada (no seu caso, a porta 8080).

Acessar via navegador web: Abra um navegador da web e acesse a seguinte URL:

```
http://localhost:8080/metrics
```

Isso deve exibir as métricas que sua aplicação está expondo.

Acessar via linha de comando: Você também pode usar ferramentas de linha de comando como curl ou wget para acessar suas métricas. No terminal, execute:

```
curl http://localhost:8080/metrics
```

```
wget -O - http://localhost:8080/metrics
```

```
http://bifrost.127.0.0.1.nip.io/metrics
```
```
http://bifrost.127.0.0.1.nip.io/
```

Isso retornará as métricas diretamente no terminal.

Esses são os métodos básicos para acessar sua aplicação localmente. Certifique-se de que a aplicação está em execução (após o comando ./bifrost) enquanto você tenta acessá-la.

```
docker build -t rampss/bifrost:latest .
docker push rampss/bifrost:latest
```