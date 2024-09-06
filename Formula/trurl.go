package main

// Trurl - Command-line tool for URL parsing and manipulation
// Homepage: https://curl.se/trurl/

import (
	"fmt"
	
	"os/exec"
)

func installTrurl() {
	// Método 1: Descargar y extraer .tar.gz
	trurl_tar_url := "https://github.com/curl/trurl/releases/download/trurl-0.15/trurl-0.15.tar.gz"
	trurl_cmd_tar := exec.Command("curl", "-L", trurl_tar_url, "-o", "package.tar.gz")
	err := trurl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trurl_zip_url := "https://github.com/curl/trurl/releases/download/trurl-0.15/trurl-0.15.zip"
	trurl_cmd_zip := exec.Command("curl", "-L", trurl_zip_url, "-o", "package.zip")
	err = trurl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trurl_bin_url := "https://github.com/curl/trurl/releases/download/trurl-0.15/trurl-0.15.bin"
	trurl_cmd_bin := exec.Command("curl", "-L", trurl_bin_url, "-o", "binary.bin")
	err = trurl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trurl_src_url := "https://github.com/curl/trurl/releases/download/trurl-0.15/trurl-0.15.src.tar.gz"
	trurl_cmd_src := exec.Command("curl", "-L", trurl_src_url, "-o", "source.tar.gz")
	err = trurl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trurl_cmd_direct := exec.Command("./binary")
	err = trurl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
