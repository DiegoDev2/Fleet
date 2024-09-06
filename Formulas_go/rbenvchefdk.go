package main

// RbenvChefdk - Treat ChefDK as another version in rbenv
// Homepage: https://github.com/docwhat/rbenv-chefdk

import (
	"fmt"
	
	"os/exec"
)

func installRbenvChefdk() {
	// Método 1: Descargar y extraer .tar.gz
	rbenvchefdk_tar_url := "https://github.com/docwhat/rbenv-chefdk/archive/refs/tags/v1.0.0.tar.gz"
	rbenvchefdk_cmd_tar := exec.Command("curl", "-L", rbenvchefdk_tar_url, "-o", "package.tar.gz")
	err := rbenvchefdk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenvchefdk_zip_url := "https://github.com/docwhat/rbenv-chefdk/archive/refs/tags/v1.0.0.zip"
	rbenvchefdk_cmd_zip := exec.Command("curl", "-L", rbenvchefdk_zip_url, "-o", "package.zip")
	err = rbenvchefdk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenvchefdk_bin_url := "https://github.com/docwhat/rbenv-chefdk/archive/refs/tags/v1.0.0.bin"
	rbenvchefdk_cmd_bin := exec.Command("curl", "-L", rbenvchefdk_bin_url, "-o", "binary.bin")
	err = rbenvchefdk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenvchefdk_src_url := "https://github.com/docwhat/rbenv-chefdk/archive/refs/tags/v1.0.0.src.tar.gz"
	rbenvchefdk_cmd_src := exec.Command("curl", "-L", rbenvchefdk_src_url, "-o", "source.tar.gz")
	err = rbenvchefdk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenvchefdk_cmd_direct := exec.Command("./binary")
	err = rbenvchefdk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rbenv")
exec.Command("latte", "install", "rbenv")
}
