### Pattern
singleTon Pattern [O]
usecase-repository Pattern [O]
layered architecture [O]
datasource Pattern [O]

### 기능 구현
질문[O]
답변[O]
덧글[O]

검색[X]
좋/싫[X]
신고[X]

### 테스트 코드
QuestionUsecase[X]
AnswerUsecase[X]
CommentUsecase[X]

### 기술 구현
handler 구현 [O]
DB 구현 [O]
Config 구현 [O]
MiddleWare 구현 - 캐싱,로깅 [X]
로그 [X]

### 리팩토링
- 의존성 관련
model 분리 - usecase와 repository [O]
외부 라이브러리에 추상 계층 추가 [X]

- 코드 관련
nil 반환 처리 & nil 체크 [X]
구조체 기본값 체크 [X]

- DB 관련
FindAll offset limit [O]
DB INS/UPT now()대신 golang time 주입 [O]
Update 개선 [X]

### 오류 수정
DB 조회 및 저장 시 배열 <-> 스트링 변환 로직 구성 [O]

-----------------------
### Go Style
error wrapping (예외 처리) [X]

