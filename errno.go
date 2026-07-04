package errnov1

type Code int

const (
	OK Code = 0 // Success

	// Basic err codes
	EPERM   Code = 1  // Operation not permitted
	ENOENT  Code = 2  // No such file or directory
	ESRCH   Code = 3  // No such process
	EINTR   Code = 4  // Interrupted system call
	EIO     Code = 5  // I/O error
	ENXIO   Code = 6  // No such device or address
	E2BIG   Code = 7  // Argument list too long
	ENOEXEC Code = 8  // Exec format error
	EBADF   Code = 9  // Bad file number
	ECHILD  Code = 10 // No child processes
	EAGAIN  Code = 11 // Try again
	ENOMEM  Code = 12 // Out of memory
	EACCES  Code = 13 // Permission denied
	EFAULT  Code = 14 // Bad address
	ENOTBLK Code = 15 // Block device required
	EBUSY   Code = 16 // Device or resource busy
	EEXIST  Code = 17 // File exists
	EXDEV   Code = 18 // Cross-device link
	ENODEV  Code = 19 // No such device
	ENOTDIR Code = 20 // Not a directory
	EISDIR  Code = 21 // Is a directory
	EINVAL  Code = 22 // Invalid argument
	ENFILE  Code = 23 // File table overflow
	EMFILE  Code = 24 // Too many open files
	ENOTTY  Code = 25 // Not a typewriter
	ETXTBSY Code = 26 // Text file busy
	EFBIG   Code = 27 // File too large
	ENOSPC  Code = 28 // No space left on device
	ESPIPE  Code = 29 // Illegal seek
	EROFS   Code = 30 // Read-only file system
	EMLINK  Code = 31 // Too many links
	EPIPE   Code = 32 // Broken pipe
	EDOM    Code = 33 // Math argument out of domain
	ERANGE  Code = 34 // Math result not representable

	// Additional err codes
	EDEADLK      Code = 35 // Resource deadlock avoided
	ENAMETOOLONG Code = 36 // File name too long
	ENOLCK       Code = 37 // No locks available
	ENOSYS       Code = 38 // Function not implemented
	ENOTEMPTY    Code = 39 // Directory not empty
	ELOOP        Code = 40 // Too many symbolic links
	ENOMSG       Code = 42 // No message of desired type
	EIDRM        Code = 43 // Identifier removed
	ENOTFOUND    Code = 44 // Not found

	// Specific err codes
	ECOMP   Code = 70 // Component error
	EENTITY Code = 71 // Entity error
	ESYSTEM Code = 72 // System error
	ECALL   Code = 73 // Call error
)

func SUCCESS(v Code) bool {
	return v == OK
}

func FAIL(v Code) bool {
	return v != OK
}

func (x Code) Int() int {
	return (int)(x)
}

func (x Code) String() string {
	switch x {
	case OK:
		return "Success"
	case EPERM:
		return "Operation not permitted"
	case ENOENT:
		return "No such file or directory"
	case ESRCH:
		return "No such process"
	case EINTR:
		return "Interrupted system call"
	case EIO:
		return "I/O error"
	case ENXIO:
		return "No such device or address"
	case E2BIG:
		return "Argument list too long"
	case ENOEXEC:
		return "Exec format error"
	case EBADF:
		return "Bad file number"
	case ECHILD:
		return "No child processes"
	case EAGAIN:
		return "Try again"
	case ENOMEM:
		return "Out of memory"
	case EACCES:
		return "Permission denied"
	case EFAULT:
		return "Bad address"
	case ENOTBLK:
		return "Block device required"
	case EBUSY:
		return "Device or resource busy"
	case EEXIST:
		return "File exists"
	case EXDEV:
		return "Cross-device link"
	case ENODEV:
		return "No such device"
	case ENOTDIR:
		return "Not a directory"
	case EISDIR:
		return "Is a directory"
	case EINVAL:
		return "Invalid argument"
	case ENFILE:
		return "File table overflow"
	case EMFILE:
		return "Too many open files"
	case ENOTTY:
		return "Not a typewriter"
	case ETXTBSY:
		return "Text file busy"
	case EFBIG:
		return "File too large"
	case ENOSPC:
		return "No space left on device"
	case ESPIPE:
		return "Illegal seek"
	case EROFS:
		return "Read-only file system"
	case EMLINK:
		return "Too many links"
	case EPIPE:
		return "Broken pipe"
	case EDOM:
		return "Math argument out of domain"
	case ERANGE:
		return "Math result not representable"
	case EDEADLK:
		return "Resource deadlock avoided"
	case ENAMETOOLONG:
		return "File name too long"
	case ENOLCK:
		return "No locks available"
	case ENOSYS:
		return "Function not implemented"
	case ENOTEMPTY:
		return "Directory not empty"
	case ELOOP:
		return "Too many symbolic links"
	case ENOMSG:
		return "No message of desired type"
	case EIDRM:
		return "Identifier removed"
	case ENOTFOUND:
		return "Not found"
	case ECOMP:
		return "Component error"
	case EENTITY:
		return "Entity error"
	case ESYSTEM:
		return "System error"
	case ECALL:
		return "Call error"
	default:
		return "Unknown error"
	}
}
