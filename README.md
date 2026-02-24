<div style="display:flex; width:100%; justify-content:center">
<img style="width:50%;" src="./assets/Logo_With_BG.png">
</div>
<br><br>

# 🌩️ API de Stockage d'Images avec Telegram

Une API simple développée en Go pour fournir un service de stockage d'images en ligne 🚀, similaire à Postimages.org ou 8upload.com, en utilisant un bot Telegram comme backend de stockage.

## ✨ Fonctionnalités

- 📤 Uploader des images.
- 🖼️ Récupérer des images via un lien.
- ✅ Lister les formats d'image supportés.

## ☁️ Utilisation

Vous pouvez utiliser l'api à l'adresse suivante :

```

  https://tgcloud.alwaysdata.net/

```

## 📚 Documentation de l'API

### `GET /supports-formats`

Retourne la liste des formats d'images acceptés pour l'upload.

**Exemple de réponse :**

```json
{
  "formats": ["image/jpeg", "image/png", "image/gif"]
}
```

### `POST /post`

Permet d'uploader une nouvelle image. L'image doit être envoyée en tant que `multipart/form-data` avec la clé `file`.

**Exemple avec cURL :**

```bash
curl -X POST -F "file=@/chemin/vers/votre/image.jpg" https://tgcloud.alwaysdata.net/post
```

**Exemple de réponse succès :**

```json
{
  "status": "success",
  "message": "Image postée avec succès",
  "id": "ID_UNIQUE_TG"
}
```

**Exemple de réponse erreur :**

```json
{
  "status": "error",
  "message": "Erreur lors de l'upload."
}
```

### `GET /img/{id}`

Récupère et affiche une image précédemment uploadée à partir de son identifiant unique.

**Exemple d'utilisation :**
Ouvrez `https://tgcloud.alwaysdata.net/img/AgACAgQAAxkDAAMraZyUGni2TrQXYkhK5ursyqon8wIAAsMOaxvNseFQ30BlvGjM3CIBAAMCAAN3AAM6BA` dans votre navigateur.

vous devriez voir l'image correspondante :

<div style="display:flex; width:100%; justify-content:center">
<img style="width:30%;" src="https://tgcloud.alwaysdata.net/img/AgACAgQAAxkDAAMraZyUGni2TrQXYkhK5ursyqon8wIAAsMOaxvNseFQ30BlvGjM3CIBAAMCAAN3AAM6BA">
</div>

## 🛠️ Prérequis pour installer et lancer le projet en local

- Go (version 1.x)
- Un bot Telegram et son token.
- L'ID d'un canal ou d'un chat Telegram où le bot a les droits de publication.

## 🚀 Installation & Lancement

1.  **Clonez le dépôt :**

    ```bash
    git clone https://github.com/AdrienMttn/TgCloud.git
    cd TgCloud
    ```

2.  **Installez les dépendances :**

    ```bash
    go mod tidy
    ```

3.  **Configuration de l'environnement :**
    Créez un fichier `.env` à la racine du projet et ajoutez les variables suivantes :

    ```env
    BOT_TOKEN="VOTRE_TOKEN_DE_BOT_TELEGRAM"
    CHAT_ID="VOTRE_ID_DE_CHAT_OU_CANAL"
    ```

4.  **Lancez le serveur :**
    ```bash
    go run main.go
    ```
    Le serveur démarrera sur `http://localhost:8100`.

---

fait avec ❤️ par [Adrien Mttn](https://github.com/AdrienMttn) | [GitHub](https://github.com/AdrienMttn) | [LinkedIn](https://www.linkedin.com/in/adrien-metton/)
