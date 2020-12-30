const express = require('express');

const app = express();

const port = 3000;

app.get('/', (req, res)=>{
	res.send('Welcome to Home Page');
});
app.get('/about', (req, res)=>{
	res.send('About page')
});

app.listen(port)