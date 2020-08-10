package pirsch

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTrackerHitTimeout(t *testing.T) {
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	req1.Header.Add("User-Agent", "valid")
	req2 := httptest.NewRequest(http.MethodGet, "/hello-world", nil)
	req2.Header.Add("User-Agent", "valid")
	store := newTestStore()
	tracker := NewTracker(store, "salt", &TrackerConfig{WorkerTimeout: time.Second * 2})
	tracker.Hit(req1, nil)
	tracker.Hit(req2, nil)
	time.Sleep(time.Second * 4)

	if len(store.hits) != 2 {
		t.Fatalf("Two requests must have been tracked, but was: %v", len(store.hits))
	}

	// ignore order...
	if store.hits[0].Path != "/" && store.hits[0].Path != "/hello-world" ||
		store.hits[1].Path != "/" && store.hits[1].Path != "/hello-world" {
		t.Fatalf("Hits not as expected: %v %v", store.hits[0], store.hits[1])
	}
}

func TestTrackerHitLimit(t *testing.T) {
	store := newTestStore()
	tracker := NewTracker(store, "salt", &TrackerConfig{
		Worker:           1,
		WorkerBufferSize: 10,
	})

	for i := 0; i < 7; i++ {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add("User-Agent", "valid")
		tracker.Hit(req, nil)
	}

	time.Sleep(time.Second) // allow all hits to be tracked
	tracker.Stop()

	if len(store.hits) != 7 {
		t.Fatalf("All requests must have been tracked, but was: %v", len(store.hits))
	}
}

func TestTrackerIgnoreHitPrefetch(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("User-Agent", "valid")
	req.Header.Set("X-Moz", "prefetch")
	tracker := NewTracker(newTestStore(), "salt", nil)

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with X-Moz header must be ignored")
	}

	req.Header.Del("X-Moz")
	req.Header.Set("X-Purpose", "prefetch")

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with X-Purpose header must be ignored")
	}

	req.Header.Set("X-Purpose", "preview")

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with X-Purpose header must be ignored")
	}

	req.Header.Del("X-Purpose")
	req.Header.Set("Purpose", "prefetch")

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with Purpose header must be ignored")
	}

	req.Header.Set("Purpose", "preview")

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with Purpose header must be ignored")
	}

	req.Header.Del("Purpose")

	if tracker.ignoreHit(req) {
		t.Fatal("Hit must not be ignored")
	}
}

func TestTrackerIgnoreHitUserAgent(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("User-Agent", "This is a bot request")
	tracker := NewTracker(newTestStore(), "salt", nil)

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with keyword in User-Agent must be ignored")
	}

	req.Header.Set("User-Agent", "This is a crawler request")

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with keyword in User-Agent must be ignored")
	}

	req.Header.Set("User-Agent", "This is a spider request")

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with keyword in User-Agent must be ignored")
	}

	req.Header.Set("User-Agent", "Visit http://spam.com!")

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with URL in User-Agent must be ignored")
	}

	req.Header.Set("User-Agent", "Mozilla/123.0")

	if tracker.ignoreHit(req) {
		t.Fatal("Hit with regular User-Agent must not be ignored")
	}

	req.Header.Set("User-Agent", "")

	if !tracker.ignoreHit(req) {
		t.Fatal("Hit with empty User-Agent must be ignored")
	}
}

func TestTrackerIgnoreHitBotUserAgent(t *testing.T) {
	tracker := NewTracker(newTestStore(), "salt", nil)

	for _, botUserAgent := range userAgentBlacklist {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("User-Agent", botUserAgent)

		if !tracker.ignoreHit(req) {
			t.Fatalf("Hit with user agent '%v' must have been ignored", botUserAgent)
		}
	}
}

type testStore struct {
	hits []Hit
}

func newTestStore() *testStore {
	return &testStore{make([]Hit, 0)}
}

func (store *testStore) Save(hits []Hit) error {
	log.Printf("Saved %d hits", len(hits))
	store.hits = append(store.hits, hits...)
	return nil
}

func (store *testStore) DeleteHitsByDay(tenantID sql.NullInt64, t time.Time) error {
	panic("implement me")
}

func (store *testStore) SaveVisitorsPerDay(day *VisitorsPerDay) error {
	panic("implement me")
}

func (store *testStore) SaveVisitorsPerHour(hour *VisitorsPerHour) error {
	panic("implement me")
}

func (store *testStore) SaveVisitorsPerLanguage(language *VisitorsPerLanguage) error {
	panic("implement me")
}

func (store *testStore) SaveVisitorsPerPage(page *VisitorsPerPage) error {
	panic("implement me")
}

func (store *testStore) Days(tenantID sql.NullInt64) ([]time.Time, error) {
	panic("implement me")
}

func (store *testStore) CountVisitorsPerDay(tenantID sql.NullInt64, t time.Time) (int, error) {
	panic("implement me")
}

func (store *testStore) CountVisitorsPerDayAndHour(tenantID sql.NullInt64, t time.Time) ([]VisitorsPerHour, error) {
	panic("implement me")
}

func (store *testStore) CountVisitorsPerLanguage(tenantID sql.NullInt64, t time.Time) ([]VisitorsPerLanguage, error) {
	panic("implement me")
}

func (store *testStore) CountVisitorsPerPage(tenantID sql.NullInt64, t time.Time) ([]VisitorsPerPage, error) {
	panic("implement me")
}

func (store *testStore) Paths(tenantID sql.NullInt64, t time.Time, t2 time.Time) ([]string, error) {
	panic("implement me")
}

func (store *testStore) Visitors(tenantID sql.NullInt64, t time.Time, t2 time.Time) ([]VisitorsPerDay, error) {
	panic("implement me")
}

func (store *testStore) PageVisits(tenantID sql.NullInt64, s string, t time.Time, t2 time.Time) ([]VisitorsPerDay, error) {
	panic("implement me")
}

func (store *testStore) VisitorLanguages(tenantID sql.NullInt64, t time.Time, t2 time.Time) ([]VisitorLanguage, error) {
	panic("implement me")
}

func (store *testStore) HourlyVisitors(tenantID sql.NullInt64, t time.Time, t2 time.Time) ([]HourlyVisitors, error) {
	panic("implement me")
}

func (store *testStore) ActiveVisitors(tenantID sql.NullInt64, t time.Time) (int, error) {
	panic("implement me")
}

func (store *testStore) CountHits(tenantID sql.NullInt64) int {
	panic("implement me")
	return 0
}

func (store *testStore) VisitorsPerDay(tenantID sql.NullInt64) []VisitorsPerDay {
	panic("implement me")
	return nil
}

func (store *testStore) VisitorsPerHour(tenantID sql.NullInt64) []VisitorsPerHour {
	panic("implement me")
	return nil
}

func (store *testStore) VisitorsPerLanguage(tenantID sql.NullInt64) []VisitorsPerLanguage {
	panic("implement me")
	return nil
}

func (store *testStore) VisitorsPerPage(tenantID sql.NullInt64) []VisitorsPerPage {
	panic("implement me")
	return nil
}
