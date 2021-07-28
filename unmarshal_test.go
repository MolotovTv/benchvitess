package benchvitess

import (
	"testing"

	_ "github.com/planetscale/vtprotobuf/vtproto"
	"google.golang.org/protobuf/proto"
)

func BenchmarkMarshal(b *testing.B) {
	o2 := &Obj2{Objects: map[uint32]*Obj1{}}
	for i := 0; i < 100000; i++ {
		o2.Objects[uint32(i)] = &Obj1{Type: "a", Value: "b"}
	}

	b.Run("marshalStd", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := proto.Marshal(o2)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("marshalVT", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := o2.MarshalVT()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkUnmarshal(b *testing.B) {
	o2 := &Obj2{Objects: map[uint32]*Obj1{}}
	for i := 0; i < 10; i++ {
		o2.Objects[uint32(i)] = &Obj1{Type: "a", Value: "b"}
	}
	vtMarshalled, err := o2.MarshalVT()
	if err != nil {
		b.Fatal(err)
	}
	stdMarshalled, err := proto.Marshal(o2)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("unMarshalStd", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			o2 := &Obj2{}
			err = proto.Unmarshal(stdMarshalled, o2)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	b.Run("unMarshalVT", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			o2 := &Obj2{}
			err := o2.UnmarshalVT(vtMarshalled)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("unMarshalVTWithPool", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			o2 := Obj2FromVTPool()
			err := o2.UnmarshalVT(vtMarshalled)
			if err != nil {
				b.Fatal(err)
			}
			o2.ReturnToVTPool()
		}
	})
}
