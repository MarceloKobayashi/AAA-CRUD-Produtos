# Link para documentação Postman

https://documenter.getpostman.com/view/34285293/2sA3QngYjA

# ps-frontend-marcelo-honda-lucas-vieira
Projeto Semestral de Desenvolvimento de Software: Implementação End-to-End  

Objetivo do Trabalho
Desenvolver uma aplicação, utilizando Vue.js para o front-end, Golang para o back-end, e MongoDB como banco de dados, com uma integrações end-to-end das funcionalidades especificadas. O projeto inclui a configuração inicial dos ambientes de desenvolvimento, a definição clara das funcionalidades via requests e responses HTTP documentadas no Postman, e o desenvolvimento incremental da aplicação ao longo de seis semanas. O projeto deve conduzido aplicando os preceitos de metodologia ágil para desenvolvimento de sistemas.

Forma de Avaliação do Projeto
Os projetos são avaliados após o final de cada sprint (semanalmente), sendo considerado na avaliação além dos aspectos técnicos a aplicabilidade da metodologia ágil. Assim, o resultado final do projeto tem relação com a melhoria continua do projeto, semana a semana, e não somente por sua entrega final e funcionalidade. Ademais os alunos são incentivados a aplicar as práticas vistas em sala de aula para garantia de manutenabilidade e extensabilidade da solução. Alguns pontos importantes:

    Projetos não documentados, não serão considerados para avaliação.
    Projetos que não evoluirem entre sprints receberam menção SR (sem-rendimento). 

 


[Sprint 0] Documentação API e Configuração Inicial
Front-End: Deve ser criado um repositório no GitLab configurado para um projeto Vue.js. O repositório deve incluir todas as dependências necessárias e uma estrutura básica de projeto pronta para desenvolvimento. O nome de repositório deve definido seguindo o template:

 

    Trabalho individual: “ps-frontend-[nome-sobrenome]”
    Trabalho em dupla: “ps-frontend-[nome-sobrenome]-[nome-sobrenome]” 

Back-End: Similarmente, deve ser criado um repositório GitLab para o projeto back-end em Golang, com as devidas configurações para runtimes e dependências iniciais.O nome de repositório deve definido seguindo o template:

    Trabalho individual: “ps-backend-[nome-sobrenome]”
    Trabalho em dupla: “ps-backend-[nome-sobrenome]-[nome-sobrenome]” 

Postman: Disponibilizar no README.md o link publicado da documentação do seu sistema, contendo as requests e responses definidas para as funcionalidades desejadas do sistema. O foco deve estar nas principais funcionalidades do sistema. Cada request deve estar claramente descrita e documentada para que o desenvolvimento posterior seja realizado com base nessas definições. Devem ser definidas no mínimo 10 requests para o projeto, não há limite máximo. Não se preocupe se caso alguma requisição precisar ser alterada durante o desenvolvimento do projeto, isso faz parte do desenvolvimento ágil.

Escrevendo Documentação do Tema: https://learning.postman.com/docs/publishing-your-api/document-a-collection/
Documentando Requests e Responses: https://learning.postman.com/docs/publishing-your-api/documenting-your-api/

[Sprint 1 - 5] Desenvolvimento Vertical (End-to-End)

     Nas semanas subsequentes, os alunos devem focar no desenvolvimento end-to-end de funcionalidades específicas da aplicação. Cada semana deverá focar na implementação de uma parte do conjunto de funcionalidades especificadas inicialmente na documentação publicado através do Postman.
     As entregas devem ser incrementais e cada uma delas deve demonstrar o progresso, incluindo:
         A(s) funcionabilidade(s) operante no front-end e back-end.
         Integração e testes demonstrando a comunicação efetiva entre o front-end e o back-end através das requests definidas.
         Atualização de documentação e código no repositório conforme o progresso. 

Como garantir que funcionalidade está adequadamente implementada?
Isso pode ser um pouco desafiador, por isso coloquei uma lista atividades que se forem feitas garantirão uma maior assertividade nas entregas.
Chamei esse framework de ONION Agile, define o seguinte fluxo de trabalho:

Para cada endpoint:

1. Defina o design do front-end para a funcionalidade
4. Implemente a camada HTTP do endpoint no back-end
5. Implemente o design no front-end
6. Codifique o comportamento interativo do front-end
7. Implemente validações no front-end
8. Implemente a lógica de processamento da request no back-end (handler)
9. Desenvolva a interação com o banco de dados
10. Implemente validações no back-end
13. Realize testes de integração usando Postman
14. Valide manualmente a implementação end-to-end e recolha feedback
15. Refatore o código para eliminar repetições utilizando padrões de projeto
16. Execute todos os testes para garantir que nenhuma funcionalidade foi afetada pela refatoração
17. Documente continuamente o processo de desenvolvimento e o código 

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
