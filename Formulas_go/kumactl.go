package main

// Kumactl - Kuma control plane command-line utility
// Homepage: https://kuma.io/

import (
	"fmt"
	
	"os/exec"
)

func installKumactl() {
	// Método 1: Descargar y extraer .tar.gz
	kumactl_tar_url := "https://github.com/kumahq/kuma/archive/refs/tags/2.8.3.tar.gz"
	kumactl_cmd_tar := exec.Command("curl", "-L", kumactl_tar_url, "-o", "package.tar.gz")
	err := kumactl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kumactl_zip_url := "https://github.com/kumahq/kuma/archive/refs/tags/2.8.3.zip"
	kumactl_cmd_zip := exec.Command("curl", "-L", kumactl_zip_url, "-o", "package.zip")
	err = kumactl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kumactl_bin_url := "https://github.com/kumahq/kuma/archive/refs/tags/2.8.3.bin"
	kumactl_cmd_bin := exec.Command("curl", "-L", kumactl_bin_url, "-o", "binary.bin")
	err = kumactl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kumactl_src_url := "https://github.com/kumahq/kuma/archive/refs/tags/2.8.3.src.tar.gz"
	kumactl_cmd_src := exec.Command("curl", "-L", kumactl_src_url, "-o", "source.tar.gz")
	err = kumactl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kumactl_cmd_direct := exec.Command("./binary")
	err = kumactl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
