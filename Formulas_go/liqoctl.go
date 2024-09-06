package main

// Liqoctl - Is a CLI tool to install and manage Liqo-enabled clusters
// Homepage: https://liqo.io

import (
	"fmt"
	
	"os/exec"
)

func installLiqoctl() {
	// Método 1: Descargar y extraer .tar.gz
	liqoctl_tar_url := "https://github.com/liqotech/liqo/archive/refs/tags/v0.10.3.tar.gz"
	liqoctl_cmd_tar := exec.Command("curl", "-L", liqoctl_tar_url, "-o", "package.tar.gz")
	err := liqoctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liqoctl_zip_url := "https://github.com/liqotech/liqo/archive/refs/tags/v0.10.3.zip"
	liqoctl_cmd_zip := exec.Command("curl", "-L", liqoctl_zip_url, "-o", "package.zip")
	err = liqoctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liqoctl_bin_url := "https://github.com/liqotech/liqo/archive/refs/tags/v0.10.3.bin"
	liqoctl_cmd_bin := exec.Command("curl", "-L", liqoctl_bin_url, "-o", "binary.bin")
	err = liqoctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liqoctl_src_url := "https://github.com/liqotech/liqo/archive/refs/tags/v0.10.3.src.tar.gz"
	liqoctl_cmd_src := exec.Command("curl", "-L", liqoctl_src_url, "-o", "source.tar.gz")
	err = liqoctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liqoctl_cmd_direct := exec.Command("./binary")
	err = liqoctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
