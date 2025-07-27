# Favicon Hash 계산기

Shodan 호환 MurmurHash3 값을 계산하는 favicon 해시 계산기입니다. Shodan 검색 쿼리에서 사용되는 해시를 생성하기 위한 Go와 Python 구현을 모두 포함합니다.

## 기능

- favicon 파일의 MurmurHash3 해시 계산
- Shodan의 favicon 해시와 완전 호환
- 10진수 및 16진수 출력 지원
- Go와 Python 구현 모두 제공

## 사용법

### Go 버전

먼저 프로그램을 빌드합니다:
```bash
go build -o favicon-hash main.go
```

favicon URL로 해시 계산:
```bash
./favicon-hash http://example.com/favicon.ico
```

16진수로 출력:
```bash
./favicon-hash -hex http://example.com/favicon.ico
```

### Python 버전

Python 스크립트 실행:
```bash
python get-shodan-favicon-hash.py
```

## 예시

```bash
./favicon-hash http://203.245.0.121/favicon.ico
```

## 구현 세부사항

- Go 버전은 Python의 mmh3 라이브러리와 정확히 호환되도록 MurmurHash3를 수동 구현
- Base64 인코딩 시 Python의 codecs.encode 동작에 맞춰 76자마다 줄바꿈 포함
- 두 구현 모두 동일한 favicon에 대해 동일한 해시 값 생성

## 종속성

Go 모듈 업데이트:
```bash
go mod tidy
```

Python 종속성 (Python 버전용):
- mmh3
- requests
- codecs