Schubladendenken
================

Install
-------

```bash
go get github.com/rrawrriw/go-drawer-lib
go get github.com/rrawrriw/go-drawer-cmd

go install github.com/rrawrriw/go-drawer-cmd

# Kopiere go-drawer-cmd executable in ein Verzeichnis das in PATH gelistet ist
cp ${GOPATH}/bin/go-drawer-cmd ~/bin/go-drawer # maybe
```

Einführung
----------

Wenn auch du gerne Dinge in Schubladen steckst dann bist du hier richtig. Schublade auf, Go Projekt rein. Damit es nicht zu einem Kutelmuttel kommt zwischen allen den vielen Versionen für alle die fantastische Go Libraries. Kannst du jetzt jedes Go Projekt in einer gesondert Verzeichnis Struktur Verwalten so wie du es bist jetzt auch, mehr oder weniger, getan hast. Mit einem kleinen Unterschied.

```bash
mkdir -p shaloin/src/github.com/monk/shaloin-lib # crazy Projekt
cd !$
```
Jetzt kommt der kleine Unterschied

```bash
eval $(go-drawer)
```
Jetzt gehts es weiter wie immer

```bash
go get github.com/monk/ohmmmm
```
*... programmiere unglaubliches Zeug  tip tip tip...*

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
export PATH=${PATH}:${GOPATH}/bin
```

Um das ganze noch einfacher zu machen, geht das überhaupt nocht, ja! Füge ich noch diese 3 Simplen Zeilen zu meiner .bashrc hinzu.

```bash
function god() {
  eval $(go-drawer);
}
```

Zack, fertig!

Die Idee kommt von dem [virtualenvwrapper](https://virtualenvwrapper.readthedocs.org/) Projekt aus der Python Welt.
