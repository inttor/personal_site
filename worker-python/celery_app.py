from celery import Celery
import os

# 读取环境变量，必须指向 服务器A 上的 Redis IP (内网IP或VPN IP)
REDIS_BROKER = os.getenv("CELERY_BROKER_URL", "redis://localhost:6379/0")
REDIS_BACKEND = os.getenv("CELERY_RESULT_BACKEND", "redis://localhost:6379/1")

app = Celery(
    'personal_site_worker',
    broker=REDIS_BROKER,
    backend=REDIS_BACKEND,
    include=['tasks.data_tasks'] # 指定要加载的 tasks
)

app.conf.update(
    task_serializer='json',
    accept_content=['json'],
    result_serializer='json',
    timezone='Asia/Shanghai',
    enable_utc=True,
)
