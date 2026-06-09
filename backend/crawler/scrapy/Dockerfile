FROM mcr.microsoft.com/playwright/python:v1.45.0-jammy

WORKDIR /app

# Install supercronic (container-friendly cron)
RUN curl -fsSLO "https://github.com/aptible/supercronic/releases/download/v0.2.29/supercronic-linux-amd64" && \
    chmod +x supercronic-linux-amd64 && \
    mv supercronic-linux-amd64 /usr/local/bin/supercronic

# Install Python deps
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Install Playwright Chromium (already in base image, this ensures it's registered)
RUN playwright install chromium

# Copy spider project
COPY ondisk/ ./ondisk/

# Crontab
COPY crontab /etc/supercronic/crontab

# Log dir
RUN mkdir -p /var/log

CMD ["supercronic", "/etc/supercronic/crontab"]
