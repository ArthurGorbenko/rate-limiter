// Setup node js server
const http = require('http');
const port = 3000;

const server = http.createServer((req, res) => {
	switch (req.url) {
		case '/': {
      console.log('/ received');
			res.writeHead(200, {'Content-Type': 'text/html'});
			res.write('Hello World!');
			res.end();
			return;
		}
		case '/up': {
      console.log('/up received');
			res.writeHead(200, {'Content-Type': 'text/html'});
			res.write('Up');
			res.end();
			return;
		}
	}
});

server.addListener('request', function (req, res) {
	console.log('Request received');
});

server.listen(port, () => {
	console.log(`Server is running at http://localhost:${port}`);
});
