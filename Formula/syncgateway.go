package main

// SyncGateway - Make Couchbase Server a replication endpoint for Couchbase Lite
// Homepage: https://docs.couchbase.com/sync-gateway/current/index.html

import (
	"fmt"
	
	"os/exec"
)

func installSyncGateway() {
	// Método 1: Descargar y extraer .tar.gz
	syncgateway_tar_url := "https://github.com/couchbase/sync_gateway.git"
	syncgateway_cmd_tar := exec.Command("curl", "-L", syncgateway_tar_url, "-o", "package.tar.gz")
	err := syncgateway_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	syncgateway_zip_url := "https://github.com/couchbase/sync_gateway.git"
	syncgateway_cmd_zip := exec.Command("curl", "-L", syncgateway_zip_url, "-o", "package.zip")
	err = syncgateway_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	syncgateway_bin_url := "https://github.com/couchbase/sync_gateway.git"
	syncgateway_cmd_bin := exec.Command("curl", "-L", syncgateway_bin_url, "-o", "binary.bin")
	err = syncgateway_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	syncgateway_src_url := "https://github.com/couchbase/sync_gateway.git"
	syncgateway_cmd_src := exec.Command("curl", "-L", syncgateway_src_url, "-o", "source.tar.gz")
	err = syncgateway_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	syncgateway_cmd_direct := exec.Command("./binary")
	err = syncgateway_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnupg")
	exec.Command("latte", "install", "gnupg").Run()
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: repo")
	exec.Command("latte", "install", "repo").Run()
}
