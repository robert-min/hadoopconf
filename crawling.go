package main

import "github.com/hadoopconf/crawling"

// configCrawling to get hdfs, core config data on https://hadoop.apache.org
func configCrawling() {
	const crawlingURL string = "https://hadoop.apache.org/docs/r2.4.1/hadoop-project-dist/hadoop-common/core-default.xml"
	const writePath string = "./doc/core-site.txt"
	crawling.Crawler(crawlingURL, writePath)
}

// configToMongo to create crawling data on mongodb
func configToMongo() {
	const writePath string = "./doc/core-site.txt"
	crawling.ParsingCreate(writePath)
}
