<h1 align="center">Navall.io - Back</h1>

<p align="center">
<a href="https://www.metropoledigital.ufrn.br/portal/"><img alt="UFRN - IMD" src="https://img.shields.io/badge/ufrn-imd-ufrn?style=for-the-badge&labelColor=%23164194&color=%230095DB&link=https%3A%2F%2Fwww.metropoledigital.ufrn.br%2Fportal%2F"></a>
<br>
<a href="https://go.dev/"><img alt="Go" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=FFFFFF"></a>
</p>

Aplicação back, desenvolvida com Go, para o jogo Navall.io, um projeto da disciplina de Jogos em Rede Multiusuários.

## Iniciando

Essas instruções lhe darão uma cópia do projeto e um caminho para executá-lo localmente para fins de desenvolvimento e teste.

### Pré-Requisitos

Esse projeto requer um ambiente de desenvolvimento Go na versão 1.22.1 ou superior.

> **Warning**
> Esse projeto usa recursos recentes da linguagem Go, como o [roteamento aprimorado](https://go.dev/blog/routing-enhancements), portanto ele **não** funcionará, em nenhuma hipótese, em versões anteriores à 1.22.

Ademais, esse projeto usa a biblioteca `mattn/go-sqlite3`, que requer a instalação de um compilador C, como o [GCC](https://gcc.gnu.org/), para funcionar. 

#### Instalação

Para instalar as dependências, basta rodar:

```sh
go mod download
```

#### Variáveis de Ambiente

Antes de executar o projeto, é necessário fornecer as variáveis de ambiente adequadas. Existe um arquivo `.env.example` na raiz do repositório com exemplos de valores. Copie o arquivo para o destino `.env`:

```sh
cp .env.example .env
```

Em seguida, edite o arquivo para preencher as variáveis de ambiente:

- `SERVER_ADDRESS`: esse campo deve conter o endereço do servidor, no formato `host:port`. Para fins de desenvolvimento, o valor padrão é `:8080`.

As demais entradas do arquivo são contextuais e, em ambiente de desenvolvimento, devem ser mantidas como estão.

#### Execução

Para rodar o projeto, basta executar o seguinte comando:

```sh
go run .
```

## Licença

Esse projeto é distribuído sob a Licença MIT. Leia o arquivo [LICENSE](LICENSE) para ter mais detalhes.
