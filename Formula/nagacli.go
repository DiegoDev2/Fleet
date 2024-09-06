package main

// NagaCli - Shader translation command-line tool
// Homepage: https://github.com/gfx-rs/naga

import (
	"fmt"
	
	"os/exec"
)

func installNagaCli() {
	// Método 1: Descargar y extraer .tar.gz
	nagacli_tar_url := "https://github.com/gfx-rs/naga/archive/refs/tags/v0.14.0.tar.gz"
	nagacli_cmd_tar := exec.Command("curl", "-L", nagacli_tar_url, "-o", "package.tar.gz")
	err := nagacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nagacli_zip_url := "https://github.com/gfx-rs/naga/archive/refs/tags/v0.14.0.zip"
	nagacli_cmd_zip := exec.Command("curl", "-L", nagacli_zip_url, "-o", "package.zip")
	err = nagacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nagacli_bin_url := "https://github.com/gfx-rs/naga/archive/refs/tags/v0.14.0.bin"
	nagacli_cmd_bin := exec.Command("curl", "-L", nagacli_bin_url, "-o", "binary.bin")
	err = nagacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nagacli_src_url := "https://github.com/gfx-rs/naga/archive/refs/tags/v0.14.0.src.tar.gz"
	nagacli_cmd_src := exec.Command("curl", "-L", nagacli_src_url, "-o", "source.tar.gz")
	err = nagacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nagacli_cmd_direct := exec.Command("./binary")
	err = nagacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
