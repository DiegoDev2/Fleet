package main

// Velero - Disaster recovery for Kubernetes resources and persistent volumes
// Homepage: https://velero.io/

import (
	"fmt"
	
	"os/exec"
)

func installVelero() {
	// Método 1: Descargar y extraer .tar.gz
	velero_tar_url := "https://github.com/vmware-tanzu/velero/archive/refs/tags/v1.14.1.tar.gz"
	velero_cmd_tar := exec.Command("curl", "-L", velero_tar_url, "-o", "package.tar.gz")
	err := velero_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	velero_zip_url := "https://github.com/vmware-tanzu/velero/archive/refs/tags/v1.14.1.zip"
	velero_cmd_zip := exec.Command("curl", "-L", velero_zip_url, "-o", "package.zip")
	err = velero_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	velero_bin_url := "https://github.com/vmware-tanzu/velero/archive/refs/tags/v1.14.1.bin"
	velero_cmd_bin := exec.Command("curl", "-L", velero_bin_url, "-o", "binary.bin")
	err = velero_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	velero_src_url := "https://github.com/vmware-tanzu/velero/archive/refs/tags/v1.14.1.src.tar.gz"
	velero_cmd_src := exec.Command("curl", "-L", velero_src_url, "-o", "source.tar.gz")
	err = velero_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	velero_cmd_direct := exec.Command("./binary")
	err = velero_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
