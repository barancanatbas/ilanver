const es = require('elasticsearch');
const esClient = new es.Client({
    host: 'localhost:9200',
});

esClient.ping({    
    requestTimeout: 1000
}, function (error) {
    if (error) {
        console.trace('Elasticsearch\'e erişilmiyor!');
    } else {
        console.log('Elasticsearch ayakta :)');
    }
});
 
module.exports = {
    esClient
}