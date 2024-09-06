package main

// Libbi - Bayesian state-space modelling on parallel computer hardware
// Homepage: https://libbi.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibbi() {
	// Método 1: Descargar y extraer .tar.gz
	libbi_tar_url := "https://github.com/lawmurray/LibBi/archive/refs/tags/1.4.5.tar.gz"
	libbi_cmd_tar := exec.Command("curl", "-L", libbi_tar_url, "-o", "package.tar.gz")
	err := libbi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbi_zip_url := "https://github.com/lawmurray/LibBi/archive/refs/tags/1.4.5.zip"
	libbi_cmd_zip := exec.Command("curl", "-L", libbi_zip_url, "-o", "package.zip")
	err = libbi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbi_bin_url := "https://github.com/lawmurray/LibBi/archive/refs/tags/1.4.5.bin"
	libbi_cmd_bin := exec.Command("curl", "-L", libbi_bin_url, "-o", "binary.bin")
	err = libbi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbi_src_url := "https://github.com/lawmurray/LibBi/archive/refs/tags/1.4.5.src.tar.gz"
	libbi_cmd_src := exec.Command("curl", "-L", libbi_src_url, "-o", "source.tar.gz")
	err = libbi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbi_cmd_direct := exec.Command("./binary")
	err = libbi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: netcdf")
exec.Command("latte", "install", "netcdf")
	fmt.Println("Instalando dependencia: qrupdate")
exec.Command("latte", "install", "qrupdate")
}
