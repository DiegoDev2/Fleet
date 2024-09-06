package main

// Mydumper - How MySQL DBA & support engineer would imagine 'mysqldump' ;-)
// Homepage: https://launchpad.net/mydumper

import (
	"fmt"
	
	"os/exec"
)

func installMydumper() {
	// Método 1: Descargar y extraer .tar.gz
	mydumper_tar_url := "https://github.com/mydumper/mydumper/archive/refs/tags/v0.16.5-1.tar.gz"
	mydumper_cmd_tar := exec.Command("curl", "-L", mydumper_tar_url, "-o", "package.tar.gz")
	err := mydumper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mydumper_zip_url := "https://github.com/mydumper/mydumper/archive/refs/tags/v0.16.5-1.zip"
	mydumper_cmd_zip := exec.Command("curl", "-L", mydumper_zip_url, "-o", "package.zip")
	err = mydumper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mydumper_bin_url := "https://github.com/mydumper/mydumper/archive/refs/tags/v0.16.5-1.bin"
	mydumper_cmd_bin := exec.Command("curl", "-L", mydumper_bin_url, "-o", "binary.bin")
	err = mydumper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mydumper_src_url := "https://github.com/mydumper/mydumper/archive/refs/tags/v0.16.5-1.src.tar.gz"
	mydumper_cmd_src := exec.Command("curl", "-L", mydumper_src_url, "-o", "source.tar.gz")
	err = mydumper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mydumper_cmd_direct := exec.Command("./binary")
	err = mydumper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: mysql-client")
	exec.Command("latte", "install", "mysql-client").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
}
