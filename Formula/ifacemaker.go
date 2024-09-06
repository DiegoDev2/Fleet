package main

// Ifacemaker - Generate interfaces from structure methods
// Homepage: https://github.com/vburenin/ifacemaker

import (
	"fmt"
	
	"os/exec"
)

func installIfacemaker() {
	// Método 1: Descargar y extraer .tar.gz
	ifacemaker_tar_url := "https://github.com/vburenin/ifacemaker/archive/refs/tags/v1.2.1.tar.gz"
	ifacemaker_cmd_tar := exec.Command("curl", "-L", ifacemaker_tar_url, "-o", "package.tar.gz")
	err := ifacemaker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ifacemaker_zip_url := "https://github.com/vburenin/ifacemaker/archive/refs/tags/v1.2.1.zip"
	ifacemaker_cmd_zip := exec.Command("curl", "-L", ifacemaker_zip_url, "-o", "package.zip")
	err = ifacemaker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ifacemaker_bin_url := "https://github.com/vburenin/ifacemaker/archive/refs/tags/v1.2.1.bin"
	ifacemaker_cmd_bin := exec.Command("curl", "-L", ifacemaker_bin_url, "-o", "binary.bin")
	err = ifacemaker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ifacemaker_src_url := "https://github.com/vburenin/ifacemaker/archive/refs/tags/v1.2.1.src.tar.gz"
	ifacemaker_cmd_src := exec.Command("curl", "-L", ifacemaker_src_url, "-o", "source.tar.gz")
	err = ifacemaker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ifacemaker_cmd_direct := exec.Command("./binary")
	err = ifacemaker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
