# Uniqrawler

```bash
# setup
## mac m1
curl https://chromedriver.storage.googleapis.com/100.0.4896.60/chromedriver_mac64_m1.zip -o chromedriver.zip
## mac intel
curl https://chromedriver.storage.googleapis.com/100.0.4896.60/chromedriver_mac64.zip -o chromedriver.zip
unzip chromedriver.zip

# run
CREDENTIAL_JSON=<token path> SHEET_ID=<sheet id> PATH=$(PWD):$PATH go run *.go
```
