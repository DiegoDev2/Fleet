package main

// Po4a - Documentation translation maintenance tool
// Homepage: https://po4a.org

import (
	"fmt"
	
	"os/exec"
)

func installPo4a() {
	// Método 1: Descargar y extraer .tar.gz
	po4a_tar_url := "https://github.com/mquinson/po4a/archive/refs/tags/v0.73.tar.gz"
	po4a_cmd_tar := exec.Command("curl", "-L", po4a_tar_url, "-o", "package.tar.gz")
	err := po4a_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	po4a_zip_url := "https://github.com/mquinson/po4a/archive/refs/tags/v0.73.zip"
	po4a_cmd_zip := exec.Command("curl", "-L", po4a_zip_url, "-o", "package.zip")
	err = po4a_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	po4a_bin_url := "https://github.com/mquinson/po4a/archive/refs/tags/v0.73.bin"
	po4a_cmd_bin := exec.Command("curl", "-L", po4a_bin_url, "-o", "binary.bin")
	err = po4a_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	po4a_src_url := "https://github.com/mquinson/po4a/archive/refs/tags/v0.73.src.tar.gz"
	po4a_cmd_src := exec.Command("curl", "-L", po4a_src_url, "-o", "source.tar.gz")
	err = po4a_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	po4a_cmd_direct := exec.Command("./binary")
	err = po4a_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
}
