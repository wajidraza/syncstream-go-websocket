from datetime import datetime
from pydantic import BaseModel, Field, EmailStr
from typing import Optional, Dict, Any

class UserSchema(BaseModel):
    id: Optional[str] = None
    email: EmailStr
    name: Optional[str] = None
    role: str = "MEMBER"
    created_at: datetime = Field(default_factory=datetime.utcnow)

class RecordEntity(BaseModel):
    id: Optional[str] = None
    title: str
    metadata: Dict[str, Any] = Field(default_factory=dict)
    status: str = "ACTIVE"
    created_at: datetime = Field(default_factory=datetime.utcnow)
