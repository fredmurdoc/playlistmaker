# Specifications

## Objectifs

Spécifier les fonctionnalités de l'outil

## Contexte

L'outil de lecture des fichiers multimedias supporte les playlist. Ces dernières permettent de jouer les morceaux de musique d'un albmu dans l'ordre. De plus, il permettent de fournir certains informations supplémentaires (:warning: définir lesquelles).
En effet, certains outils de lecture via le protocole DLNA ne peuvent pas récupérer les données d'album, de titre, voire d'ordre des morceaux.

Ainsi, pour les albums ne fournissant pas de playlist, il y a toujours un doute sur l'ordre des morceaux jouées, ce qui pose un vrai problème à la fois esthétique (pas de support des artworks), éthique (pas d'affichage de la liste des morceaux ce qui empèche de comprendre comment celui-ci a été pensé et conçu) et artistique (la lecture des titres ne se faisant pas dans l'ordre, on passe complètement à côté de la démarche musicale de l'artiste) sur l'écoute d'un album.

## Fonctionnalités

L'outil devra analyser les répertoires d'albmus de musique et détecter ceux qui ne continnent pas de fichiers de types playlist (:warning: citer les extensions de fichiers en question).
Pour les répertoires n'ayant pas de  playlists le programme devra construire les fichiers playlist.
Pour cela, il extrait les informations pour chaque fichier de musique pour en dégager : 

 - le titre
 - l'ordre dans l'album
 - l'année
 - le nom de l'artiste
 - le nom de l'album

:question: quelles sont les informations minimales pour constituer un fichier playlist ?

Si des informations manquent le programme appelera l'API de LastFM pour récupérer les informations manquantes.

:question: Faut il appeler 1 ou N API ? Quelles sont les limites d'utilisation de celles-ci ?
Quels sont les parametres obligatoires ? Est on spurs de disposer de toutes les informations ?

### Fonctionnalités supplémentaires.

Le programme fera du `best effort` :  s'il n'est pas possible de récupérer les informations nécessaires à la contruction d'une playlist (album inconnu, pas assez d'infos pour sollcitier une API), le programme devra ignorer l'album. Pour cela il devrai consigner quelque part qu'il doit ignorer cet album.

Il sera peut être envisager d'utiliser le nom du fichier pour rechercher des informations dans les APIs. 

## Choix techniques

Go a été choisi pour plusieurs raisons : 

 - le langage permet de fabriquer des livrables optimisés pour tous les types de plateformes (:question: même raspberry ?)
 - le langage supporte nativement les executions parralèles ce qui permet des gains de temps dans le scan des répertoires et la récupération d'informations.
 - le langage m'étant pas connu, c'est un bon exercice pour s'y confronter



Les inconvénients de Go : 
 - certains liens étroits avec des librairies C doivent être résolus. Ces librairires pourraient être absentes de certaines plateformes (Raspberry par exemple pour `ffmpeg`)


Les inconnus à lever : 
 - l'execution en parallèle est simplifié, par contre il  faudra etudier la conception des fonctions en amont afin de ne pas être piègé par la parallélisation : 
   * comment rassembler des informations qui sont en cours de récupération dans des traitements paralèlisés ?
 
