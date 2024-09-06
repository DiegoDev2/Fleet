package main

// Chruby - Ruby environment tool
// Homepage: https://github.com/postmodern/chruby

import (
	"fmt"
	
	"os/exec"
)

func installChruby() {
	// Método 1: Descargar y extraer .tar.gz
	chruby_tar_url := "https://github.com/postmodern/chruby/releases/download/v0.3.9/chruby-0.3.9.tar.gz"
	chruby_cmd_tar := exec.Command("curl", "-L", chruby_tar_url, "-o", "package.tar.gz")
	err := chruby_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chruby_zip_url := "https://github.com/postmodern/chruby/releases/download/v0.3.9/chruby-0.3.9.zip"
	chruby_cmd_zip := exec.Command("curl", "-L", chruby_zip_url, "-o", "package.zip")
	err = chruby_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chruby_bin_url := "https://github.com/postmodern/chruby/releases/download/v0.3.9/chruby-0.3.9.bin"
	chruby_cmd_bin := exec.Command("curl", "-L", chruby_bin_url, "-o", "binary.bin")
	err = chruby_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chruby_src_url := "https://github.com/postmodern/chruby/releases/download/v0.3.9/chruby-0.3.9.src.tar.gz"
	chruby_cmd_src := exec.Command("curl", "-L", chruby_src_url, "-o", "source.tar.gz")
	err = chruby_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chruby_cmd_direct := exec.Command("./binary")
	err = chruby_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
