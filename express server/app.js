const express = require('express');

const app = express();

const port = 3000;

app.set('views', './views');
app.engine('html', require('ejs').renderFile);
app.set('view engine', 'html');

app.get('/', (req, res)=>{
	res.render('index');
});
app.get('/about', (req, res)=>{
	res.send('About page')
});

app.listen(port)