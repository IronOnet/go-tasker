# Application CLI Tasker

`tasker` est une application en ligne de commande qui vous aide à gérer vos tâches. Il vous permet de créer, de lister, de marquer comme complétées et de supprimer des tâches. L'application stocke les tâches dans une base de données mongodb.

## Installation

Pour installer `tasker`, vous pouvez télécharger la dernière version binaire disponible sur la [page des versions GitHub](https://github.com/irononet/tasker/releases). Une fois que vous avez téléchargé le binaire, vous pouvez le déplacer dans un répertoire inclus dans la variable d'environnement `PATH` de votre système, tel que `/usr/local/bin` sur les systèmes Unix.

Alternativement, vous pouvez cloner ce dépôt et générer le binaire vous-même en exécutant la commande `go build -o tasker`.

Vous pouvez également créer un conteneur Docker avec toutes les dépendances en exécutant

```bash
sudo docker-compose up
```

## Utilisation

Voici les commandes disponibles dans `tasker` :

### `tasker add (a)`

Ajoute une nouvelle tâche à la liste. Vous serez invité à entrer la description de la tâche.

Exemple :

```
tasker add
```

### `tasker all (l)`

Liste toutes les tâches de la liste.

Exemple :

```
tasker list
```

### `tasker done`

Marque une tâche comme terminée. Vous serez invité à entrer l'ID de la tâche.

Exemple :

```
tasker done <nom_de_la_tâche>
```

### `tasker rm`

Supprime une tâche de la liste. Vous serez invité à entrer l'ID de la tâche.

Exemple :

```
tasker rm <nom_de_la_tâche>
```

### `tasker --help (-h)`

Affiche des informations d'aide sur les commandes disponibles.

Exemple :

```
tasker --help
```

## Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.# Application CLI Tasker

`tasker` est une application en ligne de commande qui vous aide à gérer vos tâches. Il vous permet de créer, de lister, de marquer comme complétées et de supprimer des tâches. L'application stocke les tâches dans une base de données mongodb.

## Installation

Pour installer `tasker`, vous pouvez télécharger la dernière version binaire disponible sur la [page des versions GitHub](https://github.com/irononet/tasker/releases). Une fois que vous avez téléchargé le binaire, vous pouvez le déplacer dans un répertoire inclus dans la variable d'environnement `PATH` de votre système, tel que `/usr/local/bin` sur les systèmes Unix.

Alternativement, vous pouvez cloner ce dépôt et générer le binaire vous-même en exécutant la commande `go build -o tasker`.

Vous pouvez également créer un conteneur Docker avec toutes les dépendances en exécutant

```bash
sudo docker-compose up
```

## Utilisation

Voici les commandes disponibles dans `tasker` :

### `tasker add (a)`

Ajoute une nouvelle tâche à la liste. Vous serez invité à entrer la description de la tâche.

Exemple :

```
tasker add
```

### `tasker all (l)`

Liste toutes les tâches de la liste.

Exemple :

```
tasker list
```

### `tasker done`

Marque une tâche comme terminée. Vous serez invité à entrer l'ID de la tâche.

Exemple :

```
tasker done <nom_de_la_tâche>
```

### `tasker rm`

Supprime une tâche de la liste. Vous serez invité à entrer l'ID de la tâche.

Exemple :

```
tasker rm <nom_de_la_tâche>
```

### `tasker --help (-h)`

Affiche des informations d'aide sur les commandes disponibles.

Exemple :

```
tasker --help
```

## Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.