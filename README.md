# data-generator-cassandra

Randomly generate data in a Cassandra database for simulating massive read and write operations. 

## Usage
To start to generate random data, go to your terminal and type:
`docker-compose build`
`docker-compose up`

This will launch 5 services:
1. A Cassandra database
2. A Golang application that writes and read data
3. A Cassandra exporter to capture metrics from Cassandra database
4. A Prometheus to scrape Cassandra metrics
5. A Grafana with a pre-configured dashboard

After a while, the data-generator will start to write and read data from Cassandra. If you want to see Cassandra metrics at Grafana, go to your web browser at `http://localhost:3000`, add Prometheus data source (`http://prometheus:9090`) and import [Criteo](https://github.com/criteo/cassandra_exporter) default dashboard with ID `6400`.