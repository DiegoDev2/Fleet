package main

// PkgConfigWrapper - Easier way to include C code in your Go program
// Homepage: https://github.com/influxdata/pkg-config

import (
	"fmt"
	
	"os/exec"
)

func installPkgConfigWrapper() {
	// Método 1: Descargar y extraer .tar.gz
	pkgconfigwrapper_tar_url := "https://github.com/influxdata/pkg-config/archive/refs/tags/v0.2.13.tar.gz"
	pkgconfigwrapper_cmd_tar := exec.Command("curl", "-L", pkgconfigwrapper_tar_url, "-o", "package.tar.gz")
	err := pkgconfigwrapper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pkgconfigwrapper_zip_url := "https://github.com/influxdata/pkg-config/archive/refs/tags/v0.2.13.zip"
	pkgconfigwrapper_cmd_zip := exec.Command("curl", "-L", pkgconfigwrapper_zip_url, "-o", "package.zip")
	err = pkgconfigwrapper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pkgconfigwrapper_bin_url := "https://github.com/influxdata/pkg-config/archive/refs/tags/v0.2.13.bin"
	pkgconfigwrapper_cmd_bin := exec.Command("curl", "-L", pkgconfigwrapper_bin_url, "-o", "binary.bin")
	err = pkgconfigwrapper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pkgconfigwrapper_src_url := "https://github.com/influxdata/pkg-config/archive/refs/tags/v0.2.13.src.tar.gz"
	pkgconfigwrapper_cmd_src := exec.Command("curl", "-L", pkgconfigwrapper_src_url, "-o", "source.tar.gz")
	err = pkgconfigwrapper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pkgconfigwrapper_cmd_direct := exec.Command("./binary")
	err = pkgconfigwrapper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
