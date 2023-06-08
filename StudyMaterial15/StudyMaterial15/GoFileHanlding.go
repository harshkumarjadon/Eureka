
package main 

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

// Reference Link
// 		https://pkg.go.dev/os
// 		https://pkg.go.dev/io
// 		https://pkg.go.dev/fmt
// 		https://pkg.go.dev/bufio


// Create File With DataFileForReading.txt 
//		In Same Code Directory Where This Code Exists

//_____________________________________________________________________

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//_____________________________________________________________________

func playWithFileReading() {
    dat, err := os.ReadFile("./DataFileForReading.txt")
    check(err)
    fmt.Print(string(dat))

    f, err := os.Open("./DataFileForReading.txt")
    check(err)

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    _, err = f.Seek(0, 0)
    check(err)

    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    f.Close()
}

//_____________________________________________________________________

func playWithFileWriting() {

    d1 := []byte("Hello Go Language!!!\n You Are Awesome!!!\n")
    err := os.WriteFile("./DataFileForWriting.txt", d1, 0644)
    check(err)

    f, err := os.Create("./DataFileForReading.txt")
    check(err)

    defer f.Close()

    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()
}

//_____________________________________________________________________

// // Reader is the interface that wraps the basic Read method.
// //
// // Read reads up to len(p) bytes into p. It returns the number of bytes
// // read (0 <= n <= len(p)) and any error encountered. Even if Read
// // returns n < len(p), it may use all of p as scratch space during the call.
// // If some data is available but not len(p) bytes, Read conventionally
// // returns what is available instead of waiting for more.
// //
// // When Read encounters an error or end-of-file condition after
// // successfully reading n > 0 bytes, it returns the number of
// // bytes read. It may return the (non-nil) error from the same call
// // or return the error (and n == 0) from a subsequent call.
// // An instance of this general case is that a Reader returning
// // a non-zero number of bytes at the end of the input stream may
// // return either err == EOF or err == nil. The next Read should
// // return 0, EOF.
// //
// // Callers should always process the n > 0 bytes returned before
// // considering the error err. Doing so correctly handles I/O errors
// // that happen after reading some bytes and also both of the
// // allowed EOF behaviors.
// //
// // Implementations of Read are discouraged from returning a
// // zero byte count with a nil error, except when len(p) == 0.
// // Callers should treat a return of 0 and nil as indicating that
// // nothing happened; in particular it does not indicate EOF.
// //
// // Implementations must not retain p.
// 	type Reader interface {
// 			Read(p []byte) (n int, err error)
// 	}

// // Writer is the interface that wraps the basic Write method.
// //
// // Write writes len(p) bytes from p to the underlying data stream.
// // It returns the number of bytes written from p (0 <= n <= len(p))
// // and any error encountered that caused the write to stop early.
// // Write must return a non-nil error if it returns n < len(p).
// // Write must not modify the slice data, even temporarily.
// //
// // Implementations must not retain p.
// 	type Writer interface {
// 			Write(p []byte) (n int, err error)
// 	}

// // Closer is the interface that wraps the basic Close method.
// //
// // The behavior of Close after the first call is undefined.
// // Specific implementations may document their own behavior.
// 	type Closer interface {
// 			Close() error
// 	}

// // Seeker is the interface that wraps the basic Seek method.
// //
// // Seek sets the offset for the next Read or Write to offset,
// // interpreted according to whence:
// // SeekStart means relative to the start of the file,
// // SeekCurrent means relative to the current offset, and
// // SeekEnd means relative to the end
// // (for example, offset = -2 specifies the penultimate byte of the file).
// // Seek returns the new offset relative to the start of the
// // file or an error, if any.
// //
// // Seeking to an offset before the start of the file is an error.
// // Seeking to any positive offset may be allowed, but if the new offset exceeds
// // the size of the underlying object the behavior of subsequent I/O operations
// // is implementation-dependent.
// 	type Seeker interface {
// 			Seek(offset int64, whence int) (int64, error)
// 	}

// // ReadWriter is the interface that groups the basic Read and Write methods.
// 	type ReadWriter interface {
// 			Reader
// 			Writer
// 	}

// // ReadCloser is the interface that groups the basic Read and Close methods.
// 	type ReadCloser interface {
// 			Reader
// 			Closer
// 	}

// // WriteCloser is the interface that groups the basic Write and Close methods.
// 	type WriteCloser interface {
// 			Writer
// 			Closer
// 	}

// // ReadWriteCloser is the interface that groups the basic Read, 
// // Write and Close methods.
// 	type ReadWriteCloser interface {
// 			Reader
// 			Writer
// 			Closer
// 	}

// // ReadSeeker is the interface that groups the basic Read and Seek methods.
// 	type ReadSeeker interface {
// 			Reader
// 			Seeker
// 	}

// // ReadSeekCloser is the interface that groups the basic Read, 
// // Seek and Close methods.
// 	type ReadSeekCloser interface {
// 			Reader
// 			Seeker
// 			Closer
// 	}

// // WriteSeeker is the interface that groups the basic Write and Seek methods.
// 	type WriteSeeker interface {
// 			Writer
// 			Seeker
// 	}

// // ReadWriteSeeker is the interface that groups the basic Read, 
// // Write and Seek methods.
// 	type ReadWriteSeeker interface {
// 			Reader
// 			Writer
// 			Seeker
// 	}

// // ReaderFrom is the interface that wraps the ReadFrom method.
// //
// // ReadFrom reads data from r until EOF or error.
// // The return value n is the number of bytes read.
// // Any error except EOF encountered during the read is also returned.
// //
// // The Copy function uses ReaderFrom if available.
// 	type ReaderFrom interface {
// 			ReadFrom(r Reader) (n int64, err error)
// 	}

// // WriterTo is the interface that wraps the WriteTo method.
// //
// // WriteTo writes data to w until there's no more data to write or
// // when an error occurs. The return value n is the number of bytes
// // written. Any error encountered during the write is also returned.
// //
// // The Copy function uses WriterTo if available.
// 	type WriterTo interface {
// 			WriteTo(w Writer) (n int64, err error)
// 	}


//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
//_____________________________________________________________________
// EXPERIMENT FOLLOWING CODE, MOMENT DONE, PLEASE RAISE YOUR HAND!!!

func main() {
	fmt.Println("\nFunction : playWithFileReading")
	playWithFileReading()

	fmt.Println("\nFunction : playWithFileWriting")
	playWithFileWriting()

	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
	// fmt.Println("\nFunction : ")
}
