# cloud-cert-renewer

Update Tencent Cloud SSL certificates directly on the certificate host.

The program runs on the machine that already serves the certificates. It checks the current public TLS certificate for each configured domain, downloads a newer Tencent Cloud certificate when the domain enters the `beforeExpired` window, replaces local certificate files atomically, runs domain-level `postCommands`, verifies the external certificate, and finally runs one round of `globalPostCommands` if every updated domain succeeded.

## Config

```yaml
alert:
  beforeExpired: 10d
  checkInterval: 12h
  notifyUrl: https://open.feishu.cn/open-apis/bot/v2/hook/xxxx

log:
  level: info

defaultProvider: tencentcloud

providerConfigs:
  tencentcloud:
    secretId: xxx
    secretKey: xxx
    autoApply:
      enabled: true
      pollInterval: 1m
      pollTimeout: 10m
      deleteDnsAutoRecord: true

globalPostCommands:
  - nginx -t
  - nginx -s reload

domains:
  - domain: doc.yourdomain.com
    certPath: /etc/nginx/ssl/doc.yourdomain.com.crt
    keyPath: /etc/nginx/ssl/doc.yourdomain.com.key
    postCommands:
      - consul kv put certs/doc.yourdomain.com.crt @{{.CertPath}}
      - consul kv put certs/doc.yourdomain.com.key @{{.KeyPath}}
```

`alert.checkInterval` controls how often the updater checks the public TLS certificate for each configured domain. If omitted, it defaults to `12h`; the minimum allowed value is `1m`. `alert.beforeExpired` controls the renewal window, not the check interval.

## Run

```sh
go run . -config=config.yaml
```

## Install

Install the latest Linux release:

```sh
curl -fsSL https://raw.githubusercontent.com/panjiang/cloud-cert-renewer/main/scripts/install.sh | sudo sh
```

Install a specific version:

```sh
curl -fsSL https://raw.githubusercontent.com/panjiang/cloud-cert-renewer/main/scripts/install.sh | sudo env VERSION=v0.1.0 sh
```

Create the runtime config:

```sh
sudo cp /etc/cloud-cert-renewer/config.yaml.example /etc/cloud-cert-renewer/config.yaml
sudo chmod 600 /etc/cloud-cert-renewer/config.yaml
sudo vi /etc/cloud-cert-renewer/config.yaml
```

Start the service after the config is ready:

```sh
sudo systemctl enable --now cloud-cert-renewer
sudo systemctl status cloud-cert-renewer
```

Upgrade to the latest release:

```sh
curl -fsSL https://raw.githubusercontent.com/panjiang/cloud-cert-renewer/main/scripts/install.sh | sudo sh
sudo systemctl restart cloud-cert-renewer
```
