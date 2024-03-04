FROM python:3.11.4-bullseye
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV PYTHONPATH="$PYTHONPATH:/usr/app"
ENV POETRY_HOME="/usr/local"
ENV GUNICORN=true
RUN curl -sSL https://install.python-poetry.org | POETRY_HOME=/usr/local python3 -
WORKDIR /usr/app
COPY . .
RUN poetry install --no-root
EXPOSE 80
ENTRYPOINT ["bash", "./entrypoint.sh" ]