#!/usr/bin/env python3
"""seed_sample_posts.py — post 테이블 샘플 데이터 삽입 (매칭 UI 검증 전용).

kta_contents 의 실제 제목을 post.txt 에 포함해 매칭 히트를 유도한다.
이 스크립트는 개발·테스트 전용이며, 실 운영 DB에는 실행하지 않는다.
"""
import os
import random
import time

import pymysql

HOST = os.environ.get("MYSQL_HOST", "127.0.0.1")
PORT = int(os.environ.get("MYSQL_PORT", "3306"))
USER = os.environ.get("MYSQL_USER", "root")
PASSWORD = os.environ.get("MYSQL_PASSWORD", "dmscks3927!")
DATABASE = os.environ.get("MYSQL_DATABASE", "rightwatch")

WEBSITES = [1, 2, 3]
QUALITY_TAGS = ["1080p", "720p", "4K", "HD", "무삭제", "BluRay"]
N_TITLES = 20


def main():
    conn = pymysql.connect(
        host=HOST, port=PORT, user=USER, password=PASSWORD,
        database=DATABASE, charset="utf8mb4",
    )
    try:
        with conn.cursor() as cur:
            cur.execute("SELECT COUNT(*) FROM post")
            if cur.fetchone()[0] > 0:
                print("post 테이블에 이미 데이터가 있습니다. 삽입을 건너뜁니다.")
                return

            cur.execute(
                "SELECT id, title FROM kta_contents WHERE title IS NOT NULL LIMIT %s",
                (N_TITLES,),
            )
            contents = cur.fetchall()

        now = time.strftime("%Y-%m-%d %H:%M:%S")
        rows = []
        for content_id, title in contents:
            quality = random.choice(QUALITY_TAGS)
            txt = f"{title} {quality} 다운로드"
            idx = f"SEED-{content_id:05d}"
            rows.append((
                random.choice(WEBSITES),
                "A", "B01", "영화", "드라마",
                idx, txt,
                None, "1000", "테스트판매자", None, None,
                f"https://ondisk.co.kr/view/{idx}",
                now,
            ))

        sql = """INSERT INTO post
            (website, cat1_code, cat2_code, cat1_title, cat2_title,
             idx, txt, lvl19, price, seller, partner, attach_file_size, item_url, last_update)
            VALUES (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)"""
        with conn.cursor() as cur:
            cur.executemany(sql, rows)
        conn.commit()
        print(f"post 테이블에 {len(rows)}건 샘플 데이터 삽입 완료")
    finally:
        conn.close()


if __name__ == "__main__":
    main()
