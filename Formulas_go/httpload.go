package main

// HttpLoad - Test throughput of a web server by running parallel fetches
// Homepage: https://www.acme.com/software/http_load/

import (
	"fmt"
	
	"os/exec"
)

func installHttpLoad() {
	// Método 1: Descargar y extraer .tar.gz
	httpload_tar_url := "https://www.acme.com/software/http_load/http_load-09Mar2016.tar.gz"
	httpload_cmd_tar := exec.Command("curl", "-L", httpload_tar_url, "-o", "package.tar.gz")
	err := httpload_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	httpload_zip_url := "https://www.acme.com/software/http_load/http_load-09Mar2016.zip"
	httpload_cmd_zip := exec.Command("curl", "-L", httpload_zip_url, "-o", "package.zip")
	err = httpload_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	httpload_bin_url := "https://www.acme.com/software/http_load/http_load-09Mar2016.bin"
	httpload_cmd_bin := exec.Command("curl", "-L", httpload_bin_url, "-o", "binary.bin")
	err = httpload_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	httpload_src_url := "https://www.acme.com/software/http_load/http_load-09Mar2016.src.tar.gz"
	httpload_cmd_src := exec.Command("curl", "-L", httpload_src_url, "-o", "source.tar.gz")
	err = httpload_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	httpload_cmd_direct := exec.Command("./binary")
	err = httpload_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
