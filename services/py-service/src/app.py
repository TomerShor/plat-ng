import logging
import os
import requests
import sys

from flask import Flask


class App:
    def __init__(self, log_level, port, host=None):
        self.app = Flask(__name__)
        self.port = port
        self.host = host or "0.0.0.0"
        self._configure_logger(log_level)
        self._setup_routes()

    def run(self, debug=False, threaded=False):
        logging.info(f"Starting at port {self.port}")
        # Make it compatible with IPv6 if Linux
        if sys.platform == "linux":
            self.app.run(host='::', port=self.port, debug=debug, threaded=threaded)
        else:
            self.app.run(host=self.host, port=self.port, debug=debug, threaded=threaded)

    def _setup_routes(self):
        @self.app.route('/')
        def hello():
            return "Hello World"

        @self.app.route('/hello-go')
        def hello_go():
            response = requests.get("http://go-service:8080/")
            return response.text

        @self.app.route('/status')
        def status():
            return "OK"

    def _configure_logger(self, level):
        log_level = logging.getLevelName(level.upper())
        logging.basicConfig(stream=sys.stdout, level=log_level)
        requests_log = logging.getLogger("requests.packages.urllib3")
        requests_log.setLevel(logging.DEBUG)
        requests_log.propagate = True
        self.app.logger.addHandler(logging.StreamHandler(sys.stdout))
        self.app.logger.setLevel(logging.DEBUG)


def parse_args():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("--port", type=int, default=os.environ.get("LISTEN_PORT", 8020))
    parser.add_argument("--log-level", type=str, default=os.environ.get("LOG_LEVEL", "debug"))
    parser.add_argument("--debugger", action="store_true", default=os.environ.get("DEBUGGER", False))
    parser.add_argument("--threaded", action="store_true", default=os.environ.get("THREADED", False))
    return parser.parse_args()


if __name__ == '__main__':
    args = parse_args()

    app = App(log_level=args.log_level, port=int(args.port))
    app.run(debug=args.debugger, threaded=args.threaded)