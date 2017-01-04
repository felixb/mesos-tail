# mesos-tail

A tail for mesos tasks.

## Install

`go install github.com/felixb/mesos-tail`

## Usage

You must specify the leading mesos master with `--master`.
Zookeeper leader detection is not implemented yet.

Run it without `--task` argument to see a list of all running tasks on your cluster.
Add `--running=false` to see completed tasks as well.

Specifing `--task` adds a substring filter to the list of tasks.
If only one task is matched, mesos-tail will start printing stdout and stderr of that task ony your local tty.
