package types

import (
	"time"
)

type MessageLoki struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Stream struct {
				Pod       string `json:"pod"`
				Stream    string `json:"stream"`
				Component string `json:"component"`
				Container string `json:"container"`
				Filename  string `json:"filename"`
				Namespace string `json:"namespace"`
				App       string `json:"app"`
				Instance  string `json:"instance"`
				Job       string `json:"job"`
				NodeName  string `json:"node_name"`
			} `json:"stream"`
			Values [][]string `json:"values"`
		} `json:"result"`
		Stats struct {
			Summary struct {
				BytesProcessedPerSecond int     `json:"bytesProcessedPerSecond"`
				LinesProcessedPerSecond int     `json:"linesProcessedPerSecond"`
				TotalBytesProcessed     int     `json:"totalBytesProcessed"`
				TotalLinesProcessed     int     `json:"totalLinesProcessed"`
				ExecTime                float64 `json:"execTime"`
				QueueTime               float64 `json:"queueTime"`
				Subqueries              int     `json:"subqueries"`
				TotalEntriesReturned    int     `json:"totalEntriesReturned"`
			} `json:"summary"`
			Querier struct {
				Store struct {
					TotalChunksRef        int `json:"totalChunksRef"`
					TotalChunksDownloaded int `json:"totalChunksDownloaded"`
					ChunksDownloadTime    int `json:"chunksDownloadTime"`
					Chunk                 struct {
						HeadChunkBytes    int `json:"headChunkBytes"`
						HeadChunkLines    int `json:"headChunkLines"`
						DecompressedBytes int `json:"decompressedBytes"`
						DecompressedLines int `json:"decompressedLines"`
						CompressedBytes   int `json:"compressedBytes"`
						TotalDuplicates   int `json:"totalDuplicates"`
					} `json:"chunk"`
				} `json:"store"`
			} `json:"querier"`
			Ingester struct {
				TotalReached       int `json:"totalReached"`
				TotalChunksMatched int `json:"totalChunksMatched"`
				TotalBatches       int `json:"totalBatches"`
				TotalLinesSent     int `json:"totalLinesSent"`
				Store              struct {
					TotalChunksRef        int `json:"totalChunksRef"`
					TotalChunksDownloaded int `json:"totalChunksDownloaded"`
					ChunksDownloadTime    int `json:"chunksDownloadTime"`
					Chunk                 struct {
						HeadChunkBytes    int `json:"headChunkBytes"`
						HeadChunkLines    int `json:"headChunkLines"`
						DecompressedBytes int `json:"decompressedBytes"`
						DecompressedLines int `json:"decompressedLines"`
						CompressedBytes   int `json:"compressedBytes"`
						TotalDuplicates   int `json:"totalDuplicates"`
					} `json:"chunk"`
				} `json:"store"`
			} `json:"ingester"`
		} `json:"stats"`
	} `json:"data"`
}

// MapServiceDeveloper - карта Имя сервиса - telegram логин программиста
var MapServiceDeveloper = make(map[string]string)

// MapSQLDeveloper - карта Имя сервиса - telegram логин программиста
var MapSQLDeveloper = make(map[string]string)

// MapTelegramUsers - карта имя в телеграм - id чата в телеграм
var MapTelegramUsers = make(map[string]int64)

// Message - структура сообщения для отправки в Телеграм
type Message struct {
	ServiceName    string
	DeveloperName  string
	LokiURL        string
	Date           time.Time
	Text           string
	IsSameLastText bool
}

// MessageVictoriaMetrics - структура лога VictoriaMetrics
type MessageVictoriaMetrics struct {
	Time                                              time.Time `json:"_time"`
	StreamID                                          string    `json:"_stream_id"`
	_Stream                                           string    `json:"_stream"`
	Msg                                               string    `json:"_msg"`
	File                                              string    `json:"file"`
	KubernetesContainerID                             string    `json:"kubernetes.container_id"`
	KubernetesContainerImage                          string    `json:"kubernetes.container_image"`
	KubernetesContainerImageID                        string    `json:"kubernetes.container_image_id"`
	KubernetesContainerName                           string    `json:"kubernetes.container_name"`
	KubernetesNamespaceLabelsKubernetesIoMetadataName string    `json:"kubernetes.namespace_labels.kubernetes.io/metadata.name"`
	KubernetesNodeLabelsBetaKubernetesIoArch          string    `json:"kubernetes.node_labels.beta.kubernetes.io/arch"`
	KubernetesNodeLabelsBetaKubernetesIoOs            string    `json:"kubernetes.node_labels.beta.kubernetes.io/os"`
	KubernetesNodeLabelsKubernetesIoArch              string    `json:"kubernetes.node_labels.kubernetes.io/arch"`
	KubernetesNodeLabelsKubernetesIoHostname          string    `json:"kubernetes.node_labels.kubernetes.io/hostname"`
	KubernetesNodeLabelsKubernetesIoOs                string    `json:"kubernetes.node_labels.kubernetes.io/os"`
	KubernetesPodIP                                   string    `json:"kubernetes.pod_ip"`
	KubernetesPodIps                                  string    `json:"kubernetes.pod_ips"`
	KubernetesPodLabelsApp                            string    `json:"kubernetes.pod_labels.app"`
	KubernetesPodLabelsPodTemplateHash                string    `json:"kubernetes.pod_labels.pod-template-hash"`
	KubernetesPodName                                 string    `json:"kubernetes.pod_name"`
	KubernetesPodNamespace                            string    `json:"kubernetes.pod_namespace"`
	KubernetesPodNodeName                             string    `json:"kubernetes.pod_node_name"`
	KubernetesPodOwner                                string    `json:"kubernetes.pod_owner"`
	KubernetesPodUID                                  string    `json:"kubernetes.pod_uid"`
	SourceType                                        string    `json:"source_type"`
	Stream                                            string    `json:"stream"`
}

// MessageLog - структура лога, с нужными колонками
type MessageLog struct {
	Date time.Time
	Text string
}
