# Hadoop Conf Web Service

## 프로젝트 개요
* Hadoop 설정 과정에서 직접 hdfs-site.xml, yarn-site.xml 등 설정 파일을 작성 및 수정하는 경우가 많다.
* 쿠버네티스 하둡 플랫폼 구축의 경우 configmap 등을 활용하면 이 불편함을 일부 해결할 수 있지만, 설정 정보를 확인하고 코드를 작성해야하는 과정은 마찬가지로 매우 불편하다.
* 본 프로젝트에서는 Hadoop 설정을 가상서버 구축과 쿠버네티스 구축 과정을 분리하여 설정과정에 도움을 줄 수 있는 웹서비스를 제공한다.
<br></br>

**Server : Golang**
<br></br>

**Client : React** 
<br></br>

**Database : Mongo DB** 

## 웹 페이지 디자인 개요
### Home
* On server / On Docker / On Kubernets 선택
* On server : 설정 완료 후 xml 설정파일 다운로드 
* On Docker : Docker 파일 다운로드
* On Kubernets : Docker 파일 + yaml 실행파일 다운로드

### Set
* Option 선택(Name + Value)
* Option - mouseOn일 때, 설명 window
* Option Value - 직접 입력(기본값) 
* 선택된 Option으로 xml 파일 미리보기

### Download
* 다운로드 링크

## To do list
**1월 9일(월)**
* hdfs GET RestAPI 생성
* set page Get RestAPI 확인
<br></br>

**1월 8일(일)**
* mongodb connection err - finish(권한 인증 부분 문제 해결)
<br></br>

**1월 7일(토)**
* Data Parsing finish
* mongodb connection err -ing 
<br></br>

**1월 6일(금)**
* Hadoop config crawling(hdfs-default, core-default 완료)
* mongodb setting 
<br></br>

**1월 5일(목)**
* 프로젝트 개요 정리
* 프로젝트 생성

