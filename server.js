const express = require('express')
const bodyParser = require('body-parser')
const app = express()
const MongoClient = require('mongodb').MongoClient
const mongoConnectionString = 'mongodb+srv://mymdb:uDLQgyfHr0rN4Z0E@mymdb.fscoq.mongodb.net/mymdb?retryWrites=true&w=majority';

MongoClient.connect(mongoConnectionString, {
    useUnifiedTopology: true
  }).then(client => {
        const db = client.db('mymdb')
        const moviesCollection = db.collection('movies')
        app.set('view engine', 'ejs')
        app.use(bodyParser.urlencoded({extended: true}))
        app.use(express.static('public'))
        app.use(bodyParser.json())
        // app.get('/', (req, res) => {
        //   res.sendFile(__dirname + "/index.html")
        // })
        app.get('/', (req, res) => {
            const cursor = db.collection('movies').find().toArray()
            .then(results => {
                res.render('index.ejs', {movies: results})
            })
            .catch(error => console.error(error))
        })
        app.post('/movies', (req, res) => {
            moviesCollection.insertOne(req.body)
            .then(result => {
                res.redirect('/')
            })
            .catch(error => console.log(error))

        })
        app.put('/movies', (req, res) => {
            moviesCollection.findOneAndUpdate(
                { title: 'Parasite' }, 
                {
                    $set: {title: req.body.title}
                }, 
                {
                    upsert: true
                })
                .then(_ => {
                    res.json('Success')
                  })
                .catch(error => console.error(error))
        })
        app.listen(3000, function(){
            console.log('listening on 3000')
        })
        console.log('Connected to Database')
    })
.catch(error => console.error(error))


