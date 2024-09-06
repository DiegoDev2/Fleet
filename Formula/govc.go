package main

// Govc - Command-line tool for VMware vSphere
// Homepage: https://github.com/vmware/govmomi/tree/master/govc

import (
	"fmt"
	
	"os/exec"
)

func installGovc() {
	// Método 1: Descargar y extraer .tar.gz
	govc_tar_url := "https://github.com/vmware/govmomi/archive/refs/tags/v0.42.0.tar.gz"
	govc_cmd_tar := exec.Command("curl", "-L", govc_tar_url, "-o", "package.tar.gz")
	err := govc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	govc_zip_url := "https://github.com/vmware/govmomi/archive/refs/tags/v0.42.0.zip"
	govc_cmd_zip := exec.Command("curl", "-L", govc_zip_url, "-o", "package.zip")
	err = govc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	govc_bin_url := "https://github.com/vmware/govmomi/archive/refs/tags/v0.42.0.bin"
	govc_cmd_bin := exec.Command("curl", "-L", govc_bin_url, "-o", "binary.bin")
	err = govc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	govc_src_url := "https://github.com/vmware/govmomi/archive/refs/tags/v0.42.0.src.tar.gz"
	govc_cmd_src := exec.Command("curl", "-L", govc_src_url, "-o", "source.tar.gz")
	err = govc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	govc_cmd_direct := exec.Command("./binary")
	err = govc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
