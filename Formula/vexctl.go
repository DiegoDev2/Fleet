package main

// Vexctl - Tool to create, transform and attest VEX metadata
// Homepage: https://openssf.org/projects/openvex/

import (
	"fmt"
	
	"os/exec"
)

func installVexctl() {
	// Método 1: Descargar y extraer .tar.gz
	vexctl_tar_url := "https://github.com/openvex/vexctl/archive/refs/tags/v0.2.6.tar.gz"
	vexctl_cmd_tar := exec.Command("curl", "-L", vexctl_tar_url, "-o", "package.tar.gz")
	err := vexctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vexctl_zip_url := "https://github.com/openvex/vexctl/archive/refs/tags/v0.2.6.zip"
	vexctl_cmd_zip := exec.Command("curl", "-L", vexctl_zip_url, "-o", "package.zip")
	err = vexctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vexctl_bin_url := "https://github.com/openvex/vexctl/archive/refs/tags/v0.2.6.bin"
	vexctl_cmd_bin := exec.Command("curl", "-L", vexctl_bin_url, "-o", "binary.bin")
	err = vexctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vexctl_src_url := "https://github.com/openvex/vexctl/archive/refs/tags/v0.2.6.src.tar.gz"
	vexctl_cmd_src := exec.Command("curl", "-L", vexctl_src_url, "-o", "source.tar.gz")
	err = vexctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vexctl_cmd_direct := exec.Command("./binary")
	err = vexctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
