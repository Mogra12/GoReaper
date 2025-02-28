# GoReaper

GoReaper é uma ferramenta de brute-force desenvolvida em Go para testar credenciais de FTP e SSH. Suporta ataques simultâneos e pode utilizar TLS para conexões seguras.

## Recursos

- Suporte a ataques de brute-force em FTP e SSH.
- Execução simultânea com limite configurável de conexões.
- Opção para conexão FTPS (FTP sobre TLS).
- Configuração flexível via linha de comando.
- Suporte para listas de usuários e senhas separadas no ataque SSH.

## Instalação

1. Clone este repositório:
   ```bash
   git clone https://github.com/seuusuario/GoReaper.git
   cd GoReaper
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Compile o programa:
   ```bash
   go build -o goreaper main.go
   ```

## Uso

### Ataque FTP

```bash
goreaper -Cn 10 -tls -w wordlist.txt -t target:21 -time 2
```

#### Parâmetros FTP:
- `-Cn`   : Define o número máximo de conexões simultâneas.
- `-tls`  : Usa FTPS (FTP sobre TLS).
- `-w`    : Caminho para o arquivo da wordlist.
- `-t`    : Endereço do servidor FTP (exemplo: hostname:21).
- `-time` : Tempo de espera em segundos entre tentativas de login (quando `lwr` está ativo).

### Ataque SSH

```bash
goreaper -U users.txt -P passwords.txt -t target:22
```

#### Parâmetros SSH:
- `-U` : Caminho para a wordlist de usuários.
- `-P` : Caminho para a wordlist de senhas.
- `-T` : Endereço do servidor SSH (exemplo: target:22).

## Contribuição

Sinta-se à vontade para abrir issues e enviar PRs para melhorias!

## Aviso Legal

**Esta ferramenta deve ser usada apenas para testes de segurança em ambientes autorizados. O uso indevido pode resultar em penalidades legais.**

