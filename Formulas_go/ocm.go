package main

// Ocm - CLI for the Red Hat OpenShift Cluster Manager
// Homepage: https://www.openshift.com/

import (
	"fmt"
	
	"os/exec"
)

func installOcm() {
	// Método 1: Descargar y extraer .tar.gz
	ocm_tar_url := "https://github.com/openshift-online/ocm-cli/archive/refs/tags/v0.1.75.tar.gz"
	ocm_cmd_tar := exec.Command("curl", "-L", ocm_tar_url, "-o", "package.tar.gz")
	err := ocm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocm_zip_url := "https://github.com/openshift-online/ocm-cli/archive/refs/tags/v0.1.75.zip"
	ocm_cmd_zip := exec.Command("curl", "-L", ocm_zip_url, "-o", "package.zip")
	err = ocm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocm_bin_url := "https://github.com/openshift-online/ocm-cli/archive/refs/tags/v0.1.75.bin"
	ocm_cmd_bin := exec.Command("curl", "-L", ocm_bin_url, "-o", "binary.bin")
	err = ocm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocm_src_url := "https://github.com/openshift-online/ocm-cli/archive/refs/tags/v0.1.75.src.tar.gz"
	ocm_cmd_src := exec.Command("curl", "-L", ocm_src_url, "-o", "source.tar.gz")
	err = ocm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocm_cmd_direct := exec.Command("./binary")
	err = ocm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
