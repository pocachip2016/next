#!/usr/bin/env python3
"""kta_contents.json → kta_contents 테이블 적재 (파라미터 바인딩, 안전한 이스케이프).

사용: python3 load_kta_contents.py
환경변수로 접속정보 재정의 가능: MYSQL_HOST, MYSQL_PORT, MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE
"""
import json
import os
import sys

import pymysql

HOST = os.environ.get("MYSQL_HOST", "127.0.0.1")
PORT = int(os.environ.get("MYSQL_PORT", "3306"))
USER = os.environ.get("MYSQL_USER", "root")
PASSWORD = os.environ.get("MYSQL_PASSWORD", "dmscks3927!")
DATABASE = os.environ.get("MYSQL_DATABASE", "rightwatch")
JSON_PATH = os.path.join(os.path.dirname(__file__), "kta_contents.json")

COLUMNS = ["genre", "title", "actors", "director", "price", "enddate", "synop"]


def main():
    with open(JSON_PATH, encoding="utf-8") as f:
        rows = json.load(f)
    print(f"loaded {len(rows)} records from {JSON_PATH}")

    conn = pymysql.connect(
        host=HOST, port=PORT, user=USER, password=PASSWORD,
        database=DATABASE, charset="utf8mb4",
    )
    try:
        with conn.cursor() as cur:
            cur.execute("SELECT COUNT(*) FROM kta_contents")
            existing = cur.fetchone()[0]
            if existing > 0:
                print(f"kta_contents already has {existing} rows — skipping load (truncate manually to reload)")
                return

            placeholders = ", ".join(["%s"] * len(COLUMNS))
            sql = f"INSERT INTO kta_contents ({', '.join(COLUMNS)}) VALUES ({placeholders})"
            values = [
                tuple(None if r.get(c) is None else str(r.get(c)) for c in COLUMNS)
                for r in rows
            ]
            cur.executemany(sql, values)
        conn.commit()
        with conn.cursor() as cur:
            cur.execute("SELECT COUNT(*) FROM kta_contents")
            print(f"inserted — kta_contents now has {cur.fetchone()[0]} rows")
    finally:
        conn.close()


if __name__ == "__main__":
    sys.exit(main())
