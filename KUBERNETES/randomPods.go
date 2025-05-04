package main

import (
	"fmt"
	"math/rand/v2"
	"os/exec"
)

func randomPodGenerator(namespaces []string, podNames []string) map[string]string {
	var cmd1 string
	var cmd2 string
	plan := make(map[string]string)
	for _, item := range namespaces {
		pod := podNames[rand.IntN(9)]
		cmd1 = fmt.Sprintf("kubectl create ns %v", item)
		cmd2 = fmt.Sprintf("kubectl run %v --image=nginx --namespace %v", pod, item)
		command, err := exec.Command("bash", "-c", cmd1).Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(command))
		}

		command2, err2 := exec.Command("bash", "-c", cmd2).Output()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			plan[item] = pod
			fmt.Println(string(command2))
		}
	}
	return plan
}

func main() {
	namespaces := []string{
		"infra-core",
		"platform-team",
		"observability",
		"vault-services",
		"backup-automation",
		"staging-environment",
		"devops-lab",
		"kube-lab",
		"sre-tools",
		"experimental-zone",
	}

	// deploymentNames := []string{
	// 	"vault-agent-deploy",
	// 	"grafana-frontend",
	// 	"elasticsearch-core",
	// 	"gitlab-runner-agent",
	// 	"zabbix-exporter",
	// 	"sre-healthchecker",
	// 	"kubewatch-logger",
	// 	"metrics-summarizer",
	// 	"vault-daily-backup",
	// 	"logshipper-core",
	// }

	podNames := []string{
		"vault-sidecar",
		"apm-ingestor",
		"node-scan-agent",
		"healthcheck-probe",
		"reboot-watcher",
		"cert-autorenewer",
		"alert-bot",
		"file-syncer",
		"job-monitor",
	}

	plan := randomPodGenerator(namespaces, podNames)
	fmt.Println(plan)
}
