package task

import (
	"strings"

	api "k8s.io/api/core/v1"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"

	"github.com/pborman/uuid"
	"github.com/turbonomic/kubeturbo/pkg/discovery/repository"
)

const (
	TaskSucceeded TaskResultState = "Succeeded"
	TaskFailed    TaskResultState = "Failed"
)

type Task struct {
	uid      string
	name     string
	nodeList []*api.Node
	podList  []*api.Pod
	pvList   []*api.PersistentVolume
	pvcList  []*api.PersistentVolumeClaim
	cluster  *repository.ClusterSummary
}

// Worker task is consisted of a list of nodes the worker must discover.
func NewTask() *Task {
	uid := uuid.NewUUID().String()
	name := strings.Split(uid, "-")[0]
	return &Task{
		uid:  uid,
		name: name,
	}
}

// Assign nodes to the task.
func (t *Task) WithNodes(nodeList []*api.Node) *Task {
	t.nodeList = nodeList
	return t
}

func (t *Task) WithNode(node *api.Node) *Task {
	t.nodeList = append(t.nodeList, node)
	return t
}

// Assign pods to the task.
func (t *Task) WithPods(podList []*api.Pod) *Task {
	t.podList = podList
	return t
}

// Assign pvs to the task.
func (t *Task) WithPVs(pvList []*api.PersistentVolume) *Task {
	t.pvList = pvList
	return t
}

// Assign pvcs to the task.
func (t *Task) WithPVCs(pvcList []*api.PersistentVolumeClaim) *Task {
	t.pvcList = pvcList
	return t
}

// Assign cluster summary to the task.
func (t *Task) WithCluster(cluster *repository.ClusterSummary) *Task {
	t.cluster = cluster
	return t
}

// Get node list from the task.
func (t *Task) NodeList() []*api.Node {
	return t.nodeList
}

// Get pod list from the task.
func (t *Task) PodList() []*api.Pod {
	return t.podList
}

// Get PV list from the task.
func (t *Task) PVList() []*api.PersistentVolume {
	return t.pvList
}

// Get PVC list from the task.
func (t *Task) PVCList() []*api.PersistentVolumeClaim {
	return t.pvcList
}

func (t *Task) Cluster() *repository.ClusterSummary {
	return t.cluster
}

func (t *Task) String() string {
	var nodes []string
	for _, node := range t.nodeList {
		nodes = append(nodes, node.GetName())
	}
	return "[id: " + t.name + ", node: " + strings.Join(nodes, ",") + "]"
}

type TaskResultState string

// A TaskResult contains a state, indicate whether the task is finished successfully; a err if there is any; a list of
// EntityDTO.
type TaskResult struct {
	workerID             string
	state                TaskResultState
	err                  error
	content              []*proto.EntityDTO
	namespaceMetrics     []*repository.NamespaceMetrics
	entityGroups         []*repository.EntityGroup
	podEntities          []*repository.KubePod
	kubeControllers      []*repository.KubeController
	containerSpecMetrics []*repository.ContainerSpecMetrics
	podVolumeMetrics     []*repository.PodVolumeMetrics
}

func NewTaskResult(workerID string, state TaskResultState) *TaskResult {
	return &TaskResult{
		workerID: workerID,
		state:    state,
	}
}

func (r *TaskResult) WorkerId() string {
	return r.workerID
}

func (r *TaskResult) State() TaskResultState {
	return r.state
}

func (r *TaskResult) Content() []*proto.EntityDTO {
	return r.content
}

func (r *TaskResult) PodEntities() []*repository.KubePod {
	return r.podEntities
}

func (r *TaskResult) NamespaceMetrics() []*repository.NamespaceMetrics {
	return r.namespaceMetrics
}

func (r *TaskResult) EntityGroups() []*repository.EntityGroup {
	return r.entityGroups
}

func (r *TaskResult) KubeControllers() []*repository.KubeController {
	return r.kubeControllers
}

func (r *TaskResult) ContainerSpecMetrics() []*repository.ContainerSpecMetrics {
	return r.containerSpecMetrics
}

func (r *TaskResult) PodVolumeMetrics() []*repository.PodVolumeMetrics {
	return r.podVolumeMetrics
}

func (r *TaskResult) Err() error {
	return r.err
}

func (r *TaskResult) WithErr(err error) *TaskResult {
	r.err = err
	return r
}

func (r *TaskResult) WithContent(entityDTOs []*proto.EntityDTO) *TaskResult {
	r.content = entityDTOs
	return r
}

func (r *TaskResult) WithPodEntities(podEntities []*repository.KubePod) *TaskResult {
	r.podEntities = podEntities
	return r
}

func (r *TaskResult) WithNamespaceMetrics(namespaceMetrics []*repository.NamespaceMetrics) *TaskResult {
	r.namespaceMetrics = namespaceMetrics
	return r
}

func (r *TaskResult) WithEntityGroups(entityGroups []*repository.EntityGroup) *TaskResult {
	r.entityGroups = entityGroups
	return r
}

func (r *TaskResult) WithKubeControllers(kubeControllers []*repository.KubeController) *TaskResult {
	r.kubeControllers = kubeControllers
	return r
}

func (r *TaskResult) WithContainerSpecMetrics(containerSpecMetrics []*repository.ContainerSpecMetrics) *TaskResult {
	r.containerSpecMetrics = containerSpecMetrics
	return r
}

func (r *TaskResult) WithPodVolumeMetrics(podVolumeMetrics []*repository.PodVolumeMetrics) *TaskResult {
	r.podVolumeMetrics = podVolumeMetrics
	return r
}
