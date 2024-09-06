package main

// RonnNg - Build man pages from Markdown
// Homepage: https://github.com/apjanke/ronn-ng

import (
	"fmt"
	
	"os/exec"
)

func installRonnNg() {
	// Método 1: Descargar y extraer .tar.gz
	ronnng_tar_url := "https://github.com/apjanke/ronn-ng/archive/refs/tags/v0.10.1.tar.gz"
	ronnng_cmd_tar := exec.Command("curl", "-L", ronnng_tar_url, "-o", "package.tar.gz")
	err := ronnng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ronnng_zip_url := "https://github.com/apjanke/ronn-ng/archive/refs/tags/v0.10.1.zip"
	ronnng_cmd_zip := exec.Command("curl", "-L", ronnng_zip_url, "-o", "package.zip")
	err = ronnng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ronnng_bin_url := "https://github.com/apjanke/ronn-ng/archive/refs/tags/v0.10.1.bin"
	ronnng_cmd_bin := exec.Command("curl", "-L", ronnng_bin_url, "-o", "binary.bin")
	err = ronnng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ronnng_src_url := "https://github.com/apjanke/ronn-ng/archive/refs/tags/v0.10.1.src.tar.gz"
	ronnng_cmd_src := exec.Command("curl", "-L", ronnng_src_url, "-o", "source.tar.gz")
	err = ronnng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ronnng_cmd_direct := exec.Command("./binary")
	err = ronnng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
}
