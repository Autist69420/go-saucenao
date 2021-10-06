# go-saucenao
Simple program to get results from SauceNAO using their api 

# Instructions
1. Download the repository 
2. CD into the folder
3. run `go get` to install needed depencies
4. Run `go run main.go <api key> <url>`
    - the `api key` is needed to make it work, you can get an api key for SauceNAO [here](https://saucenao.com/user.php?page=search-api)
    - url can be anything as long as its an image or a gif.
        - example: `https://cdn.discordapp.com/attachments/681523355576303627/895311255349395486/1633373445123.jpg`
    - (optional). Run `go build main.go` to compile it into a executable