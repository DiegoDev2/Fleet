package main

// Gemgen - Command-line tool for converting Commonmark Markdown to Gemtext
// Homepage: https://sr.ht/~kota/gemgen/

import (
	"fmt"
	
	"os/exec"
)

func installGemgen() {
	// Método 1: Descargar y extraer .tar.gz
	gemgen_tar_url := "https://git.sr.ht/~kota/gemgen/archive/v0.6.0.tar.gz"
	gemgen_cmd_tar := exec.Command("curl", "-L", gemgen_tar_url, "-o", "package.tar.gz")
	err := gemgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gemgen_zip_url := "https://git.sr.ht/~kota/gemgen/archive/v0.6.0.zip"
	gemgen_cmd_zip := exec.Command("curl", "-L", gemgen_zip_url, "-o", "package.zip")
	err = gemgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gemgen_bin_url := "https://git.sr.ht/~kota/gemgen/archive/v0.6.0.bin"
	gemgen_cmd_bin := exec.Command("curl", "-L", gemgen_bin_url, "-o", "binary.bin")
	err = gemgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gemgen_src_url := "https://git.sr.ht/~kota/gemgen/archive/v0.6.0.src.tar.gz"
	gemgen_cmd_src := exec.Command("curl", "-L", gemgen_src_url, "-o", "source.tar.gz")
	err = gemgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gemgen_cmd_direct := exec.Command("./binary")
	err = gemgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: scdoc")
	exec.Command("latte", "install", "scdoc").Run()
}
