# Annuaire Téléphonique

Un programme simple en Go pour gérer un annuaire téléphonique en ligne de commande.

## Fonctionnalités

- Ajouter une personne à l'annuaire
- Rechercher une personne par nom ou prénom
- Lister toutes les personnes de l'annuaire
- Modifier une personne de l'annuaire
- Supprimer une personne de l'annuaire

## Installation

```bash
git clone https://github.com/SwannMartineau/go/tp1-note.git
cd tp1-note
```

## Utilisation

### Ajouter une personne

```bash
go run main.go --action ajouter --nom "Dupont" --prenom "Jean" --tel "0123456789"
```

### Rechercher une personne

Par nom et prénom :

```bash
go run main.go --action rechercher --nom "Dupont" --prenom "Jean" --tel "0675487567"
```

### Modifier une personne

Par nom et prénom :

```bash
go run main.go --action modifier --nom "Dupont" --prenom "Jean" --tel "0675487567"
```

### Supprimer une personne

Par nom et prénom et tel:

```bash
go run main.go --action modifier --nom "Dupont" --prenom "Jean" --tel "0675487567"
```

### Lister toutes les personnes

```bash
go run main.go --action lister
```

## Structure du package repository

Le package `repository` contient les structures et fonctions suivantes :

### Structures

- `Person` : Représente une personne avec son nom, prénom et numéro de téléphone
- `Repository` : Structure pour gérer une collection de personnes

### Fonctions principales

- `AddPerson` : Ajoute une personne à l'annuaire
- `DeletePerson` : Supprime une personne de l'annuaire
- `ModifyPerson` : Modifie le numéro de téléphone d'une personne
- `SearchPerson` : Recherche une personne par nom ou prénom
- `ListRepository` : Affiche toutes les personnes de l'annuaire

## Tests

Pour exécuter les tests depuis le repository:

```bash
go test
```
