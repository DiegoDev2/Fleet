package main

// TidyHtml5 - Granddaddy of HTML tools, with support for modern standards
// Homepage: https://www.html-tidy.org/

import (
	"fmt"
	
	"os/exec"
)

func installTidyHtml5() {
	// Método 1: Descargar y extraer .tar.gz
	tidyhtml5_tar_url := "https://github.com/htacg/tidy-html5/archive/refs/tags/5.8.0.tar.gz"
	tidyhtml5_cmd_tar := exec.Command("curl", "-L", tidyhtml5_tar_url, "-o", "package.tar.gz")
	err := tidyhtml5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tidyhtml5_zip_url := "https://github.com/htacg/tidy-html5/archive/refs/tags/5.8.0.zip"
	tidyhtml5_cmd_zip := exec.Command("curl", "-L", tidyhtml5_zip_url, "-o", "package.zip")
	err = tidyhtml5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tidyhtml5_bin_url := "https://github.com/htacg/tidy-html5/archive/refs/tags/5.8.0.bin"
	tidyhtml5_cmd_bin := exec.Command("curl", "-L", tidyhtml5_bin_url, "-o", "binary.bin")
	err = tidyhtml5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tidyhtml5_src_url := "https://github.com/htacg/tidy-html5/archive/refs/tags/5.8.0.src.tar.gz"
	tidyhtml5_cmd_src := exec.Command("curl", "-L", tidyhtml5_src_url, "-o", "source.tar.gz")
	err = tidyhtml5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tidyhtml5_cmd_direct := exec.Command("./binary")
	err = tidyhtml5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
