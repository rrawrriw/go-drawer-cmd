go-venv
=======

Setzte standard Pfad in der .bashrc

```bash
export GOPATH=/to/path
export PATH=$PATH:${GOPATH}/bin
```

Hat man nun ein größeres Projekt mit mehreren Abhänigkeiten kann man einfach ein neues Verzeichnis erzeugen in dem ein src Verzeichns enthalten ist. Wenn man nun in diese Verzeichnis wechsel kann man simple

```bash
eval $(go-venv)
```

ausführen und setzt so GOPATH und PATH neu. Man kann so wie gewohnt die go tool chain verwenden und muss nicht gb als "Proxy" verwenden.

Um das ganze noch einwenig zu vereinfachen wäre es möglich in .bashrc folgendes einzufügen.

```bash
function gov() {
  eval $(go-venv);
}
```

```bash
go test
go get ...
```

Die Idee kommt von dem python virtualenv Projekt.

