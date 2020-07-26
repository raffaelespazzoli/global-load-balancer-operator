// Copyright (C) 2017 ScyllaDB
// Use of this source code is governed by a ALv2-style
// license that can be found at https://github.com/scylladb/go-set/LICENSE.

package clusterreferenceset

// func BenchmarkTypeSafeSetHasNonExisting(b *testing.B) {
// 	b.StopTimer()
// 	var e1 redhatcopv1alpha1.ClusterReference
// 	e := createRandomObject(e1)
// 	if v, ok := e.(redhatcopv1alpha1.ClusterReference); ok {
// 		e1 = v
// 	}
// 	b.StartTimer()
// 	s := New()
// 	for i := 0; i < b.N; i++ {
// 		s.Has(e1)
// 	}
// }

// func BenchmarkInterfaceSetHasNonExisting(b *testing.B) {
// 	b.StopTimer()
// 	var e1 redhatcopv1alpha1.ClusterReference
// 	e := createRandomObject(e1)
// 	if v, ok := e.(redhatcopv1alpha1.ClusterReference); ok {
// 		e1 = v
// 	}
// 	b.StartTimer()
// 	s := set.New(set.NonThreadSafe)
// 	for i := 0; i < b.N; i++ {
// 		s.Has(e1)
// 	}
// }

// func BenchmarkTypeSafeSetHasExisting(b *testing.B) {
// 	b.StopTimer()
// 	var e1 redhatcopv1alpha1.ClusterReference
// 	e := createRandomObject(e1)
// 	if v, ok := e.(redhatcopv1alpha1.ClusterReference); ok {
// 		e1 = v
// 	}
// 	b.StartTimer()
// 	s := New()
// 	s.Add(e1)
// 	for i := 0; i < b.N; i++ {
// 		s.Has(e1)
// 	}
// }

// func BenchmarkInterfaceSetHasExisting(b *testing.B) {
// 	b.StopTimer()
// 	var e1 redhatcopv1alpha1.ClusterReference
// 	e := createRandomObject(e1)
// 	if v, ok := e.(redhatcopv1alpha1.ClusterReference); ok {
// 		e1 = v
// 	}
// 	b.StartTimer()
// 	s := set.New(set.NonThreadSafe)
// 	s.Add(e1)
// 	for i := 0; i < b.N; i++ {
// 		s.Has(e1)
// 	}
// }

// func BenchmarkTypeSafeSetHasExistingMany(b *testing.B) {
// 	s := New()
// 	b.StopTimer()
// 	var e1 redhatcopv1alpha1.ClusterReference
// 	for i := 0; i < 10000; i++ {
// 		e := createRandomObject(e1)
// 		if v, ok := e.(redhatcopv1alpha1.ClusterReference); ok {
// 			s.Add(v)
// 			if i == 5000 {
// 				e1 = v
// 			}
// 		}
// 	}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		s.Has(e1)
// 	}
// }

// func BenchmarkInterfaceSetHasExistingMany(b *testing.B) {
// 	s := set.New(set.NonThreadSafe)
// 	b.StopTimer()
// 	var e1 redhatcopv1alpha1.ClusterReference
// 	for i := 0; i < 10000; i++ {
// 		e := createRandomObject(e1)
// 		if v, ok := e.(redhatcopv1alpha1.ClusterReference); ok {
// 			s.Add(v)
// 			if i == 5000 {
// 				e1 = v
// 			}
// 		}
// 	}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		s.Has(e1)
// 	}
// }

// func BenchmarkTypeSafeSetAdd(b *testing.B) {
// 	b.StopTimer()
// 	var e redhatcopv1alpha1.ClusterReference
// 	s := New()
// 	objs := make([]redhatcopv1alpha1.ClusterReference, 0, b.N)
// 	for i := 0; i < b.N; i++ {
// 		e := createRandomObject(e)
// 		if v, ok := e.(redhatcopv1alpha1.ClusterReference); ok {
// 			objs = append(objs, v)
// 		}
// 	}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		s.Add(objs[i])
// 	}
// }

// func BenchmarkInterfaceSetAdd(b *testing.B) {
// 	b.StopTimer()
// 	var e redhatcopv1alpha1.ClusterReference
// 	s := set.New(set.NonThreadSafe)
// 	objs := make([]redhatcopv1alpha1.ClusterReference, 0, b.N)
// 	for i := 0; i < b.N; i++ {
// 		e := createRandomObject(e)
// 		if v, ok := e.(redhatcopv1alpha1.ClusterReference); ok {
// 			objs = append(objs, v)
// 		}
// 	}
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		s.Add(objs[i])
// 	}
// }
