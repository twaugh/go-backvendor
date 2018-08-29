// Copyright (C) 2018 Tim Waugh
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package backvendor

import (
	"testing"
)

func TestDirs(t *testing.T) {
	src := GoSource("testdata/gosource")
	if src.Topdir() != "testdata/gosource" {
		t.Fatal("Topdir")
	}
	if src.Vendor() != "testdata/gosource/vendor" {
		t.Fatal("Vendor")
	}
}
