# trademiner
Trade miner is a Market Mining Application.  Trade Miner Application can work on any type of Market but currently application is extensively tested with equity option. Trade Miner Application comes along with platform to collect the data of all the options of choosen market. The trademiner application shall then post process on the collected data by running through various AI algorithms. The AI algorithm shall generate a Trade Miner Score and classifies the options into
 a. Long Term Option
 b. Short Term Option
 c. Intraday Option
 
 Sample Scoring of the option for NSE can be seen in
 
 
 #Sample
 https://docs.google.com/spreadsheets/d/1kf_2FEF39tgJ3zjMDS5Sgf_vUfJDDQr6f-WlKYvxAAM/edit#gid=0
 
#Framework.
Interested developers shall be able to add the AI algorithms, can be used for commercial or non commercial purposes. The frameworks provides necessary hooks to run the custom algorithms which can influence the trading score.

#Installation
1. Clone the Tradminer 
2. Set the GOPATH=$INSTALLPATH/trademiner
3. Run `make dep`
4. Run `make`
5. This Generates TradeMiner Application
