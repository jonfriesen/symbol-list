# Symbol List
This tool pulls security symbol lists from exchanges and compiles them into a unified JSON & CSV by date. 

## Usage

```bash
# save files to local 'some_dir' folder, will be created if missing
# -dir is optional, default is 'data'
go run cmd/retrieve/main.go -dir=some_dir
```

# Resources used
## Nasdaq / NSYE
- [nasdaq trader tools](http://www.nasdaqtrader.com/trader.aspx?id=symboldirdefs)

## TSX
- [tsx api](https://www.tsx.com/json/company-directory/search/tsx/^*)
- [tsxc api](https://www.tsx.com/json/company-directory/search/tsxv/^*)

## Cryptocurrencies
- [cryptocompare api](https://min-api.cryptocompare.com/data/all/coinlist)