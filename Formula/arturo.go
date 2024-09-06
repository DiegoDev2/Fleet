package main

// Arturo - Simple, modern and portable programming language for efficient scripting
// Homepage: https://arturo-lang.io/

import (
	"fmt"
	
	"os/exec"
)

func installArturo() {
	// Método 1: Descargar y extraer .tar.gz
	arturo_tar_url := "https://github.com/arturo-lang/arturo/archive/refs/tags/v0.9.83.tar.gz"
	arturo_cmd_tar := exec.Command("curl", "-L", arturo_tar_url, "-o", "package.tar.gz")
	err := arturo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arturo_zip_url := "https://github.com/arturo-lang/arturo/archive/refs/tags/v0.9.83.zip"
	arturo_cmd_zip := exec.Command("curl", "-L", arturo_zip_url, "-o", "package.zip")
	err = arturo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arturo_bin_url := "https://github.com/arturo-lang/arturo/archive/refs/tags/v0.9.83.bin"
	arturo_cmd_bin := exec.Command("curl", "-L", arturo_bin_url, "-o", "binary.bin")
	err = arturo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arturo_src_url := "https://github.com/arturo-lang/arturo/archive/refs/tags/v0.9.83.src.tar.gz"
	arturo_cmd_src := exec.Command("curl", "-L", arturo_src_url, "-o", "source.tar.gz")
	err = arturo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arturo_cmd_direct := exec.Command("./binary")
	err = arturo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
	fmt.Println("Instalando dependencia: nim")
	exec.Command("latte", "install", "nim").Run()
}
