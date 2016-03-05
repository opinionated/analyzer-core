# analyzer-core

This is where all the analysis code goes. Currently I am thinking that each service gets it's own package.

The analyzer core provides drivers for the pipeline to access various services and databases. The pipeline scores these articles and preloads any nessisary data (eg fetching a list of keywords). No need to worry about threading here (blocking, threadsafe etc), the pipeline will take care of that.

Again, generally we are looking to find articles that are similar to eachother. 

Feel free to add any analysis of data you want. Try to test it when / where you can. 

#Analysis
#####Alchemy
Alchemy is free and extracts a lot of data from articles:
* taxonomy - high level catagories
* keywords, entities and concepts
* document and target sentiment
* document emotions
* relations

See a live demo [here](http://www.alchemyapi.com/products/demo/alchemylanguage). 

One issue with alchemy is that we need to do the inter-data (ie strength of soccer->iraq war) relationship strengths ourselves. There is a request factory as well as a parser/filewriter in the code for alchemy. 

#####IBM Bluemix
Bluemix provides powerful machine learning tools as well as some interesting pre-canned tools:
* [concept insights](https://console.ng.bluemix.net/catalog/services/concept-insights/) - ranks concepts that are (probably) similar to the ones extracted via alchemy (ibm also owns alchemy)
* [retreive and rank](https://console.ng.bluemix.net/catalog/services/retrieve-and-rank/) - retrieves relevant texts from a database using machine learning (I think we need to train this)

See the full catalog [here](https://console.ng.bluemix.net/catalog/).

There are a couple other interesting analysis services on bluemix as well as some useful general purpose tools. We can get a key from moorthy ( I am working on it ). If we have a good ammount of credit we might want to deploy stuff to bluemix. 

We need to store data more legitimately as well, not sure where or how yet. 
