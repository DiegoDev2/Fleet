package main

// Kubescape - Kubernetes testing according to Hardening Guidance by NSA and CISA
// Homepage: https://kubescape.io

import (
	"fmt"
	
	"os/exec"
)

func installKubescape() {
	// Método 1: Descargar y extraer .tar.gz
	kubescape_tar_url := "https://github.com/kubescape/kubescape/archive/refs/tags/v3.0.16.tar.gz"
	kubescape_cmd_tar := exec.Command("curl", "-L", kubescape_tar_url, "-o", "package.tar.gz")
	err := kubescape_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubescape_zip_url := "https://github.com/kubescape/kubescape/archive/refs/tags/v3.0.16.zip"
	kubescape_cmd_zip := exec.Command("curl", "-L", kubescape_zip_url, "-o", "package.zip")
	err = kubescape_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubescape_bin_url := "https://github.com/kubescape/kubescape/archive/refs/tags/v3.0.16.bin"
	kubescape_cmd_bin := exec.Command("curl", "-L", kubescape_bin_url, "-o", "binary.bin")
	err = kubescape_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubescape_src_url := "https://github.com/kubescape/kubescape/archive/refs/tags/v3.0.16.src.tar.gz"
	kubescape_cmd_src := exec.Command("curl", "-L", kubescape_src_url, "-o", "source.tar.gz")
	err = kubescape_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubescape_cmd_direct := exec.Command("./binary")
	err = kubescape_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
