from locust import HttpUser, task, between
import random

class MyAPIUser(HttpUser):
    wait_time = between(1, 3)

    @task(10)
    def test_success_endpoint(self):
        """Load test for /success endpoint"""
        self.client.get("/success", name="/success")

    @task(3)
    def test_delay_endpoint(self):
        """Load test for /delay/{seconds} endpoint"""
        delay = random.randint(1, 5)
        self.client.get(f"/delay/{delay}", name="/delay/{seconds}")

    @task(2)
    def test_timeout_endpoint(self):
        """Load test for /timeout endpoint"""
        self.client.get("/timeout", name="/timeout")

    @task(1)
    def test_error_endpoint(self):
        """Load test for /error endpoint"""
        self.client.get("/error", name="/error")
