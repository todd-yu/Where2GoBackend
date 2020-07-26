require('dotenv').config();

const express = require('express');
const app = express();
const port = 3000;
const bodyParser = require('body-parser');
const connection = require('./database');

app.get('/', (req, res) => res.send('hello world'));

app.route('/users/:userId')
  .get(function(req, res, next) {
    connection.query(
      "SELECT * FROM `userID` WHERE id = ? LIMIT 3", req.params.userId,
      function(error, results, fields) {
        if (error) throw error;
        res.json(results);
      }
    );
  });

app.get('/status', (req, res) => res.send('Working!'));


app.set('port', process.env.PORT || port);

app.listen(port, () => console.log('listening on port 3000...'));