package node

type Node struct {
	Name            string
	Ip              string
	Cores           int
	Memory          int // this is maximum memory usable
	MemoryAllocated int
	Disk            int // this is maxium no. of disks usable
	DiskAllocated   int
	Role            string
	TaskCount       int
}
