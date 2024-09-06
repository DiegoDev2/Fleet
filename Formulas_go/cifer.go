package main

// Cifer - Work on automating classical cipher cracking in C
// Homepage: https://code.google.com/p/cifer/

import (
	"fmt"
	
	"os/exec"
)

func installCifer() {
	// Método 1: Descargar y extraer .tar.gz
	cifer_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cifer/cifer-1.2.0.tar.gz"
	cifer_cmd_tar := exec.Command("curl", "-L", cifer_tar_url, "-o", "package.tar.gz")
	err := cifer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cifer_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cifer/cifer-1.2.0.zip"
	cifer_cmd_zip := exec.Command("curl", "-L", cifer_zip_url, "-o", "package.zip")
	err = cifer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cifer_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cifer/cifer-1.2.0.bin"
	cifer_cmd_bin := exec.Command("curl", "-L", cifer_bin_url, "-o", "binary.bin")
	err = cifer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cifer_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cifer/cifer-1.2.0.src.tar.gz"
	cifer_cmd_src := exec.Command("curl", "-L", cifer_src_url, "-o", "source.tar.gz")
	err = cifer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cifer_cmd_direct := exec.Command("./binary")
	err = cifer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
