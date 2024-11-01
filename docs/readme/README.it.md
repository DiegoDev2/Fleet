![Logo](./Logo.png)

# Fleet

**Fleet** è un gestore di pacchetti sviluppato in Go, progettato per installare, gestire e configurare vari strumenti in modo semplice ed efficiente. Permette di scaricare e installare strumenti direttamente dai repository.

### Caratteristiche

- Installazione leggera ed efficiente degli strumenti.
- Comandi semplici per installare, aggiornare e rimuovere strumenti.
- Formule personalizzabili per diversi strumenti.

### Installazione

1. Clona il repository:

```bash
   git clone https://github.com/DiegoDev2/Fleet
   cd Fleet
   go build -o fleet
   sudo mv fleet /usr/local/bin
```

### Utilizzo

Per installare uno strumento, usa:

```bash
fleet install <nome-strumento>
```

Ad esempio, per installare nmap:

```bash
fleet install nmap
```


### Aggiungere Nuovi Strumenti

- Apri `lib/structTool.go`,`lib/toolStruc.go` e aggiungi il nuovo strumento all'elenco degli strumenti disponibili.
- Crea una formula per lo strumento nella directory `formulas/`.
- Utilizza il comando install per aggiungerlo al tuo sistema.


### Contribuire

Apprezziamo i contributi per migliorare Fleet. Se trovi un errore o hai un suggerimento, non esitare a inviare una pull request o aprire un issue.

### Licenza

Fleet è concesso in licenza sotto la licenza Apache 2.0. Consulta il file LICENSE per ulteriori dettagli.

![imagen](https://github.com/user-attachments/assets/30584687-bafe-415e-978d-dc568337c75d)

