FROM python:3.9.4-buster

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY ./src ./
ADD ./www /www

WORKDIR /www

CMD [ "python", "/usr/src/app/main.py" ]
