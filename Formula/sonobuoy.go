package main

// Sonobuoy - Kubernetes component that generates reports on cluster conformance
// Homepage: https://github.com/vmware-tanzu/sonobuoy

import (
	"fmt"
	
	"os/exec"
)

func installSonobuoy() {
	// Método 1: Descargar y extraer .tar.gz
	sonobuoy_tar_url := "https://github.com/vmware-tanzu/sonobuoy/archive/refs/tags/v0.57.2.tar.gz"
	sonobuoy_cmd_tar := exec.Command("curl", "-L", sonobuoy_tar_url, "-o", "package.tar.gz")
	err := sonobuoy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sonobuoy_zip_url := "https://github.com/vmware-tanzu/sonobuoy/archive/refs/tags/v0.57.2.zip"
	sonobuoy_cmd_zip := exec.Command("curl", "-L", sonobuoy_zip_url, "-o", "package.zip")
	err = sonobuoy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sonobuoy_bin_url := "https://github.com/vmware-tanzu/sonobuoy/archive/refs/tags/v0.57.2.bin"
	sonobuoy_cmd_bin := exec.Command("curl", "-L", sonobuoy_bin_url, "-o", "binary.bin")
	err = sonobuoy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sonobuoy_src_url := "https://github.com/vmware-tanzu/sonobuoy/archive/refs/tags/v0.57.2.src.tar.gz"
	sonobuoy_cmd_src := exec.Command("curl", "-L", sonobuoy_src_url, "-o", "source.tar.gz")
	err = sonobuoy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sonobuoy_cmd_direct := exec.Command("./binary")
	err = sonobuoy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
