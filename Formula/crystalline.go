package main

// Crystalline - Language Server Protocol implementation for Crystal
// Homepage: https://github.com/elbywan/crystalline

import (
	"fmt"
	
	"os/exec"
)

func installCrystalline() {
	// Método 1: Descargar y extraer .tar.gz
	crystalline_tar_url := "https://github.com/elbywan/crystalline/archive/refs/tags/v0.13.1.tar.gz"
	crystalline_cmd_tar := exec.Command("curl", "-L", crystalline_tar_url, "-o", "package.tar.gz")
	err := crystalline_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crystalline_zip_url := "https://github.com/elbywan/crystalline/archive/refs/tags/v0.13.1.zip"
	crystalline_cmd_zip := exec.Command("curl", "-L", crystalline_zip_url, "-o", "package.zip")
	err = crystalline_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crystalline_bin_url := "https://github.com/elbywan/crystalline/archive/refs/tags/v0.13.1.bin"
	crystalline_cmd_bin := exec.Command("curl", "-L", crystalline_bin_url, "-o", "binary.bin")
	err = crystalline_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crystalline_src_url := "https://github.com/elbywan/crystalline/archive/refs/tags/v0.13.1.src.tar.gz"
	crystalline_cmd_src := exec.Command("curl", "-L", crystalline_src_url, "-o", "source.tar.gz")
	err = crystalline_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crystalline_cmd_direct := exec.Command("./binary")
	err = crystalline_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: crystal")
	exec.Command("latte", "install", "crystal").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
