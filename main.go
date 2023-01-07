package main

import "github.com/hadoopconf/crawling"

func configCrawling() {
	const crawlingURL string = "https://hadoop.apache.org/docs/r2.4.1/hadoop-project-dist/hadoop-common/core-default.xml"
	const writePath string = "./doc/core-site.txt"
	crawling.Crawler(crawlingURL, writePath)
}

func main() {
	const writePath string = "./doc/hdfs-site.txt"
	crawling.Parsing(writePath)

}
