# 📡 Camada de Enlace com Código de Hamming (Go)

Esse projeto implementa uma simulação da **camada de enlace** utilizando dois processos distintos: **remetente** e **destinatário**, comunicando-se via **pipes**. O objetivo é simular um protocolo simples de framing com **código de Hamming** para detecção e correção de erros de 1 bit.

---

## 📋 Estrutura do Projeto

```
.
├── destinatario/
│   └── main.go           # Código do receptor (destinatário)
├── hamming/
│   └── hamming.go      
├── remetente/
│   └── main.go           # Código do transmissor (remetente)
├── go.mod
└── README.md
```

---

## 🔧 Como Executar

### 1. Clone o repositório

```bash
git clone https://github.com/lucasdeluccas/Ponderada_Enlace.git
cd Ponderada_Enlace
```

### 2. Compile os programas

```bash
go build -o remetente ./remetente
go build -o destinatario ./destinatario
```

### 3. Execute com pipe

```bash
./remetente 01101001 | ./destinatario
```

---

## ✉️ Funcionamento do Protocolo

### Remetente:
- Recebe uma sequência binária (`01101001`) como argumento.
- Gera o código de Hamming com bits de paridade.
- Adiciona um **cabeçalho** e um **terminador** ao frame (`01111110`).
- Envia o frame pelo `stdout`.

### Destinatário:
- Recebe o frame pelo `stdin`.
- Valida o cabeçalho e terminador.
- Simula um erro proposital (inversão de 1 bit).
- Aplica o algoritmo de Hamming para corrigir.
- Exibe:
  - Código com erro
  - Código corrigido
  - Dados extraídos

---

## 🧠 Exemplo de Execução

### Comando:
```bash
./remetente 01101001 | ./destinatario
```

### Saída:

```
Content with error: 011010010010
Corrected code:     011110010010
Extracted data:     01101001
```

---

## 📐 Protocolo de Frame

| Campo        | Valor         |
|--------------|---------------|
| Cabeçalho    | `01111110`    |
| Payload      | Dados com Hamming |
| Terminador   | `01111110`    |

---

## 🛡️ Detecção e Correção

Utilizamos **Código de Hamming (SEC - Single Error Correction)**:
- Corrige 1 bit de erro
- Detecta 2 bits de erro (sem corrigir)

---

## 🎥 Demonstração em Vídeo

📹 [Clique aqui para assistir ao vídeo explicativo](https://youtu.be/seu-video-aqui)

