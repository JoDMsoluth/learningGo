package test28

// OOD (Object Oriented Design)
// SOLID
// S : Single-responsibility principle
// O : Open–closed principle
// L : Liskov substitution principle
// I : Interface segregation principle
// D : Dependency inversion principle

// SOLID가 하는 게 좋다. -> 하지만 이상향이다. -> 현실적인 문제에 부딪히나 SOLID로 개발하려고 노력한다.
// *S : 단일 책임 원칙
// - 하나의 객체는 하나의 책임을 가져야 한다.
// - ex) 예금 잔고 객체 : 입금 기능, 출금 기능 -> 어디까지가 단일 책임인가? -> 입출금으로 하나로 둬야 되는지 따로 둬야 되는지

// ***O : 확장에는 열려있고 변경에는 닫혀있다. - 제일 중요하다
// 기능이 확장되더라도 기존 코드를 변경하지 말아야 한다.

// L : 리스코프 치환 이론 - 가장 어려움
// t.O(x)라는 함수가 있다. s.O(y)라는 함수도 있다. s는 t의 확장타입일 때, 두 함수는 같은 동작을 해야한다.
// 즉, 확장한 타입이 기존타입의 동작을 바꾸지 말라는 뜻
// GoLang은 '상속'을 지원하지 않기에 이것을 신경 쓸 필요 없다.

// I : 인터페이스 분리 법칙
// 여러개의 관계를 모아놓은 인터페이스보다 관계 하나씩 정의하는 것이 더 좋다.

// D : 의존성 역전 법칙
// 관계는 인터페이스에 의존해야지, 객체에 의존하는게 더 좋다.

type FinanceReport struct {
}

type Report struct {
}

// 1번 기능
func (r *FinanceReport) MakeReport() *Report {
	return &Report{}
}

type ReportSender interface {
	Send(*Report)
}

type EmailReportSender struct {
}

// 2번 기능 -> SendReportFile, SendReportHttp 등 확장, 추가가 된다면 기능을 따로 둬야 된다.
func (s *EmailReportSender) SandReport(r *Report) {

}

// *O 확장에는 열려있다.
type FileReportSender struct {
}

func (s *FileReportSender) SandReport(r *Report) {
}

// OOP를 넘어서 !!!!!!
// OOP의 문제점
// 1. 현실에 부딪혀 잘하기 어렵다. 올바른 지식 -> 올바른 숙달
// 2. 잘 만들더라도 새로운 프로그래머가 수천개의 클래스를 파악하는데 오래걸린다.
// 3. 현대 추세 : 프로그램을 빨리 만들고 빨리 없애자 -> OOP는 빨리 만들 수 없다.
// -> 절차적, OOP는 상태와 기능이 혼재되어 있어 오래걸리는 것이다.

// 개선법
// 1. 절차적 -> OOP -> Stateless(상태없음, 기능만)
// -> stateless 상태는 외부에서 가져오자
// ex) Micro Server, Serverless, Functional Language, MVC, MVVC

// Functional Language
// erlang, list, scala, f#, 엘릭서

// Micro Service
// 각 부분을 따로 서비스로 만들어서 한 페이지로 만든다.

// Serverless
// 여러 서버들로 나눠서 하나의 페이지에서 여러 서버를 취합해서 사용한다.

// ECS (Entity Component System)
// Player Entity (Move, Attack, Talk) - 어떤 컴포넌트들을 가지고 있느냐에 따라 Entity가 결정된다.
// Component는 Data만 가지고 있고 System은 기능만 가지고 있다.

// MVC (Model View Controller) - 데이터 뷰 기능
// Model : Data
// View : 화면
// Controller : 기능

// 결론 : 빠르게 만들고 낮은 기술부채로 제작 -> 성공하면 ㄲ -> 실패하면 프로젝트 파괴
