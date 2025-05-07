## **user**
Tabela de usuários do sistema. Cada usuário representa um profissional com acesso individual (multi-tenant). Suporta autenticação via e-mail/senha, controle de acesso por função e histórico de login.

---

### 📦 Campos

| Campo            | Tipo      | Descrição                                                                 |
|------------------|-----------|---------------------------------------------------------------------------|
| id               | uuid      | Identificador único do usuário                                            |
| name             | string    | Nome completo do profissional                                             |
| email            | string    | E-mail utilizado para login (único)                                      |
| password_hash    | string    | Hash da senha (armazenado via bcrypt ou argon2)                          |
| role             | string    | Papel do usuário (`professional`, `admin`, `secretary`, etc.)            |
| phone            | string?   | Telefone para contato (opcional)                                         |
| is_active        | bool      | Indica se a conta está ativa                                              |
| created_at       | datetime  | Data de criação da conta                                                  |
| updated_at       | datetime  | Última atualização do cadastro                                            |
| last_login_at    | datetime? | Data/hora do último login (para fins de auditoria)                       |

---

### 🔐 Regras de Negócio

1. O **e-mail deve ser único** e é utilizado como identificador de login.
2. A **senha nunca é armazenada em texto puro**, somente como hash (ex: bcrypt).
3. O campo `is_active` define se o usuário pode autenticar no sistema.
4. O campo `role` pode ser usado para controle de permissões em rotas e funcionalidades:
    - `professional`: profissional autônomo (padrão do MVP)
    - `admin`: acesso total ao sistema (futuro)
    - `secretary`, `viewer`, etc.: funções adicionais opcionais
5. O `id` do usuário pode ser utilizado como **`client_id` para isolar os dados** de pacientes, agendamentos, financeiro, etc. (multi-tenant lógico).
6. A data `last_login_at` pode ser atualizada a cada autenticação bem-sucedida para fins de auditoria e relatórios.
7. Usuários com `is_active = false` não podem realizar login, mesmo com credenciais corretas.
8. Todas as entidades do sistema (pacientes, leads, agendamentos, pagamentos, etc.) devem ter associação com o `user_id`, garantindo que **cada profissional só visualize e acesse seus próprios dados**.

---

### 📌 Observação

Este modelo é compatível com autenticação baseada em **JWT**, utilizando o `user.id` no payload do token:

```json
{
  "sub": "uuid-do-usuario",
  "role": "professional",
  "exp": 1712345678
}
```

# 📋 Estrutura de Anamnese

---

## **anamnese_template**
Representa o modelo de anamnese configurado por cada cliente.

| Campo       | Tipo    | Descrição                       |
|-------------|---------|----------------------------------|
| id          | uuid    | Identificador único              |
| title       | string  | Título do template               |
| client_id   | uuid FK | Referência ao cliente (tenant)   |

---

## **anamnese_field**
Campos personalizados que compõem o template de anamnese.

| Campo         | Tipo    | Descrição                                    |
|---------------|---------|-----------------------------------------------|
| id            | uuid    | Identificador único                           |
| field_number  | int     | Ordem de exibição                             |
| field_type    | string  | Tipo do campo (date, datetime, text, number, checkbox) |
| field_title   | string  | Título ou pergunta                            |
| field_required| bool    | Se é obrigatório                              |
| field_active  | bool    | Se está ativo ou não                          |
| client_id     | uuid FK | Referência ao cliente                         |
| anamnese_id   | uuid FK | Referência ao template (anamnese_template)   |

---

## **patient_anamnese**
Representa uma resposta preenchida para um paciente com base em um template.

| Campo        | Tipo    | Descrição                                  |
|--------------|---------|---------------------------------------------|
| id           | uuid    | Identificador único                         |
| patient_id   | uuid FK | Referência ao paciente                      |
| anamnese_id  | uuid FK | Referência ao template (anamnese_template) |
| client_id    | uuid FK | Referência ao cliente                       |
| answered_at  | datetime | Data/hora do preenchimento                 |

---

## **patient_anamnese_field**
Respostas individuais de cada campo da anamnese preenchida.

| Campo                 | Tipo    | Descrição                                   |
|-----------------------|---------|----------------------------------------------|
| id                    | uuid    | Identificador único                          |
| patient_anamnese_id   | uuid FK | Referência à resposta geral (patient_anamnese) |
| field_id              | uuid FK | Referência ao campo (anamnese_field)         |
| value                 | text    | Valor preenchido pelo paciente               |

---

> ℹ️ Observações:
> - O campo `value` é do tipo texto para permitir flexibilidade. A interpretação (data, número, booleano, etc.) deve ser feita conforme o `field_type`.
> - Esta estrutura é pensada para suportar múltiplos clientes (multi-tenant).

# 📅 Estrutura de Agendamento, Sessão e Evolução

---

## **appointment**
Representa um agendamento de atendimento entre um profissional e um paciente. Pode ou não resultar em uma sessão realizada.

| Campo            | Tipo     | Descrição                                                                 |
|------------------|----------|---------------------------------------------------------------------------|
| id               | uuid     | Identificador único do agendamento                                       |
| client_id        | uuid FK  | Identificador do cliente (tenant)                                        |
| patient_id       | uuid FK  | Identificador do paciente                                                |
| custom_repasse_type | string? | Tipo do repasse para AQUELE atendimento  `percent` ou `fixed` |
| custom_repasse_value | decimal? | Valor do repasse do atendimento |
| professional_id  | uuid FK  | Identificador do profissional responsável pelo atendimento               |
| cost_center_id | uuid FK | Identificador da origem do agendamento (Clinica X, Atendimento Proprio, etc) |
| service_title    | string   | Nome do serviço agendado (ex: "Psicoterapia Cognitiva")                  |
| start_time       | datetime | Data e hora de início do atendimento                                     |
| end_time         | datetime | Data e hora de término do atendimento                                    |
| status           | string   | Estado do agendamento: `scheduled`, `done`, `canceled`, `no_show`        |
| notes            | text     | Observações gerais do agendamento                                        |
| created_at       | datetime | Data/hora de criação do registro                                         |
| updated_at       | datetime | Data/hora da última atualização                                          |

---

## **session**
Criada automaticamente ou manualmente quando o agendamento é marcado como realizado (`done`). Representa uma sessão de fato ocorrida.

| Campo            | Tipo     | Descrição                                                                 |
|------------------|----------|---------------------------------------------------------------------------|
| id               | uuid     | Identificador único da sessão                                            |
| appointment_id   | uuid FK  | Referência ao agendamento correspondente                                 |
| client_id        | uuid FK  | Identificador do cliente (tenant)                                        |
| patient_id       | uuid FK  | Identificador do paciente                                                |
| professional_id  | uuid FK  | Identificador do profissional                                            |
| start_time       | datetime | Data/hora de início real da sessão (pode copiar do agendamento)          |
| end_time         | datetime | Data/hora de término real da sessão                                      |
| was_attended     | bool     | Indica se o paciente de fato compareceu                                  |
| created_at       | datetime | Data/hora de criação do registro                                         |

---

## **evolution**
Registro clínico (anotações, evolução terapêutica) gerado **somente se a sessão foi realizada**.

| Campo            | Tipo     | Descrição                                                                 |
|------------------|----------|---------------------------------------------------------------------------|
| id               | uuid     | Identificador único da evolução clínica                                   |
| session_id       | uuid FK  | Referência à sessão correspondente                                       |
| client_id        | uuid FK  | Identificador do cliente (tenant)                                        |
| professional_id  | uuid FK  | Identificador do profissional que escreveu a evolução                    |
| patient_id       | uuid FK  | Identificador do paciente                                                |
| content          | text     | Texto da evolução clínica (anotações livres, plano terapêutico, etc.)    |
| created_at       | datetime | Data/hora de criação do registro                                         |

---

## 🔄 Regras de negócio

- A **sessão só é criada** quando o `appointment.status = done` (ou similar como `in_progress`).
- A **evolution** só pode existir se houver uma `session`.
- Não é permitido registrar evolução para sessões ausentes ou agendamentos não realizados.

---



## 📦 Entidades e Relacionamentos

---

## **cost_center**
Define a origem do atendimento (clínica, particular, convênio, instituição etc.) e a regra de repasse associada.

| Campo           | Tipo     | Descrição                                                                 |
|------------------|----------|---------------------------------------------------------------------------|
| id               | uuid     | Identificador único                                                       |
| user_id          | uuid FK  | Profissional dono da conta (tenant)                                       |
| name             | string   | Nome da origem (ex: "Clínica X", "Particular", "APAE")                    |
| repasse_model    | string   | `clinic_pays` ou `professional_pays`                                     |
| repasse_type     | string   | `percent` ou `fixed`                                                      |
| repasse_value    | decimal  | Valor ou percentual padrão do repasse                                     |
| active           | bool     | Indica se está disponível para uso                                        |

---

## **payment**
Registro de entrada financeira. Pode ou não estar associado a agendamentos.

| Campo          | Tipo      | Descrição                                                                |
|----------------|-----------|---------------------------------------------------------------------------|
| id             | uuid      | Identificador único                                                       |
| user_id        | uuid FK   | Profissional que recebeu                                                  |
| patient_id     | uuid FK   | Paciente (pode ser null em pagamentos avulsos institucionais)            |
| cost_center_id | uuid FK   | Origem do pagamento                                                       |
| payment_date   | datetime  | Data do recebimento                                                       |
| amount         | decimal   | Valor total recebido                                                      |
| method         | string    | Forma de pagamento: `pix`, `dinheiro`, `cartão`, etc.                     |
| notes          | text      | Observações ou referência de origem (ex: "APAE Março")                   |
| created_at     | datetime  | Registro de criação                                                       |

---

## **payment_appointment**
Vincula um pagamento a uma ou mais sessões específicas. Opcional.

| Campo           | Tipo     | Descrição                                                   |
|-----------------|----------|--------------------------------------------------------------|
| id              | uuid     | Identificador único                                          |
| payment_id      | uuid FK  | Pagamento correspondente                                     |
| appointment_id  | uuid FK  | Sessão coberta por esse pagamento                            |

---

## **repasse**
Define o valor que o profissional precisa repassar à clínica ou instituição. Pode ser apenas informativo.

| Campo            | Tipo      | Descrição                                                                 |
|------------------|-----------|---------------------------------------------------------------------------|
| id               | uuid      | Identificador único                                                       |
| user_id          | uuid FK   | Profissional responsável                                                  |
| appointment_id   | uuid FK   | Sessão relacionada                                                        |
| cost_center_id   | uuid FK   | Origem da sessão                                                          |
| value            | decimal   | Valor do repasse (calculado ou fixado)                                   |
| clinic_receives  | bool      | Se `true`, o profissional precisa repassar. Se `false`, valor já retido  |
| status           | string    | `pending`, `paid`, `informational`                                       |
| paid_at          | datetime? | Data de pagamento (se aplicável)                                          |
| notes            | text      | Observações ou referência                                                 |

---

## 🔁 Regras de Negócio

---

### 🔹 Repasse padrão por origem
- Cada `cost_center` define o tipo (`percent` ou `fixed`) e valor padrão.
- O campo `repasse_model` define:
    - `clinic_pays`: clínica retém e repassa, profissional **não paga** → repasse é apenas **informativo**
    - `professional_pays`: profissional recebe e **precisa pagar** parte para a origem

---

### 🔹 Repasse personalizado por agendamento
- O agendamento (`appointment`) pode sobrescrever o repasse do cost center.
- Prioridade para cálculo do `repasse`:
    1. Se `custom_repasse_value` estiver preenchido → usar esse
    2. Senão, usar o padrão do `cost_center`

---

### 🔹 Geração de repasse
- Sempre que um `appointment` é marcado como `done`, um `repasse` é gerado.
- Se `clinic_pays`, o repasse é `informational` e `clinic_receives = false`
- Se `professional_pays`, o repasse é `pending` com `clinic_receives = true`

---

### 🔹 Pagamentos e vínculo com sessões
- Um `payment` pode cobrir uma ou várias sessões via `payment_appointment`
- É possível registrar **pagamentos avulsos** (sem sessões vinculadas), como:
    - Remuneração mensal (APAEs, escolas, convênios)
    - Acordos por carga horária
- O vínculo é **opcional** — se não houver `payment_appointment`, o pagamento é considerado **avulso**

---

### 🔹 Relatórios possíveis
- Sessões pagas vs. não pagas
- Repasses pendentes e pagos
- Total recebido por período / origem
- Total de comissões (valores "retidos" ou "repassados")
- Ganhos líquidos reais por atendimento

---


## 📦 Entidades e Relacionamentos

---

## **lead**
Pré-cadastro de uma pessoa interessada em iniciar atendimento. Pode ser convertida em paciente.

| Campo            | Tipo      | Descrição                                                                 |
|------------------|-----------|---------------------------------------------------------------------------|
| id               | uuid      | Identificador único do lead                                               |
| user_id          | uuid FK   | Profissional que fez o atendimento/conversa                               |
| full_name        | string    | Nome da pessoa                                                            |
| phone            | string?   | Telefone de contato                                                       |
| email            | string?   | E-mail (opcional)                                                         |
| birth_date       | date?     | Data de nascimento (opcional)                                             |
| contact_date     | datetime  | Data da conversa ou triagem inicial                                       |
| status           | string    | `new`, `in_analysis`, `converted`, `lost`                                 |
| was_attended     | bool      | Se já passou por triagem ou conversa inicial                              |
| converted_at     | datetime? | Data de conversão em paciente (se aplicável)                              |
| notes            | text?     | Anotações sobre a conversa, objetivos, possíveis encaminhamentos          |
| origin           | string?   | Origem do contato (Instagram, indicação, site etc.)                       |
| gdpr_block_contact | bool    | Se o lead não autoriza contato futuro (LGPD/GDPR compliance)              |
| created_at       | datetime  | Data de criação                                                           |
| updated_at       | datetime  | Última atualização                                                        |

---

## **patient**
Paciente com atendimento ativo ou anterior. Pode vir de um `lead`.

| Campo                  | Tipo      | Descrição                                                                 |
|------------------------|-----------|---------------------------------------------------------------------------|
| id                     | uuid      | Identificador único do paciente                                           |
| user_id                | uuid FK   | Profissional responsável pelo paciente                                    |
| cost_center_id         | uuid FK   | Origem padrão do atendimento do paciente                                  |
| full_name              | string    | Nome completo                                                             |
| social_name            | string?   | Nome social (opcional)                                                    |
| birth_date             | date      | Data de nascimento                                                        |
| document               | string?   | Documento (CPF, RG etc.)                                                  |
| phone                  | string?   | Telefone (WhatsApp, celular)                                              |
| email                  | string?   | E-mail de contato                                                         |
| gender                 | string?   | Gênero                                                                    |
| address                | string?   | Endereço                                                                  |
| resides_with           | string?   | Com quem o paciente reside (pai, mãe, ambos, avós, etc.)                  |
| emergency_contact_name | string?   | Nome da pessoa para contato de emergência                                 |
| emergency_contact_phone| string?   | Telefone do contato de emergência                                         |
| observation            | text?     | Observações clínicas ou administrativas                                   |
| default_repasse_type   | string?   | `percent` ou `fixed` (regra personalizada de repasse)                     |
| default_repasse_value  | decimal?  | Valor ou percentual do repasse personalizado                              |
| active                 | bool      | Se o paciente está ativo                                                  |
| created_at             | datetime  | Data de criação                                                           |
| updated_at             | datetime  | Última atualização                                                        |

---

## **patient_family**
Relação de membros familiares do paciente (útil para menores ou acompanhamento do sistema familiar).

| Campo        | Tipo      | Descrição                                                   |
|--------------|-----------|--------------------------------------------------------------|
| id           | uuid      | Identificador único                                          |
| patient_id   | uuid FK   | Vínculo com o paciente                                       |
| relationship | string    | Parentesco: `pai`, `mãe`, `cônjuge`, `filho`, `responsável` |
| name         | string    | Nome do familiar                                             |
| birth_date   | date?     | Data de nascimento (opcional)                                |
| schooling    | string?   | Escolaridade (opcional)                                      |
| occupation   | string?   | Ocupação profissional (opcional)                             |


---

## 🔁 Regras de Conversão de Lead para Paciente

1. Leads representam possíveis pacientes que passaram por triagem ou conversa inicial.
2. Leads podem ser convertidos manualmente para pacientes com um clique.
3. Ao converter:
    - Cria-se um `patient` com os dados preenchidos no `lead`.
    - O campo `lead.status` é atualizado para `converted`.
    - O campo `lead.converted_at` é preenchido.
    - O `lead` pode ser mantido para histórico e estatísticas.
4. O profissional pode registrar observações, origem, e notas durante o período de análise do lead.
5. Após convertido, o `patient` passa a participar de todos os fluxos do sistema (agendamento, sessão, evolução, pagamento, etc.)

---

## 💸 Regras de Pagamento de Leads (triagem paga)

1. Um `lead` pode realizar um atendimento prévio (ex: triagem ou conversa inicial), com ou sem cobrança.
2. Se for cobrada, o pagamento deve ser registrado na tabela `payment` como **pagamento avulso**:
    - O campo `patient_id` deve ser `null`, pois o vínculo ainda não existe.
    - O campo `lead_id` pode ser usado para vincular diretamente ao lead.
    - O campo `notes` pode complementar com texto identificador (ex: "Triagem - Maria Oliveira").
    - O campo `cost_center_id` pode apontar para um centro como "Triagem" ou "Particular".
3. O pagamento avulso **não gera agendamento** e **não cria sessão**.
4. Se o lead for convertido posteriormente, os pagamentos podem ser consultados com base no `lead_id`.
5. Pagamentos feitos por leads são incluídos nos relatórios financeiros como **recebimentos não vinculados a paciente fixo**.

---

## 🧠 Regras de Visualização e Uso

1. Leads devem ter uma tela de listagem separada da lista de pacientes.
2. A conversão para paciente deve preservar os dados e adicionar apenas os campos obrigatórios de `patient`.
3. Leads não aparecem no calendário de agendamentos (exceto se for implementado agendamento de triagem).
4. É possível filtrar os leads por:
    - Origem de contato
    - Status (`new`, `in_analysis`, `converted`, `lost`)
    - Data de contato
    - Leads com `was_attended = true` (triagem realizada)
5. O sistema deve permitir registrar observações contínuas durante o status `in_analysis`.
6. Leads com `gdpr_block_contact = true` devem ser respeitados em comunicações automáticas.
7. Leads inativos por muito tempo podem ser anonimizados, arquivados ou removidos.
8. A ficha do paciente deve permitir visualizar:
    - Dados de residência
    - Contato de emergência
    - Composição familiar (usando `patient_family`)

---

## 📊 Benefícios e Relatórios

- Relatório de leads por origem: quantos vieram de Instagram, indicação etc.
- Taxa de conversão de leads para pacientes
- Valor total recebido em triagens/consultas iniciais
- Histórico de decisões clínicas (por que não houve continuidade)
- Registro clínico inicia **apenas após conversão**, mantendo a base de pacientes limpa
- Permite segmentar pacientes por **canal de aquisição**, **interesse**, **potencial de fidelização**
- Permite controlar leads que realizaram triagem, pagaram, e decidiram não continuar
- Suporte completo a pacientes crianças ou dependentes, com estrutura familiar clara
