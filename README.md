Schubladendenken
================

Wenn auch du gerne Dinge in Schubladen steckst dann bist du hier richtig. Schublade auf, Go Projekt rein.

Damit es nicht zu einem Kutelmuttel kommt zwischen allen den vielen Versionen für alle die fantastische Go Libraries.

Kannst du jetzt jedes Go Projekt in ein Verzeichnis packen ein src Verzeichnis erstellen, genau wie du es jetzt bist auch schon gemacht hast.

Jetzt kommt der kleine Unterschied

```bash
eval $(go-venv)
```

Jetzt gehts es weiter wie immer wie du es bist auch immer gemacht hast 

```bash
go get github.com/ninja/awesome
```

*tip* *tip* *tip*

```bash
go test
```

was man halt so den ganzen Tag macht.

Was macht dieser kleine Unterschied nun? Er geht von dem aktuellen Verzichnis rückwärts die Verzeichnisstruktur nach oben findet er ein Verzeichnis mit einem src Verzeichnis darin wird GOPATH gesetzt sowie PATH um GOPATH/bin erweitert.

Denn wenn was noch wichtiger ist als Dinge in Schubladen zu stecken dann ist es nichts ändern zu müssen.

Meine Schubladen Setup
----------------------

Setzte standard Pfad in der .bashrc

```bash
export GOPATH=/to/path
export PATH=$PATH:${GOPATH}/bin
```

Um das ganze noch einfacher zu machen, geht das überhaupt nocht, ja! Füge ich noch diese 3 Simplen Zeilen zu meiner .bashrc hinzu.

```bash
function god() {
  eval $(go-venv);
}
```

Zack, fertig!

Die Idee kommt von dem [https://virtualenvwrapper.readthedocs.org/](virtualenvwrapper) Projekt aus der Python Welt.
