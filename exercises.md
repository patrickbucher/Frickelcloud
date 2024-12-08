# Aufgaben

1. **Erstellen Sie einen Fork von diesem Repository!**
2. **Klonen Sie Ihren Fork, nicht das Original-Repository!**
3. **Reichen Sie Ihre Lösungen per Pull Request ein!**

Die folgenden Aufgaben können Sie direkt in die angegebene Datei lösen. Beachten
Sie hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## Aufgabe 0) Sich mit dem Code vertraut machen

siehe `ex0.md`

Öffnen Sie die Datei `frickelbude.go`. Und betrachten Sie den Code. Die Datei enthält folgende Definitionen:

1. Preisangaben
    - `PricingCPU`: Die monatlichen Kosten für die CPUs
    - `PricingRAM`: Die monatlichen Kosten für das Memory
    - `PricingSSD`: Die monatlichen Kosten für den Storage
2. Datentypen
    - `Server`: Die Struktur enthält Angaben für CPU, RAM und SSD und kann die Konfiguration sowohl für physische wie auch virtuelle Maschinen abbilden.
    - `GuestInventory`: Die Map speichert die Konfiguration für virtuelle Maschinen unter deren Namen ab.
    - `Host`: Die Struktur kombiniert die Hardware eines physischen Servers mit einer Reihe von Gast-VMs.
    - `HostInventory`: Die Map speichert die Konfiguration für physische Maschinen (inkl. Gast-Inventar) unter deren Namen ab.
3. Rechenzentrum
    - `DataCenter`: Dieses `HostInventory` ist der Hardware-Anfangsbestand im Rechenzentrum der FrickelCloud. Es gibt zu Beginn nur drei Server, auf denen je mehrere Gäste-VMs laufen.

Beantworten Sie die folgenden Fragen um Ihr Verständnis für diese Definitionen zu überprüfen:

1. Wie heissen die Gast-VMs auf dem Host `small-1`?
2. Wie viele CPU-Kerne beanspruchen die Gast-VMs auf dem Host `medium-1`? 
3. Über wie viel unzugeordneten SSD-Speicherplatz verfügt der Host `big-1`?

Schreiben Sie Ihre Antworten in die Datei `ex0.md`.

## Aufgabe 1) Freie Ressourcen pro Host berechnen

siehe `ex1.go` und `ex1/main.go`

Diese Aufgabe bezieht sich auf den ersten Prozess (_VM-Provisionierung: Automatisches Einpassen in Hardware_) im [README](README.md#1-vm-provisionierung-automatisches-einpassen-in-hardware).

Zuerst muss man herausfinden können, wie viele freie Ressourcen auf einem Host noch zur Verfügung stehen.

### a) Verfügbare Ressourcen ermitteln

Vervollständigen Sie in `ex1.go` die Methode `AvailableRessources` folgendermassen:

1. Iterieren Sie über alle Gäste des Hosts (`h.Guests`).
2. Summieren Sie die Anzahl der CPUs, die Arbeitsspeichermenge (RAM) und die Grösse des Speichers (SSD) über alle Gäste auf.
3. Subtrahieren Sie diese Summen von der Hardware-Spezifikation des Hosts (`h.Hardware`) und geben Sie eine neue `Server`-Struktur zurück, welche diese Differenz (= freie Ressourcen) enthält.

### b) Prüfen, ob Host eine VM aufnehmen kann

Vervollständigen Sie in `ex1.go` die Methode `CanFitIn` folgendermassen:

1. Rufen Sie die soeben vervollständigte Methode `AvailableRessources` auf dem gegebenen Host (`h`) auf.
2. Prüfen Sie, ob die ermittelten verfügbaren Ressourcen (CPU, RAM, SSD) _alle_ von der gegebenen virtuellen Mashine (`vm`) benötigten Ressourcen übersteigen oder zumindest gleich gross sind und geben Sie einen entsprechenden `bool`-Wert (`true` oder `false`) zurück.

### c) Programm testen

In `ex1/main.go` wird die Methode `AvailableRessources` bereits für alle Hosts aufgerufen. Weiter wird für drei unterschiedliche VMs mithilfe der `CanFitIn`-Methode geprüft, ob Sie von den einzelnen Hosts aufgenommen werden könnte.

Testen Sie Ihre Implementierungen, indem Sie dieses Programm ausführen:

```bash
go run ex1/main.go
```

## Aufgabe 2) VM in passendes Host-Inventar aufnehmen

siehe `ex2.go` und `ex2/main.go`

Diese Aufgabe bezieht sich ebenfalls auf den ersten Prozess (_VM-Provisionierung: Automatisches Einpassen in Hardware_) im [README](README.md#1-vm-provisionierung-automatisches-einpassen-in-hardware).

Mithilfe der Methode `Host.CanFitIn` kann man prüfen, ob eine virtuelle Maschine auf dem jeweiligen Host untergebracht werden kann. Im zweiten Teil geht es darum, die VM tatsächlich auf einem passenden Host unterzubringen.

### a) Hosts mit genügend Ressourcen ermitteln

Vervollständigen Sie in `ex2.go` die Methode `FindFittingHosts` folgendermassen:

1. Iterieren Sie über das `HostInventory`.
    - Bei jedem Schleifendurchlauf erhalten Sie einen Namen und einen `Host`.
2. Prüfen Sie für jeden Host, ob er die gegebene VM aufnehmen kann.
3. Fügen Sie die Hosts mit genügend freien Ressourcen einer Map hinzu, welche Sie anschliessend zurückgeben.
    - Als Key dient der Name des Hosts, als Value der `Host` selber.

### b) VM auf Host unterbringen

Vervollständigen Sie in `ex2.go` die Methode `FitIn` folgendermassen:

1. Überprüfen Sie noch einmal, ob der Host `h` wirklich genügend Platz für die `vm` hat.
    - Geben Sie einen entsprechenden Fehler zurück, falls die verfügbaren Ressourcen zu gering sind.
2. Überprüfen Sie, ob der Host `h` nicht schon eine Gast-VM mit dem gewünschten `name` der neuen VM hat (siehe `h.Guests`).
    - Geben Sie einen Fehler zurück, falls der Name schon vergeben ist.
3. Fügen Sie die `vm` schliesslich dem Inventar des Hosts hinzu (d.h. der Map `h.Guests`).

### c) Programm testen

In `ex2/main.go` wird die Methode `FindFittingHosts` für verschiedene VMs aufgerufen. Anschliessend wird für die jeweiligen passenden Hosts die Methode `FitIn` aufgerufen, um die VM ins Inventar aufzunehmen.

Testen Sie Ihre Implementierungen, indem Sie dieses Programm ausführen:

```bash
go run ex2/main.go
```

## Aufgabe 3) VM aus dem Inventar entfernen

siehe `ex3.go` und `ex3/main.go`

Diese Aufgabe bezieht sich auf den zweiten Prozess (_VM-Ausserbetriebnahme: Automatische Freigabe der Hardware-Ressourcen_) im [README](README.md#2-vm-ausserbetriebnahme-automatische-freigabe-der-hardware-ressourcen).

### a) Die VM auf dem Inventar finden

Der Kunde braucht nicht zu wissen, auf welchem Host seine VM als Gast läuft. Will der Kunde eine VM entfernen, muss der Anbieter jedoch wissen, auf welchem Host sie sich befindet.

Vervollständigen Sie in `ex3.go` die Methode `FindHost` folgendermassen:

1. Iterieren Sie über das gegebene `HostInventory`.
    - Bei jeder Iteration erhalten Sie den Namen des Hosts und den Host selber.
2. Prüfen Sie für jeden Host, ob er in seinem Gäste-Inventar (`Guests`) eine VM mit dem gegebenen Namen `vmName` enthält.
3. Geben Sie den Namen des gefundenen Hosts zurück.
4. Kann der `vmName` auf keinem Host-Inventar gefunden werden, geben Sie einen leeren String und eine entsprechende Fehlermeldung zurück.

### b) Die VM aus dem Inventar entfernen

Vervollständigen Sie in `ex3.go` die Methode `Remove` folgendermassen:

1. Prüfen Sie noch einmal, ob die VM mit dem Namen `vmName` wirklich auf dem Inventar existiert, indem Sie die vorher vervollständigte `FindHost`-Methode aufrufen.
2. Geben Sie einen Fehler zurück, falls die VM nicht gefunden werden kann.
3. Greifen Sie mit dem gefundenen Hostname auf den Host im Inventar zu.
4. Löschen Sie die VM mit dem Namen `vmName` vom Gäste-Inventar des Hosts. (Tipp: Verwenden Sie hierzu die eingebaute [`delete`](https://pkg.go.dev/builtin#delete)-Funktion)

### c) Programm testen

In `ex3/main.go` werden bestehende VMs mithilfe der `Remove`-Methode aus dem Inventar entfernt.

Testen Sie Ihre Implementierungen, indem Sie dieses Programm ausführen:

```bash
go run ex3/main.go
```

## Aufgabe 4) Umsatz berechnen

siehe `ex4.go` und `ex4/main.go`

Diese Aufgabe bezieht sich auf den dritten Prozess (_VM-Abrechnung: Umsatzberechnung_) im [README](README#3-vm-abrechnung-umsatzberechnung).

### a) Den Umsatz pro VM berechnen

Vervollständigen Sie in `ex4.go` die Methode `CalculateRevenue` folgendermassen:

1. Ermitteln Sie die Kosten für die einzelnen VM-Komponenten (CPU, RAM, SSD) anhand der Maps `PricingCPU`, `PricingRAM` bzw. `PricingSSD`.
2. Geben Sie die Summe der Werte zurück.

### b) Den Umsatz aller VMs berechnen

Vervollständigen Sie in `ex4.go` die Methode `CalculateTotalRevenue` folgendermassen:

1. Iterieren Sie über die einzelnen Hosts des Inventars.
2. Iterieren Sie über die Guest-VMs jedes Hosts.
3. Berechnen Sie den Umsatz pro VM mithilfe der zuvor implementierten Methode `CalculateRevenue`.
4. Geben Sie die Summe als Gesamtumsatz zurück.

### c) Programm testen

In `ex4/main.go` wird der Umsatz einiger neuen VMs berechnet und ausgegeben. Weiter wird der Gesamtumsatz der FrickelCloud berechnet und ausgegeben.

Testen Sie Ihre Implementierungen, indem Sie dieses Programm ausführen:

```bash
go run ex4/main.go
```
