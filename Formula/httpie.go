package main

// Httpie - User-friendly cURL replacement (command-line HTTP client)
// Homepage: https://httpie.io/

import (
	"fmt"
	
	"os/exec"
)

func installHttpie() {
	// Método 1: Descargar y extraer .tar.gz
	httpie_tar_url := "https://github.com/httpie/cli/archive/refs/tags/3.2.3.tar.gz"
	httpie_cmd_tar := exec.Command("curl", "-L", httpie_tar_url, "-o", "package.tar.gz")
	err := httpie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpie_zip_url := "https://github.com/httpie/cli/archive/refs/tags/3.2.3.zip"
	httpie_cmd_zip := exec.Command("curl", "-L", httpie_zip_url, "-o", "package.zip")
	err = httpie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpie_bin_url := "https://github.com/httpie/cli/archive/refs/tags/3.2.3.bin"
	httpie_cmd_bin := exec.Command("curl", "-L", httpie_bin_url, "-o", "binary.bin")
	err = httpie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpie_src_url := "https://github.com/httpie/cli/archive/refs/tags/3.2.3.src.tar.gz"
	httpie_cmd_src := exec.Command("curl", "-L", httpie_src_url, "-o", "source.tar.gz")
	err = httpie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpie_cmd_direct := exec.Command("./binary")
	err = httpie_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
