# 🧺 Microservice for DigitalOcean Spaces 

Upload, remove and list objects to/from your CDN. Production-ready, zero configuration, working out of the box!

## ⚡️ Quick start

Sign in to your DigitalOcean account.

> **Don't have an account yet?** Join DigitalOcean at our [referral link](https://m.do.co/c/b41859fa9b6e)! Your profit is **$100** and we get $25. This will allow you **not** to think about paying for Spaces for 10 months, and we will send the money received to support the [Komentory](https://komentory.com/) project servers.

Next, create your own [Spaces](https://docs.digitalocean.com/products/spaces/quickstart/#create-a-space) and upload folder with permissions. Go to the `Spaces access keys` section at the [Manage Keys](https://cloud.digitalocean.com/account/api/tokens) page and create a new keys. _Save them to a notepad for later!_

OK, now we're ready to click this button:

[![Deploy do-spaces-microservice to DO](https://www.deploytodo.com/do-btn-blue-ghost.svg)](https://cloud.digitalocean.com/apps/new?repo=https://github.com/Komentory/do-spaces-microservice/tree/main)

After that, please fill all of the following environment variables (all variables are **required**):

![Screenshot](https://user-images.githubusercontent.com/11155743/130789680-e0430ed3-5667-422f-940d-3f6fffd0b539.png)

Choose name of your web service (which will be displayed in the Spaces URL, like `https://your-web-service.ondigitalocean.app`), region and a **Basic plan** (_$5/mo_) for start. Click to **Launch Basic App** and get our congratulations.

You have successfully deployed microservice to DigitalOcean! 🎉

## 📖 Next steps

Open your favorite REST API tool and send requests:

- **GET** `/v1/list` (get all objects from upload folder)
  - Request body (JSON): none
- **PUT** `/v1/upload` (put object to upload folder)
  - Request body (JSON): `{ "type": "<TYPE>", "path": "<PATH/TO/FILE>" }`
  - Supported types: `archive`, `audio`, `document`, `image`, `video`
- **DELETE** `/v1/remove` (delete object from upload folder by object key)
  - Request body (JSON): `{ "key": "<OBJECT_KEY>", "version_id": "<ID>" }`

## ⚠️ License

`do-spaces-microservice` is free and open-source software licensed under the [Apache 2.0 License](https://github.com/Komentory/do-spaces-microservice/blob/master/LICENSE).
