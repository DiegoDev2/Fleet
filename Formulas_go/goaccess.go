package main

// Goaccess - Log analyzer and interactive viewer for the Apache Webserver
// Homepage: https://goaccess.io/

import (
	"fmt"
	
	"os/exec"
)

func installGoaccess() {
	// Método 1: Descargar y extraer .tar.gz
	goaccess_tar_url := "https://tar.goaccess.io/goaccess-1.9.3.tar.gz"
	goaccess_cmd_tar := exec.Command("curl", "-L", goaccess_tar_url, "-o", "package.tar.gz")
	err := goaccess_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goaccess_zip_url := "https://tar.goaccess.io/goaccess-1.9.3.zip"
	goaccess_cmd_zip := exec.Command("curl", "-L", goaccess_zip_url, "-o", "package.zip")
	err = goaccess_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goaccess_bin_url := "https://tar.goaccess.io/goaccess-1.9.3.bin"
	goaccess_cmd_bin := exec.Command("curl", "-L", goaccess_bin_url, "-o", "binary.bin")
	err = goaccess_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goaccess_src_url := "https://tar.goaccess.io/goaccess-1.9.3.src.tar.gz"
	goaccess_cmd_src := exec.Command("curl", "-L", goaccess_src_url, "-o", "source.tar.gz")
	err = goaccess_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goaccess_cmd_direct := exec.Command("./binary")
	err = goaccess_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libmaxminddb")
exec.Command("latte", "install", "libmaxminddb")
	fmt.Println("Instalando dependencia: tokyo-cabinet")
exec.Command("latte", "install", "tokyo-cabinet")
}
