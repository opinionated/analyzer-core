# analyzer-core

This is where all the analysis code goes. Currently I am thinking that each service gets it's own package.

The analyzer core provides drivers for the pipeline to access various services and databases. The pipeline scores these articles and preloads any nessisary data (eg fetching a list of keywords). No need to worry about threading here (blocking, threadsafe etc), the pipeline will take care of that.

Again, generally we are looking to find articles that are similar to eachother. 

Feel free to add any analysis of data you want. Try to test it when / where you can. 

###things to analyze
Alchemy: 
