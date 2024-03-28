package trees

type vEB struct {
	Summary *vEB
	U       int
	Cluster []*vEB
	Min     int
	Max     int
}
