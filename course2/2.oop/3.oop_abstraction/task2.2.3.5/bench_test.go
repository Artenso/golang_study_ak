package main

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"testing"

	"github.com/howeyc/crc16"
	"github.com/sigurn/crc8"
)

func BenchmarkHash(b *testing.B) {
	b.Run("CRC64", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = crc64.Checksum([]byte("test"), &crc64.Table{})
		}
	})

	b.Run("CRC32", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = crc32.Checksum([]byte("test"), &crc32.Table{})
		}
	})

	b.Run("CRC16", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = crc16.Checksum([]byte("test"), &crc16.Table{})
		}
	})

	b.Run("CRC8", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = crc8.Checksum([]byte("test"), &crc8.Table{})
		}
	})

}

func BenchmarkSetGet(b *testing.B) {
	b.Run("CRC64", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC64())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Set(fmt.Sprintf("key-%d", i), "value")
			_, _ = m.Get(fmt.Sprintf("key-%d", i))
		}
	})

	b.Run("CRC32", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC32())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Set(fmt.Sprintf("key-%d", i), "value")
			_, _ = m.Get(fmt.Sprintf("key-%d", i))
		}
	})

	b.Run("CRC16", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC16())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Set(fmt.Sprintf("key-%d", i), "value")
			_, _ = m.Get(fmt.Sprintf("key-%d", i))
		}
	})

	b.Run("CRC8", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC8())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Set(fmt.Sprintf("key-%d", i), "value")
			_, _ = m.Get(fmt.Sprintf("key-%d", i))
		}
	})

}

func BenchmarkSet(b *testing.B) {
	b.Run("CRC64", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC64())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Set(fmt.Sprintf("key-%d", i), "value")
		}
	})

	b.Run("CRC32", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC32())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Set(fmt.Sprintf("key-%d", i), "value")
		}
	})

	b.Run("CRC16", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC16())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Set(fmt.Sprintf("key-%d", i), "value")
		}
	})

	b.Run("CRC8", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC8())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			m.Set(fmt.Sprintf("key-%d", i), "value")
		}
	})

}

func BenchmarkGet(b *testing.B) {
	b.Run("CRC64", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC64())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = m.Get(fmt.Sprintf("key-%d", i))
		}
	})

	b.Run("CRC32", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC32())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = m.Get(fmt.Sprintf("key-%d", i))
		}
	})

	b.Run("CRC16", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC16())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = m.Get(fmt.Sprintf("key-%d", i))
		}
	})

	b.Run("CRC8", func(b *testing.B) {
		m := NewHashMap(16, WithHashCRC8())
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_, _ = m.Get(fmt.Sprintf("key-%d", i))
		}
	})

}
