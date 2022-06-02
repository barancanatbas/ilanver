const {esClient} = require('./client');
const insertDoc = async (indexName, _id, data) => {
    try {
        return await esClient.index({
            index: indexName,
            id: _id,
            body: data
        });
    } catch (err) {
        console.error(err);
    }
}

module.exports = {
    insertDoc
};