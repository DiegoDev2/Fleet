package main

// Karmadactl - CLI for Karmada control plane
// Homepage: https://karmada.io/

import (
	"fmt"
	
	"os/exec"
)

func installKarmadactl() {
	// Método 1: Descargar y extraer .tar.gz
	karmadactl_tar_url := "https://github.com/karmada-io/karmada/archive/refs/tags/v1.11.0.tar.gz"
	karmadactl_cmd_tar := exec.Command("curl", "-L", karmadactl_tar_url, "-o", "package.tar.gz")
	err := karmadactl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	karmadactl_zip_url := "https://github.com/karmada-io/karmada/archive/refs/tags/v1.11.0.zip"
	karmadactl_cmd_zip := exec.Command("curl", "-L", karmadactl_zip_url, "-o", "package.zip")
	err = karmadactl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	karmadactl_bin_url := "https://github.com/karmada-io/karmada/archive/refs/tags/v1.11.0.bin"
	karmadactl_cmd_bin := exec.Command("curl", "-L", karmadactl_bin_url, "-o", "binary.bin")
	err = karmadactl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	karmadactl_src_url := "https://github.com/karmada-io/karmada/archive/refs/tags/v1.11.0.src.tar.gz"
	karmadactl_cmd_src := exec.Command("curl", "-L", karmadactl_src_url, "-o", "source.tar.gz")
	err = karmadactl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	karmadactl_cmd_direct := exec.Command("./binary")
	err = karmadactl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
