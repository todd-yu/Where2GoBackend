require('dotenv').config();

const express = require('express');
const app = express();
const port = 3000;
const bodyParser = require('body-parser');
const connection = require('./database');

app.get('/', (req, res) => res.send('hello world'));

// /articles/:year?/:month?/:day?
app.route('/users/:userId')
  .get(function(req, res, next) {
    console.log(req.params);
    connection.query("SELECT * FROM `userID` WHERE id = ? LIMIT 3", req.params.userId,
        function(error, results, fields) {
        if (error) throw error;
        res.json(results);
      }
    );
  });



// CREATE DATABASE guestbook;
// USE guestbook;
// CREATE TABLE entries (guestName VARCHAR(255), content VARCHAR(255),
// entryID INT NOT NULL AUTO_INCREMENT, PRIMARY KEY(entryID));
// INSERT INTO entries (guestName, content) values ("first guest", "I got here!");
// INSERT INTO entries (guestName, content) values ("second guest", "Me too!");
// SELECT * FROM entries;
// 
// 

app.get('/status', (req, res) => res.send('Working!'));


app.set('port', process.env.PORT || port);

app.listen(port, () => console.log('listening on port 3000...'));