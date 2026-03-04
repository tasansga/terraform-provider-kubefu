package resourcegen

import "testing"

func TestDataSourceFinalFileNameUnique(t *testing.T) {
	providerName := "k8s"
	def := Definition{
		Kind:    "Test",
		Group:   "",
		Version: "v1",
	}
	counts := dataSourceBaseNameCounts([]Definition{def}, providerName)
	got := dataSourceFinalFileName(providerName, def, counts)
	if got != "datasource_k8s_test_v1.go" {
		t.Fatalf("unexpected file name %q", got)
	}
}

func TestDataSourceFinalFileNameCollision(t *testing.T) {
	providerName := "k8s"
	defCore := Definition{
		Kind:    "Event",
		Group:   "",
		Version: "v1",
	}
	defEvents := Definition{
		Kind:    "Event",
		Group:   "events.k8s.io",
		Version: "v1",
	}
	defs := []Definition{defCore, defEvents}
	counts := dataSourceBaseNameCounts(defs, providerName)
	if counts["datasource_k8s_event_v1.go"] != 2 {
		t.Fatalf("expected collision count 2, got %d", counts["datasource_k8s_event_v1.go"])
	}
	gotCore := dataSourceFinalFileName(providerName, defCore, counts)
	if gotCore != "datasource_k8s_core_event_v1.go" {
		t.Fatalf("unexpected file name for core event %q", gotCore)
	}
	gotEvents := dataSourceFinalFileName(providerName, defEvents, counts)
	if gotEvents != "datasource_k8s_events_k8s_io_event_v1.go" {
		t.Fatalf("unexpected file name for events event %q", gotEvents)
	}
}
