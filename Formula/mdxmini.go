package main

// Mdxmini - Plays music in X68000 MDX chiptune format
// Homepage: https://github.com/mistydemeo/mdxmini/

import (
	"fmt"
	
	"os/exec"
)

func installMdxmini() {
	// Método 1: Descargar y extraer .tar.gz
	mdxmini_tar_url := "https://github.com/mistydemeo/mdxmini/archive/refs/tags/v2.0.0.tar.gz"
	mdxmini_cmd_tar := exec.Command("curl", "-L", mdxmini_tar_url, "-o", "package.tar.gz")
	err := mdxmini_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdxmini_zip_url := "https://github.com/mistydemeo/mdxmini/archive/refs/tags/v2.0.0.zip"
	mdxmini_cmd_zip := exec.Command("curl", "-L", mdxmini_zip_url, "-o", "package.zip")
	err = mdxmini_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdxmini_bin_url := "https://github.com/mistydemeo/mdxmini/archive/refs/tags/v2.0.0.bin"
	mdxmini_cmd_bin := exec.Command("curl", "-L", mdxmini_bin_url, "-o", "binary.bin")
	err = mdxmini_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdxmini_src_url := "https://github.com/mistydemeo/mdxmini/archive/refs/tags/v2.0.0.src.tar.gz"
	mdxmini_cmd_src := exec.Command("curl", "-L", mdxmini_src_url, "-o", "source.tar.gz")
	err = mdxmini_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdxmini_cmd_direct := exec.Command("./binary")
	err = mdxmini_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
