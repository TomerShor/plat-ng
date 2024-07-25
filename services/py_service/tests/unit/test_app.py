import time

import pytest
import services.py_service.src.app as app


def test_calculate_runtime():
    application = app.App(log_level="info", port=8080)
    application._start_time = time.time()
    time.sleep(1)
    run_time = application.calculate_runtime()
    assert run_time >= 1
