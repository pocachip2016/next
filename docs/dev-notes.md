# Dev Notes (rightwatch)

> 운영/개발 중 자주 쓰는 명령 모음. 원본은 `read.md`(pocachip, 2023-10-17).
> **시크릿(토큰/DB 비밀번호)은 이 문서에 절대 두지 않는다.** 자격증명은 `.env` 참조.

## crontab
```
* * * * * doitnow.sh            # 1분마다
*/10 * * * * doitnow.sh         # 10분마다
15 * * * * doitnow.sh           # 매시 15분
0 * * * * doitnow.sh            # 매시간
0 */2 * * * doitnow.sh          # 2시간마다
00 11,16 * * * <cmd>            # 11시·16시
00 09-18 * * * <cmd>            # 9~18시 매시간
00 09-18 * * 1-5 <cmd>          # 주말 제외 (Sunday=0)
30 23 * * * timeout 18000 php prices_all.php >> log.txt   # 5시간 후 자동 중지
```
- `service cron start|restart|stop`
- 특정 크론 종료: `pkill -f '<command substring>'`

## MySQL — last_update 컬럼 타입 변경 이력
```sql
-- 문자열 → DATETIME 마이그레이션
alter table rightwatch.post add column last_update2 DATETIME;
update rightwatch.post set last_update2 = str_to_date(REPLACE(last_update,'"',''), '%Y-%m-%d %H:%i:%s');
alter table rightwatch.post drop last_update;
alter table rightwatch.post change last_update2 last_update DATETIME;
```

## MySQL — 설치/외부접속 (Ubuntu)
```bash
sudo apt-get update && sudo apt-get install mysql-server
sudo ufw allow mysql
sudo systemctl enable mysql
```
```sql
-- 외부 접속 사용자 (비밀번호는 .env에서 주입)
CREATE USER 'pocachip'@'%' IDENTIFIED BY '<MYSQL_PASSWORD>';
ALTER USER 'pocachip'@'%' IDENTIFIED WITH mysql_native_password BY '<MYSQL_PASSWORD>';
GRANT ALL PRIVILEGES ON *.* TO 'pocachip'@'%';
FLUSH PRIVILEGES;
```

## 데이터 이관 (workbench export → import)
JSON 정제 후 import:
```bash
cp result_mvo.json test.json
sed -i '/^\[/d;/^\]/d;/^$/d' test.json
sed 's/,$//g' test.json > test_out.json
tr -d '\n' < test_out.json > test_online.json
cat test_online.json | jq .          # validate
```
```
# MySQL Shell (자격증명은 .env)
\c mysqlx://root:<MYSQL_ROOT_PASSWORD>@localhost
\use rightwatch
util.importJson("c:\\test_online.json", {schema:"rightwatch", table:"tmp_post", tableColumn:"json"});
```

## 콘텐츠/중복 처리 쿼리
```sql
-- 콘텐츠별 체크리스트 수 내림차순
select * from (
  select a.id, a.title, count(*) as tc
  from kta_contents a left join check_list b on a.id=b.content_id
  group by a.id
) c order by tc desc;

-- 중복 title 찾기 / 제거
select title, count(title) from kta_contents group by title having count(title)>1;
delete t1 from kta_contents t1 join kta_contents t2 on t1.title=t2.title where t1.id>t2.id;
```

## tmp_post(JSON) → 정규 테이블 적재
```sql
-- kta_contents
INSERT INTO kta_contents(genre,title,actors,director,price,enddate,synop)
SELECT json->>'$.genre', json->>'$.title', json->>'$.actors', json->>'$.director',
       json->>'$.price', json->>'$.enddate', json->>'$.synop'
FROM tmp_post;

-- post
INSERT INTO post(website,cat1_code,cat2_code,cat1_title,cat2_title,idx,txt,lvl19,price,seller,partner,attach_file_size,item_url,last_update)
SELECT json->>'$.website', json->>'$.cat1_code', json->>'$.cat2_code', json->>'$.category_title',
       json->>'$.sub_category_title', json->>'$.idx', json->>'$.txt', json->>'$.Lvl19', json->>'$.point',
       json->>'$.seller', json->>'$.partner', json->>'$.attach_file_size', json->>'$.item_url', json->>'$.time'
FROM rightwatch.tmp_post;
```

## ginbro (Go CRUD 코드 생성)
```bash
ginbro gen -u root -p '<MYSQL_ROOT_PASSWORD>' -a "127.0.0.1:3306" -d rightwatch -c utf8 -o=rightwatch
go mod init && go mod tidy && go build
```

## Scrapy 크롤러
```bash
# venv 활성화 후
scrapy crawl ondisk_update2 -a category=MVO -s MYSQL_HOST=mysql
pip3 install mysql-connector-python pymysql
```
- 스파이더 이름: `ondisk_update2`(증분, 파이프라인 DB 적재). 인자 `category`(MVO/DRA/MED/ANI).

## Docker (MySQL 단독 컨테이너 예시)
```bash
docker pull mysql
docker run --name mysql-container -e MYSQL_ROOT_PASSWORD='<MYSQL_ROOT_PASSWORD>' -d -p 3306:3306 mysql:latest
docker exec -it mysql-container bash
```
> 통합 스택은 루트 `docker-compose.yml` 사용 (자격증명은 `.env`).

## screenshots 테이블
```sql
CREATE TABLE screenshots (
  id int NOT NULL AUTO_INCREMENT,
  url TEXT,
  url_md5 CHAR(32) AS (MD5(url)),
  ct timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
select * from screenshots where url_md5 = MD5('www.google.co.kr');
```

## Angular / ngx-admin
```bash
nvm install 14.21.3
npm i --save @types/ws@8.5.4
npm start
```

## GORM SubQuery 메모
```go
db.Where("amount > (?)", db.Table("orders").Select("AVG(amount)")).Find(&orders)
db.Table("(?) as u", db.Model(&User{}).Select("name","age")).Where("age = ?", 18).Find(&User{})
```

## 코드 스타일
camelCase / kebab-case / SNAKE_CASE / PascalCase
