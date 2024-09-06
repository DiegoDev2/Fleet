package main

// K3d - Little helper to run CNCF's k3s in Docker
// Homepage: https://k3d.io

import (
	"fmt"
	
	"os/exec"
)

func installK3d() {
	// Método 1: Descargar y extraer .tar.gz
	k3d_tar_url := "https://github.com/k3d-io/k3d/archive/refs/tags/v5.7.3.tar.gz"
	k3d_cmd_tar := exec.Command("curl", "-L", k3d_tar_url, "-o", "package.tar.gz")
	err := k3d_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	k3d_zip_url := "https://github.com/k3d-io/k3d/archive/refs/tags/v5.7.3.zip"
	k3d_cmd_zip := exec.Command("curl", "-L", k3d_zip_url, "-o", "package.zip")
	err = k3d_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	k3d_bin_url := "https://github.com/k3d-io/k3d/archive/refs/tags/v5.7.3.bin"
	k3d_cmd_bin := exec.Command("curl", "-L", k3d_bin_url, "-o", "binary.bin")
	err = k3d_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	k3d_src_url := "https://github.com/k3d-io/k3d/archive/refs/tags/v5.7.3.src.tar.gz"
	k3d_cmd_src := exec.Command("curl", "-L", k3d_src_url, "-o", "source.tar.gz")
	err = k3d_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	k3d_cmd_direct := exec.Command("./binary")
	err = k3d_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
