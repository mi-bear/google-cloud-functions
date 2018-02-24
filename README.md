# google-cloud-functions-sandbox

Emulator:

```
# いずれかを選択する
# npm
$ npm install -g @google-cloud/functions-emulator

# yarn
$ yarn global add @google-cloud/functions-emulator
```

ディレクトリ構成:

```
{function-name}
 └index.js
```

## Usage

### Emulator
{function-name}ディレクトリ配下(index.jsが配置されているディレクトリ)で実行する。  
[悲報] `fish` シェルを利用している場合は, `functions` が競合するため, `bash` にでもして利用すること。  

Start Emulator:

```
$functions start
```

Deploy function:

```
$ functions deploy {function-name} --torigger-http
```

Call function:

```
$ functions call {function-name}
```

View logs:

```
$ functions logs read
```

help:

```
functions --help
```

## Deploy
### Preparation
Cloud Starageにバケットを作る。  
https://console.cloud.google.com/storage/browser  

- バケット名:
  - bear-development
- ストレージクラス: Regional
- 場所: ASIA-NORTHEAST1

### Deploy

```
$ gcloud beta functions deploy {directory-name} --trigger-http --stage-bucket kintone-development --entry-point {function-name}
```

example:

```
$ gcloud beta functions deploy bear-talk --project bear-development --trigger-http --stage-bucket bear-development --entry-point bearTalk
```

## Road to Go
https://github.com/GoogleCloudPlatform/cloud-functions-go

### Testing

http://localhost:8080/execute  

```
$ make godev
```

http://localhost:1234/execute  

```
$ go run main.go -addr=localhost:8080
```

### Deploy

```
$ make
```

`function.zip` をGoogleCloudFunctionsにアップロードする。

