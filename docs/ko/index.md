---
pageType: home

hero:
    name: LazyGophers Log
    text: 고성능 및 유연한 Go 로깅 라이브러리
    tagline: zap 기반으로 구축되었으며 풍부한 기능과 간단한 API를 제공
    actions:
        - theme: brand
          text: 빠른 시작
          link: /API
        - theme: alt
          text: GitHub에서 보기
          link: https://github.com/lazygophers/log

features:
    - title: "고성능"
      details: zap 기반으로 구축되어 객체 풀 재사용과 조건부 필드 기록을 통해 최적의 성능을 실현
      icon: 🚀
    - title: "풍부한 로그 레벨"
      details: Trace, Debug, Info, Warn, Error, Fatal, Panic 레벨을 지원
      icon: 📊
    - title: "유연한 구성"
      details: 로그 레벨, 호출자 정보, 추적 정보, 접두사, 접미사 및 출력 대상을 사용자 정의할 수 있음
      icon: ⚙️
    - title: "파일 로테이션"
      details: 매시간 로그 파일 로테이션 지원이 내장되어 있음
      icon: 🔄
    - title: "Zap 호환성"
      details: zap WriteSyncer와 원활하게 통합
      icon: 🔌
    - title: "간단한 API"
      details: 표준 로그 라이브러리와 유사한 명확한 API로 사용 및 통합이 쉬움
      icon: 🎯
---

## 빠른 시작

### 설치

```bash
go get github.com/lazygophers/log
```

### 기본 사용법

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 기본 전역 logger 사용
    log.Debug("디버그 정보")
    log.Info("일반 정보")
    log.Warn("경고 정보")
    log.Error("오류 정보")

    // 형식화된 출력 사용
    log.Infof("사용자 %s이(가) 로그인했습니다", "admin")

    // 사용자 정의 구성
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("이것은 사용자 정의 logger의 로그입니다")
}
```

## 문서

-   [API 참조](API.md) - 완전한 API 문서
-   [변경 로그](/ko/CHANGELOG) - 버전 기록
-   [기여 가이드](/ko/CONTRIBUTING) - 기여하는 방법
-   [보안 정책](/ko/SECURITY) - 보안 가이드
-   [행동 강령](/ko/CODE_OF_CONDUCT) - 커뮤니티 가이드라인

## 성능 비교

| 기능       | lazygophers/log | zap | logrus | 표준 로그 |
| ---------- | --------------- | --- | ------ | -------- |
| 성능       | 높음            | 높음 | 중간   | 낮음     |
| API 단순성 | 높음            | 중간 | 높음   | 높음     |
| 기능 풍부함 | 중간            | 높음 | 높음   | 낮음     |
| 유연성     | 중간            | 높음 | 높음   | 낮음     |
| 학습 곡선   | 낮음            | 중간 | 중간   | 낮음     |

## 라이선스

이 프로젝트는 MIT 라이선스 하에 제공됩니다 - 자세한 내용은 [LICENSE](/ko/LICENSE) 파일을 참조하세요.
