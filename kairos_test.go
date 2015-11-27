package kairos

import "testing"

type Fatalistic interface {
	Fatal(...interface{})
}

func newTestKairos(t Fatalistic) *Kairos {
	appId := "e2a8eaa7"
	appKey := "4092e4a45070bca728644e9285f084b4"

	k, err := New(uri, appId, appKey)
	if err != nil {
		t.Fatal(err)
	}

	return k
}

func TestEnroll(t *testing.T) {
	k := newTestKairos(t)
	image = "http://media.kairos.com/kairos-elizabeth.jpg"
	subjectId = "subject1"
	galleryName = "gallerytest1"
	if err := k.Enroll(image, subjectId, galleryName); err != nil {
		t.Fatal(err)
	}
}

func TestRecognize(t *testing.T) {
	k := newTestKairos(t)
	image = "http://media.kairos.com/kairos-elizabeth.jpg"
	subjectId = "subject1"
	galleryName = "gallerytest1"
	if err := k.Recognize(image, galleryName); err != nil {
		t.Fatal(err)
	}
}

func TestDetect(t *testing.T) {
	k := newTestKairos(t)
	image = "http://media.kairos.com/kairos-elizabeth.jpg"
	if err := k.Detect(image); err != nil {
		t.Fatal(err)
	}
}

func TestListGalleries(t *testing.T) {
	k := newTestKairos(t)
	if err := k.ListGalleries(); err != nil {
		t.Fatal(err)
	}
}

func TestViewGallery(t *testing.T) {
	k := newTestKairos(t)
	galleryName = "gallerytest1"
	if err := k.ViewGallery(galleryName); err != nil {
		t.Fatal(err)
	}
}

func TestRemoveSubject(t *testing.T) {
	k := newTestKairos(t)
	subjectId = "subject1"
	galleryName = "gallerytest1"
	if err := k.RemoveSubject(subjectId, galleryName); err != nil {
		t.Fatal(err)
	}
}

func TestRemoveGallery(t *testing.T) {
	k := newTestKairos(t)
	galleryName = "gallerytest1"
	if err := k.RemoveGallery(galleryName); err != nil {
		t.Fatal(err)
	}
}
