package main

// Abook - Address book with mutt support
// Homepage: https://abook.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installAbook() {
	// Método 1: Descargar y extraer .tar.gz
	abook_tar_url := "https://abook.sourceforge.io/devel/abook-0.6.1.tar.gz"
	abook_cmd_tar := exec.Command("curl", "-L", abook_tar_url, "-o", "package.tar.gz")
	err := abook_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abook_zip_url := "https://abook.sourceforge.io/devel/abook-0.6.1.zip"
	abook_cmd_zip := exec.Command("curl", "-L", abook_zip_url, "-o", "package.zip")
	err = abook_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abook_bin_url := "https://abook.sourceforge.io/devel/abook-0.6.1.bin"
	abook_cmd_bin := exec.Command("curl", "-L", abook_bin_url, "-o", "binary.bin")
	err = abook_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abook_src_url := "https://abook.sourceforge.io/devel/abook-0.6.1.src.tar.gz"
	abook_cmd_src := exec.Command("curl", "-L", abook_src_url, "-o", "source.tar.gz")
	err = abook_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abook_cmd_direct := exec.Command("./binary")
	err = abook_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
