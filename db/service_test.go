// Copyright 2018 Mark Wardle / Eldrix Ltd
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package db_test

import (
	"os"
	"testing"

	"github.com/wardle/go-terminology/db"
)

const (
	dbFilename = "../snomed.db"
)

func setUp(tb testing.TB) *db.Snomed {
	if _, err := os.Stat(dbFilename); os.IsNotExist(err) { // skip these tests if no working live snomed db
		tb.Skipf("Skipping tests against a live database. To run, create a database named %s", dbFilename)
	}
	sct, err := db.NewService(dbFilename, false)
	if err != nil {
		tb.Fatal(err)
	}
	return sct
}

func TestService(t *testing.T) {
	sct := setUp(t)
	ms, err := sct.GetConcept(24700007)
	if err != nil {
		t.Fatal(err)
	}
	parents, err := sct.GetAllParents(ms)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, p := range parents {
		if p.ID == 6118003 {
			found = true
		}
	}
	if !found {
		t.Fatal("Multiple sclerosis not correctly identified as a type of demyelinating disease")
	}
	if !sct.IsA(ms, 6118003) {
		t.Fatal("Multiple sclerosis not correctly identified as a type of demyelinating disease")
	}

	allChildrenIDs, err := sct.GetAllChildrenIDs(ms)
	if err != nil {
		t.Fatal(err)
	}
	allChildren, err := sct.GetConcepts(allChildrenIDs...)
	if err != nil {
		t.Fatal(err)
	}
	if len(allChildren) < 2 {
		t.Fatal("Did not correctly find many recursive children for MS")
	}
	for _, child := range allChildren {
		fsn, err := sct.GetFullySpecifiedName(child)
		if err != nil || fsn == nil {
			t.Fatalf("Missing FSN for concept %d : %v", child.ID, err)
		}
	}
}

func BenchmarkGetConceptAndDescriptions(b *testing.B) {
	snomed := setUp(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ms, err := snomed.GetConcept(24700007)
		if err != nil {
			b.Fatal(err)
		}
		_, err = snomed.GetDescriptions(ms)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIsA(b *testing.B) {
	snomed := setUp(b)
	ms, err := snomed.GetConcept(24700007)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	var (
		isDemyelinating  bool
		isPharmaceutical bool
	)
	for i := 0; i < b.N; i++ {
		isDemyelinating = snomed.IsA(ms, 6118003)
		isPharmaceutical = snomed.IsA(ms, 373873005)
	}
	if isDemyelinating == false || isPharmaceutical == true {
		b.Fatal("MS misclassified using IS-A hierarchy")
	}
}
