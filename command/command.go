package command

import("fmt")

type Command interface {
	Execute()
}

type Job1Command struct {
	n *Notification
}


func (j *Job1Command) Execute() {
	j.n.Job1()
}


func (j *Job2Command) Execute() {
	j.n.Job2()
}


type Job2Command struct {
	n *Notification
}


type Notification struct{}

func (*Notification) Job1() {
	fmt.Print("Notification in  job 1\n")
}

func (*Notification) Job2() {
	fmt.Print("Notification in  job 2\n")
}

func NewJob1Command(n *Notification) *Job1Command {
	return &Job1Command{
		n: n,
	}
}

func NewJob2Command(n *Notification) *Job2Command {
	return &Job2Command{
		n: n,
	}
}

type Jobs struct {
	j1 Command
	j2 Command
}

func NewJobs(j1, j2 Command) *Jobs {
	return &Jobs{
		j1: j1,
		j2: j2,
	}
}

func (j *Jobs) RunJob1() {
	j.j1.Execute()
}

func (j *Jobs) RunJob2() {
	j.j2.Execute()
}
