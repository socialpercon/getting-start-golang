from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer
from urlparse import parse_qs, urlparse

def auth_requried(myfunc):
    def checkuser(self):
        user = parse_qs(urlparse(self.path).query).get('user')
        if user:
            self.user = user[0]
            myfunc(self)
        else:
            self.wfile.write('unknown user')
    return checkuser

class myHandler(BaseHTTPRequestHandler):
    @auth_requried
    def do_GET(self):
        self.wfile.write('Hello, {0}!'.format(self.user))


if __name__ == "__main__":
    try:
        server = HTTPServer(("", 8080), myHandler)
        server.serve_forever()
    except KeyboardInterrupt:
        server.socket.close()
