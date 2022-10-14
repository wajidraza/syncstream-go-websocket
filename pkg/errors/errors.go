class BaseAppException(Exception):
    def __init__(self, message: str, status_code: int = 500, code: str = "ERROR"):
        super().__init__(message)
        self.message = message
        self.status_code = status_code
        self.code = code

class ResourceNotFoundException(BaseAppException):
    def __init__(self, resource: str):
        super().__init__(f"Requested resource '{resource}' was not found.", status_code=404, code="NOT_FOUND")
