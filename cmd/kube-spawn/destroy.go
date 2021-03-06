/*
Copyright 2017 Kinvolk GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kinvolk/kube-spawn/pkg/cluster"
)

var (
	destroyCmd = &cobra.Command{
		Use:   "destroy",
		Short: "Remove the cluster environment",
		Long: `Remove the cluster environment.
Stops the cluster if it it running`,
		Run: runDestroy,
	}
)

func init() {
	kubespawnCmd.AddCommand(destroyCmd)
}

func runDestroy(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		log.Fatalf("Command destroy doesn't take arguments, got: %v", args)
	}

	kubespawnDir := viper.GetString("dir")
	clusterName := viper.GetString("cluster-name")
	clusterDir := path.Join(kubespawnDir, "clusters", clusterName)

	kluster, err := cluster.New(clusterDir, clusterName)
	if err != nil {
		log.Fatalf("Failed to create cluster object: %v", err)
	}

	log.Printf("Destroying cluster %s ...", clusterName)

	if err := kluster.Destroy(); err != nil {
		log.Fatalf("Failed to destroy cluster: %v", err)
	}

	log.Printf("Cluster %s destroyed", clusterName)
}
