package main

// Htmlcxx - Non-validating CSS1 and HTML parser for C++
// Homepage: https://htmlcxx.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installHtmlcxx() {
	// Método 1: Descargar y extraer .tar.gz
	htmlcxx_tar_url := "https://downloads.sourceforge.net/project/htmlcxx/v0.87/htmlcxx-0.87.tar.gz"
	htmlcxx_cmd_tar := exec.Command("curl", "-L", htmlcxx_tar_url, "-o", "package.tar.gz")
	err := htmlcxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	htmlcxx_zip_url := "https://downloads.sourceforge.net/project/htmlcxx/v0.87/htmlcxx-0.87.zip"
	htmlcxx_cmd_zip := exec.Command("curl", "-L", htmlcxx_zip_url, "-o", "package.zip")
	err = htmlcxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	htmlcxx_bin_url := "https://downloads.sourceforge.net/project/htmlcxx/v0.87/htmlcxx-0.87.bin"
	htmlcxx_cmd_bin := exec.Command("curl", "-L", htmlcxx_bin_url, "-o", "binary.bin")
	err = htmlcxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	htmlcxx_src_url := "https://downloads.sourceforge.net/project/htmlcxx/v0.87/htmlcxx-0.87.src.tar.gz"
	htmlcxx_cmd_src := exec.Command("curl", "-L", htmlcxx_src_url, "-o", "source.tar.gz")
	err = htmlcxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	htmlcxx_cmd_direct := exec.Command("./binary")
	err = htmlcxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
