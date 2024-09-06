package main

// Ohcount - Source code line counter
// Homepage: https://github.com/blackducksoftware/ohcount

import (
	"fmt"
	
	"os/exec"
)

func installOhcount() {
	// Método 1: Descargar y extraer .tar.gz
	ohcount_tar_url := "https://github.com/blackducksoftware/ohcount/archive/refs/tags/4.0.0.tar.gz"
	ohcount_cmd_tar := exec.Command("curl", "-L", ohcount_tar_url, "-o", "package.tar.gz")
	err := ohcount_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ohcount_zip_url := "https://github.com/blackducksoftware/ohcount/archive/refs/tags/4.0.0.zip"
	ohcount_cmd_zip := exec.Command("curl", "-L", ohcount_zip_url, "-o", "package.zip")
	err = ohcount_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ohcount_bin_url := "https://github.com/blackducksoftware/ohcount/archive/refs/tags/4.0.0.bin"
	ohcount_cmd_bin := exec.Command("curl", "-L", ohcount_bin_url, "-o", "binary.bin")
	err = ohcount_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ohcount_src_url := "https://github.com/blackducksoftware/ohcount/archive/refs/tags/4.0.0.src.tar.gz"
	ohcount_cmd_src := exec.Command("curl", "-L", ohcount_src_url, "-o", "source.tar.gz")
	err = ohcount_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ohcount_cmd_direct := exec.Command("./binary")
	err = ohcount_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gperf")
	exec.Command("latte", "install", "gperf").Run()
	fmt.Println("Instalando dependencia: libmagic")
	exec.Command("latte", "install", "libmagic").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: ragel")
	exec.Command("latte", "install", "ragel").Run()
}
