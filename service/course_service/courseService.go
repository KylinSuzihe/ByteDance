package course_service

import (
	"demo1/models"
	"demo1/pkg/e"
	"demo1/pkg/types"
	"strconv"
)

// 作者：田思润

func CreateCourse(name string, capacity int) (int64, e.ErrNo) {
	if name == "" || capacity == 0 {
		return 0, e.ParamInvalid
	}

	var course models.Course
	course.Name = name
	course.Capacity = capacity

	errno := models.CreateCourseWithRedis(&course)

	return course.CourseId, errno
}

func GetCourse(courseId string) (*types.TCourse, e.ErrNo) {
	course, errno := models.GetCourse(courseId)
	if course == nil {
		return nil, errno
	}
	tCourse := types.TCourse{
		CourseID:  strconv.FormatInt(course.CourseId, 10),
		Name:      course.Name,
		TeacherID: course.TeacherId,
	}
	return &tCourse, errno
}

type node struct {
	to   string
	next int
}

const N int = 2000010

var (
	cnt  int
	head map[string]int
	vis  map[string]int
	f    map[string]string
	ans  map[string]string
	a    [N]node
)

func add(x string, y string) {
	cnt++
	a[cnt].to = y
	a[cnt].next = head[x]
	head[x] = cnt

}

func dfs(u string, time int) bool {
	var i int
	for i = head[u]; i != 0; i = a[i].next {
		v := a[i].to
		if vis[v]^time != 0 {
			vis[v] = time
			if (f[v] == "") || dfs(f[v], time) { //若未被匹配或者能商量，就将两人匹配
				f[v] = u
				ans[u] = v
				return true
			}
		}
	}
	return false
}

func ScheduleCourse(m map[string][]string) (map[string]string, e.ErrNo) {
	head = make(map[string]int)
	vis = make(map[string]int)
	f = make(map[string]string)
	ans = make(map[string]string)
	for i := range m {
		for j := range m[i] {
			add(i, m[i][j])
		}
	}
	cnt = 1
	for i := range m {
		dfs(i, cnt)
		cnt++
	}
	return ans, e.OK
}
