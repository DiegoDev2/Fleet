package main

// K2tf - Kubernetes YAML to Terraform HCL converter
// Homepage: https://github.com/sl1pm4t/k2tf

import (
	"fmt"
	
	"os/exec"
)

func installK2tf() {
	// Método 1: Descargar y extraer .tar.gz
	k2tf_tar_url := "https://github.com/sl1pm4t/k2tf/archive/refs/tags/v0.8.0.tar.gz"
	k2tf_cmd_tar := exec.Command("curl", "-L", k2tf_tar_url, "-o", "package.tar.gz")
	err := k2tf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	k2tf_zip_url := "https://github.com/sl1pm4t/k2tf/archive/refs/tags/v0.8.0.zip"
	k2tf_cmd_zip := exec.Command("curl", "-L", k2tf_zip_url, "-o", "package.zip")
	err = k2tf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	k2tf_bin_url := "https://github.com/sl1pm4t/k2tf/archive/refs/tags/v0.8.0.bin"
	k2tf_cmd_bin := exec.Command("curl", "-L", k2tf_bin_url, "-o", "binary.bin")
	err = k2tf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	k2tf_src_url := "https://github.com/sl1pm4t/k2tf/archive/refs/tags/v0.8.0.src.tar.gz"
	k2tf_cmd_src := exec.Command("curl", "-L", k2tf_src_url, "-o", "source.tar.gz")
	err = k2tf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	k2tf_cmd_direct := exec.Command("./binary")
	err = k2tf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
